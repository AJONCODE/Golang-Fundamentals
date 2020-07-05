package main

import (
	"fmt"
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
		return 0, fmt.Errorf("norgate math: square root of negative number: %v", f)
	}

	// implementaion
	return math.Sqrt(f), nil
}

// go run main.go
// 2020/06/24 08:09:50 norgate math: square root of negative number: -10
// exit status 1
