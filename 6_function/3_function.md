# `Function v/s Method in Go`

## 1. `Functions`:
Functions in Go are standalone blocks of code that can be called independently. They don't belong to any specific type, and you can call them directly.

Here's an example of a simple Go function:

```go
package main

import "fmt"

// add is a function that takes two integers and returns their sum.
func add(a, b int) int {
    return a + b
}

func main() {
    result := add(5, 3)
    fmt.Println(result) // Output: 8
}
```

In the code above, `add` is a function that takes two integers and returns their sum. You call it directly as `add(5, 3)`.

## 2. `Methods`:
Methods in Go are functions that are associated with a specific type. They have a receiver, which is a parameter that specifies on which type the method operates. Methods are typically used for defining behavior specific to a type.

Here's an example of a method in Go:

```go
package main

import "fmt"

// Rectangle is a user-defined type.
type Rectangle struct {
    Width  float64
    Height float64
}

// Area is a method associated with the Rectangle type.
// It has a receiver (r) of type Rectangle, which specifies
// that this method operates on instances of the Rectangle type.
// Methods are used for defining behavior specific to a type.
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    // Create an instance of the Rectangle type.
    rect := Rectangle{Width: 5, Height: 3}
    
    // Call the Area method on the rect instance.
    // Methods are called on specific instances and can
    // access and modify the instance's data.
    area := rect.Area()
    
    fmt.Println(area) // Output: 15
}
```

In this example, `Rectangle` is a user-defined type, and the `Area` method is defined for this type. You call the method on a specific instance of the `Rectangle` type, like `rect.Area()`. The receiver, `r`, inside the `Area` method refers to the specific `Rectangle` instance on which the method is called.

To summarize, the key difference between functions and methods in Go is that methods are associated with specific types and are called on instances of those types, while functions are standalone and can be called directly.