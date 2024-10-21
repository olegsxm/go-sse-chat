package models

import (
	"golang.org/x/crypto/bcrypt"
)

//easyjson:skip
type User struct {
	ID       int64
	Login    string
	Password string
}

func (u *User) SaltPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err == nil {
		u.Password = string(bytes)
	}
	return err
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) ToDTO() UserDTO {
	return UserDTO{
		ID:    u.ID,
		Login: u.Login,
	}
}

//easyjson:json
type UserDTO struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
}
