package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 5 {
			break
		}
		i++
	}
}

// 0
// 1
// 2
// 3
// 4
// 5
