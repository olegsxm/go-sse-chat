package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	services "github.com/olegsxm/go-sse-chat.git/internal/services"
)

var srv *services.Services

func New(ctx context.Context, router *echo.Group, s *services.Services) {
	srv = s
	v1 := router.Group("/v1")
	authHandlers(v1)
}
