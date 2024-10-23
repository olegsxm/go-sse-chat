package services

import "github.com/olegsxm/go-sse-chat/internal/repository"

type Services struct {
	auth    AuthService
	conv    *ConversationService
	message *MessageService
}

func (s Services) Auth() AuthService {
	return s.auth
}

func (s *Services) Conversation() *ConversationService {
	return s.conv
}

func (s *Services) Message() *MessageService {
	return s.message
}

func New(r *repository.Repository) *Services {
	return &Services{
		auth: AuthService{
			r: r.Auth(),
		},
		conv: &ConversationService{
			r: r.Conversations(),
		},
		message: &MessageService{
			r: r.Messages(),
		},
	}
}
