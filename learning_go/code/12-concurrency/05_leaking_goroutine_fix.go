package main

import (
	"fmt"
)

func countTo(max int, done <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < max; i++ {
			select {
			case ch <- i:
			case <-done:
				return
			}
		}
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	for i := range countTo(10, done) {
		if i > 5 {
			close(done) // stop the sender
			break
		}
		fmt.Println(i)
	}
}
