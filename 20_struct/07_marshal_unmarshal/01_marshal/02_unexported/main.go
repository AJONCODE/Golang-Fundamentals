package main

import (
	"encoding/json"
	"fmt"
)

type footballer struct {
	fName       string
	lName       string
	age         int
}

func main() {
	f1 := footballer{"Gerath", "Bale", 30}

	fmt.Println(f1)

	bs, _ := json.Marshal(f1)

	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}

/*
	we are not getting the data because all of the fields in struct are not exported
*/

// {Gerath Bale 30}
// [123 125]
// []uint8 
// {}