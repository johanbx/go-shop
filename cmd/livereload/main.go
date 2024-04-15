package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan string)            // broadcast channel

// Handle incoming WebSocket connections
func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		log.Printf("Received: %s", message)
	}
}

// Broadcast messages to all clients
func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// Trigger reload via HTTP endpoint
func triggerReload(w http.ResponseWriter, r *http.Request) {
	broadcast <- "reload"
	w.Write([]byte("Reload triggered"))
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	http.HandleFunc("/trigger-reload", triggerReload)
	go handleMessages()

	log.Println("HTTP server started on :5555")
	err := http.ListenAndServe(":5555", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
