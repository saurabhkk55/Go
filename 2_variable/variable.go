package main

import "fmt"

// In Go, when declaring global variables, you must initialize them at the time of declaration.
var a int = 5 // Declaring and initializing a global variable 'a' with the value 5
var b int = 5 // Declaring and initializing a global variable 'b' with the value 5

// we can just Declare a global variable, we have to initialize it as well at the same of declartion
// var b int     // Declares a global variable
// b = 5 		 // Initializes a global variable 'b'

func main() {
	a = 10         // Update the value of the global variable a
	var c int = 20 // Initializes a local variable at the time of declaration
	d := 40        // short-hand declaration syntax can only be used inside functions.
	var e int      // Declares a local variable 'e'
	e = 50         // Initializes the same local variable 'e'

	// Print the values of variables
	fmt.Println("Variable a:", a)
	fmt.Println("Variable b:", b)
	fmt.Println("Variable c:", c)
	fmt.Println("Variable d:", d)
	fmt.Println("Variable e:", e)
}
