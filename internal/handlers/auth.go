package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type credentials struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func authHandlers(router fiber.Router) {

	// SignIn Login in Chat
	//	@Summary	Login in Chat
	//	@Tags		auth
	//	@Accept		json
	//	@Produce	json
	//	@Success	200
	//	@Router		/api/auth/sign-in [post]
	router.Post("/sign-in", func(c *fiber.Ctx) error {
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

	router.Post("/sign-up", func(c *fiber.Ctx) error {
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

	router.Get("/check-login", func(c *fiber.Ctx) error {

		login := c.Query("login")

		if len(login) == 0 {
			return fiber.NewError(fiber.StatusBadRequest, "login required")
		}

		user, _ := services.Auth.FindUserByLogin(login)

		if user.ID == 0 {
			return c.SendStatus(200)
		}

		return c.SendStatus(fiber.StatusBadRequest)
	})
}
