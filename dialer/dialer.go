package dialer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"net"
	"time"
	"websocket-in-go/model"
)

type dialer struct {
	conn               net.Conn
	url                string
	maxNoOfRetries     int
	timeBetweenRetries int
}

type Dialer interface {
	Read() []byte
	DialConn(url string) (net.Conn, error)
}

func NewDialer() Dialer {
	return &dialer{
		maxNoOfRetries:     6,
		timeBetweenRetries: 4,
	}
}

func (d *dialer) DialConn(url string) (net.Conn, error) {
	d.url = url
	conn, err := d.dial()
	if err != nil {
		log.Fatal("Unable to dial first time! ", err.Error())
	}
	return conn, err
}

func (d *dialer) dial() (net.Conn, error) {
	ctx := context.TODO()
	fmt.Println("Dialing conn....")
	conn, _, _, err := ws.Dial(ctx, d.url)
	if err != nil {
		log.Println("Error dialing conn: ", err.Error())
		return nil, err
	}
	d.conn = conn
	return conn, err
}

func (d *dialer) Read() []byte {
	for {
		dataBytes, _, err := wsutil.ReadServerData(d.conn)
		if err != nil {
			d.redial()
		} else {
			deserializeMessage(dataBytes)
		}
	}
}

func (d *dialer) redial() {
	for i := 0; i < d.maxNoOfRetries; i++ {
		_, err := d.dial()
		if err != nil {
			time.Sleep(4 * time.Second)
		} else {
			return
		}
	}
	log.Fatal("Tried maxNoOfRetries: ", d.timeBetweenRetries, " Could not connect !")
}

func deserializeMessage(dataBytes []byte) {
	messageWrapper, err := topLevelDeserializing(dataBytes)
	if messageWrapper.MessageType == "A" {
		deserializeSpecificType(err, messageWrapper)
	}
}

func deserializeSpecificType(err error, messageWrapper model.MessageWrapper) {
	subMessage := model.MessageTypeA{}
	err = json.Unmarshal(messageWrapper.Content, &subMessage)
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
	fmt.Println("\n\nLevel 2 Unmarshalling...", "\n", "subMessage.Name : ", subMessage.Name, "\n", "subMessage.Place : ", subMessage.Place)
}

func topLevelDeserializing(dataBytes []byte) (model.MessageWrapper, error) {
	messageWrapper := model.MessageWrapper{}
	err := json.Unmarshal(dataBytes, &messageWrapper)
	if err != nil {
		fmt.Println("Error unmarshalling...", err.Error())
	}
	fmt.Println("Level 1 Unmarshalling...", "\n", "Type : ", messageWrapper.MessageType, "\n", "Content : ", messageWrapper.Content)
	return messageWrapper, err
}
