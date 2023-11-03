// WAP to calculate sum of marks obtained in 5 different subjects and also percentage
package main

import "fmt"

func main() {

	var eng, hindi, math, geo, science, totalMarks int
	var percentage float64

	fmt.Println("Enter marks obatins in eng, hindi, math, geo, science:")
	_, err := fmt.Scan(&eng, &hindi, &math, &geo, &science)

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Please enter positive integer only")

		return
	}

	fmt.Println("English:", eng)
	fmt.Println("Hindi:", hindi)
	fmt.Println("Math:", math)
	fmt.Println("Geography:", geo)
	fmt.Println("Science:", science)

	totalMarks = eng + hindi + math + geo + science
	fmt.Println("Your total marks out of 500:", totalMarks)

	percentage = ((float64(totalMarks) / 500) * 100)
	fmt.Println("Your percentage:", percentage)
}
