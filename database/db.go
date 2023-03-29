package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres driver
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = ""
	DBNAME   = "fga"
)

var (
	db  *sql.DB
	err error
)

func GetDB() *sql.DB {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", USER, PASSWORD, HOST, PORT, DBNAME)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}
