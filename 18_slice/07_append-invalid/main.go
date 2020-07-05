package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting[3] = "suprabadham"

	fmt.Println(greeting)
}

/*
	we can resolve this error by using append
*/

// panic: runtime error: index out of range [3] with length 3
// goroutine 1 [running]:
// main.main()
//         /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/18_slice/07_append-invalid/main.go:14 +0x1b
// exit status 2
