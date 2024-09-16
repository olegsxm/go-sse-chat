package handlers

import (
	"github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2"
)

type credentials struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func authHandlers(auth fiber.Router) {

	log.Debug("Auth navigation initialize")

	auth.Post("/sign-in", func(c *fiber.Ctx) error {
		var creds credentials

		if err := c.BodyParser(&creds); err != nil {
			return fiber.ErrBadRequest
		}

		err := validate.Struct(&creds)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		token, err := services.Auth.SignIn(creds.Login, creds.Password)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"token": token,
		})
	})

	auth.Post("/sign-up", func(c *fiber.Ctx) error {
		var creds credentials

		if err := c.BodyParser(&creds); err != nil {
			return fiber.ErrBadRequest
		}

		err := validate.Struct(&creds)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		token, err := services.Auth.SignUp(creds.Login, creds.Password)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return c.JSON(fiber.Map{
			"token": token,
		})
	})
}
