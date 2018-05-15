package main

import "fmt"

func main() {
	// Array declaration
	var a [2]int
	a[0] = 0
	a[1] = 1
	fmt.Printf("%v\n", a) // 0, 1

	// Or
	b := [2]int{}         // Zeroed
	fmt.Printf("%v\n", b) // 0, 0

	// Or
	b = [2]int{0, 1}
	fmt.Printf("%v\n", b) // 0, 1

	// Possible
	b = a // Both variable are of type [2]int

	var c [3]int
	fmt.Printf("%v\n", c)
	// c = a // Compilation error

	/*
	An array is not dynamic
	 */
}
