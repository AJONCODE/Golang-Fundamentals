package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

// ErrNorgateMath - this is idiomatic to have error variables start with err.
// We are using Err because we want it to be accessible outside the package
var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	fmt.Printf("%T\n", ErrNorgateMath)

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}

	// implementaion
	return math.Sqrt(f), nil
}

/*
	see use of errors.New in standard library:
	http://golang.org/src/pkg/bufio/bufio.go
	http://golang.org/src/pkg/io/io.go
*/

// go run main.go
// *errors.errorString
// 2020/06/24 08:01:09 norgate math: square root of negative number
// exit status 1
