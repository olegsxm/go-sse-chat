package sse

import (
	"fmt"
	"io"
	"log/slog"
	"time"
)

type EventEmitter struct {
	Id    int64
	Event string
	Data  string
}

func (e *EventEmitter) Send(w io.Writer) {
	if e.Event != "" {
		_, err := fmt.Fprintf(w, "event: %s\n", e.Event)
		if err != nil {
			slog.Error("event emitter sending event error: ", err.Error())
		}
	}

	if e.Id == 0 {
		e.Id = time.Now().Unix()
	}
	_, err := fmt.Fprintf(w, "id: %d\n", e.Id)
	if err != nil {
		slog.Error("event emitter sending Id error: ", err.Error())
	}

	_, err = fmt.Fprintf(w, "data: %s\n\n", e.Data)
	if err != nil {
		slog.Error("event emitter sending Data error: ", err.Error())
	}
}
