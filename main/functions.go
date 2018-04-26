package main

import "fmt"

// Simple function
func sum(a, b int) int {
	return a + b
}

// Multiple results
func sum2(a, b int) (int, bool) {
	n := a + b

	return n, n > 0
}

// Function first class support
func sum3(a, b int, f func(int, int) int) int {
	return f(a, b)
}

// Returns a function
func sum4() func(int, int) int {
	return func(i1 int, i2 int) int {
		return i1 + i2
	}
}

func main() {
	n1 := sum(5, 7)
	fmt.Printf("%v\n", n1)

	n2, positive := sum2(5, 7)
	fmt.Printf("value=%v, is positive=%v\n", n2, positive)

	n3, _ := sum2(5, 7) // Unused variable
	fmt.Printf("%v\n", n3)

	sum3(5, 7, sum)

	var f func(int, int) int = sum // Function assignment
	sum3(5, 7, f)

	f = sum // Function assignment with type inference
	sum3(5, 7, f)

	f2 := sum4()
	f2(5, 7)
}
