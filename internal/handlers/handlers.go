package handlers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	srv "github.com/olegsxm/go-sse-chat.git/internal/services"

	"github.com/gofiber/fiber/v2"
)

type ConstructorType struct {
}

var validate = validator.New()
var services *srv.Services

func New(app *fiber.App, s *srv.Services) *ConstructorType {
	fmt.Println("Initializing handlers")

	services = s

	api := app.Group("/api")
	auth := api.Group("/auth")

	authHandlers(auth)

	return &ConstructorType{}
}
