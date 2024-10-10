package controllers

import (
	"net/http"

	"github.com/olegsxm/go-sse-chat.git/internal/models"

	"github.com/olegsxm/go-sse-chat.git/pkg/fjson"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/render"
	"github.com/olegsxm/go-sse-chat.git/pkg/handler"
	v "github.com/olegsxm/go-sse-chat.git/pkg/validator"
)

func authHandlers() *chi.Mux {
	c := chi.NewRouter()

	c.Post("/sign-in", handler.HandleRoute(signIn))
	c.Post("/sign-up", handler.HandleRoute(signUp))

	c.Mount("/auth", c)
	return c
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
func signIn(w http.ResponseWriter, r *http.Request) error {
	c := models.AuthRequest{}

	err := fjson.ParseBody(r.Body, &c)
	if err != nil {
		return err
	}

	if e := v.ValidateStruct(c); e != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid credentials"))
		return err
	}

	res, err := uc.Auth().SignIn(c.Login, c.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	j, _ := res.MarshalJSON()

	_, e := w.Write(j)
	return e
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
