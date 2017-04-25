package gopactor

import (
	"fmt"
	"log"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type Worker struct{}

func (w *Worker) Receive(ctx actor.Context) {}

func ExampleSpawnFromInstance() {
	// Given that the Worker actor is defined elsewhere
	worker, err := SpawnFromInstance(&Worker{})
	if err != nil {
		log.Print("Failed to spawn a worker")
		return
	}

	worker.Tell("Hello, world!")
}

func ExampleSpawnFromProducer() {
	producer := func() actor.Actor {
		return &Worker{}
	}

	worker, _ := SpawnFromProducer(producer)
	worker.Tell("Hello, world!")
}

func ExampleSpawnFromFunc() {
	f := func(ctx actor.Context) {
		if msg, ok := ctx.Message().(string); ok {
			fmt.Printf("Got a message: %s\n", msg)
		}
	}

	worker, _ := SpawnFromFunc(f)
	worker.Tell("Hello, world!")
	ShouldReceiveSomething(worker)
	// Output: Got a message: Hello, world!
}

func ExampleSpawnNullActor() {
	worker, _ := SpawnFromInstance(&Worker{})
	requestor, _ := SpawnNullActor()

	worker.Request("ping", requestor)
}
