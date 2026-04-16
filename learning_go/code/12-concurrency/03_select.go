package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v           // Sends v to ch1
		v2 := <-ch2        // Blocks here, waiting for a value from ch2
		fmt.Println(v, v2) // This line is never executed
	}()

	v := 2
	var v2 int
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}

	fmt.Println(v, v2)
	// The main goroutine exits immediately after this,
	// terminating the program and killing the other goroutine.
	// The launched goroutine is stuck waiting to receive from ch2,
	// so it never reaches the fmt.Println statement.
}
