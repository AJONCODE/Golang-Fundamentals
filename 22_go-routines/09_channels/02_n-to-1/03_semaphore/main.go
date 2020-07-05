package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

// go run main.go
// 0
// 1
// 0
// 2
// 1
// 3
// 4
// 2
// 3
// 4

// go run -race main.go
// 0
// 0
// 1
// 2
// 1
// 3
// 2
// 4
// 3
// 4
