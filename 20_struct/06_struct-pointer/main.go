package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{"James", 27}
	fmt.Println(p1)
	fmt.Printf("%T \n", p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)
}

// &{James 27}
// *main.person
// James
// 27
