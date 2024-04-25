package protocol

import "time"

type Report struct {
	Producer  Producer  `json:"producer"`
	Timestamp time.Time `json:"timestamp"`
	UserData  any       `json:"user_data"`
}
