package main

import "fmt"

func main() {
	c := gen()
	receive(c)

	fmt.Println("Done")
}

func gen() <-chan int {
	c := make(chan int)

	go populate(c)

	return c
}

func populate(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}