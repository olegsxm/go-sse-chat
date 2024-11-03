package handlers

import (
	"chat/internal/models"
	"chat/internal/pkg/validator"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type authService interface {
	SignUp(ctx context.Context, login, password string) (models.AuthResponse, string, error)
	SignIn(ctx context.Context, login, password string) (models.AuthResponse, string, error)
}

type authHandlers struct {
	service authService
}

func (h *authHandlers) Mount(g *echo.Group) {
	g.POST("/sign-in", h.signIn)
	g.POST("/sign-up", h.signUp)
}

func (h *authHandlers) signIn(c echo.Context) error {
	request := new(models.AuthRequest)

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ParseError(err))
	}

	resp, refresh, err := h.service.SignIn(c.Request().Context(), request.Login, request.Password)
	if err != nil {
		// TODO create custom errors
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	c.SetCookie(&http.Cookie{
		Name:     "refresh",
		Value:    refresh,
		Expires:  time.Now().Add(120 * time.Hour),
		Secure:   true,
		HttpOnly: true,
	})

	return c.JSON(200, resp)
}

func (h *authHandlers) signUp(c echo.Context) error {
	request := new(models.AuthRequest)

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ParseError(err))
	}

	resp, refresh, err := h.service.SignUp(c.Request().Context(), request.Login, request.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	c.SetCookie(&http.Cookie{
		Name:     "refresh",
		Value:    refresh,
		Expires:  time.Now().Add(120 * time.Hour),
		Secure:   true,
		HttpOnly: true,
	})

	return c.JSON(200, resp)
}

func newAuthHandlers(service authService) *authHandlers {
	return &authHandlers{
		service: service,
	}
}
