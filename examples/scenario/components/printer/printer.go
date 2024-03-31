package printer

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/nenavizhuleto/horizon/scenario"
)

type PrinterComponent struct {
	inputs  [1]*scenario.StreamConnection[string]
	outputs [1]*scenario.StreamConnection[int]
}

func New() *PrinterComponent {
	return &PrinterComponent{
		inputs: [1]*scenario.StreamConnection[string]{
			scenario.NewStream[string](),
		},
		outputs: [1]*scenario.StreamConnection[int]{
			scenario.NewStream[int](),
		},
	}
}

func (c *PrinterComponent) Start(ctx context.Context, wg *sync.WaitGroup, name string) {
	log.Println("starting", name)

	timeout := 30 * time.Second
	timer := time.NewTimer(timeout)

	result := 0

	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				c.outputs[0].Stream() <- result
			case <-timer.C:
				return
			}
		}
	}()
loop:
	for {
		select {
		case s := <-c.inputs[0].Stream():
			timer.Reset(timeout)
			result++
			log.Println(name, "got", s)
		case <-timer.C:
			break loop
		}
	}

	log.Println("stopping", name)

	wg.Done()
}
func (c *PrinterComponent) Output(i int) scenario.Connectable { return c.outputs[i] }
func (c *PrinterComponent) Input(i int) scenario.Connectable  { return c.inputs[i] }
