// WAP to convert height from feet to integer
package main

import "fmt"

func main() {
	var heightInFeet, heightInInch int
	var heightInCm float64

	fmt.Print("Enter your height in feet (positive integer): ")
	_, err := fmt.Scan(&heightInFeet)

	if err != nil || heightInFeet <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer for height in feet.")
		return
	}

	heightInInch = heightInFeet * 12
	heightInCm = float64(heightInInch) * 2.54

	fmt.Printf("Your height in inch: %d\n", heightInInch)
	fmt.Printf("Your height in cm: %f", heightInCm)
}
