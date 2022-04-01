package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var count int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()

			atomic.AddInt64(&count, 1)
			runtime.Gosched()
			fmt.Println("Count: ", atomic.LoadInt64(&count))
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Final Count: ", count)
}
