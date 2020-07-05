package main

import "fmt"

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)
}

// 42
// 0xc0000140b0
// 0xc0000140b0
// 42
