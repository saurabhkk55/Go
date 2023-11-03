package main

import "fmt"

// Address represents a structure for storing address information.
type Address struct {
	city    string
	zipCode int
}

// Person represents a structure for storing personal information.
type Person struct {
	Name string
	Age  int
	// Embedded struct without specifying field name
	Address
	// Note:
	// In Go, you can use an embedded struct without specifying the field name if the field type is the same as the struct type.
	// This is sometimes called anonymous fields.
	// Here 'Address' is the data-type, here field name will be 'Address' by-default as we did not define it explicitly.
	// Alternatively, Embedded struct with defined field name
	// personAddress Address // Here, 'personAddress' is the field name and 'Address' is the data-type
}

func main() {
	// Create three Person instances
	var p1, p2, p3 Person

	// Initialize p1 with values
	p1.Name = "Joy"
	p1.Age = 100
	p1.Address.city = "Earth"
	p1.Address.zipCode = 101010

	// Print p1's information
	fmt.Println("Person 1 Information:", p1)
	fmt.Println("Alternative way to print Person 1 information: ")
	fmt.Println("Person 1 Information:")
	fmt.Println("Name: ", p1.Name)
	fmt.Println("Age: ", p1.Age)
	fmt.Println("City: ", p1.Address.city)
	fmt.Println("ZipCode: ", p1.Address.zipCode)
	fmt.Println()

	// Initialize p2 with values using a composite literal
	p2 = Person{
		Name: "Harry",
		Age:  30,
		Address: Address{
			city:    "India",
			zipCode: 909090,
		},
	}

	// Print p2's information
	fmt.Println("Person 2 Information:", p2)
	fmt.Println("Alternative way to print Person 2 information: ")
	fmt.Println("Name: ", p2.Name)
	fmt.Println("Age: ", p2.Age)
	fmt.Println("City: ", p2.Address.city)
	fmt.Println("ZipCode: ", p2.Address.zipCode)
	fmt.Println()

	// Initialize p3 with values using a shorter composite literal
	p3 = Person{"Ajay", 12, Address{"Delhi", 808080}}

	// Print p3's information
	fmt.Println("Person 3 Information:", p3)
	fmt.Println("Alternative way to print Person 3 information: ")
	fmt.Println("Name: ", p3.Name)
	fmt.Println("Age: ", p3.Age)
	fmt.Println("City: ", p3.Address.city)
	fmt.Println("ZipCode: ", p3.Address.zipCode)
}
