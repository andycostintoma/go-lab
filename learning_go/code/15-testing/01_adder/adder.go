package adder

func addNumbers(x, y int) int {
	return x + x // Bug: should be x + y
}
