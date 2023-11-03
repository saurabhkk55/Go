// The 'package main' statement declares that this Go file (hello_world.go) belongs to the 'main' package.
// In Go, the 'main' package is special, as it is the entry point for executing a program.

package main

// Import the "fmt" package, which provides functions for formatted input and output.
// We use it to print text to the standard output (usually the console).
import "fmt"

// The 'main' function is the entry point for the Go program.
// It is a special function, and every Go program must have a 'main' function.
func main() {
	// Use 'fmt.Println' to print the "Hello, World!" message to the standard output.
	fmt.Println("Hello, World!")
}
