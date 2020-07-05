package main

import "fmt"

func main() {
	x := 42
	fmt.Println("a - ", x)
	fmt.Println("a's memory address - ", &x)
}

// a -  42
// a's memory address -  0xc00010c000
