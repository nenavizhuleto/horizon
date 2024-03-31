package scenario

import (
	"context"
	"sync"
)

type Component interface {
	Start(context.Context, *sync.WaitGroup, string)

	Output(int) Connectable
	Input(int) Connectable
}
