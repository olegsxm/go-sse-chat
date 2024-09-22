package models

type ChatMember struct {
	ID   int64 `json:"id"`
	User int64 `json:"user"`
	Chat int64 `json:"chat"`
}
