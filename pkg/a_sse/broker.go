package a_sse

import (
	"net/http"
	"sync"
)

type Broker[T any] struct {
	newConnect     chan string
	closingClients chan string
	clients        map[string]chan T
	groups         map[string][]string

	ClientConnectedHandler    func(clientID string)
	ClientDisconnectedHandler func(clientID string)

	MessageAdapter func(msg T, clientId string) EventEmitter1

	mutex sync.Mutex
}

func (broker *Broker[T]) Stream(clientID string, w http.ResponseWriter, r http.Request) error {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	messageChan := make(chan T)

	broker.mutex.Lock()
	broker.clients[clientID] = messageChan
	broker.mutex.Unlock()

	broker.newConnect <- clientID

	defer func() {
		broker.closingClients <- clientID
	}()

	go func() {
		<-r.Context().Done()
		broker.closingClients <- clientID
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
		// CASE: New client has connected
		case clientID := <-broker.newConnect:
			// Add this client to the special group for all clients
			broker.AddToGroup(clientID, "*")
			broker.ClientConnectedHandler(clientID)

		// CASE: Client has detached and we want to stop sending them messages
		case clientID := <-broker.closingClients:
			delete(broker.clients, clientID)

			// Nasty nested loop to remove client from all groups
			for group := range broker.groups {
				for i, grpClientID := range broker.groups[group] {
					if grpClientID == clientID {
						broker.groups[group] = append(broker.groups[group][:i], broker.groups[group][i+1:]...)
					}
				}
			}

			broker.ClientDisconnectedHandler(clientID)
		}
	}
}

func (broker *Broker[T]) AddToGroup(clientID string, group string) {
	broker.groups[group] = append(broker.groups[group], clientID)
}

func (broker *Broker[T]) RemoveFromGroup(clientID string, group string) {
	for i, grpClientID := range broker.groups[group] {
		if grpClientID == clientID {
			broker.groups[group] = append(broker.groups[group][:i], broker.groups[group][i+1:]...)
		}
	}
}

func (broker *Broker[T]) SendToClient(clientID string, message T) {
	c, ok := broker.clients[clientID]

	if !ok {
		return
	}

	c <- message
}

func (broker *Broker[T]) SendToGroup(group string, message T) {
	for _, clientID := range broker.groups[group] {
		broker.clients[clientID] <- message
	}
}

func NewBroker[T any]() *Broker[T] {
	b := &Broker[T]{
		newConnect:                make(chan string),
		closingClients:            make(chan string),
		clients:                   make(map[string]chan T),
		groups:                    make(map[string][]string),
		ClientConnectedHandler:    func(clientID string) {},
		ClientDisconnectedHandler: func(clientID string) {},
	}

	go b.listen()

	return b
}
