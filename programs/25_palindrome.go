package main

import "fmt"

func main() {
	numPal()
	strPal()
}

func numPal() {
	var n, x, r, ans int

	ans = 0

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	for x = n; x > 0; x = x / 10 {
		r = x % 10
		ans = ans*10 + r
	}

	fmt.Printf("Entered  number : %d", n)
	fmt.Printf("\nReversed number : %d", ans)

	if ans == n {
		fmt.Printf("\n%d : Palindrome integer", n)
	} else {
		fmt.Printf("\n%d : Not Palindrome integer", n)
	}
}

func strPal() {
	var str, ans string
	var x int

	fmt.Print("\n\nEnter a string: ")
	fmt.Scan(&str)

	for x = len(str) - 1; x >= 0; x-- {
		ans = ans + string(str[x])
	}

	fmt.Printf("Entered  string : %s", str)
	fmt.Printf("\nReversed string : %s", ans)

	if ans == str {
		fmt.Printf("\n%s : Palindrome string", str)
	} else {
		fmt.Printf("\n%s : Not Palindrome string", str)
	}
}
