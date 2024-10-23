package services

import (
	"context"
	"github.com/olegsxm/go-sse-chat/ent"
	"github.com/olegsxm/go-sse-chat/internal/models"
)

type MessageRepo interface {
	Create(ctx context.Context, id string, conversation string, m string) (*ent.Message, error)
	Get(ctx context.Context, conversationID string) ([]*ent.Message, error)
}

type MessageService struct {
	r MessageRepo
}

func (s *MessageService) Create(ctx context.Context, id string, conversation string, msg string) (models.MessageResponse, error) {
	var res models.MessageResponse

	m, err := s.r.Create(ctx, id, conversation, msg)
	if err != nil {
		return res, err
	}

	res.ID = m.ID.String()
	res.Message = m.Message
	res.CreatedAt = m.CreatedAt.String()

	return res, err
}

func (s *MessageService) Get(ctx context.Context, conversationID string) (models.Messages, error) {
	msgs, err := s.r.Get(ctx, conversationID)

	if err != nil {
		return []models.Message{}, err
	}

	messages := make([]models.Message, len(msgs), len(msgs))

	for i, m := range msgs {
		sender, err := m.QueryUser().First(ctx)
		if err != nil {
			return []models.Message{}, err
		}

		messages[i] = models.Message{
			Id:        m.ID.String(),
			Message:   m.Message,
			CreatedAt: m.CreatedAt.UTC().String(),
			Read:      m.Read,
			Sender: models.User{
				ID:    sender.ID.String(),
				Login: sender.Login,
			},
		}
	}

	return messages, nil
}
