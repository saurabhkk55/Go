// WAP to check what type of character a user has eneterned
package main

import "fmt"

func main() {
	var my_char string
	var char byte

	fmt.Print("Enter a character: ")
	fmt.Scan(&my_char)

	// Check if the user entered more than one character
	if len(my_char) != 1 {
		fmt.Println("Please enter a single character.")
		return
	}

	// Assigning byte data type (my_char[0]) to another byte data type (char)
	char = my_char[0]

	// comparining byte data type (item) with byte data
	if char >= 'a' && char <= 'z' {
		fmt.Printf("%c: Small alphabet", char)
	} else if char >= 'A' && char <= 'Z' {
		fmt.Printf("%c: Capital alphabet", char)
	} else if char >= '0' && char <= '9' {
		fmt.Printf("%c: Digit", char)
	} else {
		fmt.Printf("%c: Special character", char)
	}
}
