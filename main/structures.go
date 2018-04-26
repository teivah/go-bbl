package main

import "fmt"

type point struct {
	// Visible only within the current package
	x int
	y int
}

type Point2 struct {
	// Visible also outside the current package
	x int
	y int
}

func functionUsingCopy(p point) {
	p.x = 10

	fmt.Printf("x=%v, y=%v\n", p.x, p.y)
}

func functionUsingPointer(p *point) {
	fmt.Printf("x=%v, y=%v\n", p.x, p.y)
}

func main() {
	/*
	Two ways to instantiate a structure
	 */

	p1 := point{1, 2}      // First way
	p1 = point{x: 1, y: 2} // It is possible to precise the member name
	fmt.Printf("%v\n", p1)

	p2 := new(point) // Second way (pointer allocation)
	p2.x = 1
	p2.y = 2
	fmt.Printf("%v\n", p2)

	functionUsingCopy(p1)
	functionUsingCopy(*p2)

	/*
	Copy benefits: immutability (passing by value)
	 */

	functionUsingPointer(p2)
	functionUsingPointer(&p1)

	/*
	Pointer benefits: performance (passing by reference)
	 */
}
