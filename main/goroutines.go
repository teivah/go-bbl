package main

import (
	"fmt"
	"time"
)

func display(s string) {
	fmt.Printf("%v\n", s)
}

func main() {
	display("hello 1")    // Current goroutine
	go display("hello 2") // Starts a new goroutine
	time.Sleep(250 * time.Millisecond)

	/*
	A goroutine is a lightweight thread (not mapped on an OS thread but a soft thread)

	The Go execution environment decides how many threads are required
	1 thread -> * goroutines

	For a CPU-bound & parallelizable operation, the optimal number of threads is ...?

	C10k problem
	Impacts: thread memory consumption, thread creation & thread switching, TLAB size in Java etc.

	Benefits: lighter, faster to start, faster to switch
	 */

	/*
	How to communicate with a goroutine?
	 */
}
