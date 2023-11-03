## `Case 7: Struct Composition`

You can use structs to build more complex data structures by combining them, creating hierarchies, or modeling relationships.

### `Example 1:`

```go
package main

import "fmt"

// Define a struct named Person to represent a person's information.
type Person struct {
    Name string
    Age  int
}

// Define a struct named Company to represent a company's information.
type Company struct {
    Name      string
    Employees []Person
}

func main() {
    // Create an instance of the Company struct and initialize its fields.
    var company Company
    company.Name = "ABC Corp"
	
	// Short-hand
    // Create a slice of Person instances (Employees) and add data.
    // employees := make([]Person, 3) // Short-hand
	
	// Full-hand
	// Create a slice of Person instances (Employees) and add data without using :=
    var employees []Person // Declare the slice without initialization // Full-hand
    employees = make([]Person, 3) // Initialize the slice with a length of 3 // Full-hand
    
    employees[0] = Person{"John", 30}   // OR, employees[0] = Person{Name: "John", Age: 30}
    employees[1] = Person{"Alice", 25}  // OR, employees[1] = Person{Name: "Alice", Age: 25}
    employees[2] = Person{"Bob", 35}    // OR, employees[2] = Person{Name: "Bob", Age: 35)
	
	// OR
    // employees[0].Name = "John"
    // employees[0].Age = 30

    // employees[1].Name = "Alice"
    // employees[1].Age = 25

    // employees[2].Name = "Bob"
    // employees[2].Age = 35


    // Assign the slice of employees to the Employees field of the Company struct.
    company.Employees = employees

    // Access and print company and employee information.
    fmt.Printf("Company Name: %s\n", company.Name)
    fmt.Println("Employees:")
    for _, employee := range company.Employees {
        fmt.Printf("Name: %s, Age: %d\n", employee.Name, employee.Age)
    }
}
```

### `Output:`

```
Company Name: ABC Corp
Employees:
Name: John, Age: 30
Name: Alice, Age: 25
Name: Bob, Age: 35
```

The output remains the same as in the previous version of the code, displaying the company name and information about its three employees.

### `Example 2:`

```go
package main

import "fmt"

// Define a struct named Person to represent a person's information.
type Person struct {
    Name string
    Age  int
}

// Define a struct named Company to represent a company's information.
type Company struct {
    Name      string
    Employees []Person
}

func main() {
    // Create an instance of the Company struct and initialize its fields.
    var company Company
    company.Name = "ABC Corp"
    
	// Short-hand
	// Create a slice of Person instances (Employees) and add data.
    // employees := []Person{
    //     {Name: "John", Age: 30},
    //     {Name: "Alice", Age: 25},
    //     {Name: "Bob", Age: 35},
    // }
	
	// Full-hand
    // Create a slice of Person instances (Employees) and add data.
    var employees []Person // Declare the slice without initialization
    employees = []Person{
        {"John", 30},
        {"Alice", 25},
        {"Bob", 35},
    }
    
    // Assign the slice of employees to the Employees field of the Company struct.
    company.Employees = employees

    // Access and print company and employee information.
    fmt.Printf("Company Name: %s\n", company.Name)
    fmt.Println("Employees:")
    for _, employee := range company.Employees {
        fmt.Printf("Name: %s, Age: %d\n", employee.Name, employee.Age)
    }
}

// Alternative Code

package main

import "fmt"

// Define a struct named Person to represent a person's information.
type Person struct {
    Name string
    Age  int
}

// Define a struct named Company to represent a company's information.
type Company struct {
    Name      string
    Employees []Person
}

func main() {
    // Create an instance of the Company struct and initialize its fields.
    var company Company
    company.Name = "ABC Corp"
    
    // Create a slice of Person instances (Employees) and add data.
    var employees []Person

    var employee1 Person
    employee1.Name = "John"
    employee1.Age = 30
    employees = append(employees, employee1)

    var employee2 Person
    employee2.Name = "Alice"
    employee2.Age = 25
    employees = append(employees, employee2)

    var employee3 Person
    employee3.Name = "Bob"
    employee3.Age = 35
    employees = append(employees, employee3)
    
    // Assign the slice of employees to the Employees field of the Company struct.
    company.Employees = employees

    // Access and print company and employee information.
    fmt.Printf("Company Name: %s\n", company.Name)
    fmt.Println("Employees:")
    for _, employee := range company.Employees {
        fmt.Printf("Name: %s, Age: %d\n", employee.Name, employee.Age)
    }
}
```

### `Example 3:`
```go
package main

import "fmt"

// Define a struct named Person to represent a person's information.
type Person struct {
    Name string
    Age  int
}

// Define a struct named Company to represent a company's information.
type Company struct {
    Name      string
    Employees []Person
}

func main() {
    // Create an instance of the Company struct and initialize its fields.
    var company Company
    company.Name = "ABC Corp"
    
	// Short-hand
    // Create a slice of Person instances (Employees) and add data.
    // employees := make([]Person, 3) // Short-hand
	
	// Full-hand
	// Create a slice of Person instances (Employees) and add data without using :=
    var employees []Person // Declare the slice without initialization // Full-hand
    employees = make([]Person, 3) // Initialize the slice with a length of 3 // Full-hand
	
    employees[0].Name = "John"
    employees[0].Age = 30

    employees[1].Name = "Alice"
    employees[1].Age = 25

    employees[2].Name = "Bob"
    employees[2].Age = 35
    
    // Assign the slice of employees to the Employees field of the Company struct.
    company.Employees = employees

    // Access and print company and employee information.
    fmt.Printf("Company Name: %s\n", company.Name)
    fmt.Println("Employees:")
    for _, employee := range company.Employees {
        fmt.Printf("Name: %s, Age: %d\n", employee.Name, employee.Age)
    }
}
```
