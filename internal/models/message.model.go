package models

import "time"

type Message struct {
	ID             int64     `json:"id"`
	Message        string    `json:"message"`
	SenderId       int64     `json:"senderId"`
	ConversationId int64     `json:"conversation"`
	Read           int64     `json:"read"`
	CreatedAt      time.Time `json:"createdAt"`
	Sender         *UserDTO  `json:"sender"`
}
