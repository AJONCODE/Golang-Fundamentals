package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
Closure helps us limit the scope of variables used by multiple functions
without closure, for two or more functions to have access to the same variable,
that variable would need to be in package scope

anonymous function
a function without a name

func expression
when we assign a function to a variable
*/

// 1
// 2
