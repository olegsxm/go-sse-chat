package repository

import (
	database "chat/internal/db"
	"chat/internal/repository/queries"
)

type Repository struct {
	auth AuthRepository
}

func (r Repository) Auth() AuthRepository {
	return r.auth
}

func New(db database.DB) Repository {
	q := queries.New(db.Pg)

	return Repository{
		auth: NewAuthRepository(q),
	}
}
