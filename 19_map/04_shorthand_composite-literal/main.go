package main

import "fmt"

func main() {

	myGreeting := map[string]string{}
	myGreeting["Tim"] = "Good morning"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)
}

// map[Jenny:Bonjour Tim:Good morning]
