package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

const (
	dbUser    = "root"
	dbPass    = "secret"
	dbAddress = "127.0.0.1:3306"
	dbase     = "bigouncefarms"
)

func TestMain(m *testing.M) {

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbAddress, dbase)
	connection, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("Could not connect to the database.")
	}

	fmt.Println("db connection is: ")
	//fmt.Println(db, database.Connection())
	fmt.Println("end db checking.")

	testQueries = New(connection)
	os.Exit(m.Run())

}
