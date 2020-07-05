package main

import "fmt"

func main() {
	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello ", name)
}

// Enter your name: Shikamaru
// Hello  Shikamaru
