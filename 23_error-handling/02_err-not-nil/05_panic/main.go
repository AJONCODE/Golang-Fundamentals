package main

import (
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}
}

/*
	Package log implements a simple logging package ...
	writes to standard error and prints the date and time of each logged message ...
	the Fatal functions call os.Exit(1) after writing the log message ...
	the Panic functions call panic after writing the log message.
*/

// go run main.go
// panic: open no-file.txt: no such file or directory
// goroutine 1 [running]:
// main.main()
//         /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/23_error-handling/02_err-not-nil/05_panic/main.go:13 +0x74
// exit status 2
