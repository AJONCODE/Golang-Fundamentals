package main

import "fmt"

func main() {
	fmt.Println(x)
	fmt.Println(y)

	x := 42 // order matters: block level scope
}

var y = 42 // order does not matter: package level scope

// # command-line-arguments
// ./main.go:6:14: undefined: x
