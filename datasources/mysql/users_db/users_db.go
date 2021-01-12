package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type dbCred struct {
	username string
	password string
	host     string
	schema   string
}

var (
	CLient *sql.DB
	err    error
)

func init() {
	// Testing in a local machine
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "password", "127.0.0.1:3306", "users_db")
	CLient, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err := CLient.Ping(); err != nil {
		panic(err)
	}
	// mysql.SetLogger()
	log.Println("Database successfully configured!")
}
