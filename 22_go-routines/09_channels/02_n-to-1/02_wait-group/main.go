package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

// go run -race main.go
// 0
// 0
// 1
// 2
// 3
// 4
// 1
// 2
// 3
// 4
