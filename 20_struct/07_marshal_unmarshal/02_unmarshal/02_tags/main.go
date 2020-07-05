package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
}

func main() {
	var p1 person
	fmt.Println("p1 person fields initialized to zero values: ")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First":"James", "Last":"Bond", "wisdom score":20}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("--------------")
	fmt.Println(p1)
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)
}

// p1 person fields initialized to zero values: 
//
//
// 0
// --------------
// {James Bond 20}
// James
// Bond
// 20
// main.person 