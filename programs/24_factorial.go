package main

import "fmt"

func main() {
	var n, ans int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	ans = factorial(n)
	fmt.Printf("Factorial of %d: %d", n, ans)
}

func factorial(x int) int {
	var r, k int
	r = 1
	for k = x; k >= 1; k-- {
		r = r * k
	}
	return r
}
