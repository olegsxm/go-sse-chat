package services

import (
	"context"
	"github.com/olegsxm/go-sse-chat/ent"
	"github.com/olegsxm/go-sse-chat/ent/message"
	"github.com/olegsxm/go-sse-chat/internal/models"
)

type crudRepo interface {
	Get(ctx context.Context, clientID string) ([]*ent.Conversation, error)
	GetByID(ctx context.Context, id string) (*ent.Conversation, error)
	Create(ctx context.Context, id string) (*ent.Conversation, error)
	Update(ctx context.Context, conv models.Conversation) error
	Delete(ctx context.Context, id string) error
}

type ConversationService struct {
	r crudRepo
}

func (c *ConversationService) FindAll(ctx context.Context, clientID string) (*models.Conversations, error) {
	all, err := c.r.Get(ctx, clientID)
	if err != nil {
		return nil, err
	}

	cs := make(models.Conversations, len(all), len(all))

	for i, v := range all {
		cs[i] = models.Conversation{
			Id:     v.ID.String(),
			Avatar: v.Avatar,
			Name:   v.Name,
		}

		msg, e := v.QueryMessages().Order(ent.Desc(message.FieldCreatedAt)).First(ctx)
		if e == nil {
			cs[i].Message = &models.Message{
				Id:        msg.ID.String(),
				Message:   msg.Message,
				Read:      msg.Read,
				CreatedAt: msg.CreatedAt.UTC().String(),
			}
		}
	}

	return &cs, nil
}

func (c *ConversationService) FindByID(ctx context.Context, id string) (models.Conversation, error) {
	cs, err := c.r.GetByID(ctx, id)
	if err != nil {
		return models.Conversation{}, err
	}

	conv := models.Conversation{
		Id:     cs.ID.String(),
		Avatar: cs.Avatar,
		Name:   cs.Name,
	}

	return conv, nil
}

func (c *ConversationService) Create(ctx context.Context, senderID string) (models.Conversation, error) {
	conv, err := c.r.Create(ctx, senderID)
	if err != nil {
		return models.Conversation{}, err
	}

	return models.Conversation{
		Id:     conv.ID.String(),
		Name:   conv.Name,
		Avatar: conv.Avatar,
	}, err
}

func (c *ConversationService) Update(ctx context.Context) (models.Conversation, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ConversationService) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
