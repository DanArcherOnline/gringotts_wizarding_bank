postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine
createdb:
	docker exec -it postgres16 createdb --username=root --owner=root gringotts_wizarding_bank
dropdb:
	docker exec -it postgres16 dropdb gringotts_wizarding_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gringotts_wizarding_bank?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gringotts_wizarding_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gringotts_wizarding_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gringotts_wizarding_bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/danarcheronline/gringotts_wizarding_bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown test server 