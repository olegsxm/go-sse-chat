package a_sse

import (
	"fmt"
	"io"
)

type EventEmitter1 struct {
	ID    string
	Event string
	Data  string
}

func (s *EventEmitter1) Write(w io.Writer) {
	if s.Event != "" {
		_, _ = fmt.Fprintf(w, "event: %s\n", s.Event)
	}

	if s.ID != "" {
		_, _ = fmt.Fprintf(w, "id: %s\n", s.ID)
	}

	_, _ = fmt.Fprintf(w, "data: %s\n\n", s.Data)

}
