package repository

import (
	"database/sql"

	"github.com/gofiber/fiber/v2/log"
)

type Repository struct {
	Auth    *authRepository
	Message *messageRepository
	Chat    *chatRepository
}

var db *sql.DB

func New(d *sql.DB) *Repository {
	db = d

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &Repository{
		Auth:    &authRepository{},
		Message: &messageRepository{},
		Chat:    &chatRepository{},
	}
}
