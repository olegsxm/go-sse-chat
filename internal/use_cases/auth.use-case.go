package use_cases

import "github.com/olegsxm/go-sse-chat.git/internal/models"

type AuthUC struct {
}

func (u AuthUC) SignIn(login, password string) (models.AuthResponse, error) {
	return models.AuthResponse{
		Token: "Token",
		User:  models.UserDTO{},
	}, nil
}
