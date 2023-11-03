// Create a Go program that defines a Rectangle struct with fields Width and Height.
// Implement methods for the Rectangle struct to calculate its area and perimeter.
// Then, create instances of the Rectangle struct and demonstrate how to use these methods.

package main

import (
	"fmt"
)

// Define a struct for a rectangle
type rectangle struct {
	length  int
	breadth int
}

// Method to calculate the area of the rectangle
func (r rectangle) area() int {
	return r.length * r.breadth
}

// Method to calculate the perimeter of the rectangle
func (r rectangle) perimeter() int {
	return 2 * (r.length + r.breadth)
}

func main() {
	// Create two rectangle instances
	var r1, r2 rectangle
	var areaAns, perimeterAns int

	// Initialize r1 with specific dimensions
	r1 = rectangle{10, 20}
	areaAns = r1.area()
	perimeterAns = r1.perimeter()

	// Print the area and perimeter of r1
	fmt.Printf("Rectangle r1:\n")
	fmt.Printf("Area: %d\n", areaAns)
	fmt.Printf("Perimeter: %d\n", perimeterAns)

	// Calculate and print the area and perimeter of r2
	calculateAndPrintAreaPerimeter(&r2)
}

// Function to calculate and print the area and perimeter of a given rectangle
func calculateAndPrintAreaPerimeter(rect *rectangle) {
	var len, bred int

	fmt.Println("Enter length for the new rectangle:")
	fmt.Scan(&len)
	fmt.Println("Enter breadth for the new rectangle:")
	fmt.Scan(&bred)

	rect.length = len
	rect.breadth = bred

	area := rect.area()
	perimeter := rect.perimeter()

	fmt.Printf("Rectangle r2:\n")
	fmt.Printf("Area: %d\n", area)
	fmt.Printf("Perimeter: %d\n", perimeter)
}
