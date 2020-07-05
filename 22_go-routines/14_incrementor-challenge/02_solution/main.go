package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := incrementor(2)

	var count int
	for n := range c {
		count++
		fmt.Println(n)
	}

	fmt.Println("Final Count:", count)
}

func incrementor(n int) chan string {
	c := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for j := 0; j < 20; j++ {
				c <- fmt.Sprint("Process: "+strconv.Itoa(i)+" printing: ", j)
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	return c
}

// go run -race main.go
// Process: 0 printing: 0
// Process: 0 printing: 1
// Process: 0 printing: 2
// Process: 0 printing: 3
// Process: 1 printing: 0
// Process: 1 printing: 1
// Process: 0 printing: 4
// Process: 0 printing: 5
// Process: 1 printing: 2
// Process: 0 printing: 6
// Process: 0 printing: 7
// Process: 1 printing: 3
// Process: 1 printing: 4
// Process: 1 printing: 5
// Process: 0 printing: 8
// Process: 0 printing: 9
// Process: 1 printing: 6
// Process: 1 printing: 7
// Process: 0 printing: 10
// Process: 0 printing: 11
// Process: 0 printing: 12
// Process: 0 printing: 13
// Process: 1 printing: 8
// Process: 0 printing: 14
// Process: 0 printing: 15
// Process: 1 printing: 9
// Process: 0 printing: 16
// Process: 0 printing: 17
// Process: 1 printing: 10
// Process: 0 printing: 18
// Process: 0 printing: 19
// Process: 1 printing: 11
// Process: 1 printing: 12
// Process: 1 printing: 13
// Process: 1 printing: 14
// Process: 1 printing: 15
// Process: 1 printing: 16
// Process: 1 printing: 17
// Process: 1 printing: 18
// Process: 1 printing: 19
// Final Count: 40
