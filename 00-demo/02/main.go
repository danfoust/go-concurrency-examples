package main

import "fmt"

func main() {
	//cs := make(chan<- int) // Simply change to bidirectional channel
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs) // Had error

	fmt.Printf("-----------\n")
	fmt.Printf("cs\t%T\n", cs)
}
