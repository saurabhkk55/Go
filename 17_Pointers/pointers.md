# `Pointers`
In Go, pointers are variables that store the memory address of another variable. They are used to manage and manipulate data indirectly by referencing the memory location rather than the actual value. Here's a basic overview of pointers in Go:

### Declaration of Pointers:

To declare a pointer, you use the `*` symbol followed by the type of the variable it will point to. For example:

```go
var ptr *int // declares a pointer to an int
```

### Getting the Address of a Variable:

You can obtain the memory address of a variable using the `&` operator:

```go
x := 10
ptr = &x // ptr now contains the memory address of x
```

### Dereferencing Pointers:

To access the value stored at the memory address pointed to by a pointer, you use the `*` operator:

```go
fmt.Println(*ptr) // prints the value stored at the memory address ptr points to
```

### Example:

```go
package main

import "fmt"

func main() {
    // Declare a variable
    x := 42

    // Declare a pointer and assign the address of the variable to it
    var ptr *int
    ptr = &x

    // Print the value and memory address
    fmt.Println("Value:", x)
    fmt.Println("Address:", &x)

    // Print the value using the pointer
    fmt.Println("Value via pointer:", *ptr)
}
```

### Pass by Reference:

Pointers can be used to achieve pass-by-reference in functions. When you pass a variable by reference using a pointer, changes made to the variable inside the function will affect the original variable.

```go
package main

import "fmt"

func modifyValue(ptr *int) {
    *ptr = 100
}

func main() {
    x := 42

    fmt.Println("Before:", x)

    modifyValue(&x)

    fmt.Println("After:", x)
}
```

In the `modifyValue` function, the parameter is a pointer to an int (`*int`). The function modifies the value at the memory address pointed to by the pointer, and the changes are reflected outside the function.

These are some basic concepts of pointers in Go. Understanding pointers is essential for working with more advanced data structures and efficient memory management.