package protocol

import "time"

type EventImage struct {
	AnalysisID string    `json:"analysis_id"`
	Timestamp  time.Time `json:"timestamp"`
	File       string    `json:"file"`
}

type EventMediaMessage struct {
	ProducerMessage[Producer]
	Recording string       `json:"recording"`
	Images    []EventImage `json:"images"`
}

func NewEventMediaMessage(producer Producer, video string, images []EventImage, options ...MessageOptions) EventMediaMessage {
	return EventMediaMessage{
		ProducerMessage: NewProducerMessage(MessageEventMedia, producer, options...),
		Recording:       video,
		Images:          images,
	}
}
