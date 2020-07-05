package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 15; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 15; i++ {
		fmt.Println("bar: ", i)
	}
	wg.Done()
}

/*
	output may vary
*/

// bar:  0
// bar:  1
// bar:  2
// bar:  3
// bar:  4
// bar:  5
// bar:  6
// bar:  7
// foo:  0
// foo:  1
// foo:  2
// foo:  3
// foo:  4
// foo:  5
// foo:  6
// foo:  7
// foo:  8
// foo:  9
// foo:  10
// foo:  11
// foo:  12
// foo:  13
// foo:  14
// bar:  8
// bar:  9
// bar:  10
// bar:  11
// bar:  12
// bar:  13
// bar:  14
