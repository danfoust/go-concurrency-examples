package main

import "fmt"

func main() {
	ch := make(chan int)

	// Send
	go foo(ch)

	// Receive | note removing goroutine makes this blocking
	bar(ch)

	fmt.Println("Done")
}

// Channels can be reassigned from general to specific
// We tend to do this to prevent weird logical issues where
// a function is both sending and receiving

// Send
func foo(ch chan<- int) {
	ch <- 42
}

// Receive
func bar(ch <-chan int) {
	fmt.Println(<-ch)
}