package main

import (
	"fmt"
)

type scalable interface {
	scale(float32)
}

/* Point */

type point struct {
	x float32
	y float32
}

func (p *point) scale(n float32) { // Implicit implementation
	p.x *= n
	p.y *= n
}

/* Car */

type car struct {
	speed float32
}

func (c *car) scale(n float32) { // Implicit implementation
	c.speed *= n
}

/* Interface as an argument */

func f(s scalable, n float32) {
	s.scale(n)
	fmt.Printf("%v\n", s)
}

func main() {
	p := point{1.1, 3.2}
	p.scale(2.0)
	fmt.Printf("%v\n", p)

	c := car{100}
	c.scale(1.5)
	fmt.Printf("%v\n", c)

	/*
	Go's philosophy:
	if it quacks like a duck and swims like a duck, then it's a duck
	 */

	var s scalable // Define a scalable

	s = &p
	f(s, 0.5)

	s = &c
	f(s, 0.5)

	/*
	Benefits: not tightly coupled to the implementation
	 */

	/*
	Structure: a collection of field
	Interface: a collection of methods
	Clear segregation between state and behavior
	 */
}
