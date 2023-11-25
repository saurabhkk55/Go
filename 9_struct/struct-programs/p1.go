package main

import "fmt"

// Define a struct type called Person that embeds the Address struct
type Person struct {
	Name string
	Age  int
	// Embedded struct without specifying field name
	Address // Equivalent to, Address Address
	// Note:
	// In Go, you can use an embedded struct without specifying the field name if the field type is the same as the struct type.
	// This is sometimes called anonymous fields.
	// Here 'Address' is the data-type, here field name will be 'Address' by-default as we did not define it explicitly.
	// Alternatively, Embedded struct with defined field name
	// personAddress Address // Here, 'personAddress' is the field name and 'Address' is the data-type
}

// Define a struct type called Address
type Address struct {
	City    string
	ZipCode int
}

func main() {
	// Declare multiple Person variables
	var p1, p2, p3, p4, p5, p6, p7, p8 Person

	// Initializing a Person 'p1' with field names using struct literal
	p1 = Person{
		Name: "Martin",
		Age:  25,
	}

	// Initialize p2 using a struct literal without specifying field names, or
	// Initializing a Person 'p2' without field names, in the order of struct fields
	p2 = Person{"Travis", 28, Address{}}

	// Initializing p3 by assigning values individually
	p3.Name = "Glenn"
	p3.Age = 31

	// Prompt the user to enter the name and age for the Person 'p4'.
	fmt.Print("Enter name for p4: ")
	fmt.Scan(&p4.Name)
	fmt.Print("Enter age for p4: ")
	fmt.Scan(&p4.Age)

	// Print details for each Person
	fmt.Println("\nP1 details:")
	fmt.Printf("Name: %s, Age: %d", p1.Name, p1.Age)
	fmt.Println("\n\nP2 details:")
	fmt.Printf("Name: %s, Age: %d", p2.Name, p2.Age)
	fmt.Println("\n\nP3 details:")
	fmt.Printf("Name: %s, Age: %d", p3.Name, p3.Age)
	fmt.Println("\n\nP4 details:")
	fmt.Printf("Name: %s, Age: %d", p4.Name, p4.Age)

	// Initializing p5 with field names using struct literal with nested Address struct
	p5 = Person{
		Name: "Kohli",
		Age:  33,
		Address: Address{
			City:    "Delhi",
			ZipCode: 110033,
		},
	}

	// Initializing p6 using a shorter ( without specifying field names ) struct literal
	p6 = Person{"Bumrah", 29, Address{"Mumbai", 45098}}

	// Initializing p7 by assigning values individually
	p7.Name = "Kane"
	p7.Age = 34
	p7.Address.City = "NewZealand"
	p7.Address.ZipCode = 989078

	// User input for p8
	fmt.Print("\n\nEnter name for p8: ")
	fmt.Scan(&p8.Name)
	fmt.Print("Enter age for p8: ")
	fmt.Scan(&p8.Age)
	fmt.Print("Enter city for p8: ")
	fmt.Scan(&p8.Address.City)
	fmt.Print("Enter zipcode for p8: ")
	fmt.Scan(&p8.Address.ZipCode)

	// Printing details for each person
	fmt.Println("\nP5 details:")
	fmt.Printf("Name: %s, Age: %d, City: %s, ZipCode: %d", p5.Name, p5.Age, p5.Address.City, p5.Address.ZipCode)
	fmt.Println("\n\nP6 details:")
	fmt.Printf("Name: %s, Age: %d, City: %s, ZipCode: %d", p6.Name, p6.Age, p6.Address.City, p6.Address.ZipCode)
	fmt.Println("\n\nP7 details:")
	fmt.Printf("Name: %s, Age: %d, City: %s, ZipCode: %d", p7.Name, p7.Age, p7.Address.City, p7.Address.ZipCode)
	fmt.Println("\n\nP8 details:")
	fmt.Printf("Name: %s, Age: %d, City: %s, ZipCode: %d", p8.Name, p8.Age, p8.Address.City, p8.Address.ZipCode)
}

/*
Output:

saura@DESKTOP-GC3SDTN MINGW64 ~/OneDrive/Desktop/GO (main)
$ go run ./9_struct/struct-programs/p1.go
Enter name for p4: Beckham
Enter age for p4: 43

P1 details:
Name: Martin, Age: 25

P2 details:
Name: Travis, Age: 28

P3 details:
Name: Glenn, Age: 31

P4 details:
Name: Beckham, Age: 43

Enter name for p8: Ronaldo
Enter age for p8: 37
Enter city for p8: UP
Enter zipcode for p8: 770077

P5 details:
Name: Kohli, Age: 33, City: Delhi, ZipCode: 110033

P6 details:
Name: Bumrah, Age: 29, City: Mumbai, ZipCode: 45098

P7 details:
Name: Kane, Age: 34, City: NewZealand, ZipCode: 989078

P8 details:
Name: Ronaldo, Age: 37, City: UP, ZipCode: 770077
*/
