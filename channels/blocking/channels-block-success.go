package main

import "fmt"

// Solution to channel blocking
// Move it to another thread/process (think goroutine)
func main() {
	ch := make(chan int)

	// Spin off another process, main thread is no longer blocked
	go func() {
		ch <- 42 // Blocking code after in scope
	}()

	// Pull a value out of the channel
	fmt.Println(<-ch) // Blocks until it receives a value
}
