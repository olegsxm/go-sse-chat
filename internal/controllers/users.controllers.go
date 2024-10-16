package controllers

import (
	"log/slog"
	"net/http"

	"github.com/olegsxm/go-sse-chat.git/pkg/cjwt"

	"github.com/labstack/echo/v4"
)

func usersControllers(g *echo.Group) {
	slog.Debug("Init users controllers")

	g.GET("/users/find", findUsers)
}

func findUsers(c echo.Context) error {
	query := c.QueryParam("query")

	if query == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "no query param")
	}

	ctxUserId := c.Get("userClaims").(cjwt.UserClaims).ID

	users, err := dependencies.Services.Users().FindUsers(query, ctxUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}
