package repository

import (
	"context"
	"github.com/olegsxm/go-sse-chat/ent"
	"github.com/olegsxm/go-sse-chat/ent/user"
)

type UserRepository struct {
	db *ent.Client
}

func (a *UserRepository) FindUser(ctx context.Context, login string) (*ent.User, error) {
	return a.db.User.Query().Where(user.Login(login)).First(ctx)
}

func (a *UserRepository) CreateUser(ctx context.Context, login, password string) (*ent.User, error) {
	return a.db.User.Create().SetLogin(login).SetPassword(password).Save(ctx)
}
