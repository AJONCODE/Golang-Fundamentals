package main

import "fmt"

func zero(z *int) {
	fmt.Println("z inside zero function: ")
	fmt.Println(z)
	*z = 0
}

func main() {
	x := 5
	fmt.Println("x inside main function: ")
	fmt.Println(&x)
	zero(&x)
	fmt.Println("final value of x: ", x)
}

// x inside main function:
// 0xc0000140b0
// z inside zero function:
// 0xc0000140b0
// final value of x:  0
