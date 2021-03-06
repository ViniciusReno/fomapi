postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

createdb: 
	docker exec -it postgres createdb --username=root --owner=root fomapi

dropdb: 
	docker exec -it postgres dropdb fomapi

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fomapi?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fomapi?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/ViniciusReno/fomapi/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock