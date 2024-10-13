package controllers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	jwt "github.com/olegsxm/go-sse-chat.git/pkg/jwt"
	validate "github.com/olegsxm/go-sse-chat.git/pkg/validator"

	"github.com/olegsxm/go-sse-chat.git/internal/models"

	_ "github.com/go-chi/render"
)

func authControllers(g *echo.Group) {
	slog.Debug("Init auth controllers")
	g.POST("/auth/sign-in", signIn)
	g.POST("/auth/sign-up", signUp)
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
	r := models.AuthRequest{}

	if err := ctx.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := validate.ValidateStruct(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := dependencies.Services.Auth().SignIn(r.Login, r.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	token, err := jwt.CreateToken(
		&jwt.UserClaims{ID: user.ID, Login: user.Login},
		dependencies.Config.JWTSecret,
	)
	if err != nil {
		slog.Error("CreateToken err:", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	resp := models.AuthResponse{
		Token: token,
		User:  user.ToDTO(),
	}

	return ctx.JSON(200, resp)
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
func signUp(ctx echo.Context) error {
	r := models.AuthRequest{}

	if err := ctx.Bind(&r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := validate.ValidateStruct(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	u, err := dependencies.Services.Auth().SignIn(r.Login, r.Password)
	if err == nil && u.ID != 0 {
		slog.Error(fmt.Sprintf("Login: %s; Password: %s", r.Login, r.Password))
		return echo.NewHTTPError(http.StatusInternalServerError, "Error signing up")
	}

	if u.ID != 0 {
		return echo.NewHTTPError(http.StatusConflict, "User already exists")
	}

	u, err = dependencies.Services.Auth().SignUp(r.Login, r.Password)
	if err != nil {
		slog.Error(err.Error(), fmt.Sprintf("Login: %s; Password: %s", r.Login, r.Password))
		return echo.NewHTTPError(http.StatusInternalServerError, "Error signing up")
	}

	token, err := jwt.CreateToken(
		&jwt.UserClaims{ID: u.ID, Login: u.Password},
		dependencies.Config.JWTSecret,
	)
	if err != nil {
		slog.Error("CreateToken err:", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	resp := models.AuthResponse{
		Token: token,
		User:  u.ToDTO(),
	}

	return ctx.JSON(200, resp)
}
