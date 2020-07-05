package main

import "fmt"

func main() {
	var x [26]byte

	fmt.Println("Length: ", len(x))
	fmt.Printf("Type: %T \n", x)
	fmt.Println(x)
	for i := 0; i <= 25; i++ {
		x[i] = byte(i)
	}
	for _, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
	}
}

// Length:  26
// Type: [26]uint8
// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
// 0 - uint8 - 0
// 1 - uint8 - 1
// 2 - uint8 - 10
// 3 - uint8 - 11
// 4 - uint8 - 100
// 5 - uint8 - 101
// 6 - uint8 - 110
// 7 - uint8 - 111
// 8 - uint8 - 1000
// 9 - uint8 - 1001
// 10 - uint8 - 1010
// 11 - uint8 - 1011
// 12 - uint8 - 1100
// 13 - uint8 - 1101
// 14 - uint8 - 1110
// 15 - uint8 - 1111
// 16 - uint8 - 10000
// 17 - uint8 - 10001
// 18 - uint8 - 10010
// 19 - uint8 - 10011
// 20 - uint8 - 10100
// 21 - uint8 - 10101
// 22 - uint8 - 10110
// 23 - uint8 - 10111
// 24 - uint8 - 11000
// 25 - uint8 - 11001
