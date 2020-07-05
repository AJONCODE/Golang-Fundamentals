package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	c := make(chan int)

	go func() {
		wg.Add(1)
		for i := 0; i < 5; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 0; i < 5; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

// go run -race main.go
// 0
// 0
// 1
// 1
// 2
// 2
// 3
// 3
// 4
// 4
// ==================
// WARNING: DATA RACE
// Write at 0x00c0000160b8 by goroutine 9:
//   internal/race.Write()
//       /usr/local/go/src/internal/race/race.go:41 +0x114
//   sync.(*WaitGroup).Wait()
//       /usr/local/go/src/sync/waitgroup.go:128 +0x115
//   main.main.func3()
//       /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/22_go-routines/09_channels/02_n-to-1/01_race-condition/main.go:30 +0x38
// Previous read at 0x00c0000160b8 by goroutine 7:
//   internal/race.Read()
//       /usr/local/go/src/internal/race/race.go:37 +0x1e8
//   sync.(*WaitGroup).Add()
//       /usr/local/go/src/sync/waitgroup.go:71 +0x1fb
//   main.main.func1()
//       /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/22_go-routines/09_channels/02_n-to-1/01_race-condition/main.go:14 +0x45
// Goroutine 9 (running) created at:
//   main.main()
//       /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/22_go-routines/09_channels/02_n-to-1/01_race-condition/main.go:29 +0x110
// Goroutine 7 (running) created at:
//   main.main()
//       /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/22_go-routines/09_channels/02_n-to-1/01_race-condition/main.go:13 +0xb8
// ==================
// Found 1 data race(s)
// exit status 66
