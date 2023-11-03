// WAP to enter unit price, quantity and then calculate the total cost price.
package main

import "fmt"

func main() {

	var unitPrice, quantity, totalCost int

	fmt.Print("How many items/quantity do you want? ")
	fmt.Scan(&quantity)

	fmt.Print("\nWhat is the price for only 1 item? ")
	fmt.Scan(&unitPrice)

	totalCost = quantity * unitPrice

	fmt.Printf("\nTotal cost: %d", totalCost)
}
