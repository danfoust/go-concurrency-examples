package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// Send
	go send(even, odd, quit)

	// Receive
	receive(even, odd, quit)

	fmt.Println("Done")
}

func send(even, odd chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(q)
}

func receive(even, odd <-chan int, q <-chan bool) {
	for {
		select {
		case v := <- even:
			fmt.Println("From e channel: ", v)
		case v := <- odd:
			fmt.Println("From o channel: ", v)
		case i, ok := <- q:
			if !ok {
				fmt.Println("From comma !ok channel: ", i, ok)
				return
			} else {
				fmt.Println("From comma ok channel: ", i)
			}
		}
	}
}

