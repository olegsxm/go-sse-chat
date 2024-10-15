package repository

import "database/sql"

type Storage interface {
	Sql() *sql.DB
}

type Repository struct {
	authRepository  *AuthRepository
	chatRepository  *ChatRepository
	usersRepository *UsersRepository
}

func (r *Repository) Auth() *AuthRepository {
	return r.authRepository
}
func (r *Repository) Chat() *ChatRepository {
	return r.chatRepository
}
func (r *Repository) Users() *UsersRepository {
	return r.usersRepository
}

var st Storage

func New(storage Storage) Repository {
	st = storage
	return Repository{
		authRepository:  &AuthRepository{},
		chatRepository:  &ChatRepository{},
		usersRepository: &UsersRepository{},
	}
}
