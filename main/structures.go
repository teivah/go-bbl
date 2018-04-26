package main

import "fmt"
import "bbl-golang/structures"

type a struct {
	foo int
	b   b
}

type b struct {
	bar int
	c   c
}

type c struct {
	baz int
}

func functionUsingCopy(p structures.Point) {
	p.X = 10

	fmt.Printf("x=%v, y=%v\n", p.X, p.Y)
}

func functionUsingPointer(p *structures.Point) {
	fmt.Printf("x=%v, y=%v\n", p.X, p.Y)
}

type pointWithAltitude struct {
	structures.Point
	altitude float32
}

func main() {
	/*
	Two ways to instantiate a structure
	 */

	p1 := structures.Point{1, 2}      // First way
	p1 = structures.Point{X: 1, Y: 2} // It is possible to precise the member name
	fmt.Printf("%v\n", p1)

	p2 := new(structures.Point) // Second way (Pointer allocation)
	p2.X = 1
	p2.Y = 2
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

	/*
	No inheritance, composition only
	 */
	p := pointWithAltitude{}
	p.X = 1
	p.Y = 2
	p.altitude = 32.4

	/*
	NPE?
	 */
	var v1 a // Zeroed
	var v2 a // Zeroed
	v2.b.c.baz = v1.b.c.baz
	fmt.Printf("%v\n", v2.b.c.baz)
}
