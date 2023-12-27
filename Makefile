DB_URL=postgresql://root:secret@localhost:5432/gojo?sslmode=disable


postgres:
	docker run --name postgresGOJO -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.1-alpine3.19

queue:
	docker run --name queueGOJO -p 6370:6379 -d redis:7.2.3-alpine3.19

cache:
	docker run --name cacheGOJO -p 6380:6379 -d redis:7.2.3-alpine3.19

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
	@read -p "Path: " Path; \
	if [ "$$Path" != "" ]; then \
		read -p "Name: " Name; \
		if [ "$$Name" != "" ]; then \
			godepgraph $$Path | dot -Tpng -o "graph/$$Name.png"; \
		fi; \
	fi

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

proto:
	rm -rf pb/*.go
	rm -rf pb/*/*.go
	rm -f doc/swagger/*.swagger.json
	rm -f doc/statik/*.go
	protoc --proto_path=proto --proto_path=proto/uspb --proto_path=proto/nfpb --proto_path=proto/ampb --proto_path=proto/aspb --proto_path=proto/shpb --proto_path=. \
	--go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=gojo \
	proto/*.proto proto/uspb/*.proto proto/nfpb/*.proto  proto/ampb/*.proto proto/aspb/*.proto proto/shpb/*.proto
	find . -name '*.proto' | xargs clang-format -i --verbose
	statik -src=./doc/swagger -dest=./doc

v1:
	protoc --proto_path=proto --proto_path=proto/v1 --proto_path=. \
	--go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/v1/*/*.proto

evans:
	evans --host localhost --port 9090 -r repl

db:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

build: fmt
	go clean -x
	go clean -cache -x
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -v -ldflags "-w -s -extldflags '-static'" \
	-gcflags="-S -m" -trimpath -mod=readonly -buildmode=pie \
	-a -installsuffix nocgo -o gojo .

cgo: fmt
	go clean -x
	go clean -cache -x
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
	go build -v -ldflags "-w -s -extldflags '-static'" \
	-gcflags="-S -m" -trimpath -mod=readonly -buildmode=pie \
	-a -installsuffix cgo -o gojo .

restart:
	docker stop queueGOJO cacheGOJO postgresGOJO
	docker start queueGOJO cacheGOJO postgresGOJO

dcs:
	docker stop queueGOJO cacheGOJO postgresGOJO
	docker compose build --no-cache
	docker compose up

dcd:
	docker compose down -v

fmt:
	find . -name '*.proto' | xargs clang-format -i --verbose
	find . -name "*.go" -print0 | xargs -0 gofmt -w

.PHONY: postgres queue cache createdb dropdb mgup mgdown mgup1 mgdown1 nmg sqlc graph test server mock proto v1 evans db build cgo restart dcs dcd fmt