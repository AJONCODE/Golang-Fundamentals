package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FName       string
	LName       string
	Age         int
	notExported string
}

func main() {
	p1 := Person{"Gerath", "Bale", 30, "Not Exported"}

	fmt.Println(p1)

	bs, _ := json.Marshal(p1)

	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}

/*
	we are not getting "notExported" data because "notExported" field in struct is not exported
*/

// {Gerath Bale 30 Not Exported}
// [123 34 70 78 97 109 101 34 58 34 71 101 114 97 116 104 34 44 34 76 78 97 109 101 34 58 34 66 97 108 101 34 44 34 65 103 101 34 58 51 48 125]
// []uint8 
// {"FName":"Gerath","LName":"Bale","Age":30}