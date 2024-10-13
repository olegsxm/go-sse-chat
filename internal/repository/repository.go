package repository

import "database/sql"

type Storage interface {
	Sql() *sql.DB
}

type Repository struct {
	authRepository *AuthRepository
	chatRepository *ChatRepository
}

func (r *Repository) Auth() *AuthRepository {
	return r.authRepository
}
func (r *Repository) Chat() *ChatRepository {
	return r.chatRepository
}

var st Storage

func New(storage Storage) Repository {
	st = storage
	return Repository{
		authRepository: &AuthRepository{},
		chatRepository: &ChatRepository{},
	}
}
