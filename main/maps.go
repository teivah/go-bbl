package main

import "fmt"

func main() {
	// Initializing a map
	m := make(map[int]string)
	m[0] = "zero"
	m[1] = "one"
	fmt.Printf("%v\n", m)

	// Other way
	m = map[int]string{
		0: "zero",
		1: "one",
	}
	fmt.Printf("%v\n", m)

	for k, v := range m {
		fmt.Printf("key=%v, value=%v\n", k, v)
	}
}
