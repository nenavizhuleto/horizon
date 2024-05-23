package system

import "encoding/json"

type MessageType string

const (
	MessageNotification = MessageType("system.notification")
	MessageAction       = MessageType("system.action")
)

type RawMessage = Message[json.RawMessage]

type Message[B any] struct {
	Tags []string    `json:"tags"`
	Type MessageType `json:"type"`
	Body B           `json:"body"`
}
