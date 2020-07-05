package main

import "fmt"

func main() {

	name := "Shikamaru"
	fmt.Println("1. main func: ")
	fmt.Println(name) // Shikamaru

	changeMe(name)

	fmt.Println("2. main func: ")
	fmt.Println(name) // Shikamaru
}

func changeMe(z string) {
	fmt.Println("1. changeMe func: ")
	fmt.Println(z) // Shikamaru

	z = "Kakashi"

	fmt.Println("2. changeMe func: ")
	fmt.Println(z) // Kakashi
}

// 1. main func:
// Shikamaru
// 1. changeMe func:
// Shikamaru
// 2. changeMe func:
// Kakashi
// 2. main func:
// Shikamaru
