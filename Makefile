DB_URL=postgresql://root:secret@localhost:5432/gojo?sslmode=disable
postgres:
	sudo docker run --name postgresGOJO -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.0-alpine3.18

createdb:
	sudo docker exec -it postgresGOJO createdb --username=root --owner=root  gojo

dropdb:
	sudo docker exec -it postgresGOJO dropdb gojo

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
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
	go test -v -cover -short ./... -race

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/gojo.go github.com/dj-yacine-flutter/gojo/db/database Gojo
##mockgen -package mockwk -destination worker/mock/distributor.go github.com/dj-yacine-flutter/gojo/worker TaskDistributor

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	rm -f doc/statik/*.go
	protoc --proto_path=proto  --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=gojo \
    proto/*.proto 
	statik -src=./doc/swagger -dest=./doc


evans:
	evans --host localhost --port 9090 -r repl

redis:
	sudo docker run --name redisGOJO -p 6379:6379 -d redis:7.2-alpine3.18

db:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

.PHONY: postgres createdb dropdb migrateup migratedown migratedrop sqlc graph test server mock proto evans redis new_migration db