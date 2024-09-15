package sse_chat

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/olegsxm/go-sse-chat.git/internal/handlers"

	"github.com/olegsxm/go-sse-chat.git/internal/pkg/server"

	"go.uber.org/fx"
)

const idleTimeout = 5 * time.Second

func Run() {
	fx.New(
		fx.Provide(
			server.NewHttpServer(":3000"),
			handlers.New,
		),
		fx.Invoke(func(app *fiber.App, h *handlers.Handlers) {
			server.Run(app)
			_ = h
		}),
	).Run()
}
