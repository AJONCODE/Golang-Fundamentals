package string

import (
	"fmt"
	"testing"
)

func TestMakeUppercase(t *testing.T) {
	expected := "SHIKAMARU NARA!"
	actual := MakeUppercase("shikamaru nara")

	if actual != expected {
		fmt.Errorf("unable to achieve the result: Expected: %s, Actual: %s", expected, actual)
	}
}

// o test
// PASS
// ok      github.com/AJONCODE/Golang-Fundamentals/24_testing/02_string 0.002s
