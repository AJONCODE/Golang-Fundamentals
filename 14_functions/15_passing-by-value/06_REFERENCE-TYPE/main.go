package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // []
	changeMe(m)
	fmt.Println(m) // [Shikamaru]
}

func changeMe(z []string) {
	z[0] = "Shikamaru"
	fmt.Println(z) // [Shikamaru]
}

// []
// [Shikamaru]
// [Shikamaru]
