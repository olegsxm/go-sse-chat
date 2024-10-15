package models

import "database/sql"

type ConversationDTO struct {
	ID      int64          `json:"id"`
	Avatar  string         `json:"avatar"`
	Name    sql.NullString `json:"name"`
	Sender  *User          `json:"sender"`
	Message Message        `json:"message"`
}

type Conversation struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Created int64  `json:"created"`
}
