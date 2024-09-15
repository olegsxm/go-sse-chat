package sse_chat

import (
	"github.com/gofiber/fiber/v2"

	"github.com/olegsxm/go-sse-chat.git/internal/handlers"
	"github.com/olegsxm/go-sse-chat.git/internal/pkg/validator"

	"github.com/olegsxm/go-sse-chat.git/internal/pkg/server"

	"go.uber.org/fx"
)

func Run() {
	fx.New(
		fx.Provide(
			server.NewHttpServer(":3000"),
			validator.New,
			handlers.New,
		),
		fx.Invoke(func(app *fiber.App, h *handlers.ConstructorType) {
			server.Run(app)
			_ = h
		}),
	).Run()
}
