package main

import "fmt"

func zero(z *int) {
	*z = 0
}

func main() {
	x := 42

	zero(&x)
	fmt.Println(x)
}

// 0
