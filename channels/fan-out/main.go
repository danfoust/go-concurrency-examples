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

	fmt.Println("Done")
}

func populate(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func fanOutIn(ch1, ch2 chan int) {
	defer close(ch2)

	var wg sync.WaitGroup
	for v := range ch1 {
		wg.Add(1)
		go func(v2 int) {
			ch2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))

	return n + rand.Intn(1000)
}