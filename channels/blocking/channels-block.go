package main

import "fmt"

// Returns: Fatal error, all goroutines are asleep
// Know that channels BLOCK
func main() {
	ch := make(chan int)

	// Put a value into the channel
	ch <- 42 // BLOCKING

	// Pull a value out of the channel
	fmt.Println(<-ch)
}
