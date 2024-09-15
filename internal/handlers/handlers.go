package handlers

import (
	"fmt"

	"github.com/olegsxm/go-sse-chat.git/internal/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

type ConstructorType struct {
}

var validate *validator.XValidator

func New(app *fiber.App, v *validator.XValidator) *ConstructorType {
	fmt.Println("Initializing handlers")
	validate = v

	api := app.Group("/api")
	auth := api.Group("/auth")

	authHandlers(auth)

	return &ConstructorType{}
}
