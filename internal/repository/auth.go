package repository

import (
	"chat/internal/models"
	"chat/internal/repository/queries"
	"context"
)

type AuthRepository struct {
	queries *queries.Queries
}

func (a AuthRepository) FindUserByLogin(ctx context.Context, login string) (models.User, error) {
	u, err := a.queries.GetUserByLogin(ctx, login)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Id:       u.ID,
		Login:    u.Login,
		Password: u.Password,
		Salt:     u.Salt,
	}

	return user, nil
}

func (a AuthRepository) CreateUser(ctx context.Context, login string, hash string, salt []byte) (models.User, error) {
	u, err := a.queries.CreateUser(ctx, queries.CreateUserParams{
		Login:    login,
		Password: hash,
		Salt:     salt,
	})

	if err != nil {
		return models.User{}, err
	}

	return models.User{Id: u.ID, Login: u.Login}, nil
}

func NewAuthRepository(q *queries.Queries) AuthRepository {
	return AuthRepository{
		queries: q,
	}
}
