package main

import "fmt"

type person struct {
	fName string
	lName string
	team  int
}

func (p person) fullName() string {
	return p.fName + " " + p.lName
}

func main() {
	p1 := person{
		"Shikamaru",
		"Nara",
		10,
	}

	p2 := person{
		"Naruto",
		"Uzumaki",
		7,
	}

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}

// Shikamaru Nara
// Naruto Uzumaki
