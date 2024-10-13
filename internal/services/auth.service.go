package services

import (
	"errors"
	"log/slog"

	"github.com/olegsxm/go-sse-chat.git/internal/models"
	"github.com/olegsxm/go-sse-chat.git/internal/repository"
)

type AuthService struct {
	repository *repository.Repository
}

func (s *AuthService) SignIn(login, password string) (models.User, error) {
	u, e := s.repository.Auth().FindUserByLogin(login)

	if e != nil {
		slog.Error("AuthService sign-in error: ", e.Error())
		return models.User{}, e
	}

	if u.VerifyPassword(password) {
		return u, nil
	}

	return u, errors.New("invalid password")
}

func (s *AuthService) SignUp(login, password string) (models.User, error) {
	user := models.User{
		Login:    login,
		Password: password,
	}

	if err := user.SaltPassword(); err != nil {
		return models.User{}, err
	}

	id, err := s.repository.Auth().CreateUser(login, user.Password)
	if err != nil {
		return models.User{}, err
	}

	user.ID = id

	return user, nil
}

func newAuthService(r *repository.Repository) *AuthService {
	return &AuthService{
		repository: r,
	}
}
