package handlers

import (
	"chat/internal/models"
	"chat/internal/pkg/jwt"
	"context"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"strings"
)

type conversationService interface {
	CreateConversationWithMessage(ctx context.Context, message, senderID, participantID string) (models.ConversationDTO, error)
}

type conversationHandler struct {
	service conversationService
	secret  string
}

func (h *conversationHandler) Mount(g *echo.Group) {
	g.POST("/create/private", h.createPrivateConversation, protectMiddleware)
	g.POST("/create-message", h.createMessage, protectMiddleware)
}

func (h *conversationHandler) createPrivateConversation(c echo.Context) error {
	token := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
	claims := jwt.ChatClaims{}

	if err := jwt.Parse(token, &claims, h.secret); err != nil {
		slog.Error("Error parsing token", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	req := models.CreatePrivateConversationRequest{}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	conversation, err := h.service.CreateConversationWithMessage(c.Request().Context(), req.Message, claims.Id, req.Participant)
	if err != nil {
		// TODO Custom errors
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, conversation)
}

func (h *conversationHandler) createMessage(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "Not implemented")
}

func newConversationHandler(service conversationService, secret string) *conversationHandler {
	return &conversationHandler{
		service: service,
		secret:  secret,
	}
}
