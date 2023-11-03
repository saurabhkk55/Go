# `Variable Declaration in Go`

In the Go programming language, you can define variables using the following syntax:

```go
var variableName dataType
```

Here's what each part of the syntax means:

- `var`: This is the keyword used to declare a variable.
- `variableName`: This is the name you choose for your variable, which should be a valid identifier (a sequence of letters, digits, and underscores) and should follow Go's naming conventions, typically using camelCase or PascalCase.
- `dataType`: This specifies the data type of the variable, indicating the type of value the variable will hold, such as `int`, `float64`, `string`, and so on.

Let's explore this with some examples:

```go
var age int // Declares an integer variable named 'age'.
var name string // Declares a string variable named 'name'.
var temperature float64 // Declares a floating-point variable named 'temperature'.
```
<div style="border: 5px solid white;"></div>
<br>

You can also initialize a variable at the time of declaration like this:

```go
var score int = 100 // Declares an integer variable named 'score' and assigns it the value 100.
var message string = "Hello, World!" // Declares a string variable named 'message' and assigns it a string value.
```

Go also allows you to use the short variable declaration `:=` when you want Go to infer the data type based on the value you provide:

```go
age := 30 // Infers the data type as 'int' and assigns the value 30 to the 'age' variable.
name := "Alice" // Infers the data type as 'string' and assigns the string value to the 'name' variable.
```

**`Note`**: The short-hand declaration syntax can only be used inside functions.

<div style="border: 5px solid white;"></div>
<br>

You can declare multiple variables of the same type without initializing them on the same line. This is useful when you want to declare several variables but assign values to them later in your code. Here's the syntax for declaring multiple variables:

```go
var variable1, variable2, variable3 dataType
```

For example:

```go
var x, y, z int
```

In this example, three integer variables `x`, `y`, and `z` are declared, but they are not initialized with values. You can assign values to these variables later in your code, like this:

```go
x = 10
y = 20
z = 30
```

Or, you can initialize them all in one statement:

```go
x, y, z = 10, 20, 30
```

<div style="border: 5px solid white;"></div>
<br>

You can declare and initialize multiple variables of the same type on the same line, separated by commas. This is a convenient way to declare and set the initial values for several variables of the same type in a single statement. Here's the syntax:

```go
var variable1, variable2, variable3 dataType = value1, value2, value3
```

Here's an example:

```go
var x, y, z int = 10, 20, 30
```

In this example, three integer variables `x`, `y`, and `z` are declared and initialized with values 10, 20, and 30, respectively. All of them have the same data type, which is `int`.

You can also use the short variable declaration `:=` to declare and initialize multiple variables of the same type in one line, as long as they are all newly declared variables:

```go
a, b, c := 1, 2, 3
```

## **`Note`**:
1. In Go, global variables must be initialized at the time of declaration.
2. Assign a new value to a global variable within a function, it updates the variable's value.
3. short variable declaration ( `:=` )is possible only inside the function.
4. **`Global Variables`**: In Go, global variables (or package-level variables) are defined at the package level and can be accessed by all the files within the same package. These variables are declared outside of any function in a package.

5. **`Exported Variables`**: To make a global variable accessible outside of the package (i.e., to other packages), you need to export it. In Go, variable names are exported if they start with an uppercase letter. This is known as the "exported" or "public" variable. Variables starting with a lowercase letter are not exported and are only accessible within the package they are defined in.

    **So, to clarify**:

    - A global variable declared with an uppercase initial letter (e.g., `VariableName`) is considered an exported global variable and can be accessed from other packages.
    - A global variable declared with a lowercase initial letter (e.g., `variableName`) is not exported and is only accessible within the package it's defined in.

    **Here's an example to illustrate this:**

    ```go
    // In package mypackage

    package mypackage

    var VariableName string // Exported, can be accessed from other packages
    var variableName string // Not exported, only accessible within mypackage
    ```

    In another package, you can access `VariableName`, but you won't be able to access `variableName`.

    ```go
    // In another package

    package anotherpackage

    import "path/to/mypackage"

    func main() {
        value := mypackage.VariableName // Accessing the exported global variable
        // This is valid

        value2 := mypackage.variableName // This is invalid, variableName is not exported
        // You will get a compilation error
    }
    ```
6. **`characters`**:
    In Go, there is no dedicated data type for individual characters like some other programming languages. Instead, you typically use the `rune` data type to work with Unicode characters. A `rune` is an alias for `int32` and represents a Unicode code point, which can be used to represent individual characters.

    Here's how you can declare a `rune` variable in Go:

    ```go
    var my_char rune // it is as same as "var my_char int32"
	// that is why we cant use it for all the characters

	fmt.Print("Enter a character: ")
	fmt.Scan(&my_char)

	fmt.Printf("MyChar= %c", my_char)
	fmt.Printf("\nMyChar Type= %T", my_char)
	fmt.Printf("\nMyChar value= %v", my_char)
    // Output
    // $ go run prog_12.go
    // Enter a character: a
    // MyChar=
    // MyChar Type= int32
    // MyChar value= 0

    // $ go run prog_12.go 
    // Enter a character: 7
    // MyChar= 7
    // MyChar Type= int32
    // MyChar value= 7

    // $ go run prog_12.go 
    // Enter a character: 45
    // MyChar= 45
    // MyChar Type= int32
    // MyChar value= 45
    ```

    You can also use the **`short-hand declaration syntax`**, which is available inside functions:

    ```go
    charName := 'a'
    char1, char2 := 'a', 'b'
    ```
    ### **`NOTE`**: 
    Working on a single character (like 'a', 'A', '1', '@', etc.) with the rune data type is not suitable as `rune` is nothing but `int32` and will accept only integer values within its range. To work with a single character, you should use a string with index 0 for the first character and then convert it into a `rune` or dont convert it into `rune` just use the first element of string at index 0. The same information is mentioned in below code:

    ```go
    // WAP to check what type of character a user has eneterned
    package main

    import "fmt"

    func main() {
        var my_char string

        fmt.Print("Enter a character: ")
        fmt.Scan(&my_char)

        // Check if the user entered more than one character
        if len(my_char) != 1 {
            fmt.Println("Please enter a single character.")
            return
        }

        // Convert the string to a rune(int32) to work with individual characters
        char := rune(my_char[0]) // Or, char := int32(my_char[0]) // both are same
        // char := my_char[0] // It will also work fine and correct

        if char >= 'a' && char <= 'z' {
            fmt.Printf("%c: Small alphabet", char)
        } else if char >= 'A' && char <= 'Z' {
            fmt.Printf("%c: Capital alphabet", char)
        } else if char >= '0' && char <= '9' {
            fmt.Printf("%c: Digit", char)
        } else {
            fmt.Printf("%c: Special character", char)
        }
    }
    ```
