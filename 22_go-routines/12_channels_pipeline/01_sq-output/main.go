package main

import "fmt"

func main() {
	// Set up the pipeline
	c := gen(2, 3)

	out := sq(c)

	// Consume the output.
	for n := range out {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)

	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()

	return out
}

func sq(c chan int) chan int {
	out := make(chan int)

	go func() {
		for n := range c {
			out <- n * n
		}
		close(out)
	}()

	return out
}

// go run main.go
// 4
// 9
