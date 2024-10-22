package router

import (
	"github.com/olegsxm/go-sse-chat/pkg/http_errors"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.

		switch err.(type) {
		case http_errors.ChatHTTPError:
			w.WriteHeader(err.(http_errors.ChatHTTPError).Code())
			_, _ = w.Write([]byte(err.Error()))
		default:
			w.WriteHeader(500)
		}
	}
}
