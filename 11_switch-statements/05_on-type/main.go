package main

import "fmt"

//  switch on types
//  -- normally we switch on value of variable
//  -- go allows you to switch on type of variable

type Contact struct {
	greeting string
	name     string
}

// SwitchOnType works with interfaces
// we'll learn more about interfaces later
func SwitchOnType(x interface{}) { // this is an assert; asserting, "x is of this type"
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("Shikamaru")
	var t = Contact{
		greeting: "Senpai",
		name:     "Kakashi",
	}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}

// int
// string
// contact
// string
// string
