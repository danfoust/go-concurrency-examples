package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	count := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()

			// Prevent multiple go-routines from accessing
			// code block between lock and unlock
			mu.Lock()

			v := count
			runtime.Gosched()
			v++
			count = v

			mu.Unlock()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Count: ", count)
}
