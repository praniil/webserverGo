package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("FROM CLIENT SIDE")
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	fmt.Println(u.String())

	//connection to websocket server
	connection, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial: ", err)
	}

	defer connection.Close()

	//send message to server
	message := []byte("Hello, i am Pranil your client")
	err = connection.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatal("write: ", err)
	}

	//reading the servers response
	_, response, err := connection.ReadMessage()
	if err != nil {
		log.Fatal("read: ", err)
	}
	fmt.Printf("received response: %s \n", response)
}
