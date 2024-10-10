package handlers

import "C"
import (
	"github.com/go-chi/chi/v5"
	"github.com/olegsxm/go-sse-chat.git/pkg/handler"
	"net/http"
)

func authHandlers() *chi.Mux {
	c := chi.NewRouter()

	c.Route("/v1", func(r chi.Router) {
		r.Get("/sign-in", handler.HandleRoute(signIn))
	})

	c.Mount("/auth", c)

	return c
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Router       /accounts/{id} [get]
func signIn(w http.ResponseWriter, r *http.Request) error {
	_, e := w.Write([]byte("Hello, World!"))

	return e
}
