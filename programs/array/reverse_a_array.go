// Reverse an Array:
// Write a Go function that takes an array of integers and reverses the elements in-place.
// For example, if the input array is [1, 2, 3, 4, 5], the function should modify the array to [5, 4, 3, 2, 1].
package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}

	var brr [5]int
	var x, counter int

	for x = len(arr) - 1; x >= 0; x-- {
		brr[counter] = arr[x]
		counter++
	}

	fmt.Print("Reversed Error : ", brr)
}
