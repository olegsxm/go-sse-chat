package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
)

type IUseCase interface {
	Auth()
}

func New(ctx context.Context, mux *chi.Mux, cases IUseCase) {
	mux.Mount("/api/v1", authHandlers())
}
