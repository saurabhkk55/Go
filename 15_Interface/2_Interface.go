package main

import (
	"fmt"
	"math"
)

// Define a struct named "Circle" with Radius as its attributes
type Circle struct {
	Radius float64
}

// Implement the "area()" method for the "Circle" struct
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Define a struct named "Square" with Width and Height as its attributes
type Square struct {
	Width  float64
	Height float64
}

// Implement the "area()" method for the "Square" struct
func (s Square) Area() float64 {
	return s.Width * s.Height
}

// Sizer is an interface that defines the common method Area().
type Sizer interface {
	Area() float64
}

func main() {
	// Create instances of Circle and Square
	var c Circle = Circle{Radius: 10}
	var s Square = Square{Height: 10, Width: 5}

	// Determine the smaller shape using the Less function
	l := Less(c, s)
	fmt.Printf("%+v is the smallest\n", l)
}

// Less compares two shapes based on their areas and returns the smaller one.
func Less(s1, s2 Sizer) Sizer {
	// Use the Area() method from the Sizer interface to compare areas
	if s1.Area() < s2.Area() {
		return s1
	}
	return s2
}

// Output:
// {Width:5 Height:10} is the smallest
