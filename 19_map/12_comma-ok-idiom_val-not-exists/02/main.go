package main

import "fmt"

func main() {

	ninjaTeam := map[string]int{
		"sasuke":    7,
		"shikamaru": 10,
	}

	fmt.Println(ninjaTeam)

	if val, exists := ninjaTeam["kakashi"]; exists {
		delete(ninjaTeam, "kakashi")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}
}

// map[sasuke:7 shikamaru:10]
// That value doesn't exist.
// val:  0
// exists:  false
