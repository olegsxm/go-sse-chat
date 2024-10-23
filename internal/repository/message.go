package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/olegsxm/go-sse-chat/ent"
	"github.com/olegsxm/go-sse-chat/ent/conversation"
	"github.com/olegsxm/go-sse-chat/ent/message"
)

type MessageRepository struct {
	ent *ent.Client
}

func (m *MessageRepository) Create(ctx context.Context, id string, conversation string, msg string) (*ent.Message, error) {
	return m.ent.Message.Create().
		SetMessage(msg).
		SetConversationID(uuid.MustParse(conversation)).
		SetUserID(uuid.MustParse(id)).
		Save(ctx)

}

func (m *MessageRepository) Get(ctx context.Context, conversationID string) ([]*ent.Message, error) {
	var convUuid = uuid.MustParse(conversationID)
	return m.ent.Message.Query().
		Where(
			message.HasConversationWith(conversation.IDEQ(convUuid)),
		).
		All(ctx)
}
