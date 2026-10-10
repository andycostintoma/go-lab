package main

import (
	"log"
	_2_farewell_workflow2 "temporal101/02-farewell-workflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "greeting-tasks", worker.Options{})

	w.RegisterWorkflow(_2_farewell_workflow2.GreetSomeone)
	w.RegisterActivity(_2_farewell_workflow2.GreetInSpanish)
	w.RegisterActivity(_2_farewell_workflow2.FarewellInSpanish)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
