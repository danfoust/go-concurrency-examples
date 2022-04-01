package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan int)

	go func() {
		const GOROUTINES = 10
		for i := 0; i < GOROUTINES; i++ {
			for i2 := 0; i2 < 10; i2++ {
				ch<- i2
			}
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	end := time.Now()
	fmt.Println("DONE.  Took ", end.Sub(start))
}