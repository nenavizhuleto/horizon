package scenario

import (
	"context"
	"fmt"
	"sync"
)

type Script struct {
	name string

	mx         sync.Mutex // components mutex
	components map[string]Component

	connections []Connection
}

func NewScript(name string) *Script {
	return &Script{
		name:        name,
		components:  make(map[string]Component),
		connections: make([]Connection, 0),
	}
}

func (s *Script) String() string {
	return fmt.Sprintf("\ncomponents: %+v\nconnections: %+v\n", s.components, s.connections)
}

func (s *Script) NewComponent(name string, handler Component) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.components[name] = handler
}

func (s *Script) Connect(from_component string, output int, to_component string, input int) {
	s.mx.Lock()
	defer s.mx.Unlock()

	var exists bool

	c1, exists := s.components[from_component]
	c2, exists := s.components[to_component]

	if !exists {
		panic("unknown component")
	}

	c1.Output(output).Connect(c2.Input(input))

	{
		from_connector := NewConnector(from_component, output)
		to_connector := NewConnector(to_component, input)
		connection := NewConnection(from_connector, to_connector)
		s.connections = append(s.connections, connection)
	}
}

func (s *Script) Run() {

	ctx := context.Background()

	var wg sync.WaitGroup
	for name, component := range s.components {
		wg.Add(1)
		go component.Start(ctx, &wg, name)
	}

	wg.Wait()
}
