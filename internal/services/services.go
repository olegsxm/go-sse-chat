package services

import (
	"chat/internal/pkg/config"
	"chat/internal/repository"
)

type Repository interface{}

type Services struct {
	auth AuthService
}

func (s *Services) Auth() AuthService {
	return s.auth
}

func (s *Services) Conversation() any {
	return nil
}

func NewServices(r repository.Repository, cfg config.Config) Services {
	return Services{
		auth: newAuthService(r.Auth(), cfg.Security.JWTKey, cfg.Security.PasswordPepper),
	}
}
