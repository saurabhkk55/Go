// WAP to get all the even number from a particular range
package main

import "fmt"

func main() {
	var a, b, x int

	fmt.Println("Enter 2 numbers:")
	fmt.Scan(&a, &b)

	if a%2 == 0 {
		a = a
	} else {
		a = a + 1
	}

	fmt.Println("Even numbers are:")

	for x = a; x <= b; x = x + 2 {
		fmt.Print(x, " ")
	}
}
