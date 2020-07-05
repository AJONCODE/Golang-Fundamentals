package main

import "fmt"

func main() {
	fmt.Println(greet("Shikamaru ", "Nara"))
}

func greet(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

/*
	Try not to use it (less readable)
*/

// Shikamaru Nara
