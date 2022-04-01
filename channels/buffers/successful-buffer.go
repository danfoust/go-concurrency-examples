package main

import "fmt"

// Best practice to avoid buffered channels
// Code should be written with concurrency to avoid
// The issues of exceeding buffer and creating deadlock
func main() {
	// Buffered channel allows for x number of values to sit in channel
	// regardless of whether or not something is ready to pull it off
	ch := make(chan int, 1)

	ch <- 42 // Not blocking due to buffer

	fmt.Println(<-ch)
}