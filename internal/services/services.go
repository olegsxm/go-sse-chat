package services

import "chat/internal/repository"

type Repository interface{}

type Services struct {
	auth AuthService
}

func (s *Services) Auth() AuthService {
	return s.auth
}

func NewServices(r repository.Repository) Services {
	return Services{
		auth: newAuthService(r.Auth()),
	}
}
