package main

import "fmt"

func main() {
	done := make(chan bool)
	values := []string{"a", "b", "c"}

	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	/*
		for _ = range values {
			<-done
		}
	*/
	for range values {
		<-done
	}
}

/*
Some confusion may arise when using closures with concurrency.
One might mistakenly expect to see a, b, c as the output.
What you'll probably see instead is c, c, c. This is because
each iteration of the loop uses the same instance of the variable v,
so each closure shares that single variable. When the closure runs,
it prints the value of v at the time fmt.Println is executed,
but v may have been modified since the goroutine was launched.
To help detect this and other problems before they happen,
run go vet.
SOURCE:
https://golang.org/doc/faq#closures_and_goroutines
*/

// go vet
// # github.com/AJONCODE/Golang-Fundamentals/22_go-routines/09_channels/08_closures/01_no-closure-binding
// ./main.go:11:16: loop variable v captured by func literal
// go run -race main.go
// ==================
// WARNING: DATA RACE
// Read at 0x00c00011a1e0 by goroutine 8:
//   main.main.func1()
//       /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/22_go-routines/09_channels/08_closures/01_no-closure-binding/main.go:11 +0x3c
// Previous write at 0x00c00011a1e0 by main goroutine:
//   main.main()
//       /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/22_go-routines/09_channels/08_closures/01_no-closure-binding/main.go:9 +0x119
// Goroutine 8 (running) created at:
//   main.main()
//       /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/22_go-routines/09_channels/08_closures/01_no-closure-binding/main.go:10 +0x163
// ==================
// ==================
// WARNING: DATA RACE
// Read at 0x00c00011a1e0 by goroutine 7:
//   main.main.func1()
//       /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/22_go-routines/09_channels/08_closures/01_no-closure-binding/main.go:11 +0x3c
// Previous write at 0x00c00011a1e0 by main goroutine:
//   main.main()
//       /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/22_go-routines/09_channels/08_closures/01_no-closure-binding/main.go:9 +0x119
// Goroutine 7 (running) created at:
//   main.main()
//       /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/22_go-routines/09_channels/08_closures/01_no-closure-binding/main.go:10 +0x163
// ==================
// c
// c
// c
// Found 2 data race(s)
// exit status 66
