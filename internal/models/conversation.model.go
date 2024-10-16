package models

import "database/sql"

type ConversationDTO struct {
	ID      int64    `json:"id"`
	Avatar  string   `json:"avatar"`
	Name    string   `json:"name"`
	Message *Message `json:"message"`
}

func (d *ConversationDTO) ToDTO(c Conversation) {
	d.ID = c.ID
	if c.Name.Valid {
		d.Name = c.Name.String
	} else {
		d.Name = ""
	}
}

type Conversation struct {
	ID      int64
	Name    sql.NullString
	Created int64
}

func (c *Conversation) ToDTO() ConversationDTO {
	d := ConversationDTO{}

	d.ToDTO(*c)

	return d
}

type ConversationParticipants struct {
	ConversationId int64
	UserId         int64
}

type NewConversationRequest struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}
