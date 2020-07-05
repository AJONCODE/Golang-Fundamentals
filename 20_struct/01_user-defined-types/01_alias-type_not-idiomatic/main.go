package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)
}

// main.foo 44
