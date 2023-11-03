## `Case 5: Pointer to Struct`

When you pass a struct as an argument to a function, it's passed by value, which means a copy of the struct is created. To modify the original struct, you can pass a pointer to it.

Demonstrates passing a pointer to a struct to modify the original struct:

```go
package main

import "fmt"

// Define a struct named Person to represent a person's information.
type Person struct {
    Name string
    Age  int
}

// Define a function that takes a pointer to a Person and modifies its fields.
func modifyPerson(p *Person) {
    p.Name = "Charlie"
    p.Age = 40
}

func main() {
    // Create a Person instance and initialize its fields.
    var person Person
    person.Name = "John"
    person.Age = 30

    // Call the modifyPerson function with a pointer to the Person instance.
    modifyPerson(&person)

    // Print the modified person's information.
    fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}
```

**`Output:`**

```go
Name: Charlie, Age: 40
```
