package protocol

type Camera struct {
	Name    string   `json:"name"`
	Group   string   `json:"group"`
	Modules []string `json:"modules"`
}

type ProducerOptions struct {
	Camera *Camera `json:"camera,omitempty"`
}

type Producer struct {
	ProducerOptions
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProducerMessage[T any] struct {
	Message
	Producer T `json:"producer"`
}

func NewProducer(id, name string, options ...ProducerOptions) Producer {
	return Producer{
		ID:              id,
		Name:            name,
		ProducerOptions: Options(options...),
	}
}

func NewProducerMessage[T any](t MessageType, producer T, options ...MessageOptions) ProducerMessage[T] {
	return ProducerMessage[T]{
		Message:  NewMessage(t, options...),
		Producer: producer,
	}
}
