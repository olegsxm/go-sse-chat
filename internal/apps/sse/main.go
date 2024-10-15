package sse

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/olegsxm/go-sse-chat.git/internal/config"
	"github.com/olegsxm/go-sse-chat.git/internal/controllers"
	"github.com/olegsxm/go-sse-chat.git/internal/db"
	"github.com/olegsxm/go-sse-chat.git/internal/repository"
	"github.com/olegsxm/go-sse-chat.git/internal/services"
	"golang.org/x/net/http2/h2c"

	"golang.org/x/net/http2"

	_ "github.com/swaggo/http-swagger/v2"
)

func New(ctx context.Context, cfg *config.AppConfig) *http.Server {
	slog.Info("Sse Chat Running")

	h2s := &http2.Server{
		IdleTimeout: 10 * time.Second,
	}

	e := echo.New()
	// TODO Custom JSON Bind & Serialize

	if !cfg.Production {
		slog.Debug("Using cors")
		e.Use(middleware.CORS())

		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	st := db.New()
	repos := repository.New(&st)

	srv := services.New(&repos)
	api := e.Group("/api")

	controllers.New(controllers.Dependencies{
		ctx,
		api,
		&srv,
		cfg,
	})

	addr := cfg.Server.DevAddress

	if cfg.Production {
		addr = cfg.Server.Address
	}

	server := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(e, h2s),
	}

	return server

}
