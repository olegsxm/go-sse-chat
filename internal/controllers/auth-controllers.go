package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/olegsxm/go-sse-chat.git/pkg/fjson"
	validate "github.com/olegsxm/go-sse-chat.git/pkg/validator"
	"log/slog"
	"net/http"

	"github.com/olegsxm/go-sse-chat.git/internal/models"

	_ "github.com/go-chi/render"
)

func authHandlers(g *echo.Group) {
	slog.Debug("Init auth handlers")
	g.POST("/auth/sign-in", signIn)
}

//	SignIn godoc
//
// @Summary		Login at chat
// @Description	Login in Chat
// @Tags			Auth
// @Accept			json
// @Produce		json
//
// @Param			data	body	models.AuthRequest	true	"Login Request"
//
// @Router			/auth/sign-in [post]
func signIn(ctx echo.Context) error {
	c := models.AuthRequest{}

	if err := fjson.ParseBody(ctx.Request().Body, &c); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	if err := validate.ValidateStruct(c); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(200, echo.Map{
		"ok": 1,
	})
}

//	SignUp godoc
//
// @Summary		Register at chat
// @Description	SignUp in Chat
// @Tags			Auth
// @Accept			json
// @Produce		json
//
// @Param			data	body	models.AuthRequest	true	"Login Request"
//
// @Router			/auth/sign-up [post]
func signUp(w http.ResponseWriter, r *http.Request) error {
	return nil
}
