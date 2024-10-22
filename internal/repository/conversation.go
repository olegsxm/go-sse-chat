package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/olegsxm/go-sse-chat/ent"
	"github.com/olegsxm/go-sse-chat/ent/conversation"
	"github.com/olegsxm/go-sse-chat/internal/models"
)

type ConversationRepository struct {
	db *ent.Client
}

func (r *ConversationRepository) Get(ctx context.Context) ([]*ent.Conversation, error) {

	return r.db.Conversation.Query().
		WithMessages(func(query *ent.MessageQuery) {}).
		//Where(conversation.HasMessages()).
		Limit(50).
		All(ctx)
}

func (r *ConversationRepository) GetByID(ctx context.Context, id string) (*ent.Conversation, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return r.db.Conversation.Get(ctx, uid)
}

func (r *ConversationRepository) Create(ctx context.Context, senderID string) (*ent.Conversation, error) {
	return r.db.Conversation.Create().
		AddUser(&ent.User{
			ID: uuid.MustParse(senderID),
		}).
		Save(ctx)
}

func (r *ConversationRepository) Update(ctx context.Context, conv models.Conversation) error {
	uid, err := uuid.Parse(conv.Id)
	if err != nil {
		return err
	}

	_, err = r.db.Conversation.Update().Where(conversation.ID(uid)).
		SetNillableName(conv.Name).
		SetNillableAvatar(conv.Avatar).
		Save(ctx)

	return err
}

func (r *ConversationRepository) Delete(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = r.db.Conversation.Delete().Where(conversation.ID(uid)).Exec(ctx)

	return err
}
