package main

import (
	"fmt"
	"math"
)

// interface. only method signatures
type shape interface {
	area() float64
}

// function that receives a shape
func printArea(s shape) {
	fmt.Printf("Area of shape is %.2f\n", s.area())
}

//----------------------------------
// Rectangle
type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

//----------------------------------
// Circle
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// interface implementation is implicit in Go, you don't need to specify it.
// Just have a method with the same signature
func main() {
	r := rectangle{10, 20}
	printArea(r)

	c := circle{10}
	printArea(c)
}
