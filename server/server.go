package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

// first connection through http and then upgraded to websocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("Received message: %s \n", msg)

		err = ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

const Port = 8080

func main() {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins
	}).Handler
	fmt.Println("FROM SERVER SIDE")
	http.HandleFunc("/ws", handleConnections)
	fmt.Println("Websocket server started on: 8080")
	err := http.ListenAndServe(fmt.Sprintf(":%d", Port), corsHandler(http.DefaultServeMux))
	// log.Fatal(http.ListenAndServe("localhost:8080", nil))
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
