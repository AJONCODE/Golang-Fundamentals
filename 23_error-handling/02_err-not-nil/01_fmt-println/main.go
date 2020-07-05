package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")

	if err != nil {
		fmt.Println("error happened:", err)
	}
}

/*
	Println formats using the default formats for its operands and writes to standard output.
*/

// go run main.go
// error happened:  open no-file.txt: no such file or directory
