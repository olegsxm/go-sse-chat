package models

import (
	"github.com/google/uuid"
	"time"
)

type Conversation struct {
	Id      uuid.UUID `json:"id"`
	Name    *string   `json:"name"`
	Private bool      `json:"private"`
	Created time.Time `json:"created"`
}

type ConversationDTO struct {
	Id      string      `json:"id"`
	Name    *string     `json:"name"`
	Private bool        `json:"private"`
	Created time.Time   `json:"created"`
	Message *MessageDTO `json:"message"`
}
