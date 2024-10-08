package sse

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/olegsxm/go-sse-chat.git/pkg/handler"

	"golang.org/x/net/http2/h2c"

	"github.com/go-chi/chi/v5"

	"golang.org/x/net/http2"
)

func Run(context.Context) error {
	slog.Info("Sse Chat Running")
	h2s := &http2.Server{
		IdleTimeout: 10 * time.Second,
	}

	mux := chi.NewRouter()

	mux.Get("/", handler.HandleRoute(func(writer http.ResponseWriter, request *http.Request) error {
		_, err := writer.Write([]byte("Hello, World!"))
		return err
	}))

	server := &http.Server{
		Addr:    ":443",
		Handler: h2c.NewHandler(mux, h2s),
	}

	if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
