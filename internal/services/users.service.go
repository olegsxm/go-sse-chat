package services

import (
	"github.com/olegsxm/go-sse-chat.git/internal/models"
	"github.com/olegsxm/go-sse-chat.git/internal/repository"
)

type UsersService struct {
	r *repository.Repository
}

func (s *UsersService) FindUsers(query string, excludedID int64) ([]models.UserDTO, error) {
	users, err := s.r.Users().FindUsers(query, excludedID)

	if err != nil {
		return nil, err
	}

	result := make([]models.UserDTO, len(users))

	for i, user := range users {
		result[i] = user.ToDTO()
	}

	return result, err
}

func newUsersService(r *repository.Repository) *UsersService {
	return &UsersService{
		r: r,
	}
}
