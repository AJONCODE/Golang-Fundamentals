package main

import "fmt"

func main() {
	var x [26]string

	fmt.Println("Length: ", len(x))
	fmt.Printf("Type %T \n", x)
	fmt.Println(x)
	for i := 65; i <= 90; i++ {
		x[i-65] = string(i)
	}
	for _, v := range x {
		fmt.Printf("%v - %T - %v \n", v, v, []byte(v))
	}
}

// Length:  26
// Type [26]string
// [                         ]
// A - string - [65]
// B - string - [66]
// C - string - [67]
// D - string - [68]
// E - string - [69]
// F - string - [70]
// G - string - [71]
// H - string - [72]
// I - string - [73]
// J - string - [74]
// K - string - [75]
// L - string - [76]
// M - string - [77]
// N - string - [78]
// O - string - [79]
// P - string - [80]
// Q - string - [81]
// R - string - [82]
// S - string - [83]
// T - string - [84]
// U - string - [85]
// V - string - [86]
// W - string - [87]
// X - string - [88]
// Y - string - [89]
// Z - string - [90]
