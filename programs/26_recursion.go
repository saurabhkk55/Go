// Write a Go function SumOfNaturals that calculates
//
//	the sum of the first n natural numbers using recursion.
package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter a positive integer: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Print("No natural numbers to sum")
	} else {
		fmt.Print(addNatural(n))
	}

}

func addNatural(ans int) int {

	// Base case: If ans is 1
	if ans == 1 {
		return 1
	}

	// Recursive case: ans + addNatural(ans-1)
	ans = ans + addNatural(ans-1)
	return ans
}
