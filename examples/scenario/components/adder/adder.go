package adder

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nenavizhuleto/horizon/scenario"
)

type AdderComponent struct {
	inputs  [2]*scenario.StreamConnection[int]
	outputs [1]*scenario.StreamConnection[string]
}

func New() *AdderComponent {
	return &AdderComponent{
		inputs: [2]*scenario.StreamConnection[int]{
			scenario.NewStream[int](),
			scenario.NewStream[int](),
		},
		outputs: [1]*scenario.StreamConnection[string]{
			scenario.NewStream[string](),
		},
	}
}

func (c *AdderComponent) Start(ctx context.Context, wg *sync.WaitGroup, name string) {
	log.Println("starting", name)

	timeout := 30 * time.Second
	timer := time.NewTimer(timeout)

	result := 0
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				c.outputs[0].Stream() <- fmt.Sprintf("result is %d\n", result)
			case <-timer.C:
				return
			}
		}
	}()
loop:
	for {
		select {
		case a := <-c.inputs[0].Stream():
			timer.Reset(timeout)
			result += a
		case b := <-c.inputs[1].Stream():
			timer.Reset(timeout)
			result += b
		case <-timer.C:
			break loop
		}
	}

	log.Println("stopping", name)

	wg.Done()
}
func (c *AdderComponent) Output(i int) scenario.Connectable { return c.outputs[i] }
func (c *AdderComponent) Input(i int) scenario.Connectable  { return c.inputs[i] }
