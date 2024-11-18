init:
	go mod init github.com/Viczdera/bank
	
createdb:
	docker exec -it bank_postgres12 createdb --username=root --owner=root s_bank

dropdb:
	docker exec -it bank_postgres12 dropdb s_bank

postgres:
	docker run --name bank_postgres12 -p 8080:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

migrateup:
	migrate -path db/migrations  -database "postgres://root:secret@localhost:8080/s_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations  -database "postgres://root:secret@localhost:8080/s_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY:
	postgres createdb dropdb migrateup migratedown