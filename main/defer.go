package main

import (
	"fmt"
	"os"
)

func f1() {
	defer fmt.Println("A")
	defer fmt.Println("B")
	fmt.Println("C")
}

func f2() {
	src, _ := os.Open(".gitignore")

	defer src.Close() // Regardless of the execution results, the resource will be closed

	// ...
}

func main() {
	f1()
	f2()
}
