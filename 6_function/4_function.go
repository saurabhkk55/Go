package main

import "fmt"

func factorial(n int) int {
	// Base case: If n is 0 or 1, the factorial is 1.
	if n == 0 || n == 1 {
		return 1
	}

	// Recursive case: Calculate n * factorial of (n-1).
	return n * factorial(n-1)
}

func main() {
	var n, result int
	n = 5
	result = factorial(n)
	fmt.Printf("The factorial of %d is %d\n", n, result)
}
