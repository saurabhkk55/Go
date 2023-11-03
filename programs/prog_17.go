// 1	4	7	10	13	... n
package main

import "fmt"

func main() {
	var n, x, ans int

	n = 20
	ans = 0

	for x = 0; x < n; x++ {
		ans = 3*x + 1
		fmt.Print(ans, " ")
	}
}
