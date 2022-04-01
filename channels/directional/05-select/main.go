package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// Send
	go send(even, odd, quit)

	// Receive
	receive(even, odd, quit)

	fmt.Println("Done")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <- e:
			fmt.Println("From e channel: ", v)
		case v := <- o:
			fmt.Println("From o channel: ", v)
		case v := <- q:
			fmt.Println("From q channel: ", v)
			return
		}
	}
}

