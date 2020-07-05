package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)

	go incrementor("foo: ")
	go incrementor("bar: ")

	wg.Wait()

	fmt.Println("Final counter: ", counter)
}

func incrementor(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)

		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, " counter: ", counter)
	}
	wg.Done()
}

// go run main.go
// bar:  0  Counter:  1
// foo:  0  Counter:  2
// bar:  1  Counter:  3
// bar:  2  Counter:  4
// foo:  1  Counter:  5
// bar:  3  Counter:  6
// bar:  4  Counter:  7
// foo:  2  Counter:  8
// foo:  3  Counter:  9
// bar:  5  Counter:  10
// foo:  4  Counter:  11
// foo:  5  Counter:  12
// bar:  6  Counter:  13
// foo:  6  Counter:  14
// bar:  7  Counter:  15
// foo:  7  Counter:  16
// bar:  8  Counter:  17
// foo:  8  Counter:  18
// foo:  9  Counter:  19
// bar:  9  Counter:  20
// Final counter:  20

// go run -race main.go
// foo:  0  Counter:  1
// bar:  0  Counter:  2
// foo:  1  Counter:  3
// foo:  2  Counter:  4
// bar:  1  Counter:  5
// foo:  3  Counter:  6
// foo:  4  Counter:  7
// bar:  2  Counter:  8
// bar:  3  Counter:  9
// foo:  5  Counter:  10
// bar:  4  Counter:  11
// bar:  5  Counter:  12
// foo:  6  Counter:  13
// bar:  6  Counter:  14
// foo:  7  Counter:  15
// bar:  7  Counter:  16
// foo:  8  Counter:  17
// bar:  8  Counter:  18
// bar:  9  Counter:  19
// foo:  9  Counter:  20
// Final counter:  20
