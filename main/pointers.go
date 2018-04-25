package main

import "fmt"

func main() {
	a := 7

	p := &a
	fmt.Printf("Reference=%v, value=%v\n", p, *p)

	/*
	Btw, Go has a Garbage Collector!
	No need to use malloc() or free()
	 */
}
