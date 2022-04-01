package main

import "fmt"

// Arrow has to be on same side of channel as in declaration/signature
// *Note on confusing terminology
// This is a receiver channel that can only receive values
// But it can also called a "send-only" channel - as you can "only send" it values
// It cannot send values
func main() {
	// Can only receive values
	ch := make(chan<- int, 1)

	ch <- 42

	fmt.Printf("%T\n", ch)
}
