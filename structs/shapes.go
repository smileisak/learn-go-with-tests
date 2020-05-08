package structs

import "math"

// Shape interface to abstract a Shape
type Shape interface {
	Area() float64
}

// Rectangle is a type to hold a rectangle attributes
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of a rectancle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Circle is a type to hold a circle attributes
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle is a type to hold a circle attributes
type Triangle struct {
	Base   float64
	Height float64
}

// Area calculates the area of a Triangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

// Perimeter calculates a perimeter of a Rectagle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
