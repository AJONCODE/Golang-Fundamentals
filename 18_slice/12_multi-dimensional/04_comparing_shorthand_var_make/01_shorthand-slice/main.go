package main

import (
	"fmt"
)

func main() {
	student := []string{}
	students := [][]string{}
	student[0] = "Todd"
	// student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}

// panic: runtime error: index out of range [0] with length 0
// goroutine 1 [running]:
// main.main()
//         /home/ajoncode/goworkspace/src/github.com/AJONCODE/Golang-Fundamentals/18_slice/12_multi-dimensional/04_comparing_shorthand_var_make/01_shorthand-slice/main.go:10 +0x18
// exit status 2
