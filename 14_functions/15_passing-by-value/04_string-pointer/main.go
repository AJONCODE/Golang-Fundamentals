package main

import "fmt"

func main() {
	name := "Shikamaru"
	fmt.Println("1. main func: ")
	fmt.Println(name)
	fmt.Println(&name)

	changeMe(&name)

	fmt.Println("2. main func: ")
	fmt.Println(name)
	fmt.Println(&name)
}

func changeMe(z *string) {
	fmt.Println("1. changeMe func: ")
	fmt.Println(z)
	fmt.Println(*z)

	*z = "Kakashi"

	fmt.Println("2. changeMe func: ")
	fmt.Println(z)
	fmt.Println(*z)
}

// 1. main func:
// Shikamaru
// 0xc0000961e0
// 1. changeMe func:
// 0xc0000961e0
// Shikamaru
// 2. changeMe func:
// 0xc0000961e0
// Kakashi
// 2. main func:
// Kakashi
// 0xc0000961e0
