postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root  e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12  createdb --username=root --owner=root bigouncefarms
dropdb:
	docker exec -it postgres12 dropdb bigouncefarms

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bigouncefarms?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bigouncefarms?sslmode=disable" -verbose down
sqlc:
	sqlc generate

mysql:
	docker run --name mysql8.0 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:8.0

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
