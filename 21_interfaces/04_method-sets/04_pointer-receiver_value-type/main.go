package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area: ", s.area())
}

func main() {
	c := circle{7}

	info(c)
}

// # command-line-arguments
// ./main.go:27:6: cannot use c (type circle) as type shape in argument to info:
//         circle does not implement shape (area method has pointer receiver)
