package controllers

import (
	"context"

	"github.com/olegsxm/go-sse-chat.git/internal/interfaces"

	"github.com/go-chi/chi/v5"
)

type IUseCase interface {
	Auth() interfaces.IAuth
}

var uc IUseCase

func New(ctx context.Context, router chi.Router, cases IUseCase) {
	uc = cases

	router.Mount("/v1", authHandlers())
}
