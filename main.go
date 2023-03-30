package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
)

func main() {
	url := "ws://localhost:8080/talk-to-server"
	fmt.Println("Starting WebSocket Client...")

	fmt.Println("Dialing to WebSocket Server...\n\n")
	ctx := context.Background()
	conn, _, _, err := ws.Dial(ctx, url)
	if err != nil {
		log.Fatal(err.Error())
	}
	dataBytes, _, err := wsutil.ReadServerData(conn)
	if err != nil {
		fmt.Println("Error reading from websocket connection ! ", err.Error())
	}

	messageWrapper := MessageWrapper{}
	err = json.Unmarshal(dataBytes, &messageWrapper)
	if err != nil {
		fmt.Println("Error unmarshalling...", err.Error())
	}

	fmt.Println("Level 1 Unmarshalling...", "\n", "Type : ", messageWrapper.MessageType, "\n", "Content : ", messageWrapper.Content)

	subMessage := SubMessage{}
	err = json.Unmarshal(messageWrapper.Content, &subMessage)
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	fmt.Println("\n\nLevel 2 Unmarshalling...", "\n", "subMessage.Name : ", subMessage.Name, "\n", "subMessage.Place : ", subMessage.Place)
}

type MessageWrapper struct {
	MessageType string          `json:"message_type"`
	Content     json.RawMessage `json:"content"`
}

type SubMessage struct {
	Name  string `json:"name"`
	Place string `json:"place"`
}