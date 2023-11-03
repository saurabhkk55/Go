// WAP to enter unit price, quantity and then calculate the total cost price after 14 discount.
package main

import "fmt"

func main() {
	// Declare variables as unsigned ( positive or >=0 ) integers only
	var unitPrice, quantity, totalCost, discount uint

	// Set the discount percentage
	discount = 14

	fmt.Print("How many items do you want? ")
	fmt.Scan(&quantity)

	fmt.Print("What is the price for only 1 item? ")
	fmt.Scan(&unitPrice)

	// Calculate the total cost before the discount
	totalCost = quantity * unitPrice

	// Calculate the discount amount as a float64
	discountAmount := float64(totalCost) * (float64(discount) / 100)

	// Print the discount amount as a float64
	fmt.Println("Applied discount =", discountAmount)

	// Convert and print the discount amount as an integer
	fmt.Println(int(discountAmount))

	// Apply the discount to the total cost
	totalCost = totalCost - uint(discountAmount)

	// Print the total cost after the discount
	fmt.Printf("Total cost after %d%% discount is %d\n", discount, totalCost)
}
