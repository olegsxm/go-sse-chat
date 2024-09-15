package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type credentials struct {
	Login string `json:"login"`
}

func authHandlers(auth fiber.Router) {

	fmt.Println("Auth navigation initialize")

	auth.Post("/sign-in", func(c *fiber.Ctx) error {
		var creds credentials

		if err := c.BodyParser(&creds); err != nil {
		}

		return c.JSON(fiber.Map{
			"success": true,
		})
	})
}
