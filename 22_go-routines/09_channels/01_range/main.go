package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

// 0
// 1
// 2
// 3
// 4
