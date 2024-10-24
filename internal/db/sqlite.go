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
	slog.Info("Creating new sqlite db file")

	createUsersTableSQL := `
		CREATE TABLE IF NOT EXISTS users (
			id INT PRIMARY KEY,
			login TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL
		)
`

	createConversationsTableSQL := `
		CREATE TABLE IF NOT EXISTS conversations (
		    id INT PRIMARY KEY,
		    name Text DEFAULT NULL,
		    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		) 
	`
	createMessagesTableSQL := `CREATE TABLE IF NOT EXISTS messages (
    	id INTEGER primary key autoincrement ,
    	conversation_id INT not null,
    	sender_id INT NOT NULL,
    	message TEXT NOT NULL,
    	read INT default 0,
    	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    	FOREIGN KEY (conversation_id) REFERENCES conversations(id),
  		FOREIGN KEY (sender_id) REFERENCES users(id)
	)`

	createConversationParticipantsTableSQL := `
		CREATE TABLE IF NOT EXISTS conversation_participants (
		  conversation_id INT NOT NULL,
		  user_id INT NOT NULL,
		  PRIMARY KEY (conversation_id, user_id),
		  FOREIGN KEY (conversation_id) REFERENCES conversations(id),
		  FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`

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

	_, e = tx.Exec(createConversationsTableSQL)
	if e != nil {
		_ = tx.Rollback()
		panic(e)
	}

	_, e = tx.Exec(createMessagesTableSQL)
	if e != nil {
		_ = tx.Rollback()
		panic(e)
	}

	_, e = tx.Exec(createConversationParticipantsTableSQL)
	if e != nil {
		_ = tx.Rollback()
		panic(e)
	}

	_ = tx.Commit()
}
