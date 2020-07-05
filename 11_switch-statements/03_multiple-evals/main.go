package main

import "fmt"

func main() {
	switch "Shikamaru" {
	case "Sasuke", "Naruto", "Sakura":
		fmt.Println("Team 7")
	case "Shikamaru", "Choji", "Ino":
		fmt.Println("Team 10")
	case "Kiba", "Hinata", "Shino":
		fmt.Println("Team 8")
	case "Neji", "Lee", "Tenten":
		fmt.Println("Team 11")
	default:
		fmt.Println("Konohagakure village")
	}
}

// Team 10
