package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	Login    string
	Password string
	Salt     []byte
}

func (user *User) ToDTO() UserDTO {
	return UserDTO{
		Id:    user.Id.String(),
		Login: user.Login,
	}
}

type UserDTO struct {
	Id    string `json:"id"`
	Login string `json:"login"`
}
