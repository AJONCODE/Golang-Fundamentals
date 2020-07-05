package main

import (
	"fmt"
	"log"
	"math"
)

// NorgateMathError - using struct for more informative custom error
// We are using "N" in "NorgateMathError" because we want it to be accessible outside the package
type NorgateMathError struct {
	lat, long string
	err       error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, &NorgateMathError{"50.2289 N", "99.4656 W", nme}
	}

	// implementaion
	return math.Sqrt(f), nil
}

/*
	see use of structs with error type in standard library:
	http://www.goinggo.net/2014/11/error-handling-in-go-part-ii.html

	http://golang.org/pkg/net/#OpError
	http://golang.org/src/pkg/net/dial.go
	http://golang.org/src/pkg/net/net.go
	http://golang.org/src/pkg/encoding/json/decode.go
*/

// go run main.go
// 2020/06/24 08:21:25 a norgate math error occured: 50.2289 N 99.4656 W norgate math redux: square root of negative number: -10
// exit status 1
