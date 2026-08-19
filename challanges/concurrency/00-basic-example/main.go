package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1. Create a channel to communicate safely between goroutines
	messageChan := make(chan string)
	var wg sync.WaitGroup

	// 2. Spawn a concurrent worker goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Worker: Starting background task...")
		time.Sleep(500 * time.Millisecond)     // Simulate background work
		messageChan <- "Worker: Task completed!" // Send data to the channel
	}()

	// 3. Block and wait to receive data from the channel in the main thread
	fmt.Println("Main: Waiting for worker...")
	result := <-messageChan
	fmt.Println("Main: Received message ->", result)

	// 4. Wait for the background worker to finish cleanly
	wg.Wait()
	fmt.Println("Main: All goroutines finished. Exiting.")
}
