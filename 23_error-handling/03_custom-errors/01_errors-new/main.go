package main

import (
	"errors"
	"log"
	"math"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}

	// implementaion
	return math.Sqrt(f), nil
}

// go run main.go
// 2020/06/24 07:46:05 norgate math: square root of negative number
// exit status 1
