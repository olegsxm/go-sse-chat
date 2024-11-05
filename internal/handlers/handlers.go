package handlers

import (
	"chat/internal/pkg/config"
	"chat/internal/pkg/jwt"
	"chat/internal/services"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var protectMiddleware echo.MiddlewareFunc

type Dependencies struct {
	Api      *echo.Group
	Services services.Services
	Config   config.Config
}

type Handlers struct {
	deps Dependencies
}

func (h *Handlers) Mount() {
	newAuthHandlers(h.deps.Services.Auth()).
		Mount(h.deps.Api.Group("/auth"))

	newConversationHandler(h.deps.Services.Conversation(), h.deps.Config.Security.JWTKey).
		Mount(h.deps.Api.Group("/conversation"))
}

func New(d Dependencies) *Handlers {
	protectMiddleware = echojwt.WithConfig(jwt.NewEchoJwtConfig(d.Config.Security.JWTKey))

	return &Handlers{
		deps: d,
	}
}
