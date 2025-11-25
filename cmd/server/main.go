package main

import (
	"fmt"
	"net/http"

	"github.com/Kent0710/go-websocket/internal/websocket"
)

func main() {
	fmt.Println("Websocket server starting on :8080")

	// Register the WebSocket handler
	http.HandleFunc("/ws", websocket.WSHandler)

	// Star the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
