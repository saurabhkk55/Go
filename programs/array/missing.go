// Find Missing Number in Array:
// Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, write a Go function to find the missing number.
// For example, if the input array is [3, 0, 1], the function should return 2.
package main

import (
	"fmt"
	"sort"
)

func main() {
	var k = [3]int{3, 0, 1}
	fmt.Print(missing(k))
}

func missing(arr [3]int) int {

	// Convert the array to a slice
	slice := arr[:]

	// Sort the slice in ascending order
	sort.Ints(slice)

	// Copy the sorted elements back into the array
	copy(arr[:], slice)

	fmt.Println("Sorted array (ascending):", arr)
	fmt.Printf("Sorted array type: %T\n", arr)

	var x, counter int

	counter = 0
	for x = 0; x < len(arr); x++ {
		if arr[x] != counter {
			break
		}
		counter++
	}
	return counter // returning an integer
}
