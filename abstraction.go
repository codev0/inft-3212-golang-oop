package main

import (
	"fmt"
	"math"
)

// Shape is an interface that abstracts different geometric shapes.
// It declares an Area method that any concrete type must implement.
type Shape interface {
	Area() float64
}

// Circle struct represents a circle with a radius.
type Circle struct {
	Radius float64
}

// Rectangle struct represents a rectangle with width and height.
type Rectangle struct {
	Width, Height float64
}

// Area method for Circle to satisfy the Shape interface.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area method for Rectangle to satisfy the Shape interface.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// PrintArea takes a Shape interface as its parameter and prints its area,
// demonstrating how an interface can be used to abstract the details of the shape.
// %0.2f default width, precision 2 for float64
func PrintArea(s Shape) {
	fmt.Printf("The area of the shape is %0.2f\n", s.Area())
}

func abstraction() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}

	PrintArea(circle)
	PrintArea(rectangle)
}
