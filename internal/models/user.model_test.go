package models

import (
	"testing"
)

func TestUserModel(t *testing.T) {
	user := User{
		Login:    "login",
		Password: "password",
	}

	err := user.SaltPassword()
	if err != nil {
		t.Error(err)
		return
	}

	if user.Password == "password" {
		t.Error("password should not match")
	}
}
