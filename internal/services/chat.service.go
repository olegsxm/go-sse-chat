package services

import (
	"github.com/olegsxm/go-sse-chat.git/internal/models"
	"github.com/olegsxm/go-sse-chat.git/internal/repository"
)

type ChatService struct {
	r *repository.Repository
}

func (s *ChatService) CreateConversation(from int64, to int64) (models.ConversationDTO, error) {
	var err error
	dto := models.ConversationDTO{}

	conversation, err := s.r.Chat().CreateConversation(from, to)
	if err != nil {
		return dto, err
	}

	dto.ToDTO(conversation)

	return dto, err
}

func (s *ChatService) GetConversation() ([]models.ConversationDTO, error) {
	return s.r.Chat().GetConversations()
}

func newChatService(r *repository.Repository) *ChatService {
	return &ChatService{r: r}
}
