package handlers

import "github.com/labstack/echo/v4"

type conversationService interface {
}

type conversationHandler struct {
	service conversationService
}

func (h *conversationHandler) Mount(g *echo.Group) {

}

func newConversationHandler(service conversationService) *conversationHandler {
	return &conversationHandler{
		service: service,
	}
}
