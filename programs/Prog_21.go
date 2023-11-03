package main

import "fmt"

func main() {
	var x, y, k, item int

	for x = 4; x > 0; x-- {
		for y = x; y < 4; y++ {
			fmt.Print("  ")
		}
		for k = x; k > 0; k-- {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
	//  Output
	// * * * *
	//   * * *
	//     * *
	// 	     *

	fmt.Print("\n")

	for x = 4; x > 0; x-- {
		for y = x - 1; y > 0; y-- {
			fmt.Print("  ")
		}
		for k = x; k <= 4; k++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
	//  Output
	// 	     *
	// 	   * *
	//   * * *
	// * * * *

	fmt.Print("\n")

	for x = 4; x > 0; x-- {
		for k = x; k < 4; k++ {
			fmt.Print(" ")
		}
		for y = x; y > 0; y-- {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
	//  Output
	// * * * *
	//  * * *
	//   * *
	//    *

	fmt.Print("\n")

	for x = 4; x > 0; x-- {
		for y = 1; y < x; y++ {
			fmt.Print(" ")
		}
		for k = x; k <= 4; k++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
	//  Output
	// 	  *
	//   * *
	//  * * *
	// * * * *

	fmt.Print("\n")

	item = 0
	for x = 4; x > 0; x-- {
		for k = x; k < 4; k++ {
			fmt.Print("  ")
		}
		for y = 0; y < x; y++ {
			item++
			fmt.Print(item, " ")
		}
		fmt.Print("\n")
	}
	//  Output
	// 1 2 3 4
	// 	 5 6 7
	// 	   8 9
	// 		 10

	fmt.Print("\n")

	for x = 4; x > 0; x-- {
		item = 0
		for k = x; k > 1; k-- {
			fmt.Print("  ")
		}
		for y = x; y <= 4; y++ {
			item++
			fmt.Print(item, " ")
		}
		fmt.Print("\n")
	}
	//  Output
	// 	     1
	// 	   1 2
	//   1 2 3
	// 1 2 3 4
}
