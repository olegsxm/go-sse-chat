package use_cases

import "github.com/olegsxm/go-sse-chat.git/internal/interfaces"

type IRepository interface {
	Auth()
}

type UseCases struct {
	auth interfaces.IAuth
}

func (c UseCases) Auth() interfaces.IAuth {
	return c.auth
}

func New(r IRepository) UseCases {
	return UseCases{
		auth: AuthUC{},
	}
}
