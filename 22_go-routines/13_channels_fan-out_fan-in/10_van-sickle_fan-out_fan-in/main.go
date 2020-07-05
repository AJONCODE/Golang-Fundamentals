package main

import "fmt"

func main() {
	in := gen()
	out := make(chan int)

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is closed
	// Distribute work across multiple functions (ten goroutines) that all read from in.
	fanOut(in, 10, out)

	// FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 through c9 onto a single channel
	go func() {
		for n := range out {
			fmt.Println(n)
		}
	}()
}

func gen() <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()

	return out
}

func fanOut(in <-chan int, n int, out chan<- int) {
	for i := 0; i < n; i++ {
		factorial(in, out)
	}
}

func factorial(in <-chan int, out chan<- int) {
	go func() {
		for n := range in {
			out <- fact(n)
		}
	}()
}

func fact(n int) int {
	total := 1

	for i := n; i > 1; i-- {
		total *= n
	}

	return total
}

/*
	Mike Van Sickle
*/

// go run -race main.go
// 9
// 64
// 625
// 7776
// 117649
// 2097152
// 43046721
// 1000000000
// 25937424601
// 743008370688
// 9
// 64
// 625
// 7776
// 117649
// 2097152
// 43046721
// 1000000000
// 25937424601
// 743008370688
// 9
// 64
// 625
// 7776
// 117649
// 2097152
// 43046721
// 1000000000
// 25937424601
// 743008370688
// 9
// 64
// 625
// 7776
// 117649
// 2097152
// 43046721
// 1000000000
// 25937424601
// 743008370688
// 9
// 64
// 7776
// 625
// 117649
// 2097152
// 43046721
// 1000000000
// 25937424601
// 743008370688
// 9
// 64
// 625
// 7776
// 117649
// 2097152
// 43046721
// 1000000000
// 25937424601
// 743008370688
// 9
// 64
// 625
// 7776
// 117649
// 2097152
// 43046721
// 1000000000
// 25937424601
// 743008370688
// 9
// 64
// 625
// 7776
// 2097152
// 117649
// 43046721
// 1000000000
// 743008370688
// 25937424601
// 9
// 64
// 625
// 7776
// 117649
// 2097152
// 43046721
// 1000000000
// 25937424601
// 743008370688
// 9
// 64
// 743008370688
// 7776
// 117649
// 625
// 43046721
// 2097152
// 25937424601
// 1000000000
