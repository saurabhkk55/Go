// Rotate an Array:
// Implement a Go function that rotates an array of integers to the right by a given number of positions.
// For example, if the input array is [1, 2, 3, 4, 5] and the number of positions is 2, the function should modify the array to [4, 5, 1, 2, 3].// rotates an array of integers to the right by a given number of positions
package main

import "fmt"

func main() {
	var r int
	r = 2

	var k = [5]int{1, 2, 3, 4, 5}
	fmt.Print(rotate(k, r))
}

func rotate(arr [5]int, rotate int) [5]int {
	var index, counter, size int
	size = len(arr)

	var brr [5]int

	for index = 0; index < len(arr); index++ {
		counter = index + rotate
		if counter >= len(arr) {
			counter = counter - size
			brr[counter] = arr[index]
		} else {
			brr[counter] = arr[index]
		}
	}
	return brr // returning an array
}
