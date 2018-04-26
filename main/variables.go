package main

import "fmt"

var global int

func main() {
	n := 5 // First assignment
	fmt.Printf("%v\n", n)

	n = 6 // Second assignment
	fmt.Printf("%v\n", n)

	a, b, c := true, "string", 0 // Multiple assignments
	fmt.Printf("%v %v %v\n", a, b, c)

	var v int32 = 7 // Var keyword
	fmt.Printf("%v\n", v)

	// A variable is strongly typed
	x := 5
	fmt.Printf("%v\n", x)
	// x = "some string" is not going to work

	// No generics
	var i interface{}
	i = 32
	fmt.Printf("%v\n", i)
	i = "sss"
	fmt.Printf("%v\n", i)

	// Switch on the type
	switch i.(type) {
	case int:
		// ...
	case string:
		// ...
	default:
		// ...
	}

	// Casting
	var s string
	// s = i // Not going to work
	s = i.(string)
	fmt.Printf("%v\n", s)

	// Global variable
	global = 7
}
