## `Case 6: Struct Initialization with new vs. Struct Literal`

- You can initialize a struct using the `new` function to create a pointer to a zeroed struct.
- Alternatively, you can use a struct literal to create and initialize a struct in a more common and idiomatic way.

```go
package main

import "fmt"

// Define a struct named Person to represent a person's information.
type Person struct {
    Name string
    Age  int
}

func main() {
    // Initialize a struct using `new` (less common approach).
    var p *Person // Declare a pointer to a Person
    p = new(Person) // Allocate memory for a new Person struct and get a pointer to it
    p.Name = "David"
    p.Age = 35

    // Initialize a struct using a struct literal (preferred and more common approach).
    var person Person
    person.Name = "Alice"
    person.Age = 30

    // Print the information of the struct initialized with `new`.
    fmt.Printf("Person (initialized with new) - Name: %s, Age: %d\n", p.Name, p.Age)

    // Print the information of the struct initialized with a struct literal.
    fmt.Printf("Person (initialized with a struct literal) - Name: %s, Age: %d\n", person.Name, person.Age)
}
```

**`Output:`**

```go
Person (initialized with new) - Name: David, Age: 35
Person (initialized with a struct literal) - Name: Alice, Age: 30
```
