package math

// Adder adds the numbers
func Adder(xs ...int) int {
	res := 0
	for _, v := range xs {
		res += v
	}
	return res
}
