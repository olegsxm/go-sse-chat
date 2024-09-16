package db

import (
	"database/sql"
	"os"

	"github.com/gofiber/fiber/v2/log"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var err error

func NewSqlDb() *sql.DB {
	db, err = sql.Open("sqlite3", "chat.db")

	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	return db
}

func InitDb(db *sql.DB) {
	file, err := os.ReadFile("./sql/init-user-table.sql")
	if err != nil {
		log.Fatal("Error init user table file error: ", err)
	}

	initUserTableQuery := string(file)

	_, e := db.Exec(initUserTableQuery)

	if e != nil {
		log.Fatal("Error init user table error: ", e)
	}
}
