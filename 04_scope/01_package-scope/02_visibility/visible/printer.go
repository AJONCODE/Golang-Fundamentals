package visible

import "fmt"

func PrintVar() {
	fmt.Println(MyName)
	fmt.Println(yourName)
}

/*
MyName and yourName both can be accessed because we are inside visible package
*/
