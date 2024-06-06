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

type Message[B any] struct {
	Producer Producer    `json:"producer"`
	Consumer Consumer    `json:"consumer"`
	Type     MessageType `json:"type"`
	Body     B           `json:"body"`
}

type Typable interface {
	Type() MessageType
}

func join(types ...MessageType) (result MessageType) {
	for _, t := range types {
		result = result + MessageTypeSeparator + t
	}
	return
}

func NewMessage[T Typable](producer Producer, body T) Message[T] {
	return Message[T]{
		Producer: producer,
		Type:     body.Type(),
		Body:     body,
	}
}

type RawMessage Message[json.RawMessage]
