DB_URL=postgresql://root:secret@localhost:5432/gojo?sslmode=disable

run:
	clear
	go run .

postgres:
	docker run --name postgresGOJO -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.1-alpine3.19

queue:
	docker run --name queueGOJO -p 6370:6379 -d redis:7.2.3-alpine3.19

cache:
	docker run --name cacheGOJO -p 6380:6379 -d redis:7.2.3-alpine3.19

meili:
	docker run --name meiliGOJO -p 7700:7700 -d getmeili/meilisearch:v1.6.0-rc.5

createdb:
	docker exec -it postgresGOJO createdb --username=root --owner=root  gojo

dropdb:
	docker exec -it postgresGOJO dropdb gojo

mgup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

mgdown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

mgup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

mgdown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

nmg:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

graph:
	goda graph github.com/dj-yacine-flutter/gojo... | dot -Tsvg -o gojo.svg

test:
	go test -v -cover -count=1 -short ./... -race

server: fmt
	go clean
	go clean -cache -x
	clear
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/gojo.go github.com/dj-yacine-flutter/gojo/db/database Gojo
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/dj-yacine-flutter/gojo/worker TaskDistributor
	mockgen -package mockpg -destination ping/mock/key_genrator.go github.com/dj-yacine-flutter/gojo/ping KeyGenrator

v1: fmt
	rm -rf pb/v1/*.go
	rm -rf pb/v1/*/*.go
	rm -f doc/v1/swagger/*.swagger.json
	rm -f doc/v1/statik/*.go
	protoc --proto_path=proto --proto_path=proto/v1 --proto_path=. \
	--go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/v1/swagger --openapiv2_opt=allow_merge=true,merge_file_name=gojo \
	proto/v1/*/*.proto
	statik -src=./doc/v1/swagger -dest=./doc/v1

evans:
	evans --host localhost --port 9090 -r repl

db:
	dbml2sql --postgres -o doc/v1/schema.sql doc/v1/db.dbml

build: fmt
	go clean -x
	go clean -cache -x
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build -v \
	-tags netgo \
	-ldflags "-w -s -extldflags '-static'" \
	-gcflags="-S -m" \
	-trimpath -mod=readonly -buildmode=pie \
	-a -installsuffix nocgo -o gojo .

cgo: fmt
	go clean -x
	go clean -cache -x
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
	go build -v -ldflags "-w -s -extldflags '-static'" \
	-gcflags="-S -m" -trimpath -mod=readonly -buildmode=pie \
	-a -installsuffix cgo -o gojo .

restart:
	docker stop queueGOJO cacheGOJO postgresGOJO meiliGOJO
	docker start queueGOJO cacheGOJO postgresGOJO meiliGOJO

dcs:
	docker stop queueGOJO cacheGOJO postgresGOJO meiliGOJO
	docker compose build --no-cache
	docker compose up

dcd:
	docker compose down -v

fmt:
	find . -name '*.proto' | xargs clang-format -i --verbose
	find . -name "*.go" -print0 | xargs -0 gofmt -w

clean: dcd
	docker buildx prune --all --force
	docker volume prune --all --force
	docker network prune  --force
	docker system df
	go clean -x
	go clean -cache -x

.PHONY: run postgres queue cache meili createdb dropdb mgup mgdown mgup1 mgdown1 nmg sqlc graph test server mock v1 evans db build cgo restart dcs dcd fmt clean