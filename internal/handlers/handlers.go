package handlers

import (
	"github.com/gofiber/fiber/v2/log"

	"github.com/go-playground/validator/v10"
	srv "github.com/olegsxm/go-sse-chat.git/internal/services"

	"github.com/gofiber/fiber/v2"
)

type ConstructorType struct {
}

var validate = validator.New()
var services *srv.Services

func New(app *fiber.App, s *srv.Services) *ConstructorType {
	log.Debug("Initializing handlers")
	services = s
	api := app.Group("/api")

	auth := api.Group("/auth")
	messages := api.Group("/messages")
	chat := api.Group("/chat")
	user := api.Group("/user")

	authHandlers(auth)
	messageHandler(messages)
	chatHandlers(chat)
	userHandlers(user)

	return &ConstructorType{}
}
