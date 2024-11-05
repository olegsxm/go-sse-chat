package repository

import (
	"chat/internal/models"
	"chat/internal/repository/queries"
	"context"
	"github.com/google/uuid"
)

type MessageRepository struct {
	queries *queries.Queries
}

func (m MessageRepository) Create(ctx context.Context, message string, conversationID, userID uuid.UUID) (models.Message, error) {
	msg, err := m.queries.CreateMessage(ctx, queries.CreateMessageParams{
		Message:        message,
		SenderID:       userID,
		ConversationID: conversationID,
	})

	if err != nil {
		return models.Message{}, err
	}

	res := models.Message{
		Id:             msg.ID,
		Message:        msg.Message,
		ConversationId: msg.ConversationID,
		SenderId:       msg.SenderID,
		CreatedAt:      msg.CreatedAt.Time,
	}

	return res, nil
}

func NewMessageRepository(queries *queries.Queries) MessageRepository {
	return MessageRepository{
		queries: queries,
	}
}
