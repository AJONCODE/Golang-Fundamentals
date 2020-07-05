package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementor("foo: ")
	go incrementor("bar: ")
	wg.Wait()
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter: ", counter)
	}
	wg.Done()
}

/*
	A data race is when two or more goroutines attempt to read and write to the same resource at the same time.
	Race conditions can create bugs that appear totally random or can never surface as they corrupt data.
	Atomic functions and mutexes are a way to synchronize the access of shared resources between goroutines.
*/

// go  run main.go
// bar:  0 Counter:  1
// bar:  1 Counter:  2
// foo:  0 Counter:  1
// foo:  1 Counter:  2
// foo:  2 Counter:  3
// bar:  2 Counter:  2
// foo:  3 Counter:  4
// foo:  4 Counter:  5
// foo:  5 Counter:  6
// bar:  3 Counter:  3
// bar:  4 Counter:  4
// bar:  5 Counter:  5
// foo:  6 Counter:  7
// foo:  7 Counter:  8
// bar:  6 Counter:  6
// bar:  7 Counter:  7
// foo:  8 Counter:  9
// bar:  8 Counter:  8
// foo:  9 Counter:  10
// bar:  9 Counter:  9
// bar:  10 Counter:  10
// foo:  10 Counter:  11
// bar:  11 Counter:  11
// bar:  12 Counter:  12
// foo:  11 Counter:  12
// bar:  13 Counter:  13
// foo:  12 Counter:  13
// foo:  13 Counter:  14
// foo:  14 Counter:  15
// bar:  14 Counter:  14
// bar:  15 Counter:  15
// foo:  15 Counter:  16
// bar:  16 Counter:  16
// bar:  17 Counter:  17
// foo:  16 Counter:  17
// foo:  17 Counter:  18
// foo:  18 Counter:  19
// foo:  19 Counter:  20
// bar:  18 Counter:  18
// bar:  19 Counter:  19

// go run -race main.go
// bar:  0 Counter:  1
// ==================
// WARNING: DATA RACE
// Write at 0x000000658640 by goroutine 7:
//   main.incrementor()
//       /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/22_go-routines/06_race-condition/main.go:25 +0xb5
// Previous write at 0x000000658640 by goroutine 8:
//   main.incrementor()
//       /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/22_go-routines/06_race-condition/main.go:25 +0xb5
// Goroutine 7 (running) created at:
//   main.main()
//       /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/22_go-routines/06_race-condition/main.go:15 +0x74
// Goroutine 8 (running) created at:
//   main.main()
//       /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/22_go-routines/06_race-condition/main.go:16 +0xa1
// ==================
// foo:  0 Counter:  1
// bar:  1 Counter:  2
// bar:  2 Counter:  3
// bar:  3 Counter:  4
// foo:  1 Counter:  2
// bar:  4 Counter:  5
// bar:  5 Counter:  6
// bar:  6 Counter:  7
// foo:  2 Counter:  3
// foo:  3 Counter:  4
// foo:  4 Counter:  5
// bar:  7 Counter:  8
// bar:  8 Counter:  9
// foo:  5 Counter:  6
// foo:  6 Counter:  7
// bar:  9 Counter:  10
// foo:  7 Counter:  8
// bar:  10 Counter:  11
// foo:  8 Counter:  9
// foo:  9 Counter:  10
// bar:  11 Counter:  12
// foo:  10 Counter:  11
// foo:  11 Counter:  12
// foo:  12 Counter:  13
// bar:  12 Counter:  13
// bar:  13 Counter:  14
// bar:  14 Counter:  15
// foo:  13 Counter:  15
// bar:  15 Counter:  16
// foo:  14 Counter:  17
// bar:  16 Counter:  17
// foo:  15 Counter:  18
// foo:  16 Counter:  19
// bar:  17 Counter:  18
// bar:  18 Counter:  19
// foo:  17 Counter:  20
// foo:  18 Counter:  21
// foo:  19 Counter:  22
// bar:  19 Counter:  20
// Found 1 data race(s)
// exit status 66
