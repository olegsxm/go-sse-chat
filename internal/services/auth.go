package services

import (
	"context"
	"github.com/olegsxm/go-sse-chat/ent"
	"github.com/olegsxm/go-sse-chat/internal/models"
	"github.com/olegsxm/go-sse-chat/pkg/http_errors"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

type userRepository interface {
	FindUser(ctx context.Context, login string) (*ent.User, error)
	CreateUser(ctx context.Context, login, password string) (*ent.User, error)
}

type AuthService struct {
	r userRepository
}

func (s AuthService) CreateUser(ctx context.Context, login, password string) (models.User, error) {
	saltPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("Error salt password: ", err.Error(), " password: ", password)
		return models.User{}, http_errors.UnknownError{}
	}

	eu, err := s.r.CreateUser(ctx, login, string(saltPass))
	if err != nil {
		slog.Error("Error creating user: ", err.Error(), " password: ", password, " login: ", login)

		if ent.IsConstraintError(err) {
			return models.User{}, http_errors.UserAlreadyExists{}
		}

		return models.User{}, err
	}

	return models.User{ID: eu.ID.String(), Login: eu.Login}, nil
}

func (s AuthService) FindUser(context context.Context, login, password string) (models.User, error) {
	user, err := s.r.FindUser(context, login)

	if err != nil {
		if ent.IsNotFound(err) {
			return models.User{}, http_errors.UserNotFound{}
		}

		slog.Error(err.Error())

		return models.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return models.User{}, http_errors.InvalidCredentials{}
	}

	return models.User{user.ID.String(), user.Login}, nil
}
