package services

import (
	"github.com/olegsxm/go-sse-chat.git/internal/models"
)

type AuthService struct {
}

func (s AuthService) SignIn(login, password string) (models.User, error) {
	user := models.User{}
	return user, nil
}

func (s AuthService) SignUp(login, password string) (models.User, error) {
	user := models.User{}
	return user, nil
}

func newAuthService() AuthService {
	return AuthService{}
}
