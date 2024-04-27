package protocol

import "time"

type Value Detection

type ValueDetectionMessage = DetectionMessage[Value]

func NewValueDetectionMessage(producer Producer, ts time.Time, values []Value, options ...MessageOptions) ValueDetectionMessage {
	return NewDetectionMessage(MessageValueDetection, producer, ts, values, options...)
}
