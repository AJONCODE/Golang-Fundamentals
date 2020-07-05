package main

import "fmt"

type person struct {
	fName string
	lName string
	age   int
}

type footballer struct {
	person
	club     string
	position string
}

func main() {
	f1 := footballer{
		person: person{
			fName: "Gerath",
			lName: "Bale",
			age:   30,
		},
		club:     "Real Madrid C.F.",
		position: "RWF",
	}

	f2 := footballer{
		person: person{
			fName: "Son",
			lName: "Heung-min",
			age:   27,
		},
		club:     "Tottenham Hotspur F.C.",
		position: "LWF",
	}

	fmt.Println(f1.fName, f1.lName, f1.age, f1.club, f1.position)
	fmt.Println(f2.fName, f2.lName, f2.age, f2.club, f2.position)
}

// Gerath Bale 30 Real Madrid C.F. RWF
// Son Heung-min 27 Tottenham Hotspur F.C. LWF
