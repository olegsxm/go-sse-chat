package controllers

import (
	"context"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/olegsxm/go-sse-chat.git/pkg/jwt"

	"github.com/labstack/echo/v4"
	"github.com/olegsxm/go-sse-chat.git/internal/config"
	"github.com/olegsxm/go-sse-chat.git/internal/services"
)

var dependencies Dependencies

var protectMiddleware echo.MiddlewareFunc

type Dependencies struct {
	Ctx      context.Context
	Router   *echo.Group
	Services *services.Services
	Config   *config.AppConfig
}

func New(deps Dependencies) {
	dependencies = deps
	protectMiddleware = echojwt.WithConfig(jwt.NewEchoJwtConfig(dependencies.Config.JWTSecret))

	v1 := dependencies.Router.Group("/v1")

	authControllers(v1)
	chatControllers(v1)
	usersControllers(v1)
}
