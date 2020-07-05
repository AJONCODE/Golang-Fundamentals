package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Ali", "Sancho", "Messi", "Bale", "Ronaldo"}

	sort.Strings(s)

	fmt.Println(s)
}

// [Ali Bale Messi Ronaldo Sancho]
