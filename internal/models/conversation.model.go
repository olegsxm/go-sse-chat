package models

//easyjson:json
type Conversation struct {
	Id      string   `json:"id"`
	Name    *string  `json:"name"`
	Avatar  *string  `json:"avatar"`
	Message *Message `json:"message"`
}

//easyjson:json
type Conversations []Conversation
