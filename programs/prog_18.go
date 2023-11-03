// 1  -3  6  -10  15 ... n
package main

import "fmt"

func main() {
	var x, n, postive, negative int

	n = 20
	postive = 0

	for x = 1; x <= n; x++ {
		postive = postive + x

		if x%2 == 0 {
			negative = postive * -1
			fmt.Print(negative, " ")
		} else {
			fmt.Print(postive, " ")
		}
	}
}
