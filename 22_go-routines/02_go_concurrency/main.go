package main

import "fmt"

func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo: ", i)
	}
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar: ", i)
	}
}

/*
	Nothing happened ?
		- We have 3 go routines that are running here: main, foo and bar and then main exits
*/
