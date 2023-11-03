// Find the Maximum Subarray Sum:
// Given an array of integers, write a Go function to find the maximum sum of a contiguous subarray within the array.
// Your function should return the maximum sum. For example, for the input array [−2, 1, −3, 4, −1, 2, 1, −5, 4], the function should return 6,
// which corresponds to the subarray [4, -1, 2, 1].
package main

import (
	"fmt"
	"math"
)

func main() {
	var size int
	var k []int

	// Ask the user for the size of the slice.
	fmt.Print("Size: ")
	fmt.Scan(&size)

	// Create a slice with the specified size.
	k = make([]int, size)

	fmt.Println("Enter elements:")

	// Prompt the user to enter elements for the slice.
	for n := 0; n < size; n++ {
		fmt.Printf("Enter element at index %d: ", n)
		fmt.Scan(&k[n])
	}

	// Calculate the maximum subarray sum using the subArraySum function.
	result := subArraySum(k)

	// Display the maximum subarray sum.
	fmt.Printf("Maximum subarray sum: %d\n", result)
}

// subArraySum calculates the maximum subarray sum using the Kadane's algorithm.
func subArraySum(arr []int) int {
	var ans, min int
	min = math.MinInt64

	for x := 0; x < len(arr); x++ {
		ans = 0
		for y := x; y < len(arr); y++ {
			ans = ans + arr[y]
			if ans > min {
				min = ans
			}
		}
	}
	return min // returning an integer
}
