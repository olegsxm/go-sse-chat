package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID       int64
	Login    string
	Password string
	Salt     string
	Created  time.Time
	Updated  sql.NullTime
}

func (u *User) ToDTO() UserDTO {
	return UserDTO{
		ID:      u.ID,
		Login:   u.Login,
		Created: u.Created,
		Updated: u.Updated,
	}
}

//easyjson:json
type UserDTO struct {
	ID      int64        `json:"id"`
	Login   string       `json:"login"`
	Created time.Time    `json:"created"`
	Updated sql.NullTime `json:"updated"`
}
