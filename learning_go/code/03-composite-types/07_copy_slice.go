package main

import (
	"fmt"
)

func main() {
	// Example 1: Basic copy
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println("Example 1:", y, num) // Output: [1 2 3 4] 4

	// Example 2: Copy subset of a slice
	x = []int{1, 2, 3, 4}
	y = make([]int, 2)
	num = copy(y, x)
	fmt.Println("Example 2:", y, num) // Output: [1 2] 2

	// Example 3: Copy from the middle of the source slice
	x = []int{1, 2, 3, 4}
	y = make([]int, 2)
	copy(y, x[2:])
	fmt.Println("Example 3:", y) // Output: [3 4]

	// Example 4: Copy with overlapping sections
	x = []int{1, 2, 3, 4}
	num = copy(x[:3], x[1:])
	fmt.Println("Example 4:", x, num) // Output: [2 3 4 4] 3
}
