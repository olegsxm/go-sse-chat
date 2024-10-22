package repository

import (
	"github.com/olegsxm/go-sse-chat/internal/db"
)

type Repository struct {
	auth *UserRepository
	conv *ConversationRepository
}

func (r *Repository) Auth() *UserRepository {
	return r.auth
}

func (r *Repository) Conversations() *ConversationRepository {
	return r.conv
}

func New(d *db.Db) *Repository {
	return &Repository{
		auth: &UserRepository{
			db: d.SQL(),
		},
		conv: &ConversationRepository{
			db: d.SQL(),
		},
	}
}
