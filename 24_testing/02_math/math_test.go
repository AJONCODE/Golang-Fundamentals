package math

import (
	"testing"
)

func TestAdder(t *testing.T) {
	expected := 11 // pass
	// expected := 5 // fail
	actual := Adder(4, 7)

	if actual != expected {
		t.Errorf("Adder function does not add up: Expected: %d, Actual: %d", expected, actual)
		// t.Fatal("4 + 7 did not equal 11")
	}
}

func BenchmarkAdder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Adder(4, 7)
	}
}

// go test
// PASS
// ok      github.com/AJONCODE/Golang-Fundamentals/24_testing/01_math   0.004s

// go test -bench=.
// goos: linux
// goarch: amd64
// pkg: github.com/AJONCODE/Golang-Fundamentals/24_testing/02_math
// BenchmarkAdder-4        257969908                4.54 ns/op
// PASS
// ok      github.com/AJONCODE/Golang-Fundamentals/24_testing/02_math   1.651s
