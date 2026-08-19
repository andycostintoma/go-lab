package main

import (
	"fmt"
	"runtime"
	"time"
)

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {
	fmt.Println("goroutines before:", runtime.NumGoroutine())

	for i := range countTo(10) {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("goroutines after break:", runtime.NumGoroutine())

	// Give it a bit of time to observe
	for {
		fmt.Println("still alive goroutines:", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}
