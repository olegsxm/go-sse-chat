package controllers

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

func chatControllers(g *echo.Group) {
	slog.Debug("Init chat controllers")

	g.GET("/chat/conversations", func(c echo.Context) error {
		return nil
	}, protectMiddleware)
}
