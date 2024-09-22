package handlers

import (
	"strings"

	"github.com/olegsxm/go-sse-chat.git/internal/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

func chatHandlers(router fiber.Router) {

	router.Get("/", func(c *fiber.Ctx) error {
		//lastChatId, _ := strconv.Atoi(c.Query("last-chat-id"))

		//services.Chat.GetChats(lastChatId)

		token := strings.Replace(c.Get("Authorization"), "Bearer ", "", 1)

		id := jwt.DecodeToken(token)["id"].(int)

		chats := services.Chat.GetChats(id, 0)

		return c.JSON(chats)
	})

}
