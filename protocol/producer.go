package protocol

type ProducerType string

type Producer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// IP-Camera / Videofile
type MotionProducer struct {
	Producer
}

func NewMotionProducer(id, name string) MotionProducer {
	return MotionProducer{
		Producer: NewProducer(id, name),
	}
}

// Neural networks
type ObjectProducer struct {
	Producer
}

func NewObjectProducer(id, name string) ObjectProducer {
	return ObjectProducer{
		Producer: NewProducer(id, name),
	}
}

// Accelerometers, Light, Sound, Pressuer, Temperature, Humidity, Gas
type ValueProducer struct {
	Producer
}

func NewValueProducer(id, name string) ValueProducer {
	return ValueProducer{
		Producer: NewProducer(id, name),
	}
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
