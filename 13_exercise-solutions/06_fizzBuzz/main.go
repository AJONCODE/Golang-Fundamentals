package main

import "fmt"

func main() {
	for i := 1; i <= 15; i++ {
		if i%15 == 0 {
			fmt.Println(i, " -- FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, " -- FIZZ")
		} else if i%5 == 0 {
			fmt.Println(i, " -- BUZZ")
		} else {
			fmt.Println(i)
		}
	}
}

// 1
// 2
// 3  -- FIZZ
// 4
// 5  -- BUZZ
// 6  -- FIZZ
// 7
// 8
// 9  -- FIZZ
// 10  -- BUZZ
// 11
// 12  -- FIZZ
// 13
// 14
// 15  -- FizzBuzz
