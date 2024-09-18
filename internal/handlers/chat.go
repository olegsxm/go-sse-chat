package handlers

import "github.com/gofiber/fiber/v2"

const ()

func chatHandlers(router fiber.Router) {

	router.Get("/", func(c *fiber.Ctx) error {
		services.Chat.GetChats()
		return nil
	})

}
