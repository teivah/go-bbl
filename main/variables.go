package main

import "fmt"

var global int

const pi float32 = 3.1415

func main() {
	// First assignment
	n := 5
	fmt.Printf("%v\n", n)

	// A variable is statically typed
	// n = "some string" // Compilation error

	// Second assignment
	n = 6
	fmt.Printf("%v\n", n)

	// Var keyword, used to force the type (e.g. int16, int32 or int64)
	var v int32 = 7
	fmt.Printf("%v\n", v)

	// Any type
	var i interface{}
	i = 32
	i = "sss"
	fmt.Printf("%v\n", i)

	// Global variable
	global = 7

	// Constant
	fmt.Printf("%v\n", pi)
}
