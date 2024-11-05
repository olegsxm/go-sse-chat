package sse

import (
	"github.com/goccy/go-json"
	"net/http"
	"sync"
)

type Broker struct {
	connect            chan string
	connections        map[string]chan any
	closingConnections chan string
	MessageAdapter     func(msg any, clientID string) EventEmitter
	ConnectionHandler  func(clientID string)
	mutex              sync.Mutex
}

func (b *Broker) Stream(clientID string, w http.ResponseWriter, r http.Request) error {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	// TODO warning insecure header value
	w.Header().Set("Access-Control-Allow-Origin", "*")

	messageChan := make(chan any)

	b.mutex.Lock()
	b.connections[clientID] = messageChan
	b.mutex.Unlock()

	b.connect <- clientID

	defer func() {
		b.closingConnections <- clientID
	}()

	go func() {
		<-r.Context().Done()
		b.closingConnections <- clientID
	}()

	for {
		msg := <-messageChan

		emitter := b.MessageAdapter(msg, clientID)

		emitter.Send(w)
		w.(http.Flusher).Flush()
	}
}

func (b *Broker) Listen() {
	for {
		select {
		case clientID := <-b.connect:
			b.ConnectionHandler(clientID)

		case clientID := <-b.closingConnections:
			delete(b.connections, clientID)

		}
	}
}

func (b *Broker) SendMessage(clientID string, message any) {
	c, ok := b.connections[clientID]
	if !ok {
		return
	}
	c <- message
}

func NewBroker() *Broker {
	broker := Broker{
		connect:            make(chan string),
		connections:        make(map[string]chan any),
		closingConnections: make(chan string),

		MessageAdapter: func(msg any, clientID string) EventEmitter {
			marshal, err := json.Marshal(msg)

			if err != nil {
				return EventEmitter{
					Event: "error",
					Data:  "{}",
				}
			}

			return EventEmitter{
				Event: "message",
				Data:  string(marshal),
			}
		},
		ConnectionHandler: func(clientID string) {

		},
	}

	go broker.Listen()

	return &broker
}
