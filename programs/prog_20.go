// Patterns
package main

import "fmt"

func main() {
	var x, y, item int

	item = 0
	for x = 0; x < 4; x++ {
		for y = 0; y < 4; y++ {
			item = item + 1
			fmt.Print(item, " ")
		}
		fmt.Print("\n")
	}
	// Output
	// 1 2 3 4
	// 5 6 7 8
	// 9 10 11 12
	// 13 14 15 16

	fmt.Print("\n")

	for x = 0; x < 4; x++ {
		for y = 0; y < 4; y++ {
			item = (2 * y) + 1
			fmt.Print(item, " ")
		}
		fmt.Print("\n")
	}
	// Output
	// 1 3 5 7
	// 1 3 5 7
	// 1 3 5 7
	// 1 3 5 7

	fmt.Print("\n")

	item = 0
	for x = 4; x > 0; x-- {
		for y = 0; y < x; y++ {
			item++
			fmt.Print(item, " ")
		}
		fmt.Print("\n")
	}
	// Output
	// 1 2 3 4
	// 5 6 7
	// 8 9
	// 10

	fmt.Print("\n")

	for x = 4; x > 0; x-- {
		for y = 0; y < x; y++ {
			item = (2 * y) + 1
			fmt.Print(item, " ")
		}
		fmt.Print("\n")
	}
	// Output
	// 1 3 5 7
	// 1 3 5
	// 1 3
	// 1

	fmt.Print("\n")

	for x = 4; x > 0; x-- {
		for y = 1; y <= 4; y++ {
			item = 2 * y
			fmt.Print(item, " ")
		}
		fmt.Print("\n")
	}
	// Output
	// 2 4 6 8
	// 2 4 6 8
	// 2 4 6 8
	// 2 4 6 8

	fmt.Print("\n")

	for x = 4; x > 0; x-- {
		for y = 1; y <= x; y++ {
			item = 2 * y
			fmt.Print(item, " ")
		}
		fmt.Print("\n")
	}
	// Output
	// 2 4 6 8
	// 2 4 6
	// 2 4
	// 2
}
