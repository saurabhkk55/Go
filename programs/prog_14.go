// WAP to convert lowerCase into UpperCase and vice, print ASCII of digit and special character
package main

import "fmt"

func main() {
	var str, convert string
	var x int
	var item byte

	fmt.Print("Enter a character: ")
	fmt.Scan(&str)

	// Check if the user entered more than one character
	if len(str) != 1 {
		fmt.Println("Please enter a single character.")
		return
	}

	// Assigning byte data type (str[0]) to another byte data type (item)
	item = str[0] // 1

	// Converting byte into int
	// to get the ASCII value of 'item' variable
	x = int(item) // 1
	// OR
	// converting byte into int
	// to get the ASCII value of 'item' variable
	// x = int(str[0]) // 2

	// comparining byte data type (item) with byte data
	if item >= 'a' && item <= 'z' {
		x = x - 32
		// converting int 'x' into string 'convert' but vice-versa is not possible
		convert = string(x)
		fmt.Printf("ASCII value: %d and Corresponding character: %s", x, convert)
	} else if item >= 'A' && item <= 'Z' {
		x = x + 32
		// converting int 'x' into string 'convert' but vice-versa is not possible
		convert = string(x)
		fmt.Printf("ASCII value: %d and Corresponding character: %s", x, convert)
	} else if item >= '0' && item <= '9' {
		// converting int 'x' into string 'convert' but vice-versa is not possible
		convert = string(x)
		fmt.Printf("ASCII value of character %s is %d", convert, x)
	} else {
		// converting int 'x' into string 'convert' but vice-versa is not possible
		convert = string(x)
		fmt.Printf("ASCII value of character %s is %d", convert, x)
	}
}
