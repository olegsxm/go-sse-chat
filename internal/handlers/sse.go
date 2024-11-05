package handlers

import (
	"chat/internal/pkg/jwt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"strings"
)

type SseHandler struct {
	broker Broker
	secret string
}

func (h *SseHandler) Mount(g *echo.Group) {
	g.GET("/sse", h.connect)
}

func (h *SseHandler) connect(c echo.Context) error {
	token := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
	if token == "" {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	claims := jwt.ChatClaims{}

	if err := jwt.Parse(token, &claims, h.secret); err != nil {
		slog.Error("Error parsing token", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	return h.broker.Stream(claims.Id, c.Response().Writer, *c.Request())
}

func newSseHandler(broker Broker, secret string) *SseHandler {
	return &SseHandler{
		broker: broker,
		secret: secret,
	}
}
