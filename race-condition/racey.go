package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
 	How NOT to program
 	Do not try this in prod
	Or ever really
	Besides for educational purposes and stuff
*/
func main() {
	count := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := count

			runtime.Gosched()
			v++
			count = v

			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Count: ", count)
}
