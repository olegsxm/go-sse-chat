package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/olegsxm/go-sse-chat.git/internal/models"
)

func messageHandler(router fiber.Router) {
	router.Post("/send", func(ctx *fiber.Ctx) error {

		message := models.Message{}

		if err := ctx.BodyParser(&message); err != nil {
			log.Error("parse message error ", err.Error())
			return err
		}

		if err := validate.Struct(message); err != nil {
			log.Error("validate message error ", err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "invalid message")
		}

		message, err := services.Message.CreateMessage(message)
		if err != nil {
			fmt.Println(err)
			return fiber.NewError(fiber.StatusBadRequest, "invalid message")
		}

		return ctx.JSON(message)
	})
}
