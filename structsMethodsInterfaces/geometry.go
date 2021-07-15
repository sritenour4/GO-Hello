package structsMethodsInterfaces

import "math"

// Shape is implemented by anything that can tell us its Area.
type Shape interface {
	Area() float64
}

// Rectangle has the dimensions of a rectangle.
type Rectangle struct {
	Width float64
	Height float64
}

// Triangle has the dimensions of a triangle.
type Triangle struct {
	Base float64
	Height float64
}

// Circle represents a circle.
type Circle struct {
	Radius float64
}

// Area returns the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area returns the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area returns the area of the triangle.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter returns the perimeter of a rectangle.
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}


