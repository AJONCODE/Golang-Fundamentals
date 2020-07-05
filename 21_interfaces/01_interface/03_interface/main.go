package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

// implements the shape interface
func (s square) area() float64 {
	return s.side * s.side
}

// implements the shape interface
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{
		side: 5,
	}

	c := circle{
		radius: 7,
	}

	fmt.Println("Square: ")
	info(s)
	fmt.Println("Circle: ")
	info(c)
}

// Square:
// {5}
// 25
// Circle:
// {7}
// 153.93804002589985
