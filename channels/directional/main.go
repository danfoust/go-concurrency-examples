package main

import "fmt"

func main() {
	ch := make(chan int)
	chr := make(<-chan int) // Receive (called sent-only...)
	chs := make(chan<- int) // Send

	fmt.Printf("%T\n", ch)
	fmt.Printf("%T\n", chr)
	fmt.Printf("%T\n", chs)

	// Specific to general doesn't assign
	// ch = chr
	// ch = chs

	// Specific to specific doesn't assign
	// chr = chs

	// General to specific assigns
	//chr = ch
}
