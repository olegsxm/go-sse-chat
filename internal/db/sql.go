package db

import (
	"database/sql"

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
	var err error

	_, err = db.Exec(create_users_table_query)

	if err != nil {
		log.Fatal("Error init user table error: ", err)
	}

	_, err = db.Exec(create_messages_table_query)

	if err != nil {
		log.Fatal("Error init messages table error: ", err)
	}

	_, err = db.Exec(create_chats_table_query)
	if err != nil {
		log.Fatal("Error init chats table error: ", err)
	}

	_, err = db.Exec(create_chat_members_table)
	if err != nil {
		log.Fatal("Error init chats members table error: ", err)
	}
}
