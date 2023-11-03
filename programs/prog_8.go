// WAP to reverse a number

package main

import "fmt"

func main() {

	var n, x, d, r int

	r = 0

	fmt.Print("Eneter number: ")
	_, err := fmt.Scan(&n)

	if err != nil {
		fmt.Printf("\nError: %v. Please enter an integer.", err)
	}

	for x = n; x > 0; x = x / 10 {
		d = x % 10
		r = r*10 + d
	}

	fmt.Printf("Reversed number: %d", r)
}
