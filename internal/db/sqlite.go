package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func newSqliteConnection() *sql.DB {
	d, err := sql.Open("sqlite3", "./chat.db")
	if err != nil {
		log.Fatal(err)
	}
	return d
}
