package main

import "fmt"

func main() {
	n := 5
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 50; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func() {
			for n := range c {
				fmt.Println(n)
			}
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}
}

// go run -race main.go
// 0
// 1
// 3
// 2
// 4
// 5
// 7
// 6
// 9
// 8
// 11
// 13
// 12
// 14
// 17
// 15
// 10
// 18
// 20
// 23
// 24
// 19
// 21
// 26
// 27
// 29
// 30
// 31
// 32
// 33
// 34
// 35
// 22
// 37
// 16
// 28
// 25
// 36
// 41
// 43
// 44
// 45
// 46
// 47
// 39
// 49
// 38
// 40
// 42
// 48
