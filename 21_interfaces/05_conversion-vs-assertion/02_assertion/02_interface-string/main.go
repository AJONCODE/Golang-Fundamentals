package main

import "fmt"

func main() {
	var name interface{} = "Sydney"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}

/*
	ASSERTION (can only be performed on interface)
		str, ok := name.(string)
		we are checking if "name" interface to of type string and if yes then giving the required string as well
*/

// string
