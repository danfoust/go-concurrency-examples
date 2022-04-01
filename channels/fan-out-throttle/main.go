package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go populate(ch1)
	go fanOutIn(ch1, ch2)

	for v := range ch2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(ch1, ch2 chan int) {
	defer close(ch2)

	const goroutines = 10
	var wg sync.WaitGroup
		wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range ch1 {
				func(v2 int) {
					ch2 <- timeConsumingWork(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
