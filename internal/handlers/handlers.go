package handlers

import "github.com/gofiber/fiber/v2"

type HandlerConstructorType int

func New(app *fiber.App) HandlerConstructorType {
	api := app.Group("/api")

	authHandlers(api)

	return 0
}
