package main

import "fmt"

func display(s string) {
	fmt.Printf("%v\n", s)
}

func main() {
	display("hello")    // Current goroutine
	go display("hello") // Starts a new goroutine

	/*
	A goroutine is a lightweight thread (not mapped on an OS thread but a soft thread)
	The Go execution environment decides how many threads are required
	For a CPU-bound & parallelizable operation, the optimal number of threads is ...?

	C10k problem
	Impacts: thread switching, TLAB size in Java etc.

	Benefits: lighter, faster to start, faster to switch
	 */

	/*
	How to communicate with a goroutine?
	 */
}
