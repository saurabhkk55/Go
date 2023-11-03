# `Array in Go`

In Go, an array is a fixed-size collection of elements of the same data type.<br>
The size of an array is determined at the time of declaration and cannot be changed during the program's execution.<br>
Here's an explanation of arrays in Go with code examples and comments:

```go
package main

import "fmt"

func main() {
    // Declaring an array with a fixed size
    var fruits [4]string

    // Initializing elements of the array
    fruits[0] = "Apple"
    fruits[1] = "Banana"
    fruits[2] = "Orange"
    fruits[3] = "Grapes"

    // Accessing elements of the array using index
    fmt.Println("Fruit at index 0:", fruits[0]) // Apple
    fmt.Println("Fruit at index 2:", fruits[2]) // Orange

    // Getting the length of the array
    fmt.Println("Length of fruits array:", len(fruits)) // 4

    // Declare and initialize an array in a single line without shorthand
    var numbers = [4]int{1, 2, 3, 4}

    // Declaring and initializing an array in one line using shorthand
    num := [5]int{1, 2, 3, 4, 5}

    // Iterating through an array using a for loop
    fmt.Println("Numbers array:")
    for i := 0; i < len(numbers); i++ {
        fmt.Printf("Element at index %d: %d\n", i, numbers[i])
    }

    // Shorter way to iterate through an array using range
    fmt.Println("Fruits array:")
    for index, fruit := range fruits {
        fmt.Printf("Fruit at index %d: %s\n", index, fruit)
    }
}
```

# `Good learnings:`

- In Go, the `range` keyword is used to iterate over various data structures, including arrays, slices, maps, and channels.
- When you use `range` in a `for` loop, it returns two values for each iteration: the index (or key) and the value associated with that index.
- You can use both values or just one, depending on your needs. In your example:

```go
for index, fruit := range fruits {
    fmt.Printf("Fruit at index %d: %s\n", index, fruit)
}
```

Here's what's happening:

- `index` is a variable that holds the index of the current element in the array.
- `fruit` is a variable that holds the value of the current element in the array.

<div style="border: 5px solid white;"></div>
<br>

You can use any variable names you like in place of `index` and `fruit`. The key concept is that `range` returns two values, and you can assign them to variables of your choice. For example, you could use `i` and `f` like this:

```go
for i, f := range fruits {
    fmt.Printf("Fruit at index %d: %s\n", i, f)
}
```
In this case, `i` holds the index, and `f` holds the fruit name. You can use as many variables as necessary for your specific use case. If you only need one of the values (e.g., just the index or just the value), you can omit the variable you don't need. For instance, if you only want the values, you can do:

<div style="border: 5px solid white;"></div>
<br>

```go
for _, fruit := range fruits {
    fmt.Printf("Fruit: %s\n", fruit)
}
```

In this case, the underscore `_` is used as a placeholder for the index value, indicating that you're not interested in using it.

## `Take input from a user to initialize an array`
To take input from a user to initialize an array in Go, you can use the `fmt` package to read user input from the standard input (usually the keyboard). Here's an example of how you can do this:

```go
package main

import (
    "fmt"
)

func main() {
    // Create an array
    var myArray [5]int

    // Prompt the user to enter each element of the array
    for i := 0; i < len(myArray); i++ {
        fmt.Printf("Enter element %d: ", i)
        _, err := fmt.Scan(&myArray[i])
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
    }

    // Display the resulting array
    fmt.Println("Your array:", myArray)
}
```

In this example, we create an array with a fixed size of 5 elements. We then use a loop to prompt the user for each element of the array, and we store those values in the array.

Error handling is important when working with user input to handle cases where the user enters invalid input. In this code, we check for errors when reading user input and display an error message if any issues occur.

Please note that arrays in Go have a fixed size, and you must specify the size at the time of declaration. If you want a dynamic data structure that can grow or shrink, you should use slices instead.