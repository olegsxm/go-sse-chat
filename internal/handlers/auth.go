package handlers

import "github.com/gofiber/fiber/v2"

func authHandlers(api fiber.Router) {
	auth := api.Group("/auth")

	auth.Get("/sign-in", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
		})
	})
}
