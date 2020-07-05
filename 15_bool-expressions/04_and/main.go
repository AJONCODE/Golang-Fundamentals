package main

import "fmt"

func main() {

	name := "Shikamaru"
	team := 10

	if name == "Shikamaru" && team == 7 {
		fmt.Println("This did not run")
	} else if name == "Shikamaru" && team == 10 {
		fmt.Println("Shikamru, Choji, Ino")
	} else {
		fmt.Println("Nothing ran")
	}

}

// Shikamru, Choji, Ino
