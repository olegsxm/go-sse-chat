package sse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Broker[T any] struct {
	connect     chan string
	connections map[string]chan T

	closingConnections chan string

	ConnectionHandler func(clientID string)
	DisconnectHandler func(clientID string)

	MessageAdapter func(msg T, clientID string) EventEmitter

	mutex sync.Mutex
}

func NewBroker[T any]() *Broker[T] {
	broker := Broker[T]{
		connect:            make(chan string),
		connections:        make(map[string]chan T),
		closingConnections: make(chan string),

		ConnectionHandler: DefaultConnectHandler,
		DisconnectHandler: DefaultDisconnectHandler,
		MessageAdapter: func(msg T, clientID string) EventEmitter {
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
	}

	go broker.listen()

	return &broker
}

func (broker *Broker[T]) Stream(clientID string, w http.ResponseWriter, r http.Request) error {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	messageChan := make(chan T)

	broker.mutex.Lock()
	broker.connections[clientID] = messageChan
	broker.mutex.Unlock()

	broker.connect <- clientID

	defer func() {
		broker.closingConnections <- clientID
	}()

	go func() {
		<-r.Context().Done()
		broker.closingConnections <- clientID
	}()

	for {
		msg := <-messageChan

		sse := broker.MessageAdapter(msg, clientID)

		sse.Write(w)
		w.(http.Flusher).Flush()
	}
}

func (broker *Broker[T]) listen() {
	for {
		select {
		case clientID := <-broker.connect:
			broker.ConnectionHandler(clientID)

		case clientID := <-broker.closingConnections:
			delete(broker.connections, clientID)
			broker.DisconnectHandler(clientID)
		}
	}
}

func (broker *Broker[T]) SendMessage(clientID string, message T) {
	c, ok := broker.connections[clientID]
	if !ok {
		return
	}

	c <- message
}

func DefaultConnectHandler(clientID string) {
	fmt.Println("Connect user with id = ", clientID)
}

func DefaultDisconnectHandler(clientID string) {
	fmt.Println("Client ", clientID, " disconnected")
}
