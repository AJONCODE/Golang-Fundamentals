package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		1: "Good morning!",
		2: "Bonjour!",
		3: "Buenos dias!",
		4: "Bongiorno!",
	}

	fmt.Println(myGreeting)

	delete(myGreeting, 3)

	fmt.Println(myGreeting)
}

// map[1:Good morning! 2:Bonjour! 3:Buenos dias! 4:Bongiorno!]
// map[1:Good morning! 2:Bonjour! 4:Bongiorno!]
