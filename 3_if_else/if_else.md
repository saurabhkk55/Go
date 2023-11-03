# `Conditional Statements in Go`

In Go, the syntax for the `if` statement is straightforward, and it can be combined with `else` and `else if` to handle different conditions. 

**Here is the basic syntax of the `if` statement in Go:**

```go
if condition {
    // Code to execute if the condition is true
} else {
    // Code to execute if the condition is false
}
```

You can also use `else if` to handle multiple conditions:

```go
if condition1 {
    // Code to execute if condition1 is true
} else if condition2 {
    // Code to execute if condition2 is true
} else {
    // Code to execute if neither condition1 nor condition2 is true
}
```

<div style="border: 5px solid white;"></div>
<br>

Here are some examples of how to use `if`, `else if`, and `else` in Go:

## 1. Simple `if` statement:

```go
package main

import "fmt"

func main() {
    x := 10

    if x > 5 {
        fmt.Println("x is greater than 5")
    }
}
```

<div style="border: 5px solid white;"></div>
<br>

## 2. `if` with `else`:

```go
package main

import "fmt"

func main() {
    x := 3

    if x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is not greater than 5")
    }
}
```

<div style="border: 5px solid white;"></div>
<br>

## 3. `if`, `else if`, and `else`:

```go
package main

import "fmt"

func main() {
    x := 7

    if x < 5 {
        fmt.Println("x is less than 5")
    } else if x == 5 {
        fmt.Println("x is equal to 5")
    } else {
        fmt.Println("x is greater than 5")
    }
}
```
<div style="border: 5px solid white;"></div>
<br>

## 4. `Nested if` statements:

```go
package main

import "fmt"

func main() {
    x := 10
    y := 5

    if x > 5 {
        if y > 3 {
            fmt.Println("Both x and y are greater than their respective thresholds.")
        } else {
            fmt.Println("x is greater than 5, but y is not greater than 3.")
        }
    } else {
        fmt.Println("x is not greater than 5.")
    }
}
```
<div style="border: 5px solid white;"></div>
<br>

These are some basic examples of how you can use `if`, `else if`, and `else` statements in Go to control the flow of your program based on different conditions. You can combine and nest these statements to handle more complex logic.