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

type MediaMessage struct {
	EventID   string    `json:"event_id"`
	Recording Recording `json:"recording"`
	Media     []Media   `json:"media"`
}

func NewMediaMessage(producer Camera, message MediaMessage) Message[MediaMessage] {
	return Message[MediaMessage]{
		Producer: producer,
		Type:     MessageMedia,
		Body:     message,
	}
}
