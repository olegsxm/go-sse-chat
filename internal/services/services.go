package services

import (
	"github.com/olegsxm/go-sse-chat.git/internal/repository"
)

type Services struct {
	auth  *AuthService
	chat  *ChatService
	users *UsersService
}

func (s *Services) Auth() *AuthService {
	return s.auth
}

func (s *Services) Chat() *ChatService { return s.chat }

func (s *Services) Users() *UsersService { return s.users }

func New(r *repository.Repository) Services {
	return Services{
		auth:  newAuthService(r),
		chat:  newChatService(r),
		users: newUsersService(r),
	}
}
