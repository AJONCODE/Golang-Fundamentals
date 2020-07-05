package main

import "fmt"

func main() {
	age := 44

	fmt.Println("1. main func: ")
	fmt.Println(age)
	fmt.Println(&age)

	changeMe(&age)

	fmt.Println("2. main func: ")
	fmt.Println(age)
	fmt.Println(&age)
}

func changeMe(z *int) {

	fmt.Println("1. changeMe func: ")
	fmt.Println(z)
	fmt.Println(*z)

	*z = 24

	fmt.Println("2. changeMe func: ")
	fmt.Println(z)
	fmt.Println(*z)
}

// 1. main func:
// 44
// 0xc00010c000
// 1. changeMe func:
// 0xc00010c000
// 44
// 2. changeMe func:
// 0xc00010c000
// 24
// 2. main func:
// 24
// 0xc00010c000
