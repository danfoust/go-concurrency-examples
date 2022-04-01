package main

import "fmt"

func main() {
	// Buffered channel allows for x number of values to sit in channel
	// regardless of whether or not something is ready to pull it off
	ch := make(chan int, 1)

	ch <- 42 // Not blocking due to buffer
	ch <- 43 // BLOCKING, exceeds buffer

	fmt.Println(<-ch)
}