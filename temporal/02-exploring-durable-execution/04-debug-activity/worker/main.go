package main

import (
	"log"
	"temporal102/04-debug-activity"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, debug_activity.TaskQueueName, worker.Options{})

	w.RegisterWorkflow(debug_activity.PizzaWorkflow)
	w.RegisterActivity(debug_activity.GetDistance)
	w.RegisterActivity(debug_activity.SendBill)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}

}
