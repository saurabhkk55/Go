package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	fibonacci(0, 1, n)
}

func fibonacci(p int, q int, m int) {
	var t, x int

	if m == 0 {
		fmt.Print("N/A")
	} else if m == 1 {
		fmt.Print("0 ")
	} else if m == 2 {
		fmt.Print("0 1 ")
	} else {
		fmt.Print("0 1 ")
		for x = 3; x <= m; x++ {
			t = p + q
			p = q
			q = t
			fmt.Print(t, " ")
		}
	}
}
