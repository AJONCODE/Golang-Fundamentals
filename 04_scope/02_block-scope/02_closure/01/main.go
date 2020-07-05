package main

import "fmt"

func main() {
	x := "Shikamaru Nara"

	fmt.Println(x)
	{
		fmt.Println(x)

		y := "Laziness is the mother of all bad habits, but ultimately she is a mother and we should respect her"
		fmt.Println(y)
	}
	// fmt.Println(y) // outside the scope of y
}

// Shikamaru Nara
// Shikamaru Nara
// Laziness is the mother of all bad habits, but ultimately she is a mother and we should respect her
