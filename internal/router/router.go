package router

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	_ "github.com/go-chi/jwtauth/v5"
	"github.com/olegsxm/go-sse-chat/internal/services"
	"net/http"
)

var tokenAuth *jwtauth.JWTAuth

type controller interface {
	Mount(*chi.Mux)
}

func New(s *services.Services) *chi.Mux {
	//TODO  get secret from config
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(cors.AllowAll().Handler)
	r.Use(middleware.GetHead)

	api := chi.NewRouter()

	api.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			next.ServeHTTP(w, r)
		})
	})

	h := []controller{
		&auth{s.Auth()},
		&conversations{s.Conversation()},
	}

	for _, h := range h {
		h.Mount(api)
	}

	r.Mount("/api", api)

	return r
}

func cancelableContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithCancel(r.Context())

		// TODO THINK ABOUT THAT
		go func() {
			select {
			case <-ctx.Done():
				cancel()
			}
		}()

		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
