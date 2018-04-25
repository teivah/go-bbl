package main

import "fmt"

func main() {
	// Array declaration
	var a[2]int
	a[0] = 0
	a[1] = 1
	fmt.Printf("%v\n", a)

	// Or
	b := [2]int{}
	b[0] = 0
	b[1] = 1
	fmt.Printf("%v\n", b)

	// Or
	b = [2]int{0, 1}
	fmt.Printf("%v\n", b)
}
