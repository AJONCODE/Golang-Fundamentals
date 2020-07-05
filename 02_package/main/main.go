package main

import (
	"fmt"

	"github.com/AJONCODE/Golang-Fundamentals/02_package/bearfromalaska"
	"github.com/AJONCODE/Golang-Fundamentals/02_package/stringutil"
	alpha "github.com/AJONCODE/Golang-Fundamentals/02_package/wolf"
)

// alpha is used as an alias to use WolfName from /02_package/wolf

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.MyName)
	fmt.Println(bearfromalaska.BearName)
	fmt.Println(alpha.WolfName)
}

// Hello, Go!
// Ankit
// winniepooh
// Alpha
