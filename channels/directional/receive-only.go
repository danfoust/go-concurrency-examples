package main

import "fmt"

func main() {
	// Can only receive values
	ch := make(<-chan int, 1)

	// Arrow has to be on same side of channel as in declaration/signature
	ch <- 42

	fmt.Printf("%T\n", ch)
}
