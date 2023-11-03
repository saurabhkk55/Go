// Remove Duplicates from an Array:
// Write a Go function that removes duplicates from an array of integers.
// After running the function, the array should contain only unique elements in the same order as they appear in the original array.
package main

import "fmt"

func main() {
	// var k = [9]int{1, 2, 3, 2, 4, 5, 6, 5, 7}
	// var k = [7]int{3, 5, 5, 8, 8, 3, 7}
	var k = [10]int{9, 2, 2, 2, 6, 6, 9, 4, 1, 3}

	fmt.Print(removeDuplicate(k))
}

func removeDuplicate(arr [10]int) []int {

	var flag int

	var brr []int
	brr = append(brr, arr[0])

	for x := 1; x < len(arr); x++ {
		flag = 0
		for y := 0; y < len(brr); y++ {
			if arr[x] == brr[y] {
				flag = 1
			}
		}
		if flag == 0 {
			brr = append(brr, arr[x])
		}
	}
	return brr // returning a slice
}
