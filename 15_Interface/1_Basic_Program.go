package main

import (
	"fmt"
	"math"
)

// Define an interface named "shape" with a method "area()"
type shape interface {
	area()
}

// Define a struct named "rectangle" with length and breadth as its attributes
type rectangle struct {
	length  int
	breadth int
}

// Implement the "area()" method for the "rectangle" struct
func (rec rectangle) area() {
	var rectangle_area int = rec.length * rec.breadth
	fmt.Printf("Area of rectangle = %v\n", rectangle_area)
}

// Define a struct named "circle" with radius as its attribute
type circle struct {
	radius int
}

// Implement the "area()" method for the "circle" struct
func (cir circle) area() {
	var circle_area float64 = math.Pi * math.Pow(float64(cir.radius), 2)
	fmt.Printf("Area of circle = %v\n", circle_area)
}

func main() {
	// Without using Interface
	fmt.Println("Without using Interface. Using Structs only.")
	var r rectangle
	var c circle

	// Create instances of rectangle and circle
	r = rectangle{10, 40}
	c = circle{10}

	// Call the "area()" method for circle and rectangle directly
	c.area()
	r.area()

	// Using Interface
	fmt.Println("\nUsing Interface.")
	var s1, s2 shape

	// Create instances of circle and rectangle and assign to the shape interface
	s1 = circle{10}
	s2 = rectangle{10, 40}

	// Call the "area()" method through the shape interface
	s1.area()
	s2.area()
}

/*
Output:
Without using Interface. Using Structs only.
Area of circle = 314.1592653589793
Area of rectangle = 400

Using Interface.
Area of circle = 314.1592653589793
Area of rectangle = 400
*/
