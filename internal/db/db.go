package db

import (
	"database/sql"

	"github.com/gofiber/fiber/v2/log"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func New() *sql.DB {
	db, err := sql.Open("sqlite3", "chat.db")

	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	return db
}
