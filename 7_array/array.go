package main

import "fmt"

func main() {
	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"
	fruits[3] = "Grapes"

	fmt.Println("Fruit at index 0:", fruits[0])
	fmt.Println("Fruit at index 2:", fruits[2])
	fmt.Printf("Length of fruits array : %d\n", len(fruits))

	// Declare and initialize an array in a single line without shorthand
	var numbers = [4]int{1, 2, 3, 4}
	fmt.Println("Number at index 0:", numbers[0])

	// Declaring and initializing an array in one line using shorthand
	num := [5]int{10, 20, 20, 40, 50}
	fmt.Println("Num at index 0:", num[0])

	// Iterating through an array using a for loop
	fmt.Println("numbers array: ")
	var x int
	for x = 0; x < len(numbers); x++ {
		fmt.Printf("Index : %d and Value : %d\n", x, numbers[x])
	}

	// Shorter way to iterate through an array using range
	fmt.Println("Fruits array:")
	for my_index, my_fruit := range fruits {
		fmt.Printf("Fruit at index %d: %s\n", my_index, my_fruit)
	}
}
