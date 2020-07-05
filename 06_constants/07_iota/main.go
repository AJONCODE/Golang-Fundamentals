package main

import "fmt"

const (
	_  = iota             // 0
	Kb = 1 << (iota * 10) // 1 << (1 * 10)
	Mb = 1 << (iota * 10) // 1 << (2 * 10)
)

func main() {
	fmt.Println("binary \t \t decimal")
	fmt.Printf("%b \t", Kb)
	fmt.Printf("%d \n", Kb)
	fmt.Printf("%b \t", Mb)
	fmt.Printf("%d \n", Mb)
}

// binary           decimal
// 10000000000     1024
// 100000000000000000000   1048576

/*
bit shifting
	1 << 10
		would be 10000000000
	1 << 20
		would be 100000000000000000000
*/
