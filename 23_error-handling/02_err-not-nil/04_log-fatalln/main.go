package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

/*
	Package log implements a simple logging package ...
	writes to standard error and prints the date and time of each logged message ...
	the Fatal functions call os.Exit(1) after writing the log message ...
	the Panic functions call panic after writing the log message.
*/

/*
	Fatalln is equivalent to Println() followed by a call to os.Exit(1).
*/

// go run main.go
// 2020/06/23 10:15:59 open no-file.txt: no such file or directory
// exit status 1
