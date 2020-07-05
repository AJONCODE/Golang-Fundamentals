package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	studyGroup := people{"Zeno", "Ali", "Sancho", "Messi", "Bale", "Ronaldo"}

	fmt.Println("Before sorting: ", studyGroup)

	sort.Sort(studyGroup)

	fmt.Println("After sorting: ", studyGroup)
}

// Before sorting:  [Zeno Ali Sancho Messi Bale Ronaldo]
// After sorting:  [Ali Bale Messi Ronaldo Sancho Zeno]
