package main

import "fmt"

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 10
	fmt.Println(a)
	fmt.Println(*b)
}

// 42
// 0xc0000140b0
// 0xc0000140b0
// 42
// 10
// 10

/*
this is useful
we can pass a memory address instead of a bunch of values (we can pass a reference)
and then we can still change the value of whatever is stored at that memory address
this makes our programs more performant

we don't have to pass around large amounts of data
*/
