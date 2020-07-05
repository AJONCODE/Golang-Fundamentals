package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	fmt.Println(myGreeting)
}

// map[Jenny:Bonjour! Tim:Good morning!]
