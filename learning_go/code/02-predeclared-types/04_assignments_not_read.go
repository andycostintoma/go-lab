package main

import "fmt"

// on the Go Playground at https://oreil.ly/8JLA6
func main() {
	x := 10 // this assignment isn't read!
	x = 20
	fmt.Println(x)
	x = 30 // this assignment isn't read!
}
