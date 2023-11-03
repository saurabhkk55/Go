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
	var k = [9]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Print(sub_array_sum(k))
}

func sub_array_sum(arr [9]int) int {
	var ans, x, y, min int
	min = math.MinInt64

	for x = 0; x < len(arr); x++ {
		ans = 0
		for y = x; y < len(arr); y++ {
			ans = ans + arr[y]
			if ans > min {
				min = ans
			}
		}
	}
	return min
}
