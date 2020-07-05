package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar: ", i)
		time.Sleep(time.Duration(2 * time.Millisecond))
	}
	wg.Done()
}

/*
	we are just using "time.Sleep()" to visually confirm the concurrency via "go"
*/

// bar:  0
// foo:  0
// bar:  1
// foo:  1
// bar:  2
// foo:  2
// bar:  3
// bar:  4
// foo:  3
// foo:  4
