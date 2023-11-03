## `Case 9: Custom String Representation`
You can implement the Stringer interface to customize the string representation of a struct.


```go
package main

import (
    "fmt"
)

// Person represents a person with a name and age.
type Person struct {
    Name string
    Age  int
}

// String method is used to customize the string representation of the Person struct.
func (p Person) String() string {
    return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func main() {
    // Create a Person instance with a name and age.
    var person Person
    person.Name = "John Doe"
    person.Age = 30

    // Print the custom string representation of the Person instance.
    fmt.Println(person)
}
```

In this code:

1. We define a `Person` struct with `Name` and `Age` fields.

2. We implement the `String` method for the `Person` struct. This method is a part of the `Stringer` interface, and it customizes the string representation of the struct. It uses `fmt.Sprintf` to format the string as desired.

3. In the `main` function, we create a `Person` instance and then print it using `fmt.Println`. The custom `String` method is called to display the custom string representation of the `Person` struct.

This will output:

```
Name: John Doe, Age: 30
```
