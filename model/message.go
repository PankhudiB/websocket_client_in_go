package model

import "encoding/json"

type MessageWrapper struct {
	MessageType string          `json:"message_type"`
	Content     json.RawMessage `json:"content"`
}

type MessageTypeA struct {
	Name  string `json:"name"`
	Place string `json:"place"`
}
