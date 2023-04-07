package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
)

func main() {
	url := "ws://localhost:8080/talk-to-server"
	fmt.Println("Starting WebSocket Client...")
	fmt.Println("Dialing to WebSocket Server...\n\n")

	conn, _, err := dialConn(url)
	dataBytes := read(conn)
	deserializeMessage(err, dataBytes)
}

func deserializeMessage(err error, dataBytes []byte) {
	messageWrapper, err := topLevelDeserializing(err, dataBytes)
	if messageWrapper.MessageType == "A" {
		deserializeSpecificType(err, messageWrapper)
	}
}

func deserializeSpecificType(err error, messageWrapper MessageWrapper) {
	subMessage := MessageTypeA{}
	err = json.Unmarshal(messageWrapper.Content, &subMessage)
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
	fmt.Println("\n\nLevel 2 Unmarshalling...", "\n", "subMessage.Name : ", subMessage.Name, "\n", "subMessage.Place : ", subMessage.Place)
}

func topLevelDeserializing(err error, dataBytes []byte) (MessageWrapper, error) {
	messageWrapper := MessageWrapper{}
	err = json.Unmarshal(dataBytes, &messageWrapper)
	if err != nil {
		fmt.Println("Error unmarshalling...", err.Error())
	}
	fmt.Println("Level 1 Unmarshalling...", "\n", "Type : ", messageWrapper.MessageType, "\n", "Content : ", messageWrapper.Content)
	return messageWrapper, err
}

func read(conn net.Conn) []byte {
	dataBytes, _, err := wsutil.ReadServerData(conn)
	if err != nil {
		fmt.Println("Error reading from websocket connection ! ", err.Error())
	}
	return dataBytes
}

func dialConn(url string) (net.Conn, ws.Handshake, error) {
	ctx := context.Background()
	conn, _, _, err := ws.Dial(ctx, url)
	if err != nil {
		log.Fatal(err.Error())
	}
	return conn, _, err
}

type MessageWrapper struct {
	MessageType string          `json:"message_type"`
	Content     json.RawMessage `json:"content"`
}

type MessageTypeA struct {
	Name  string `json:"name"`
	Place string `json:"place"`
}
