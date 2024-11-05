package handlers

import (
	"chat/internal/pkg/config"
	"chat/internal/pkg/jwt"
	"chat/internal/pkg/sse"
	"chat/internal/services"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

var protectMiddleware echo.MiddlewareFunc

type Broker interface {
	Stream(clientID string, w http.ResponseWriter, r http.Request) error
	SendMessage(clientID string, message any)
}

type Dependencies struct {
	Api           *echo.Group
	Services      services.Services
	Config        config.Config
	MessageBroker *sse.Broker
}

type Handlers struct {
	deps Dependencies
}

func (h *Handlers) Mount() {
	newAuthHandlers(h.deps.Services.Auth()).
		Mount(h.deps.Api.Group("/auth"))

	newSseHandler(h.deps.MessageBroker, h.deps.Config.Security.JWTKey).
		Mount(h.deps.Api.Group("/broker"))

	newConversationHandler(h.deps.Services.Conversation(), h.deps.MessageBroker, h.deps.Config.Security.JWTKey).
		Mount(h.deps.Api.Group("/conversation"))
}

func New(d Dependencies) *Handlers {
	protectMiddleware = echojwt.WithConfig(jwt.NewEchoJwtConfig(d.Config.Security.JWTKey))

	return &Handlers{
		deps: d,
	}
}
