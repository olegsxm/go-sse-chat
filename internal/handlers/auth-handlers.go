package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/olegsxm/go-sse-chat.git/pkg/handler"
	"net/http"
)

func authHandlers() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/auth/sign-in", handler.HandleRoute(signIn))

	return r
}

func signIn(w http.ResponseWriter, r *http.Request) error {
	_, e := w.Write([]byte("Hello, World!"))

	return e
}
