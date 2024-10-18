package controllers

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/olegsxm/go-sse-chat.git/pkg/cjwt"

	"github.com/olegsxm/go-sse-chat.git/internal/models"

	"github.com/labstack/echo/v4"
)

func chatControllers(g *echo.Group) {
	slog.Debug("Init chat controllers")

	g.GET("/chat/conversations", getConversations, protectMiddleware)
	g.POST("/chat/conversations", createConversation, protectMiddleware)
	g.GET("/chat/conversation/:conversationId/messages", getMessages, protectMiddleware)
	g.POST("/chat/conversation/:conversationId/create-message", createMessage, protectMiddleware)
}

func getConversations(c echo.Context) error {
	ctxUserId := c.Get("userClaims").(cjwt.UserClaims).ID

	cs, e := dependencies.Services.Chat().GetConversation(ctxUserId)

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

func createMessage(c echo.Context) error {
	ctxUser := c.Get("userClaims").(cjwt.UserClaims)
	conversationId, err := strconv.ParseInt(c.Param("conversationId"), 10, 64)

	if err != nil {
		return echo.ErrBadRequest
	}

	message := models.Message{
		SenderId:       ctxUser.ID,
		ConversationId: conversationId,
		CreatedAt:      time.Now(),
	}

	if err := c.Bind(&message); err != nil {
		slog.Error(err.Error())
		return echo.ErrBadRequest
	}

	m, err := dependencies.Services.Chat().CreateMessage(message)
	if err != nil {
		slog.Error(err.Error())
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, m)
}

func getMessages(c echo.Context) error {
	ctxUser := c.Get("userClaims").(cjwt.UserClaims)
	conversationId, err := strconv.ParseInt(c.Param("conversationId"), 10, 64)

	if err != nil {
		return echo.ErrBadRequest
	}

	messages, err := dependencies.Services.Chat().GetMessages(conversationId, ctxUser.ID)
	if err != nil {
		slog.Error(err.Error())
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, messages)
}
