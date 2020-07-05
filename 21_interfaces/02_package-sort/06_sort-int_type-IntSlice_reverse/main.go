package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{5, 23, 11, 17, 7, 10, 9, 4, 6, 8}

	sort.Sort(sort.Reverse(sort.IntSlice(i)))

	fmt.Println(i)
}

// [23 17 11 10 9 8 7 6 5 4]
