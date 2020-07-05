package main

import "fmt"

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
}

// 42
// 0xc0000140b0
// 0xc0000140b0

/*
The above code makes "b" pointer to memory address where an int is stored
"b" is of type "int pointer"

"*int"
	the * is part of the type
	b is of type "*int"
*/
