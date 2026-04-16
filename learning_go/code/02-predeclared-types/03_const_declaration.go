package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

// this code will not compile
// ./ex3.go:23:2: cannot assign to x (constant 10 of type int64)
// ./ex3.go:24:2: cannot assign to y (untyped string constant "hello")
// on the Go Playground at https://oreil.ly/FdG-W
func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	x = x + 1
	y = "bye"

	fmt.Println(x)
	fmt.Println(y)
}
