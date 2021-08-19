postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

createdb: 
	docker exec -it postgres createdb --username=root --owner=root fomapi

dropdb: 
	docker exec -it postgres dropdb fomapi

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/postgres?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

rundocker:
	cd docker && docker compose build && docker compose up

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server