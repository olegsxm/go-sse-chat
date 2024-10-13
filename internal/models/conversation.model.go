package models

import "database/sql"

type ConversationDTO struct {
	ID      int64          `json:"id"`
	Avatar  string         `json:"avatar"`
	Name    sql.NullString `json:"name"`
	Sender  User           `json:"sender"`
	Message Message        `json:"message"`
}
