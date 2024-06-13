package protocol

import (
	"encoding/json"
)

type MessageType string

const MessageTypeSeparator = "."

const (
	MessageDetection = MessageType("detection")
	MessageFrame     = MessageType("frame")
	MessageMedia     = MessageType("media")
	MessageEvent     = MessageType("event")
	MessageAnalysis  = MessageType("analysis")
)

type RawMessage Message[json.RawMessage]

type Message[B any] struct {
	Producer Producer    `json:"producer"`
	Consumer Consumer    `json:"consumer"`
	Type     MessageType `json:"type"`
	Body     B           `json:"body"`
}

func NewMessage[T Typable](producer Producer, body T) Message[T] {
	return Message[T]{
		Producer: producer,
		Type:     body.Type(),
		Body:     body,
	}
}
