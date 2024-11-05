package models

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	Id             uuid.UUID `json:"id"`
	Message        string    `json:"message"`
	ConversationId uuid.UUID `json:"conversationId"`
	SenderId       uuid.UUID `json:"senderId"`
	CreatedAt      time.Time `json:"createdAt"`
}

type MessageDTO struct {
	Id             string    `json:"id"`
	Message        string    `json:"message"`
	ConversationId string    `json:"conversationId"`
	CreatedAt      time.Time `json:"createdAt"`
	Sender         UserDTO   `json:"sender"`
}

type MessageRequest struct {
	Message string `json:"message"`
}
