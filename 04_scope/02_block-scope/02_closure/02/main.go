package main

import "fmt"

var x = 0

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
Closure helps us limit the scope of variables used by multiple functions
without closure, for two or more functions to have access to the same variable,
that variable would need to be in package scope
*/

// 1
// 2
