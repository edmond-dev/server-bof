#--------------------------------------------------------------------------------
#================================|   MYSQL   |===================================
#--------------------------------------------------------------------------------
mysql:
	docker run --name mysql80 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:latest

mysqlbash:
	docker exec -it mysql80 bash

createmysqldb:
	docker exec -it mysql80 mysql --user=root --password=secret --execute="CREATE DATABASE bigouncefarms;"

createmysqltestdb:
	docker exec -it mysql80 mysql --user=root --password=secret --execute="CREATE DATABASE bigouncefarms_test;"

dropmysqldb:
	docker exec -it mysql80 mysql --user=root --password=secret --execute="DROP DATABASE bigouncefarms;"

migratemysqlup:
	migrate -path db/migration -database "mysql://root:secret@tcp(127.0.0.1:3306)/bigouncefarms"  up

migratemysqldown:
	migrate -path db/migration -database "mysql://root:secret@tcp(127.0.0.1:3306)/bigouncefarms"  down

sqlc:
	sqlc generate

test:
	go test  -v -cover ./db/...


.PHONY: createmysqldb dropmysqldb migratemysqlup migratemysqldown sqlc test createmysqltestdb
