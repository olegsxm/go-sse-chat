package router

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/olegsxm/go-sse-chat/internal/models"
	"log/slog"
	"net/http"
)

type crudService interface {
	FindAll(ctx context.Context) (*models.Conversations, error)
	FindByID(ctx context.Context, id string) (models.Conversation, error)
	Create(ctx context.Context, id string) (models.Conversation, error)
	Update(ctx context.Context) (models.Conversation, error)
	Delete(ctx context.Context, id string) error
}

type conversations struct {
	s crudService
}

func (h *conversations) Mount(c *chi.Mux) {
	cr := chi.NewRouter()

	cr.Route("/", func(r chi.Router) {
		// TODO uncomment
		r.Use(jwtauth.Verifier(tokenAuth), jwtauth.Authenticator(tokenAuth))

		r.Method("Get", "/", Handler(h.getConversations))
		r.Method("Get", "/{id}", Handler(h.getConversation))
		r.Method("Post", "/", Handler(h.createConversation))
		r.Method("Patch", "/{id}", Handler(h.updateConversation))
		r.Method("Delete", "/{id}", Handler(h.deleteConversation))
	})

	c.Mount("/conversations", cr)
}

func (h *conversations) getConversations(w http.ResponseWriter, r *http.Request) error {
	all, err := h.s.FindAll(r.Context())
	if err != nil {
		return err
	}

	json, err := all.MarshalJSON()
	if err != nil {
		return err
	}
	_, err = w.Write(json)
	return err
}

func (h *conversations) getConversation(w http.ResponseWriter, r *http.Request) error {
	var err error

	c, err := h.s.FindByID(r.Context(), chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	json, err := c.MarshalJSON()

	_, err = w.Write(json)
	return err
}

func (h *conversations) createConversation(w http.ResponseWriter, r *http.Request) error {
	var err error

	_, m, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return err
	}

	var senderID string

	if _, ok := m["id"]; ok {
		senderID = m["id"].(string)
	}

	chi.URLParam(r, "to")

	c, err := h.s.Create(r.Context(), senderID)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	json, err := c.MarshalJSON()
	if err != nil {
		return err
	}

	_, err = w.Write(json)
	return err
}

func (h *conversations) updateConversation(w http.ResponseWriter, r *http.Request) error {
	var err error
	_, err = w.Write([]byte("conversation"))
	return err
}

func (h *conversations) deleteConversation(w http.ResponseWriter, r *http.Request) error {
	var err error
	_, err = w.Write([]byte("conversation"))
	return err
}
