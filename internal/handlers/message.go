package handlers

import (
	"fmt"
	"strconv"

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

	router.Get("/", func(ctx *fiber.Ctx) error {
		// TODO check user access

		chatId, err := strconv.Atoi(ctx.Query("chat-id"))

		if err != nil {
			return fiber.ErrBadRequest
		}

		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			limit = 0
		}

		lastMessageId, err := strconv.Atoi(ctx.Query("last-id"))
		if err != nil {
			lastMessageId = 0
		}

		messages, e := services.Message.GetMessages(chatId, limit, lastMessageId)
		if e != nil {
			log.Error(e)
			return fiber.ErrInternalServerError
		}

		return ctx.JSON(messages)
	})
}
