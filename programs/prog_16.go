// 1	3	6	10	15	...	n
package main

import "fmt"

func main() {
	var n int = 20
	var ans, x int
	ans = 0

	for x = 1; x <= n; x++ {
		ans = ans + x
		fmt.Print(ans, " ")
	}

}
