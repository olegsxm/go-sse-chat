package services

import (
	"errors"

	"github.com/olegsxm/go-sse-chat.git/internal/pkg/jwt"

	"github.com/olegsxm/go-sse-chat.git/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
}

func (s *authService) SignIn(login, password string) (string, error) {
	user, err := repository.Auth.FindUserByLogin(login)

	if err != nil {
		return "", err
	}

	if !s.checkPassword(user.Password, password) {
		return "", errors.New("invalid credentials")
	}

	return jwt.CreateToken(user.ID, login)

}

func (s *authService) SignUp(l string, p string) (string, error) {
	user, err := s.FindUserByLogin(l)

	if err != nil {
		return "", err
	}

	if user.ID != 0 {
		return "", errors.New("user already exists")
	}

	user.Login = l
	user.Password, err = s.hashPassword(p)

	if err != nil {
		return "", errors.New("error hashing password")
	}

	userId, err := repository.Auth.CreateUser(user)
	if err != nil {
		return "", err
	}

	user.ID = userId

	token, err := jwt.CreateToken(userId, user.Login)

	if err != nil {
		return "", errors.New("error creating token")
	}

	return token, err
}

func (s *authService) FindUserByLogin(login string) (models.User, error) {
	return repository.Auth.FindUserByLogin(login)
}

/* Password Utils */
func (s *authService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *authService) checkPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
