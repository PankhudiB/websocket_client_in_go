package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket-in-go/dialer"
)

func main() {
	url := "ws://localhost:8080/talk-to-server"
	fmt.Println("Starting WebSocket Client...")
	fmt.Println("Dialing to WebSocket Server...\n\n")

	d := dialer.NewDialer()
	d.DialConn(url)
	go d.Read()

	httpEngine := gin.Default()
	err := http.ListenAndServe(":8081", httpEngine)
	if err != nil {
		fmt.Println("Error starting server!")
		return
	}
}
