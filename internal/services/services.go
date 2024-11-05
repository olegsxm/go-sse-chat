package services

import (
	"chat/internal/pkg/config"
	"chat/internal/repository"
)

type Repository interface{}

type Services struct {
	auth         AuthService
	conversation ConversationService
}

func (s *Services) Auth() AuthService {
	return s.auth
}

func (s *Services) Conversation() ConversationService {
	return s.conversation
}

func NewServices(r repository.Repository, cfg config.Config) Services {
	return Services{
		auth: newAuthService(r.Auth(), cfg.Security.JWTKey, cfg.Security.PasswordPepper),
		conversation: NewConversationService(conversationServiceDependencies{
			conversationRepository: r.Conversation(),
			participantRepository:  r.Participant(),
			messageRepository:      r.Message(),
		}),
	}
}
