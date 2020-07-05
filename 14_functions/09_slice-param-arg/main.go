package main

import "fmt"

func main() {
	data := []float64{1, 4, 6, 7, 8, 9, 10, 11, 13, 17, 21, 23}
	n := average(data)
	fmt.Println(n)
}

func average(sf []float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}

// [1 4 6 7 8 9 10 11 13 17 21 23]
// []float64
// 10.833333333333334
