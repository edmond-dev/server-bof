package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"server-bof/config"
)

var DB *sql.DB

func Connection() *sql.DB {

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASS"),
		config.GetEnv("DB_ADDRESS"),
		config.GetEnv("DB"))
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
	}

	return db
}

func MysqlConnection() {
	DB = Connection()
}
