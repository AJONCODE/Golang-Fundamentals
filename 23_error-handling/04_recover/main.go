package main

import "fmt"

func main() {
	doThings()
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("We panicked but everyone's fine")
		fmt.Println(r)
	}
}

func doThings() {
	defer recoverFromPanic()

	for i := 0; i < 5; i++ {
		if i == 2 {
			panic("PANIC: i is equal to 2")
		}
		fmt.Println(i)
	}
}

// go run main.go
// 0
// 1
// We panicked but everyone's fine
// PANIC: i is equal to 2
