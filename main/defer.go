package main

import (
	"fmt"
	"os"
)

func f1() {
	defer fmt.Println("First call")
	fmt.Println("Second call")
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
