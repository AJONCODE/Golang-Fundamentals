package main

import (
	"fmt"

	"github.com/AJONCODE/Golang-Fundamentals/04_scope/01_package-scope/02_visibility/visible"
)

func main() {
	fmt.Println(visible.MyName)

	visible.PrintVar()
}

// Ankit
// Ankit
// Shikamaru

/*
fmt.Println(visible.yourName)
	this will give error :
		cannot refer to unexported name visible.yourName
		undefined: visible.yourName
*/
