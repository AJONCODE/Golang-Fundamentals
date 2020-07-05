package main

import "fmt"

type ninja struct {
	name string
	team int
}

func main() {
	n1 := ninja{"Shikamaru", 10}
	fmt.Println("1. main func: ")
	fmt.Println(n1)
	fmt.Println(&n1.name)

	changeMe(&n1)

	fmt.Println("2. main func: ")
	fmt.Println(n1)
	fmt.Println(&n1.name)
}

func changeMe(z *ninja) {
	fmt.Println("1. changeMe func: ")
	fmt.Println(z)
	fmt.Println(&z.name)

	z.name = "Choji"

	fmt.Println("2. changeMe func: ")
	fmt.Println(z)
	fmt.Println(&z.name)

}

// 1. main func:
// {Shikamaru 10}
// 0xc00000c080
// 1. changeMe func:
// &{Shikamaru 10}
// 0xc00000c080
// 2. changeMe func:
// &{Choji 10}
// 0xc00000c080
// 2. main func:
// {Choji 10}
// 0xc00000c080
