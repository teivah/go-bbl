package main

import "fmt"

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}
	var slice []string

	/*
	Slicing an existing array (reference to an array)
	 */

	slice = array[0:5]
	slice = array[:]  // low:high
	slice = array[3:] // 3:high

	// A slice is a pointer to an array (plus a length and a capacity)
	array[4] = "m"
	fmt.Printf("array=%v, slice=%v\n", array, slice) // a, b, c, d, m and d, m

	slice = append(slice, "x") // Adding a new element to a slice
	fmt.Printf("%v\n", slice)  // d, m, x

	/*
	Creating a new slice
	 */

	slice2 := make([]int, 5) // type, length (and capacity)
	fmt.Printf("%v\n", slice2)

	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", len(slice3)) // 5
	slice3 = append(slice3, 6)
	fmt.Printf("%v\n", len(slice3)) // 6

	var slice4 []int
	fmt.Printf("%v\n", len(slice4)) // 0
}
