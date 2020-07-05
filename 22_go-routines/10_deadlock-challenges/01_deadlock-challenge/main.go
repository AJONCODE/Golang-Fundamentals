package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

/*
This results in a deadlock.
Can you determine why?
And what would you do to fix it?
*/

// go run main.go
// fatal error: all goroutines are asleep - deadlock!
// goroutine 1 [chan send]:
// main.main()
//         /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/22_go-routines/10_deadlock-challenges/01_deadlock-challenge/main.go:9 +0x59
// exit status 2
