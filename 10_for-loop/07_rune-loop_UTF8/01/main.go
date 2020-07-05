package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

	fmt.Println("\n rune example: ")

	x := 'A'
	fmt.Println(x)
	fmt.Printf("%T \n", x)

	fmt.Println("\n string example: ")

	y := "A"
	fmt.Println(y)
	fmt.Printf("%T \n", y)
}

/*
NOTE:
Some operating systems (Windows) might not print characters where i < 256
If you have this issue, you can use this code:
fmt.Println(i, " - ", string(i), " - ", []int32(string(i)))

UTF-8 is the text coding scheme used by Go.
UTF-8 works with 1 - 4 bytes.
A byte is 8 bits.
[]byte deals with bytes, that is, only 1 byte (8 bits) at a time.
[]int32 allows us to store the value of 4 bytes, that is, 4 bytes * 8 bits per byte = 32 bits.
*/

// 65  -  A  -  [65]
// 66  -  B  -  [66]
// 67  -  C  -  [67]
// 68  -  D  -  [68]
// 69  -  E  -  [69]
// 70  -  F  -  [70]
// 71  -  G  -  [71]
// 72  -  H  -  [72]
// 73  -  I  -  [73]
// 74  -  J  -  [74]
// 75  -  K  -  [75]
// 76  -  L  -  [76]
// 77  -  M  -  [77]
// 78  -  N  -  [78]
// 79  -  O  -  [79]
// 80  -  P  -  [80]
// 81  -  Q  -  [81]
// 82  -  R  -  [82]
// 83  -  S  -  [83]
// 84  -  T  -  [84]
// 85  -  U  -  [85]
// 86  -  V  -  [86]
// 87  -  W  -  [87]
// 88  -  X  -  [88]
// 89  -  Y  -  [89]
// 90  -  Z  -  [90]

//  rune example:
// 65
// int32

//  string example:
// A
// string
