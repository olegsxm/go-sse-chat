package services

import (
	"chat/internal/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"log/slog"
)

type conversationRepository interface {
	CreateConversation(ctx context.Context) (models.Conversation, error)
	GetConversations(ctx context.Context) ([]models.Conversation, error)
	GetConversationById(ctx context.Context, id string) (models.Conversation, error)
	SetConversationName(ctx context.Context, id, name string) error
	CreateConversationWithMessage(ctx context.Context, senderUid uuid.UUID, participantUid uuid.UUID, message string) (models.Conversation, models.Message, error)
}

type participantRepository interface {
	Create(ctx context.Context, conversation, userID uuid.UUID) error
	GetConversationParticipants(ctx context.Context, conversationID uuid.UUID) ([]models.UserDTO, error)
}

type messageRepository interface {
	Create(ctx context.Context, message string, conversationID, userID uuid.UUID) (models.Message, error)
}

type ConversationService struct {
	conversationRepository conversationRepository
	participantRepository  participantRepository
	messageRepository      messageRepository
}

func (c ConversationService) CreateConversationWithMessage(ctx context.Context, message, senderID, participantID string) (models.ConversationDTO, error) {
	senderUid, err := uuid.Parse(senderID)
	if err != nil {
		return models.ConversationDTO{}, err
	}

	participantUid, err := uuid.Parse(participantID)
	if err != nil {
		return models.ConversationDTO{}, err
	}

	conv, m, err := c.conversationRepository.CreateConversationWithMessage(ctx, senderUid, participantUid, message)
	if err != nil {
		return models.ConversationDTO{}, err
	}

	// TODO Set dialog name
	res := models.ConversationDTO{
		Id:      conv.Id.String(),
		Name:    conv.Name,
		Private: conv.Private,
		Created: conv.Created,
		Message: &models.MessageDTO{
			Id:             m.Id.String(),
			Message:        m.Message,
			ConversationId: m.ConversationId.String(),
			CreatedAt:      m.CreatedAt,
			Sender: models.UserDTO{
				Id: m.SenderId.String(),
			},
		},
	}

	return res, nil
}

func (c ConversationService) SetConversationName(ctx context.Context, conversationID, name string) error {
	err := c.conversationRepository.SetConversationName(ctx, conversationID, name)

	if err != nil {
		slog.Error("set conversation name error", err.Error()) // TODO Create custom error
	}

	return err
}

func (c ConversationService) CreateMessage(ctx context.Context, conversationID, senderID, message string) (models.MessageDTO, []models.UserDTO, error) {
	var err error
	cUid, err := uuid.Parse(conversationID)
	if err != nil {
		return models.MessageDTO{}, nil, errors.New("conversation id parse error")
	}

	senderUid, err := uuid.Parse(senderID)
	if err != nil {
		slog.Error("conversation id parse error", err.Error())
		return models.MessageDTO{}, nil, errors.New("sender id parse error")
	}

	m, e := c.messageRepository.Create(ctx, message, cUid, senderUid)
	if e != nil {
		slog.Error("conversation message error", e.Error())
		return models.MessageDTO{}, nil, errors.New("message create error")
	}

	res := models.MessageDTO{
		Id:             m.Id.String(),
		Message:        m.Message,
		ConversationId: m.ConversationId.String(),
		CreatedAt:      m.CreatedAt,
		Sender: models.UserDTO{
			Id: m.SenderId.String(),
		},
	}

	participants, err := c.participantRepository.GetConversationParticipants(ctx, cUid)
	if err != nil {
		slog.Error("conversation participants error", err.Error())
		return models.MessageDTO{}, nil, err
	}
	return res, participants, nil
}

// CONSTRUCTOR
type conversationServiceDependencies struct {
	conversationRepository conversationRepository
	participantRepository  participantRepository
	messageRepository      messageRepository
}

func NewConversationService(d conversationServiceDependencies) ConversationService {
	return ConversationService{
		conversationRepository: d.conversationRepository,
		participantRepository:  d.participantRepository,
		messageRepository:      d.messageRepository,
	}
}
