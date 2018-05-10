package main

import "fmt"

var global int

const pi float32 = 3.1415

func main() {
	// First assignment
	n := 5
	fmt.Printf("%v\n", n)

	// Second assignment
	n = 6
	fmt.Printf("%v\n", n)

	// Multiple assignments
	a, b, c := true, "string", 0
	fmt.Printf("%v %v %v\n", a, b, c)

	// Var keyword, used to force the type (e.g. int16, int32 or int64)
	var v int32 = 7
	fmt.Printf("%v\n", v)

	// A variable is statically typed
	x := 5
	fmt.Printf("%v\n", x)
	// x = "some string" // Compilation error

	// Any type
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
	// s = i // Compilation error
	s = i.(string)
	fmt.Printf("%v\n", s)

	// Global variable
	global = 7

	// Constant
	fmt.Printf("%v\n", pi)
}
