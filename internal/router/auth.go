package router

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/olegsxm/go-sse-chat/internal/models"
	"github.com/olegsxm/go-sse-chat/pkg/http_errors"
	"github.com/olegsxm/go-sse-chat/pkg/utils"
	"log/slog"
	"net/http"
	"time"
)

type authService interface {
	CreateUser(ctx context.Context, login, password string) (models.User, error)
	FindUser(context context.Context, login, password string) (models.User, error)
}

type auth struct {
	service authService
}

func (a *auth) Mount(c *chi.Mux) {
	auth := chi.NewRouter()

	auth.Route("/", func(r chi.Router) {
		r.With(cancelableContext).Method("POST", "/sign-in", Handler(a.signIn))
		r.With().Method("POST", "/sign-up", Handler(a.signUp))
	})

	c.Mount("/auth", auth)
}

func (a *auth) signIn(w http.ResponseWriter, r *http.Request) error {
	var authRequest models.AuthRequest

	if err := utils.HttpBind(r.Body, &authRequest); err != nil {
		return err
	}

	user, err := a.service.FindUser(r.Context(), authRequest.Login, authRequest.Password)
	if err != nil {
		return err
	}

	tokenString, err := genToken(user)
	if err != nil {
		return err
	}

	authResponse := models.AuthResponse{
		Token: tokenString,
		User:  user,
	}

	json, err := authResponse.MarshalJSON()
	if err != nil {
		slog.Error("Error marshalling auth response")
		return err
	}

	_, _ = w.Write(json)

	return nil
}

func (a *auth) signUp(w http.ResponseWriter, r *http.Request) error {
	var authRequest models.AuthRequest

	if err := utils.HttpBind(r.Body, &authRequest); err != nil {
		return http_errors.BadRequest{}
	}

	user, err := a.service.CreateUser(r.Context(), authRequest.Login, authRequest.Password)
	if err != nil {
		return err
	}

	tokenString, err := genToken(user)
	if err != nil {
		return err
	}

	authResponse := models.AuthResponse{
		Token: tokenString,
		User:  user,
	}

	json, err := authResponse.MarshalJSON()
	if err != nil {
		slog.Error("Error marshalling auth response")
		return http_errors.InternalServerError{}
	}

	_, _ = w.Write(json)

	return nil
}

func genToken(u models.User) (string, error) {
	claims := map[string]interface{}{
		"id":    u.ID,
		"login": u.Login,
		"exp":   time.Now().Add(time.Hour * 10).Unix(),
	}
	_, tokenString, err := tokenAuth.Encode(claims)

	return tokenString, err
}
