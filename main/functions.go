package main

import "fmt"

// Simple function
func sum(a, b int) int {
	return a + b
}

// Multiple results
func sum2(a, b int) (int, bool) {
	n := sum(a, b)

	return n, n > 0
}

// Function first class support
func delegate(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	n1 := sum(5, 7)
	fmt.Printf("%v\n", n1)

	n2, positive := sum2(5, 7)
	fmt.Printf("value=%v, is positive=%v\n", n2, positive)

	n3, _ := sum2(5, 7) // Unused variable
	fmt.Printf("%v\n", n3)

	delegate(5, 7, sum)

	var f func(int, int) int = sum // Function assignment
	delegate(5, 7, f)

	f2 := sum // Function assignment with type inference
	delegate(5, 7, f2)
}
