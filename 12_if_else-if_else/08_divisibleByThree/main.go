package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

// 3
// 6
// 9
