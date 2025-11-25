package main

import (
	"fmt"
	"net/http"

	"github.com/Kent0710/go-websocket/internal/websocket"
)

func main() {
	// Register the WebSocket handler
	http.HandleFunc("/ws", websocket.WSHandler)
	go websocket.HandleMessages()

	fmt.Println("Websocket server starting on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
