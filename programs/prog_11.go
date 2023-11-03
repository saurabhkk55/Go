// WAP to check a year is leap year or not
package main

import "fmt"

func main() {
	var year int

	fmt.Print("Enter a year: ")
	_, err := fmt.Scan(&year)

	if err != nil {
		fmt.Printf("%v. Please enter a valid year.", err)
		return
	}

	if ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0) {
		fmt.Printf("%d is a leap year.", year)
	} else {
		fmt.Printf("%d is not a leap year.", year)
	}
}
