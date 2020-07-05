package main

import "fmt"

func main() {
	n := 5
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 4; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for x := range c {
		fmt.Println(x)
	}
}

// go run -race main.go
// 0
// 0
// 0
// 0
// 0
// 1
// 1
// 1
// 1
// 1
// 2
// 2
// 3
// 2
// 2
// 2
// 3
// 3
// 3
// 3
