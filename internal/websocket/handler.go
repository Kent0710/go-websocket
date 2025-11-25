package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// UpgradeHandler upgrades an HTTP connection to a WebSocket connection.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}

	defer conn.Close()

	// Listen for incoming messages
	for {
		// Read message from the client
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		fmt.Println("Received message:", string(message))
		// Send message back to the client
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}
