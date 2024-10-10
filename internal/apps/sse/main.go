package sse

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/olegsxm/go-sse-chat.git/internal/config"
	"github.com/olegsxm/go-sse-chat.git/internal/handlers"
	"github.com/olegsxm/go-sse-chat.git/internal/repository"
	"github.com/olegsxm/go-sse-chat.git/internal/use_cases"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"golang.org/x/net/http2/h2c"

	"github.com/go-chi/chi/v5"

	"golang.org/x/net/http2"

	_ "github.com/swaggo/http-swagger/v2"
)

func Run(ctx context.Context) *http.Server {
	slog.Info("Sse Chat Running")

	cfg, err := config.New()

	if err != nil || cfg == nil {
		slog.Error("Error loading config")
	}

	r := repository.New()
	us := use_cases.New(&r)

	h2s := &http2.Server{
		IdleTimeout: 10 * time.Second,
	}

	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.StripSlashes)

	mux.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(cfg.Swagger.Url),
	))

	handlers.New(ctx, mux, &us)

	server := &http.Server{
		Addr:    cfg.Server.Address,
		Handler: h2c.NewHandler(mux, h2s),
	}

	return server

}
