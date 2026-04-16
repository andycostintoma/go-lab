package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MyStruct represents the structure of each JSON object in the array
type MyStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// process is a dummy function that processes each JSON object
func process(elem MyStruct) {
	fmt.Printf("Processed: %+v\n", elem)
}

func main() {
	// Simulated JSON array as input (could be from a file, network stream, etc.)
	jsonData := `[{"name": "Alice", "age": 30}, {"name": "Bob", "age": 25}, {"name": "Charlie", "age": 35}]`

	// Create a JSON decoder from the input data
	reader := strings.NewReader(jsonData)
	dec := json.NewDecoder(reader)

	// Read the opening `[` of the JSON array
	_, err := dec.Token()
	if err != nil {
		panic(err)
	}

	// Process each element in the array
	for dec.More() {
		var elem MyStruct
		err := dec.Decode(&elem)
		if err != nil {
			panic(err)
		}
		process(elem)
	}

	// Read the closing `]` of the JSON array
	_, err = dec.Token()
	if err != nil {
		panic(err)
	}
}
