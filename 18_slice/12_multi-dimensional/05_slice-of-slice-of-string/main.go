package main

import (
	"fmt"
)

func main() {
	var records [][]string
	// student 1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"
	// store the record
	records = append(records, student1)
	// student 2
	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "96.00"
	// store the record
	records = append(records, student2)
	// print
	fmt.Println(records)
	fmt.Printf("Type: %T \n", records)
}

// [[Foster Nathan 100.00 74.00] [Gomez Lisa 92.00 96.00]]
// Type: [][]string
