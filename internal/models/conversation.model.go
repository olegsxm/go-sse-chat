package models

//easyjson:json
type Conversation struct {
	Id     string  `json:"id"`
	Name   *string `json:"name"`
	Avatar *string `json:"avatar"`
}

//easyjson:json
type Conversations []Conversation
