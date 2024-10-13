package repository

import "database/sql"

type Storage interface {
	Sql() *sql.DB
}

type Repository struct {
	authRepository *AuthRepository
}

func (r *Repository) Auth() *AuthRepository {
	return r.authRepository
}

var st Storage

func New(storage Storage) Repository {
	st = storage
	return Repository{
		authRepository: &AuthRepository{},
	}
}
