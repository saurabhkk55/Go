## `Case 4: Method Attachments`

You can associate methods with a struct type, turning it into a receiver for those methods. This is a way to add behavior to your data.

Defining a method attached to the struct `Person`:

```go
package main

import "fmt"

// Define a struct named Person to represent a person's information.
type Person struct {
    Name string
    Age  int
}

// Define a method named PrintInfo associated with the Person struct.
// This method takes a Person receiver and prints the person's name and age.
func (p Person) PrintInfo() {
    fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func main() {
    // Create a Person instance.
    var person Person
    person.Name = "John"
    person.Age = 30

    // Call the PrintInfo method on the Person instance to print the person's information.
    person.PrintInfo()
}
```

**`Output:`**

```go
Name: John, Age: 30
```
