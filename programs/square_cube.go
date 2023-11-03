// Program to find Square and Cube of a number
package main

import "fmt"

func main() {

	var num, cube, square int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	square = num * num
	cube = num * num * num

	fmt.Printf("Type of %v and %v are %T and %T respectively\n", num, cube, num, cube)
	fmt.Println("square of", num, "is", square)
	fmt.Printf("Cube of %v is %v", num, cube)
}
