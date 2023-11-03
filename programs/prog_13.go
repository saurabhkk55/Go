package main

import "fmt"

func main() {
	var my_var int

	my_var = int('a')
	fmt.Printf("ASCII value of 'a' = %d\n", my_var)

	// Subtract 32 from the ASCII value of 'a' to get the uppercase 'A'.
	my_var = int('a') - 32
	fmt.Printf("ASCII value of 'a' after subtracting 32 = %d\n", my_var)

	// Convert the ASCII value back into a character.
	var char string
	char = string(my_var)
	fmt.Printf("Character with ASCII value %d = %s\n", my_var, char)

	fmt.Println("###########################################################")

	// Declare a string 's' with the value "saurabh".
	var s string
	s = "saurabh"

	// Get the ASCII value of the first character 's' (index 0).
	my_var = int(s[0])
	fmt.Printf("ASCII value of the first character in 's' = %d\n", my_var)

	// Subtract 32 from the ASCII value to get the uppercase character.
	my_var = int(s[0]) - 32
	char = string(my_var)
	fmt.Printf("Character with ASCII value %d (after subtracting 32) = %s\n", my_var, char)
}
