package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"zero":  "Good morning!",
		"one":   "Bonjour!",
		"two":   "Buenos dias!",
		"three": "Bongiorno!",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, "two")
	fmt.Println(myGreeting)
}

// map[one:Bonjour! three:Bongiorno! two:Buenos dias! zero:Good morning!]
// map[one:Bonjour! three:Bongiorno! zero:Good morning!]
