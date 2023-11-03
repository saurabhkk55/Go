// WAP to swap 2 number
package main

import "fmt"

func main() {

	var a, b, t int

	fmt.Println("Enter 2 number:")
	fmt.Scan(&a, &b)

	fmt.Printf("Before Swapping number_1 = %d and number_2 = %d\n", a, b)

	t = a
	a = b
	b = t

	fmt.Printf("After  Swapping number_1 = %d and number_2 = %d", a, b)
}
