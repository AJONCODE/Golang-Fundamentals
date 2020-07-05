package main

import "fmt"

func main() {
	switch "Sasuke" {
	case "Sasuke":
		fmt.Println("Sasuke Uchiha")
		fallthrough
	case "Kakashi":
		fmt.Println("Kakashi Hatake")
		fallthrough
	case "Shikamaru":
		fmt.Println("Shikamaru Nara")
	case "Naruto":
		fmt.Println("Naruto Uzumaki")
		fallthrough
	default:
		fmt.Println("Naruto Shippuden")
	}
}

// Sasuke Uchiha
// Kakashi Hatake
// Shikamaru Nara
