package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Printf("%T - %v \n", p1, p1)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Printf("%T - %v \n", p2, p2)
	fmt.Println(p2.first, p2.last, p2.age)
}

// main.person - {James Bond 20}
// James Bond 20
// main.person - {Miss Moneypenny 18}
// Miss Moneypenny 18
