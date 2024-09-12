package server

import (
	"context"
	"time"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/fx"
)

const idleTimeout = 5 * time.Second

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewHttpServer(port string) func(lc fx.Lifecycle) *fiber.App {

	return func(lc fx.Lifecycle) *fiber.App {
		app := fiber.New(fiber.Config{
			IdleTimeout: idleTimeout,
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
					Success: false,
					Message: err.Error(),
				})
			},
		})

		app.Use(cors.New())

		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				return app.Listen(port)
			},
			OnStop: func(ctx context.Context) error {
				return app.Shutdown()
			},
		})

		return app
	}
}
