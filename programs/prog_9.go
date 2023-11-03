// WAP to calculate sum of two equal-digit numbers like this:
// 123 and 349 = (1+3)+(2+4)+(3+9) = 4+6+12 = 22

package main

import "fmt"

func main() {
	var a, b, r, r1, r2 int
	r = 0
	fmt.Println("Enter two numbers:")
	_, err := fmt.Scan(&a, &b)

	if err != nil {
		fmt.Printf("%v. Please enter integers only.", err)
		return
	}

	for x, y := a, b; x > 0 && y > 0; x, y = x/10, y/10 {
		r1 = x % 10
		r2 = y % 10
		r = r + r1 + r2
	}

	fmt.Printf("\nSum of each digits: %d", r)
}
