package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres driver
)

const (
	HOST   = "localhost"
	PORT   = 5432
	USER   = "postgres"
	DBNAME = "fga"
)

var (
	db  *sql.DB
	err error
)

func GetDB() *sql.DB {
	psqlInfo := fmt.Sprintf("postgres://postgres:%s@%s/%s?sslmode=disable", USER, HOST, DBNAME)

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
