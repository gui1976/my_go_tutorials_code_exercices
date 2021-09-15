package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	height, width float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.height * s.width
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	my_circle := circle{radius: 6}
	my_square := square{height: 2, width: 3}
	info(my_circle)
	info(my_square)
}
