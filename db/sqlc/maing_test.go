package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"testing"
)

const (
	dbUser    = "root"
	dbPass    = "secret"
	dbAddress = "127.0.0.1:3306"
	dbName    = "bigouncefarms"
	DbDriver  = "mysql"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	//dbAddr := config.GetEnv("DB_ADDRESS")
	//dbName := config.GetEnv("DB_NAME")
	//dbUser := config.GetEnv("DB_USER")
	//dbPass := config.GetEnv("DB_PASS")
	//driverName := config.GetEnv("DB_DRIVER")

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbAddress, dbName)
	connection, err := sql.Open(DbDriver, dns)
	if err != nil {
		log.Fatal("Could not connect to the database.")
	}

	testQueries = New(connection)
	os.Exit(m.Run())

}
