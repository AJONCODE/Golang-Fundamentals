package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{5, 23, 11, 17, 7, 10, 9, 4, 6, 8}

	sort.Ints(i)

	fmt.Println(i)
}

// [4 5 6 7 8 9 10 11 17 23]
