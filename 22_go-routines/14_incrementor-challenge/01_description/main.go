package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go incrementor("1")
	go incrementor("2")
	wg.Wait()
	fmt.Println("Final Counter:", count)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		atomic.AddInt64(&count, 1)
		fmt.Println("Process: "+s+" printing:", i)
	}
	wg.Done()
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/

// go run main.go
// Process: 2 printing: 0
// Process: 2 printing: 1
// Process: 2 printing: 2
// Process: 2 printing: 3
// Process: 1 printing: 0
// Process: 1 printing: 1
// Process: 1 printing: 2
// Process: 1 printing: 3
// Process: 1 printing: 4
// Process: 1 printing: 5
// Process: 1 printing: 6
// Process: 1 printing: 7
// Process: 1 printing: 8
// Process: 1 printing: 9
// Process: 1 printing: 10
// Process: 1 printing: 11
// Process: 1 printing: 12
// Process: 1 printing: 13
// Process: 1 printing: 14
// Process: 1 printing: 15
// Process: 1 printing: 16
// Process: 1 printing: 17
// Process: 1 printing: 18
// Process: 1 printing: 19
// Process: 2 printing: 4
// Process: 2 printing: 5
// Process: 2 printing: 6
// Process: 2 printing: 7
// Process: 2 printing: 8
// Process: 2 printing: 9
// Process: 2 printing: 10
// Process: 2 printing: 11
// Process: 2 printing: 12
// Process: 2 printing: 13
// Process: 2 printing: 14
// Process: 2 printing: 15
// Process: 2 printing: 16
// Process: 2 printing: 17
// Process: 2 printing: 18
// Process: 2 printing: 19
// Final Counter: 40
