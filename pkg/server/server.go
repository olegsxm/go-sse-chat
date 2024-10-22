package server

import (
	"github.com/go-chi/chi/v5"
	"golang.org/x/net/http2"
	"log"
	"log/slog"
	"net/http"
)

func New(mux *chi.Mux) *http.Server {
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := http2.ConfigureServer(server, &http2.Server{})
	if err != nil {
		slog.Error("Configure server error")
		log.Fatal(err)
	}

	return server
}
