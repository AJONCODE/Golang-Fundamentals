package string

import "strings"

// MakeUppercase transforms a string to all caps with an exclamation point
func MakeUppercase(s string) string {
	return strings.ToUpper(s) + "!"
}
