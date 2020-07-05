package main

import "fmt"

func main() {

	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
}

// [Monday Tuesday Wednesday Thursday Friday]
// 5
// 5
