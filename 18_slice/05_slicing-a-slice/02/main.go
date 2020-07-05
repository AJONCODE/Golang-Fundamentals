package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	fmt.Print("[1:2] ")
	fmt.Println(greeting[1:2])
	fmt.Print("[:2] ")
	fmt.Println(greeting[:2])
	fmt.Print("[5:] ")
	fmt.Println(greeting[5:])
	fmt.Print("[:] ")
	fmt.Println(greeting[:])
}

// [1:2] [Bonjour!]
// [:2] [Good morning! Bonjour!]
// [5:] [Selamat pagi! Gutten morgen!]
// [:] [Good morning! Bonjour! dias! Bongiorno! Ohayo! Selamat pagi! Gutten morgen!]
