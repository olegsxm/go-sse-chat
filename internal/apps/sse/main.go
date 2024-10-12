package sse

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/olegsxm/go-sse-chat.git/internal/controllers"
	"github.com/olegsxm/go-sse-chat.git/internal/services"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log/slog"
	"net/http"
	"time"

	"github.com/olegsxm/go-sse-chat.git/internal/config"
	"golang.org/x/net/http2/h2c"

	"golang.org/x/net/http2"

	_ "github.com/swaggo/http-swagger/v2"
)

func New(ctx context.Context) *http.Server {
	slog.Info("Sse Chat Running")

	cfg, err := config.New()

	if err != nil || cfg == nil {
		slog.Error("Error loading config")
		return nil
	}

	h2s := &http2.Server{
		IdleTimeout: 10 * time.Second,
	}

	e := echo.New()

	if !cfg.Production {
		slog.Debug("Using cors")
		e.Use(middleware.CORS())
	}

	srv := services.New()
	api := e.Group("/api")

	controllers.New(ctx, api, &srv)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	server := &http.Server{
		Addr:    cfg.Server.Address,
		Handler: h2c.NewHandler(e, h2s),
	}

	return server

}
