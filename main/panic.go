package main

import (
	"os"
	"fmt"
)

func anImportantFunction(filename string) string {
	src, err := os.Open(filename)

	if err != nil {
		panic(fmt.Errorf("the file %v cannot be opened", filename))
	}

	defer src.Close()
	// ...
	return ""
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recover from panic: %v", r)
		}
	}()

	anImportantFunction("somethingWhichDoesNotExist")
}
