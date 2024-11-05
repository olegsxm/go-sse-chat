package repository

import (
	"chat/internal/repository/queries"
	"context"
	"github.com/google/uuid"
)

type ParticipantRepository struct {
	queries *queries.Queries
}

func (p ParticipantRepository) Create(ctx context.Context, conversation, userID uuid.UUID) error {
	return p.queries.CreateParticipant(ctx, queries.CreateParticipantParams{
		ConversationID: conversation,
		UserID:         userID,
	})
}

func NewParticipantRepository(queries *queries.Queries) ParticipantRepository {
	return ParticipantRepository{
		queries: queries,
	}
}
