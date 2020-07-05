package main

import "fmt"

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

func main() {
	foo()
	bar()
}

// foo:  0
// foo:  1
// foo:  2
// foo:  3
// foo:  4
// bar:  0
// bar:  1
// bar:  2
// bar:  3
// bar:  4
