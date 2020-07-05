package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world()
	hello()
}

/*
	right before main exits, world func executes.

	defer is often used to close a file
*/

// hello world
