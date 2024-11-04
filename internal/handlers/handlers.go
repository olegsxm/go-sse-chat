package handlers

import (
	"chat/internal/services"
	"github.com/labstack/echo/v4"
)

type Dependencies struct {
	Api      *echo.Group
	Services services.Services
}

type Handlers struct {
	deps Dependencies
}

func (h *Handlers) Mount() {
	newAuthHandlers(h.deps.Services.Auth()).Mount(h.deps.Api.Group("/auth"))
	newConversationHandler(h.deps.Services.Conversation())
}

func New(d Dependencies) *Handlers {
	return &Handlers{
		deps: d,
	}
}
