package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"server-bof/config"
)

var DB *sql.DB

func MysqlConnection() {
	dbAddr := config.GetEnv("DB_ADDRESS")
	dbName := config.GetEnv("DB_NAME")
	dbUser := config.GetEnv("DB_USER")
	dbPass := config.GetEnv("DB_PASS")
	driverName := config.GetEnv("DB_DRIVER")

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbAddr, dbName)
	db, err := sql.Open(driverName, dns)

	if err != nil {
		log.Println(err)
		log.Println(dns)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
	}

	DB = db
}
