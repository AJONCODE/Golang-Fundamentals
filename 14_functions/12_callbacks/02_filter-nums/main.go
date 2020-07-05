package main

import "fmt"

func main() {
	xs := filter([]int{1, 4, 6, 7, 8, 9, 10, 11, 17, 21, 23}, func(n int) bool {
		return n > 1
	})

	fmt.Println(xs)
}

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int

	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}

	return xs
}

// [4 6 7 8 9 10 11 17 21 23]
