package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened:", err)
	}
}

/*
	Package log implements a simple logging package ...
	writes to standard error and prints the date and time of each logged message ...
	the Fatal functions call os.Exit(1) after writing the log message ...
	the Panic functions call panic after writing the log message.
*/

/*
	log.Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
*/

// go run main.go
// 2020/06/23 09:55:59 err happened: open no-file.txt: no such file or directory
