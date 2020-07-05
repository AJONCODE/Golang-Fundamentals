package main

import "fmt"

func main() {
	var x [26]string
	fmt.Printf("%T \n", x)
	fmt.Println(len(x))
	fmt.Println(x)
	for i := 65; i <= 90; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[25])
}

// [26]string
// 26
// [                         ]
// [A B C D E F G H I J K L M N O P Q R S T U V W X Y Z]
// A
// Z
