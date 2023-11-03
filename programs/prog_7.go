// WAP to covert age in days to age in year,months and days

package main

import "fmt"

func main() {
	// var ageInDays, ageYears, ageMonths, ageDays uint // age can't be negative

	var ageInDays, a, b, c, d uint

	fmt.Print("Enter your age in days: ")
	_, err := fmt.Scan(&ageInDays)

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Please enter positive interger value only")
	}

	if ageInDays >= 365 {
		a = ageInDays / 365
		fmt.Printf("Age in years: %d\n", a)
	}

	b = ageInDays % 365

	if b >= 31 {
		c = b / 31
		fmt.Printf("Age in months: %d\n", c)

	}

	if (b > 0 && b < 31) || (c > 0) {
		d = b % 31
		fmt.Printf("Age in days: %d", d)
	}
}
