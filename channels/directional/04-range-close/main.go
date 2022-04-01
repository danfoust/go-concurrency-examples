package main

import "fmt"

func main() {
	ch := make(chan int)

	// Send
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch) // Must close to prevent deadlock fatal error
	}()

	// Receive
	for v := range ch { // Pulls off channel until channel closes, then closes loop
		fmt.Println(v)
	}

	fmt.Println("Done")
}
