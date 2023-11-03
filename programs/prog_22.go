package main

import "fmt"

func main() {
	var x, y, n, counter int

	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&n)

	if err != nil {
		fmt.Printf("%v. Enter a integer only.", err)
		return
	}

	fmt.Printf("Prime numbers from 1 to %d are: ", n)

	for x = 1; x <= n; x++ {
		counter = 0
		for y = 1; y <= x; y++ {
			if x%y == 0 {
				counter++
			}
		}
		if counter == 2 {
			fmt.Print(x, " ")
		}
	}
}
