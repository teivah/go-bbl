package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func sum2(a, b int) (int, bool) {
	n := a + b

	return n, n > 0
}

func sum3(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func sum4() func(int, int) int {
	return func(i1 int, i2 int) int {
		return i1 + i2
	}
}

func main() {
	// Simple function
	n1 := sum(5, 7)
	fmt.Printf("%v\n", n1)

	// Multiple results
	n2, positive := sum2(5, 7)
	fmt.Printf("value=%v, is positive=%v\n", n2, positive)

	// Unused variable
	n3, _ := sum2(5, 7)
	fmt.Printf("%v\n", n3)

	// Function first class support
	sum3(5, 7, sum)

	// Function assignment, the type is optional
	var f func(int, int) int = sum

	// Closure
	rate := 5
	f = func(i int, i2 int) int {
		return i * i2 * rate
	}
	sum3(5, 7, f)

	// Returns a function
	f2 := sum4()
	f2(5, 7)
}
