package protocol

type Producer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProducerMessage[T any] struct {
	Message
	Producer T `json:"producer"`
}

func NewProducer(id, name string) Producer {
	return Producer{
		ID:   id,
		Name: name,
	}
}

func NewProducerMessage[T any](t MessageType, producer T) ProducerMessage[T] {
	return ProducerMessage[T]{
		Message:  new_message(t),
		Producer: producer,
	}
}
