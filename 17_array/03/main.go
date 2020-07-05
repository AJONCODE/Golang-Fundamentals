package main

import "fmt"

func main() {
	var x [10]int

	fmt.Println("length: ", len(x))
	fmt.Println(x)
	for i := 0; i < 10; i++ {
		x[i] = i
	}
	for _, v := range x {
		fmt.Printf("%v - %T - %b \n", v, v, v)
	}
}

// length:  10
// [0 0 0 0 0 0 0 0 0 0]
// 0 - int - 0
// 1 - int - 1
// 2 - int - 10
// 3 - int - 11
// 4 - int - 100
// 5 - int - 101
// 6 - int - 110
// 7 - int - 111
// 8 - int - 1000
// 9 - int - 1001
