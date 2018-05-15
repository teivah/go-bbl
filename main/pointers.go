package main

import "fmt"

func main() {
	a := 7

	p := &a
	fmt.Printf("Reference=%v, value=%v\n", p, *p)
	fmt.Printf("Reference=%v, value=%v\n", &a, a)

	/*
   	Btw, Go has a Garbage Collector!
	*/
}
