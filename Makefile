DB_URL=postgresql://root:secret@localhost:5432/gojo?sslmode=disable


postgres:
	docker run --name postgresGOJO -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.0-alpine3.18

redis:
	docker run --name redisGOJO -p 6379:6379 -d redis:7.2-alpine3.18

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

server:
	go fmt
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
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

db:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

build:
	go clean -x
	go clean -cache -x
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -v -ldflags "-w -s -extldflags '-static'" \
	-gcflags="-S -m" -trimpath -mod=readonly -buildmode=pie -a -installsuffix nocgo \
	-o gojo .

restart:
	docker stop redisGOJO postgresGOJO
	docker start redisGOJO postgresGOJO

dcs:
	docker stop redisGOJO postgresGOJO
	docker compose build --no-cache
	docker compose up

dcd:
	docker compose down -v

.PHONY: postgres redis createdb dropdb mgup mgdown mgup1 mgdown1 nmg sqlc graph test server mock proto evans db build restart dcs dcd