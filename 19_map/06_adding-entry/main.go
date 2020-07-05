package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	myGreeting["Harleen"] = "Howdy"

	fmt.Println(myGreeting)
}

// map[Harleen:Howdy Jenny:Bonjour! Tim:Good morning!]
