package structs

import "math"

// Rectangle describes a rectangle
type Rectangle struct {
	width  float64
	height float64
}

// Circle describes a circle
type Circle struct {
	radius float64
}

// Triangle describes a Triangle
type Triangle struct {
	base   float64
	height float64
}

// Interface for a shape
type Shape interface {
	area() float64
}

// Perimeter returns the perimeter of a retangle
func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.height)
}

// Area returns the area of a rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Area returns the area of a circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Area returns the area of a triangle
func (t Triangle) area() float64 {
	return (t.base * t.height) * 0.5
}
