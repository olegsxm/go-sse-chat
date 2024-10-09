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

func signIn(w http.ResponseWriter, r *http.Request) error {
	_, e := w.Write([]byte("Hello, World!"))

	return e
}
