package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is closed
	// Distribute work across multiple functions (ten goroutines) that all read from in.
	xc := fanOut(in, 10)

	// FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 through c9 onto a single channel

	// TROUBLESHOOTING
	fmt.Printf("%T \n", xc)
	fmt.Println("*******************", len(xc))
	for _, v := range xc {
		fmt.Println("********", <-v)
	}

	for n := range merge(xc...) {
		fmt.Println(n)
	}

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

func fanOut(in <-chan int, n int) []<-chan int {
	var xc []<-chan int // the length of slice needed to be zero
	for i := 0; i < n; i++ {
		xc = append(xc, factorial(in))
	}
	return xc
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// go run main.go
// []<-chan int
// ******************* 10
// ******** 6
// ******** 24
// ******** 120
// ******** 720
// ******** 5040
// ******** 40320
// ******** 362880
// ******** 3628800
// ******** 39916800
// ******** 479001600
// 120
// 5040
// 40320
// 362880
// 3628800
// 6
// 5040
// 720
// 479001600
// 24
// 39916800
// 24
// 120
// 720
// 6
// 40320
// 362880
// 3628800
// 5040
// 720
// 40320
// 39916800
// 6
// 479001600
// 120
// 24
// 39916800
// 479001600
// 3628800
// 40320
// 24
// 120
// 720
// 6
// 362880
// 5040
// 720
// 5040
// 6
// 24
// 40320
// 120
// 362880
// 3628800
// 39916800
// 479001600
// 362880
// 39916800
// 479001600
// 720
// 6
// 5040
// 24
// 120
// 3628800
// 120
// 362880
// 39916800
// 24
// 40320
// 3628800
// 479001600
// 6
// 720
// 5040
// 362880
// 6
// 24
// 5040
// 3628800
// 40320
// 39916800
// 479001600
// 120
// 720
// 39916800
// 362880
// 3628800
// 479001600
// 40320
// 6
// 24
// 720
// 120
// 40320
// 5040
// 362880
// 39916800
// 3628800
// 479001600
