package controllers

import (
	"log/slog"
	"net/http"

	"github.com/olegsxm/go-sse-chat.git/internal/models"

	"github.com/labstack/echo/v4"
)

func chatControllers(g *echo.Group) {
	slog.Debug("Init chat controllers")

	g.GET("/chat/conversations", getConversations, protectMiddleware)
	g.POST("/chat/conversations", createConversation, protectMiddleware)
}

func getConversations(c echo.Context) error {
	cs, e := dependencies.Services.Chat().GetConversation()

	if e != nil {
		slog.Error(e.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, cs)
}

func createConversation(c echo.Context) error {
	data := models.NewConversationRequest{
		From: getUserClaims(c).ID,
	}

	if err := c.Bind(&data); err != nil {
		slog.Error("data binding err ", err)
		return c.JSON(http.StatusBadRequest, nil)
	}

	conv, e := dependencies.Services.Chat().CreateConversation(data.From, data.To)
	if e != nil {
		slog.Error("create conversation error", e.Error())
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, conv)
}
