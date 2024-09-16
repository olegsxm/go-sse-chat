package sse_chat

import (
	"database/sql"
	"os"

	"github.com/olegsxm/go-sse-chat.git/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/olegsxm/go-sse-chat.git/internal/handlers"
	"github.com/olegsxm/go-sse-chat.git/internal/pkg/validator"
	"github.com/olegsxm/go-sse-chat.git/internal/repository"
	"github.com/olegsxm/go-sse-chat.git/internal/services"

	"github.com/olegsxm/go-sse-chat.git/internal/pkg/server"

	"go.uber.org/fx"
)

func Run() {
	fx.New(
		fx.Provide(
			server.NewHttpServer,
			validator.New,
			db.New,
			repository.New,
			services.New,
			handlers.New,
		),
		fx.Invoke(func(app *fiber.App, h *handlers.ConstructorType, d *sql.DB) {
			server.Run(os.Getenv("SSE_SERVER_PORT"), app)
			defer d.Close()
		}),
	).Run()
}
