package shapes

import (
	"math"
)

// interface. only method signatures
type Shape interface {
	Area() float64
}

//----------------------------------
// Rectangle
type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

//----------------------------------
// Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
