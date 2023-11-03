// WAP to swap 2 number without using 3rd variable
package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Enter 2 number:")
	fmt.Scan(&a, &b)

	fmt.Printf("Before swapping a = %d and b = %d\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("After  swapping a = %d and b = %d", a, b)
}
