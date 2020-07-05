package main

import "fmt"

func main() {
	switch "Shikamaru" {
	case "Sasuke":
		fmt.Println("Sasuke Uchiha")
	case "Kakashi":
		fmt.Println("Kakashi Hatake")
	case "Shikamaru":
		fmt.Println("Shikamaru Nara")
	case "Naruto":
		fmt.Println("Naruto Uzumaki")
	default:
		fmt.Println("Naruto Shippuden")
	}
}

/*
no default fallthrough
fallthrough is optional
-- you can specify fallthrough by explicitly stating it
-- if we specify fallthrough then it implies the next case will run
-- break isn't needed like in other languages
*/

// Shikamaru Nara
