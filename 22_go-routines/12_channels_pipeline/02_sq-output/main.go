package main

import "fmt"

func main() {
	// Set up the pipeline and consume the output.
	for n := range sq(gen(2, 3)) {
		fmt.Println(n)
	}

	for n := range sq(sq(gen(2, 3))) {
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

// go run -race main.go
// 4
// 9
// 16
// 81
