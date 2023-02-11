package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func main() {
	url := "ws://localhost:8080/talk-to-server"
	fmt.Println("Starting WebSocket Client...")

	fmt.Println("Dialing to WebSocket Server...\n\n")
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	//Sending message
	if err := conn.WriteMessage(websocket.TextMessage, []byte("Hello from client!\n")); err != nil {
		log.Fatal(err)
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Message:", string(message))
	}
}
