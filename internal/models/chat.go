package models

import (
	"time"
)

type Chat struct {
	Id        int64      `json:"id"`
	Name      *string    `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ChatResponse struct {
	Chat
	CompanionName string `json:"companion_name"`
}
