package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	app *fiber.App
}

func New(app *fiber.App) *Handlers {
	fmt.Println("Initializing handlers")

	api := app.Group("/api")
	auth := api.Group("/auth")

	authHandlers(auth)

	return &Handlers{
		app: app,
	}
}
