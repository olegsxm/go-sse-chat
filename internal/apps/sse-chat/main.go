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
		),
		fx.Provide(handlers.New),
		fx.Invoke(func(h handlers.Handlers, f *fiber.App) {
			_ = h
		}),
	).Run()
}
