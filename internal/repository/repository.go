package repository

import (
	"github.com/olegsxm/go-sse-chat/internal/db"
)

type Repository struct {
	auth *UserRepository
	conv *ConversationRepository
	msg  *MessageRepository
}

func (r *Repository) Auth() *UserRepository {
	return r.auth
}

func (r *Repository) Conversations() *ConversationRepository {
	return r.conv
}

func (r *Repository) Messages() *MessageRepository {
	return r.msg
}

func New(d *db.Db) *Repository {
	return &Repository{
		auth: &UserRepository{
			ent: d.SQL(),
		},
		conv: &ConversationRepository{
			ent: d.SQL(),
		},
		msg: &MessageRepository{
			ent: d.SQL(),
		},
	}
}
