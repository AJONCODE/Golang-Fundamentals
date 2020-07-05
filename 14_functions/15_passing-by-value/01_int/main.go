package main

import "fmt"

func main() {
	age := 44
	changeMe(age)
	fmt.Println("Main func: ")
	fmt.Println(&age)
	fmt.Println(age)
}

func changeMe(z int) {
	fmt.Println("changeMe func: ")
	fmt.Println(&z)
	fmt.Println(z)
	z = 24
}

/*
	when changeMe is called on line 8
	the value 44 is being passed as an argument
*/

// changeMe func:
// 0xc0000b6018
// 44
// Main func:
// 0xc0000b6010
// 44
