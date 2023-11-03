// WAP to arrange 3 numbers in asscending order

package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("Enter 3 number:")
	_, err := fmt.Scan(&a, &b, &c)

	if err != nil {
		fmt.Printf("%v. Please enter positive integers only.", err)
		return
	}

	fmt.Printf("Before Sorting: a = %d b = %d c = %d\n", a, b, c)
	fmt.Print("After sorting in asscending order: ")

	if a > b {
		if b > c {
			fmt.Printf("a, b, c: %d %d %d", a, b, c)
			return
		} else {
			if c > a {
				fmt.Printf("c, a, b: %d %d %d", c, a, b)
				return
			}
			fmt.Printf("a, c, b: %d %d %d", a, c, b)
			return
		}
	} else {
		if a > c {
			fmt.Printf("b, a, c: %d %d %d", b, a, c)
			return
		} else {
			if c > b {
				fmt.Printf("c, b, a: %d %d %d", c, b, a)
				return
			} else {
				fmt.Printf("b, c, a: %d %d %d", b, c, a)
				return
			}
		}
	}
}
