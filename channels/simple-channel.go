package main

import "fmt"

func incr(x int) int {
	return x * 5
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- incr(5)
	}()
	fmt.Println(<-ch)
}
