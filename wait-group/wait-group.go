package main

import (
	"fmt"
	"sync"
)

/*
	Concurrency: is about code DESIGN (patterns).

	You can write concurrent code that is not run in parallel

	Parallelism: is how code is RUN

	Be sure to pass in a pointer to wait group
*/

func loop(n int, wg *sync.WaitGroup) {
	defer wg.Done() // 3

	for i := 0; i < 10; i++ {
		fmt.Println("Routine: ", n+1, " | Step: ", i+1)
	}
}

func main() {
	var wg sync.WaitGroup // 1

	// We're simply using an iterator here, but it is best advised to evaluate
	// any go routine params before adding to wg.Add() & go f(), as to prevent
	// a panic
	for i := 0; i < 9; i++ {
		wg.Add(1) // 2

		// Do not pass wg by value!
		go loop(i, &wg)
	}
	fmt.Println("Waiting...")
	wg.Wait() // 4
	fmt.Println("Exiting...")
}
