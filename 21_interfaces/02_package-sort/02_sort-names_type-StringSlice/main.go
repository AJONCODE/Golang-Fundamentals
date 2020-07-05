package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []string{"Ali", "Sancho", "Messi", "Bale", "Ronaldo"}

	y := []string{"Ali", "Sancho", "Messi", "Bale", "Ronaldo"}

	// 1. here we are converting "[]slice" into type "StringSlice" and then using StringSlice's "Sort" method
	sort.StringSlice(x).Sort()

	// 2. when we convert "[]slice" into type "StringSlice", we can see in "https://godoc.org/sort" that
	// "StringSlice" have built-in methods "Len", "Less" and "Swap" and these three methods are needed for us
	// to implement "Interface" interface. Hence we can use Interface's "Sort" method
	sort.Sort(sort.StringSlice(y))

	fmt.Println("Sorted x: ", x)
	fmt.Println("Sorted y: ", y)
}

/*
	Since we can't attach method here (as we do not have a type)

		// works properly
		type t []string
		func (z t) something() {
			fmt.Println(z)
		}

		// we'll get unresolved type error
		s := []string{"AJ", "CJ"}
		func (z s) something() {
			fmt.Println(z)
		}
*/

// Sorted x:  [Ali Bale Messi Ronaldo Sancho]
// Sorted y:  [Ali Bale Messi Ronaldo Sancho]
