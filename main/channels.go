package main

import "fmt"

func displayThruChannel(c <-chan int) { // <- is optional, in this case it indicates we cannot send data through this channel
	for v := range c {
		fmt.Printf("I've just received %v\n", v)
	}
}

func main() {
	/*
	Go main philosophy:
	Don't communicate by sharing memory, share memory by communicating
	 */

	c := make(chan int, 10) // Buffered channel, 10 maximum messages can be buffered before to block the sender
	go displayThruChannel(c)

	c <- 5
	c <- 7

	close(c) // The best practice is to close a channel from the sender

	c <- 8 // We can't send value to a closed channel

	/*
	Benefits of communicating by sharing memory:
	- No lock
	- No false sharing
	 */
}
