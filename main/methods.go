package main

import "fmt"

type customer struct {
	name     string
	lastname string
}

func display1(c customer) string { // Simple function
	return c.name + " " + c.lastname
}

func (c customer) display2() string { // Method with a receiver argument (point)
	return c.name + " " + c.lastname
}

func modifyName1(c customer, name string) {
	c.name = name
}

func modifyName2(c *customer, name string) {
	c.name = name
}

func (c customer) modifyName3(name string) { // Value receiver
	c.name = name
}

func (c *customer) modifyName4(name string) { // Pointer receiver
	c.name = name
}

func main() {
	c := customer{"Teiva", "Harsanyi"}

	s := display1(c)
	fmt.Printf("%v\n", s)

	// Method on a structure
	s = c.display2()
	fmt.Printf("%v\n", s)

	/*
	A method is just a function with a receiver argument
	 */

	c = customer{"Teiva", "Harsanyi"}
	modifyName1(c, "Bob")
	fmt.Printf("%v\n", c)

	c = customer{"Teiva", "Harsanyi"}
	modifyName2(&c, "Bob")
	fmt.Printf("%v\n", c)

	c = customer{"Teiva", "Harsanyi"}
	c.modifyName3("Bob")
	fmt.Printf("%v\n", c)

	c = customer{"Teiva", "Harsanyi"}
	c.modifyName4("Bob")
	fmt.Printf("%v\n", c)

	/*
	Value receiver: immutability, may be used in read-only methods
	 */

	/*
	Pointer receiver: performance, to be used in update methods
	 */
}
