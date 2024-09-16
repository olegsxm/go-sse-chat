package services

import repo "github.com/olegsxm/go-sse-chat.git/internal/repository"

type Services struct {
	Auth *authService
}

var repository *repo.Repository

func New(repo *repo.Repository) *Services {
	repository = repo

	return &Services{
		Auth: &authService{},
	}
}
