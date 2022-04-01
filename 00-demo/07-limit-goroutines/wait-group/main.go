package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()
	ch := make(chan int)

	go func() {
		const GOROUTINES = 10
		for i := 0; i < GOROUTINES; i++ {
			wg.Add(1)
			go func() {
				for i2 := 0; i2 < 10; i2++ {
					ch<- i2
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	end := time.Now()
	fmt.Println("DONE.  Took ", end.Sub(start))
}