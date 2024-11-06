package repository

import (
	"chat/internal/models"
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

func (p ParticipantRepository) GetConversationParticipants(ctx context.Context, conversationID uuid.UUID) ([]models.UserDTO, error) {
	u, err := p.queries.GetConversationParticipants(ctx, conversationID)
	if err != nil {
		return nil, err
	}

	users := make([]models.UserDTO, len(u))

	for i, u := range u {
		users[i] = models.UserDTO{
			Id:    u.ID.String(),
			Login: u.Login,
		}
	}

	return users, nil
}

func NewParticipantRepository(queries *queries.Queries) ParticipantRepository {
	return ParticipantRepository{
		queries: queries,
	}
}
