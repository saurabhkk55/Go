package main

import "fmt"

// Define a struct type "Person" to represent a person with a name and age.
type Person struct {
	Name string
	Age  int
}

func main() {
	// Declare four variables of type "Person" to store information about individuals.
	var p1, p2, p3, p4 Person

	// Set values for the fields of the first person (p1).
	p1.Name = "Joy"
	p1.Age = 100

	// Print the information of the first person (p1).
	fmt.Printf("Name: %s, Age: %d\n\n", p1.Name, p1.Age)

	/////////////////////////////////////////////////////////////

	// Prompt the user to enter the name and age for the second person (p2).
	fmt.Print("Enter Name for p2: ")
	fmt.Scan(&p2.Name)
	fmt.Print("Enter Age for p2: ")
	fmt.Scan(&p2.Age)

	// Print the information of the second person (p2).
	fmt.Printf("Name: %s, Age: %d\n\n", p2.Name, p2.Age)

	/////////////////////////////////////////////////////////////

	// Initialize the third person (p3) using a struct literal.
	p3 = Person{
		Name: "Go",
		Age:  2009,
	}

	// Print the information of the third person (p3).
	fmt.Printf("Name: %s, Age: %d\n", p3.Name, p3.Age)
	fmt.Println("p3 info: ", p3)
	fmt.Println()

	/////////////////////////////////////////////////////////////

	// Initialize the fourth person (p4) using a struct literal without field names.
	p4 = Person{"Hero", 1}

	// Print the information of the fourth person (p4).
	fmt.Printf("Name: %s, Age: %d\n", p4.Name, p4.Age)
	fmt.Print("p4 info: ", p4)
}
