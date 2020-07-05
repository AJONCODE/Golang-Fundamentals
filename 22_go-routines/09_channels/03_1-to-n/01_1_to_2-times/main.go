package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 25; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	<-done
	<-done
}

// go run -race main.go
// 0
// 1
// 2
// 3
// 4
// 5
// 7
// 8
// 9
// 10
// 6
// 12
// 11
// 13
// 15
// 16
// 14
// 18
// 17
// 20
// 21
// 22
// 23
// 19
// 24
