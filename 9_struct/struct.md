
# `Struct`
- A struct in Go is a user-defined type that allows you to group together related data fields of different data-types into a single unit. 
- Structs are useful for representing real-world entities, such as a person, a product, or a customer order.

**`Syntax`**

To declare a struct, you use the `type` keyword followed by the name of the struct and a list of fields enclosed in curly braces. Each field has a name and a type.

```go
type StructName struct {
    Field1 Type1
    Field2 Type2
    // Add more fields as needed
}
```

- `StructName` is the name of the struct type.
- `Field1`, `Field2`, and so on are the names of the struct's fields.
- `Type1`, `Type2`, and so on are the types of data that each field can hold.

**`Example:`**
```go
type Person struct {
  Name string
  Age int
}
```

This code declares a struct called Person with two fields: Name of type string and Age of type int.

## `Case 1: Basic Structs`

Defining a simple struct named `Person`:

```go
package main

import "fmt"

// Define a struct type "Person" to represent a person with a name and age.
type Person struct {
	Name string
	Age  int
}

func main() {
	// Declare two variables of type "Person" to store information about two individuals.
	var p1, p2 Person

	// Set values for the fields of the first person (p1).
	p1.Name = "Joy"
	p1.Age = 100

	// Print the information of the first person (p1).
	fmt.Printf("Name: %s, Age: %d\n", p1.Name, p1.Age)

	// Prompt the user to enter the name and age for the second person (p2).
	fmt.Print("Enter Name for p2: ")
	fmt.Scan(&p2.Name)
	fmt.Print("Enter Age for p2: ")
	fmt.Scan(&p2.Age)

	// Print the information of the second person (p2).
	fmt.Printf("Name: %s, Age: %d\n", p2.Name, p2.Age)
}
// Output
// Name: Joy, Age: 100
// Enter Name for p2: John
// Enter Age for p2: 24
// Name: John, Age: 24
```

## `Case 2: Struct Initialization`

Initializing a struct `Person` with and without field names:
- You can initialize a struct with field names and values.
- You can also omit the field names if the values are provided in the order of the struct fields.

```go
package main

import "fmt"

// Define a struct named Person
type Person struct {
    Name string
    Age  int
}

func main() {
    // Initializing a Person with field names
    var person1 Person
    person1 = Person{
        Name: "Alice",
        Age:  30,
    }

    // Initializing a Person without field names, in the order of struct fields
    var person2 Person
    person2 = Person{"Bob", 25}

    // Accessing and printing the values of the fields
    fmt.Printf("Person 1 - Name: %s, Age: %d\n", person1.Name, person1.Age)
    fmt.Printf("Person 2 - Name: %s, Age: %d\n", person2.Name, person2.Age)
}
```

## `Case 3: Embedded Structs`

You can embed one struct into another to create more complex data structures. This is a form of composition.

Defining embedded structs, `Address` and `Contact`:

```go
package main

import "fmt"

// Define the Address struct to represent a person's address
type Address struct {
    Street  string
    City    string
    ZipCode string
}

// Define the Person struct to represent basic personal information
type Person struct {
    Name string // Field to store the person's name (string)
    Age  int    // Field to store the person's age (int)
}

// Define the Contact struct, which embeds both Person and Address
type Contact struct {
    Person  // Embedded Person struct for personal information
    Address // Embedded Address struct for contact address
    Email   string // Field to store the person's email address (string)
}

func main() {
    // Create a Contact instance
    var contact Contact
    contact.Name = "John"   // Set the Name field of the embedded Person struct
    contact.Age = 30    // Set the Age field of the embedded Person struct
    contact.Email = "john@example.com"  // Set the Email field of the Contact struct
    contact.Address.Street = "123 Main St"  // Set the Street field of the embedded Address struct
    contact.Address.City = "Anytown"    // Set the City field of the embedded Address struct
    contact.Address.ZipCode = "12345"   // Set the ZipCode field of the embedded Address struct

    // Access and print the values of the embedded structs and Contact fields
    fmt.Printf("Name: %s, Age: %d\n", contact.Name, contact.Age)
    fmt.Printf("Email: %s\n", contact.Email)
    fmt.Printf("Address: %s, %s, %s\n", contact.Address.Street, contact.Address.City, contact.Address.ZipCode)
}
```
