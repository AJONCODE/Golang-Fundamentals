package main

import "fmt"

func main() {
	myFootballers := map[int]string{
		1:  "David de Gea",
		2:  "Toby Alderweireld",
		4:  "Medhi Benatia",
		6:  "PAUL POGBA",
		7:  "Heung-Min Son",
		8:  "Arturo Vidal",
		9:  "Mauro Icardi",
		10: "Thiago Alcântara",
		11: "Gareth Bale",
		12: "Marcelo Vieira da Silva",
		14: "Radja Nainggolan",
		17: "Jérôme Boateng",
		27: "David Alaba",
	}

	for key, value := range myFootballers {
		fmt.Println(value, ": ", key)
	}
}

// David de Gea :  1
// Toby Alderweireld :  2
// Medhi Benatia :  4
// Heung-Min Son :  7
// Arturo Vidal :  8
// Mauro Icardi :  9
// Thiago Alcântara :  10
// Gareth Bale :  11
// Marcelo Vieira da Silva :  12
// PAUL POGBA :  6
// Radja Nainggolan :  14
// Jérôme Boateng :  17
// David Alaba :  27
