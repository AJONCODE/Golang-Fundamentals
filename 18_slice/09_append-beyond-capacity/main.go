package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting = append(greeting, "Suprabadham")
	greeting = append(greeting, "Zǎo'ān")
	greeting = append(greeting, "Ohayou gozaimasu")
	greeting = append(greeting, "gidday")

	fmt.Println(greeting)
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}

// [Good morning! Bonjour! buenos dias! Suprabadham Zǎo'ān Ohayou gozaimasu gidday]
// 7
// 10