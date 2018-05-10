package calculator

import "fmt"

func add(a, b int) int {
	return a + b
}

func race() {
	c := make(chan bool)
	m := make(map[string]string)

	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()

	m["2"] = "b" // Second conflicting access.

	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
