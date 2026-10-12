package main

import (
	"log"
	"temporal102/03-testing-code"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, testing_code.TaskQueueName, worker.Options{})

	w.RegisterWorkflow(testing_code.SayHelloGoodbye)
	w.RegisterActivity(testing_code.TranslateTerm)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
