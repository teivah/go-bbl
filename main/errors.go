package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		fmt.Println("b is equals to zero")
		return 0, errors.New("b cannot be equals to zero")
	}

	return a / b, nil
}

func main() {
	s, err := divide(1, 2)
	if err != nil {
		// Do something with the error
		// return
	}
	fmt.Printf("%v\n", s)

	s, err = divide(1, 0)
	if err != nil {
		// Do something with the error
		// return
	}
	fmt.Printf("%v\n", s)

	/*
	No throws, no catch
	 */
}
