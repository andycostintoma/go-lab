package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"temporal102/03-testing-code"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "translation-workflow",
		TaskQueue: testing_code.TaskQueueName,
	}

	if len(os.Args) <= 2 {
		log.Fatalln("Must specify name and language code as command-line arguments")
	}

	input := testing_code.TranslationWorkflowInput{
		Name:         os.Args[1],
		LanguageCode: os.Args[2],
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, testing_code.SayHelloGoodbye, input)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result testing_code.TranslationWorkflowOutput
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalln("Unable to format result in JSON format", err)
	}
	log.Printf("Workflow result: %s\n", string(data))
}
