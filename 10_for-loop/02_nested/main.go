package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, " - ", j)
		}
	}
}

// 0  -  0
// 0  -  1
// 0  -  2
// 1  -  0
// 1  -  1
// 1  -  2
// 2  -  0
// 2  -  1
// 2  -  2
