package models

import (
	"database/sql"
	"time"
)

type Chat struct {
	Id        int64          `json:"id"`
	Name      sql.NullString `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}
