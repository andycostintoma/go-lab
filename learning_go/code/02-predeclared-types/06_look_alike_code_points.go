package main

import "fmt"

// this code compiles, but lookalike characters are a very bad idea.
// on the Go Playground at https://oreil.ly/hrvb6
func main() {
	ａ := "hello"   // Unicode U+FF41
	a := "goodbye" // standard lowercase a (Unicode U+0061), different from the line above!
	fmt.Println(ａ)
	fmt.Println(a)
}
