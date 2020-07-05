package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}
}

/*
	Package log implements a simple logging package ...
	writes to standard error and prints the date and time of each logged message ...
	the Fatal functions call os.Exit(1) after writing the log message ...
	the Panic functions call panic after writing the log message.
*/

/*
	Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
*/

/*
	Since we have created a "log.txt" file in init block, the "log.txt" will catch all our log messages.
*/
