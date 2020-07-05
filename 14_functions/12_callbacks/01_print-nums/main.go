package main

import "fmt"

func main() {
	visit([]int{1, 4, 6, 7, 8, 9, 10, 11, 17, 21, 23}, func(n int) {
		fmt.Println(n)
	})
}

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

/*
	callback: passing a func as an argument
*/

// 1
// 4
// 6
// 7
// 8
// 9
// 10
// 11
// 17
// 21
// 23
