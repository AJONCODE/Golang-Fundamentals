package main

import "fmt"

func main() {
	ns := []int{1, 4, 6, 7, 8, 9, 10}
	os := []int{11, 13, 17, 21, 23}

	ns = append(ns, os...)

	fmt.Println(ns)
	fmt.Println(len(ns))
	fmt.Println(cap(ns))
}

// [1 4 6 7 8 9 10 11 13 17 21 23]
// 12
// 14
