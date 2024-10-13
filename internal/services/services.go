package services

import (
	"github.com/olegsxm/go-sse-chat.git/internal/repository"
)

type Services struct {
	auth *AuthService
	chat *ChatService
}

func (s *Services) Auth() *AuthService {
	return s.auth
}

func (s *Services) Chat() *ChatService { return s.chat }

func New(r *repository.Repository) Services {
	return Services{
		auth: newAuthService(r),
		chat: newChatService(r),
	}
}
