package main

import "fmt"

func main() {

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	fmt.Println(food)

}

// # command-line-arguments
// ./main.go:13:14: undefined: food
