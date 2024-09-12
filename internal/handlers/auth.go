package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func authHandlers(auth fiber.Router) {

	fmt.Println("Auth navigation initialize")

	auth.Get("/sign-in", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
		})
	})
}
