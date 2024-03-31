package scenario

import "fmt"

type Connectable interface {
	Connect(Connectable) error
}

type Connector struct {
	Name string
	At   int
}

type Connection struct {
	From Connector
	To   Connector
}

func NewConnector(name string, at int) Connector {
	return Connector{
		Name: name,
		At:   at,
	}
}

func NewConnection(from Connector, to Connector) Connection {
	return Connection{
		From: from,
		To:   to,
	}
}

type StreamConnection[T any] struct {
	stream chan T
}

func NewStream[T any]() *StreamConnection[T] {
	return &StreamConnection[T]{
		stream: make(chan T),
	}
}

func (s *StreamConnection[T]) Connect(to Connectable) error {
	switch connector := to.(type) {
	case *StreamConnection[T]:
		go func() {
			for v := range s.stream {
				connector.stream <- v
			}
		}()
		return nil
	default:
		return fmt.Errorf("coudn't find valid connector")
	}
}

func (s *StreamConnection[T]) Stream() chan T {
	return s.stream
}
