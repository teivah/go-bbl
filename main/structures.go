package main

import "fmt"
import "bbl-golang/structures"

type pointWithAltitude struct {
	structures.Point
	altitude float32
}

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

func byValue(p structures.Point) {
	p.X = 10
}

func byReference(p *structures.Point) {
	p.X = 10
}

func main() {
	/*
	A structure is a typed collection of field
	*/

	/*
	Two ways to instantiate a structure
	 */

	// First way
	structureInstance := structures.Point{1, 2}
	structureInstance = structures.Point{X: 1, Y: 2} // Naming members

	// Second way (Pointer allocation)
	structurePointer := new(structures.Point)
	structurePointer.X = 1
	structurePointer.Y = 2

	// Passing by value
	byValue(structureInstance)
	byValue(*structurePointer)

	// Passing by reference
	byReference(structurePointer)
	byReference(&structureInstance)

	/*
	Passing by value: immutability (copy)
	Passing by reference: performance (pointer)
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
