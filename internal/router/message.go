package router

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/olegsxm/go-sse-chat/internal/models"
	"github.com/olegsxm/go-sse-chat/pkg/utils"
	"net/http"
)

type messageService interface {
	Create(ctx context.Context, id string, conversation string, m string) (models.MessageResponse, error)
	Get(ctx context.Context, conversationID string) (models.Messages, error)
}

type message struct {
	service messageService
}

func (m *message) Mount(c *chi.Mux) {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth), jwtauth.Authenticator(tokenAuth))

		r.Method("Get", "/conversation/{id}", Handler(m.get))
		r.Method("Post", "/", Handler(m.create))
	})

	c.Mount("/message", router)
}

func (mh *message) create(w http.ResponseWriter, r *http.Request) error {
	var mr models.MessageRequest

	if err := utils.HttpBind(r.Body, &mr); err != nil {
		return err
	}

	_, m, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return err
	}

	var senderID string

	_ = senderID

	if _, ok := m["id"]; ok {
		senderID = m["id"].(string)
	}

	msg, err := mh.service.Create(r.Context(), senderID, mr.Conversation, mr.Message)
	if err != nil {
		return err
	}

	json, err := msg.MarshalJSON()
	if err != nil {
		return err
	}

	_, _ = w.Write(json)

	return nil
}

func (m *message) get(w http.ResponseWriter, r *http.Request) error {
	conversationID := chi.URLParam(r, "id")

	if conversationID == "" {
		return errors.New("conversation id is required")
	}

	messages, err := m.service.Get(r.Context(), conversationID)
	if err != nil {
		return err
	}

	j, err := messages.MarshalJSON()
	if err != nil {
		return err
	}

	_, _ = w.Write(j)

	return nil
}
