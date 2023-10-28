DB_URL=postgresql://root:secret@localhost:5432/gojo?sslmode=disable
postgres:
	sudo docker run --name postgresGOJO -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.0-alpine3.18

createdb:
	sudo docker exec -it postgresGOJO createdb --username=root --owner=root  gojo

dropdb:
	sudo docker exec -it postgresGOJO dropdb gojo

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
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/gojo.go github.com/dj-yacine-flutter/gojo/db/database Gojo
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/dj-yacine-flutter/gojo/worker TaskDistributor

proto:
	rm -rf pb/*.go
	rm -rf pb/*/*.go
	rm -f doc/swagger/*.swagger.json
	rm -f doc/statik/*.go
	protoc --proto_path=proto --proto_path=proto/uspb --proto_path=proto/nfpb --proto_path=proto/ampb --proto_path=proto/aspb --proto_path=. \
	--go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=gojo \
	proto/*.proto proto/uspb/*.proto proto/nfpb/*.proto  proto/ampb/*.proto proto/aspb/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

redis:
	sudo docker run --name redisGOJO -p 6379:6379 -d redis:7.2-alpine3.18

db:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

build:
	GOOS=linux GOARCH=amd64 go build -v -ldflags "-s -w" -gcflags="-S -m" -trimpath -mod=readonly -buildmode=pie -a -o gojo .

restart:
	sudo docker stop redisGOJO postgresGOJO
	sudo docker start redisGOJO postgresGOJO

dcs:
	sudo docker stop redisGOJO postgresGOJO
	sudo docker compose build
	sudo docker compose up

dcd:
	sudo docker compose down
	docker volume rm gojo_data-volume

.PHONY: postgres createdb dropdb mgup mgdown mgup1 mgdown1 nmg sqlc graph test server mock proto evans redis  db build restart dcs dcd