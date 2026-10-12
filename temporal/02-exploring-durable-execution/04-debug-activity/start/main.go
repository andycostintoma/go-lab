package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"temporal102/04-debug-activity"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	order := *createPizzaOrder()

	workflowID := fmt.Sprintf("pizza-workflow-order-%s", order.OrderNumber)

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: debug_activity.TaskQueueName,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, debug_activity.PizzaWorkflow, order)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result debug_activity.OrderConfirmation
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalln("Unable to format order confirmation as JSON", err)
	}
	log.Printf("Workflow result: %s\n", string(data))
}

func createPizzaOrder() *debug_activity.PizzaOrder {
	customer := debug_activity.Customer{
		CustomerID: 12983,
		Name:       "María García",
		Email:      "maria1985@example.com",
		Phone:      "415-555-7418",
	}

	address := debug_activity.Address{
		Line1:      "701 Mission Street",
		Line2:      "Apartment 9C",
		City:       "San Francisco",
		State:      "CA",
		PostalCode: "94103",
	}

	p1 := debug_activity.Pizza{
		Description: "Large, with mushrooms and onions",
		Price:       1500,
	}

	p2 := debug_activity.Pizza{
		Description: "Small, with pepperoni",
		Price:       1200,
	}

	p3 := debug_activity.Pizza{
		Description: "Medium, with extra cheese",
		Price:       1300,
	}

	items := []debug_activity.Pizza{p1, p2, p3}

	order := debug_activity.PizzaOrder{
		OrderNumber: "Z1238",
		Customer:    customer,
		Items:       items,
		Address:     address,
		IsDelivery:  true,
	}

	return &order
}
