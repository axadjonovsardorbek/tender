package platform

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketHub struct {
	Clients   map[*websocket.Conn]bool
	Broadcast chan []byte
	Mutex     sync.Mutex
}

func NewWebSocketHub() *WebSocketHub {
	return &WebSocketHub{
		Clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan []byte),
	}
}

func (hub *WebSocketHub) Run() {
	for {
		message := <-hub.Broadcast
		hub.Mutex.Lock()
		for client := range hub.Clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("WebSocket error: %v", err)
				client.Close()
				delete(hub.Clients, client)
			}
		}
		hub.Mutex.Unlock()
	}
}

func (hub *WebSocketHub) AddClient(conn *websocket.Conn) {
	hub.Mutex.Lock()
	hub.Clients[conn] = true
	hub.Mutex.Unlock()
	log.Println("New WebSocket client connected.")
}

func (hub *WebSocketHub) RemoveClient(conn *websocket.Conn) {
	hub.Mutex.Lock()
	if _, ok := hub.Clients[conn]; ok {
		delete(hub.Clients, conn)
		conn.Close()
		log.Println("WebSocket client disconnected.")
	}
	hub.Mutex.Unlock()
}
