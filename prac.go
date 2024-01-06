package main

import (
	"fmt"
	"math"
)

func main() {
	var str string

	str = "saurabh"
	str = str + "k"

	// x := str[0]
	// y := string(str[0])
	// z := len(str) - 1

	// fmt.Println("s: ", string(str[0]))
	// fmt.Println("x: ", x)
	// fmt.Println("y: ", y)
	// fmt.Println("z: ", string(str[z]))
	// fmt.Println("str: ", str)

	// for i := 0; i < len(str); i++ {
	// 	fmt.Print(string(str[i]), "=")
	// }

	// var p string
	// var n int
	// n = 5

	// var myMap map[string]int
	// myMap = map[string]int{"I": 1, "V": 2, "X": 3, "L": 50, "C": 100, "D": 500, "M": 1000}
	// s := "LVIII"
	// var i, v, p int
	// p = 0
	// for i = 0; i < len(s); i++ {
	// 	v = myMap[string(s[i])]
	// 	p = p + v
	// 	fmt.Println("V= ", v)
	// 	fmt.Println("P= ", p)
	// }
	// fmt.Println("P= ", p)

	s1 := "saurabh"
	s2 := "Lucky"

	n := math.Min(float64(len(s1)), float64(len(s2)))
	// n := math.Min(len(s1), len(s2))
	n = 0
	n = n - 1
	fmt.Println(n)
}
