package protocol

import "time"

type Recording struct {
	File string `json:"file"`
	URL  string `json:"url"`
}

type Media struct {
	AnalysisID string    `json:"analysis_id"`
	Timestamp  time.Time `json:"timestamp"`
	File       string    `json:"file"`
	URL        string    `json:"url"`
}

type MediaMessageBody struct {
	EventID   string    `json:"event_id"`
	Recording Recording `json:"recording"`
	Media     []Media   `json:"media"`
}

func (m MediaMessageBody) Type() MessageType {
	return MessageMedia
}

type MediaImageMessageBody struct {
	EventID string `json:"event_id"`
	Image   Media  `json:"image"`
}

func (m MediaImageMessageBody) Type() MessageType {
	return join(MessageMedia, "image")
}

type MediaRecordingMessageBody struct {
	EventID   string    `json:"event_id"`
	Recording Recording `json:"recording"`
}

func (m MediaRecordingMessageBody) Type() MessageType {
	return join(MessageMedia, "recording")
}
