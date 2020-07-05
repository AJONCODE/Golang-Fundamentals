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
		errNorgateMath := fmt.Errorf("norgate math: square root of negative number: %v", f)
		return 0, errNorgateMath
	}

	// implementaion
	return math.Sqrt(f), nil
}

// go run main.go
// 2020/06/24 08:13:06 norgate math: square root of negative number: -10
// exit status 1
