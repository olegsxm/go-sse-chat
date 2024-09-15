package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewHttpServer(port string) func() *fiber.App {
	return func() *fiber.App {
		fmt.Println("Server is starting...")

		app := fiber.New(fiber.Config{
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

		go func() {
			if err := app.Listen(port); err != nil {
				log.Panic(err)
			}
		}()

		app.Get("/ping", func(c *fiber.Ctx) error {
			return c.SendString("pong")
		})

		return app
	}
}

func Run(app *fiber.App) *int {
	c := make(chan os.Signal, 1)

	// Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")

	return nil
}
