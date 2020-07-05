package main

import "fmt"

func zero(x int) {
	fmt.Println("x inside zero function: ")
	fmt.Printf("%p \n", &x)
	fmt.Println(&x)
	x = 0
}

func main() {
	x := 5
	fmt.Println("x inside main function: ")
	fmt.Printf("%p \n", &x)
	fmt.Println(&x)
	zero(x)
	fmt.Println("final value of x: ", x)
}

// x inside main function:
// 0xc0000140b0
// 0xc0000140b0
// x inside zero function:
// 0xc0000140b8
// 0xc0000140b8
// final value of x:  5
