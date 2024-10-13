package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/olegsxm/go-sse-chat.git/internal/config"
	services "github.com/olegsxm/go-sse-chat.git/internal/services"
)

var dependencies Dependencies

type Dependencies struct {
	Ctx      context.Context
	Router   *echo.Group
	Services *services.Services
	Config   *config.AppConfig
}

func New(deps Dependencies) {
	dependencies = deps
	v1 := dependencies.Router.Group("/v1")
	authHandlers(v1)
}
