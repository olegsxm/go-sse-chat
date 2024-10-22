package db

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/olegsxm/go-sse-chat/ent"
	"log"
	"log/slog"
)

func newSqlClient() *ent.Client {
	client, err := ent.Open("sqlite3", "chat.db?_fk=1")
	if err != nil {
		slog.Error("Open sqlite3 failed:", err.Error())
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	return client
}
