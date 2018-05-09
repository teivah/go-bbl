package main

import (
	"fmt"
)

func displayThroughChannel(c <-chan int) { // <- is optional, in this case it indicates we cannot send data through this channel
	for v := range c {
		fmt.Printf("I've just received %v\n", v)
	}
}

func generateInts(n int) (<-chan int) {
	out := make(chan int, 1024)
	go func() {
		for i := 0; n-i > 0; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func map1(input <-chan int) (<-chan int) {
	out := make(chan int, 1024)

	go func() {
		for v := range input {
			out <- v * 2
		}
		close(out)
	}()

	return out
}

func map2(input <-chan int) (<-chan int) {
	out := make(chan int, 1024)

	go func() {
		for v := range input {
			out <- v + 1
		}
		close(out)
	}()

	return out
}

func main() {
	/*
	Go main philosophy:
	Don't communicate by sharing memory, share memory by communicating
	 */

	c := make(chan int, 1024) // Buffered channel, 10 maximum messages can be buffered before to block the sender goroutine (not the thread)
	go displayThroughChannel(c)

	c <- 5
	c <- 7

	close(c) // The best practice is to close a channel from the sender

	// c <- 8 // We can't send value to a closed channel, it will panic

	/*
	Benefits of communicating by sharing memory:
	- No lock
	- No false sharing
	 */

	/*
	Pipelines example
	*/

	source := generateInts(5)
	result := map2(map1(source))

	for v := range result {
		fmt.Printf("%v\n", v)
	}

	/*
	Small benchmark on 100k elements:
	- Java stream: 3'869'619 ns (3.8 ms)
	- Java stream: 2'952'112 ns (2.9 ms)
	- Go: 2'703 ns (0,0027 ms), 1000 times faster
	*/
}
