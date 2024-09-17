package models

import "time"

type Message struct {
	ID        int64     `json:"id"`
	Message   string    `json:"message"`
	Sender    int64     `json:"sender"`
	Recipient int64     `json:"recipient"`
	ChatId    int64     `json:"chatId"`
	CreatedAt time.Time `json:"createdAt"`
}
