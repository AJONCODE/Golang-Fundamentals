package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	// FAN OUT
	// Distribute the sq work across two go routines that both read from in
	c1 := sq(in)
	c2 := sq(in)

	// FAN IN
	// Consume the merged output from multiple channels.
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	fmt.Printf("Type of nums: %T\n", nums)
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
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

func merge(cs ...chan int) chan int {
	fmt.Printf("Type of cs: %T\n", cs)
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

/*
type of variadic paramenter is a slice of the same type

FAN OUT
Multiple functions reading from the same channel until that channel is closed
FAN IN
A function can read from multiple inputs and proceed until all are closed by
multiplexing the input channels onto a single channel that's closed when
all the inputs are closed.
PATTERN
there's a pattern to our pipeline functions:
-- stages close their outbound channels when all the send operations are done.
-- stages keep receiving values from inbound channels until those channels are closed.
source:
https://blog.golang.org/pipelines
*/

/*
CHALLENGE #1
When know HOW to do fan out / fan in, but do we know WHY?
Why would we want to do fan out / fan in?
*/

// go run main.go
// Type of nums: []int
// Type of cs: []chan int
// 4
// 9
