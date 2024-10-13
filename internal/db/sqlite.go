package db

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbName = "./chat.db"
)

func newSqliteConnection() *sql.DB {
	slog.Info("Creating new sqlite connection")

	if stat, _ := os.Stat(dbName); stat == nil {
		_, err := os.Create(dbName)
		if err != nil {
			panic("Failed to create new sqlite db")
		}
	}

	d, err := sql.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	}

	initSQLDB(d)

	return d
}

func initSQLDB(db *sql.DB) {
	slog.Debug("Creating new sqlite db file")

	createUsersTableSQL := `CREATE TABLE IF NOT EXISTS users (id integer PRIMARY KEY, login TEXT, password TEXT)`

	if err := db.Ping(); err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	_, e := tx.Exec(createUsersTableSQL)
	if e != nil {
		_ = tx.Rollback()
		panic(e)
	}

	_ = tx.Commit()
}
