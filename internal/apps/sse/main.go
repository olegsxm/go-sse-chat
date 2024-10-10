package sse

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/olegsxm/go-sse-chat.git/pkg/middlewares"

	"github.com/go-chi/cors"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/olegsxm/go-sse-chat.git/internal/config"
	"github.com/olegsxm/go-sse-chat.git/internal/controllers"
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
	uc := use_cases.New(&r)

	h2s := &http2.Server{
		IdleTimeout: 10 * time.Second,
	}

	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.StripSlashes)

	if !cfg.Production {
		slog.Debug("Using cors")
		mux.Use(cors.AllowAll().Handler)
	}

	api := mux.Route("/api", func(r chi.Router) {
		r.Use(middlewares.SetResponseHeaders)
		controllers.New(ctx, r, uc)
	})

	mux.Mount("/", api)

	mux.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(cfg.Swagger.Url),
	))

	server := &http.Server{
		Addr:    cfg.Server.Address,
		Handler: h2c.NewHandler(mux, h2s),
	}

	return server

}
