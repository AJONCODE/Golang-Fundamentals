package main

import "fmt"

func main() {
	i := 0

	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 10 {
			break
		}
	}
}

// 1
// 3
// 5
// 7
// 9
// 11
