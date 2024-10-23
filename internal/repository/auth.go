package repository

import (
	"context"
	"github.com/olegsxm/go-sse-chat/ent"
	"github.com/olegsxm/go-sse-chat/ent/user"
)

type UserRepository struct {
	ent *ent.Client
}

func (a *UserRepository) FindUser(ctx context.Context, login string) (*ent.User, error) {
	return a.ent.User.Query().Where(user.Login(login)).First(ctx)
}

func (a *UserRepository) CreateUser(ctx context.Context, login, password string) (*ent.User, error) {
	return a.ent.User.Create().SetLogin(login).SetPassword(password).Save(ctx)
}