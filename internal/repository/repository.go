package repository

import (
	database "chat/internal/db"
	"chat/internal/repository/queries"
)

type Repository struct {
	auth         AuthRepository
	conversation ConversationRepository
	participant  ParticipantRepository
	message      MessageRepository
}

func (r Repository) Auth() AuthRepository {
	return r.auth
}

func (r Repository) Conversation() ConversationRepository {
	return r.conversation
}

func (r Repository) Participant() ParticipantRepository {
	return r.participant
}

func (r Repository) Message() MessageRepository {
	return r.message
}

func New(db database.DB) Repository {
	q := queries.New(db.Pg)

	return Repository{
		auth:         NewAuthRepository(q),
		conversation: NewConversationRepository(q, db.Pg),
		participant:  NewParticipantRepository(q),
		message:      NewMessageRepository(q),
	}
}
