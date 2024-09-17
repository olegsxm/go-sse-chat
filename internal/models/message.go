package models

import "time"

type Message struct {
	ID        int64     `json:"id"`
	Message   string    `json:"message" validate:"required"`
	Sender    int64     `json:"sender" validate:"required"`
	Recipient int64     `json:"recipient" validate:"required"`
	ChatId    int64     `json:"chatId"`
	CreatedAt time.Time `json:"createdAt"`
}
