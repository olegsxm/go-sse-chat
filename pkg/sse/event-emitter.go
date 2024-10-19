package sse

import (
	"fmt"
	"io"
	"log/slog"
	"time"
)

type EventEmitter struct {
	ID    int64
	Event string
	Data  string
}

func (s *EventEmitter) Write(w io.Writer) {
	if s.Event != "" {
		_, err := fmt.Fprintf(w, "event: %s\n", s.Event)
		if err != nil {
			slog.Error(err.Error())
		}
	}

	if s.ID == 0 {
		s.ID = time.Now().Unix()
		_, err := fmt.Fprintf(w, "id: %s\n", s.ID)
		if err != nil {
			slog.Error(err.Error())
		}
	}

	_, err := fmt.Fprintf(w, "data: %s\n\n", s.Data)
	if err != nil {
		slog.Error(err.Error())
	}
}
