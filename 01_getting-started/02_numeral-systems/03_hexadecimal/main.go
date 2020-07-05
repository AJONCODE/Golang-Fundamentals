package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %X \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#x \n", 42, 42, 42)
}

// 42 - 101010 - 2a
// 42 - 101010 - 0x2a
// 42 - 101010 - 2A
// 42       101010          0x2a
