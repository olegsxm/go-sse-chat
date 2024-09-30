package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func userHandlers(router fiber.Router) {
	router.Get("/find", func(ctx *fiber.Ctx) error {

		login, ok := ctx.Queries()["login"]

		if !ok {
			log.Error("Error on query 'login' param", ctx.Queries()["login"])
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		users, err := services.Auth.SearchUser(login)

		if err != nil {
			return fiber.ErrBadRequest
		}

		return ctx.JSON(users)
	})
}
