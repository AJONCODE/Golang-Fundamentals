package main

import "fmt"

func main() {
	var x [6]int
	fmt.Printf("%T \n", x)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[4])
	x[4] = 13
	fmt.Println(x[4])
	fmt.Println(x)
}

// [6]int
// [0 0 0 0 0 0]
// 6
// 0
// 13
// [0 0 0 0 13 0]
