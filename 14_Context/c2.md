## In this section, you will create a program with a function that receives a context as a parameter. You will also call that function using an empty context you create with the `context.TODO` and `context.Background` functions.

# First: Create an empty context using `context.TODO()`:

## 1. `Creating a Go File:`
Inside the 'contexts' directory, you create a Go file named 'main.go' where your program code will reside.

   ```bash
   nano main.go
   ```

## 2. `Writing the Go Code:`
You open 'main.go' using an editor (like nano) and write the following code:

   ```go
   package main

   import (
       "context"
       "fmt"
   )

   func doSomething(ctx context.Context) {
       fmt.Println("Doing something!")
   }

   func main() {
       ctx := context.TODO()
       doSomething(ctx)
   }
   ```

   This code defines a function `doSomething` that takes a `context.Context` parameter but doesn't use it for now. The `main` function creates an empty context using `context.TODO()` and then calls the `doSomething` function with that context.

## 3. `Running the Program:`
You run the program using the `go run` command:

   ```bash
   go run main.go
   ```

## 4. `Understanding Output:`
The program prints "Doing something!" to the console, indicating that the `doSomething` function was called successfully.
   
## Learnings:

## 1. `context.TODO():`
   - `context.TODO()` is a function in Go that creates an empty or starting context. It's used as a placeholder or a temporary solution when you need a context but you're not sure which specific context to use. 

## 2. `Empty Context:`
   - A context in Go is an object that carries deadlines, cancellations, and other request-scoped values across API boundaries and between processes. An empty context essentially means a context with no associated deadlines, cancellations, or specific values.

## 3. `Placeholder Usage:`
   - When you're writing code and need to pass a context, but you haven't decided on the type of context or specific values it should carry, you can use `context.TODO()` as a temporary placeholder. It doesn't provide any additional information or functionality but serves as a basic, empty context to satisfy the function's parameter requirements.

## 4. `Not Sure Which Context to Use:`
   - Sometimes, in the early stages of development or when integrating with different libraries, you might not be certain about the type of context to use. In such cases, `context.TODO()` can be used until you have a clearer understanding of your application's requirements.

In summary, `context.TODO()` is a way to create a simple, empty context that can be used temporarily when you need a context object but don't yet have all the details about what it should contain. It acts as a placeholder until you can provide a more specific context with the necessary information.

This program is a basic introduction to using contexts in Go. It creates a simple function named `doSomething` that takes a context as a parameter, the `main()` function creates an empty context using `context.TODO()`, and then calls the `doSomething` function with that context. The primary purpose here is to set up the structure and demonstrate the use of a context, even though it's not fully utilized in the example.

In this example, the `doSomething` function only prints a message. But the important part is that it receives a context, which allows it to be more flexible and adaptable in the future. We can later add code to the function that uses the context to achieve more complex tasks.


# Second: Create an empty context using `context.Background()`:

## 1. `Updating the Program:`
   - Open your `main.go` file and replace the line that creates an empty context using `context.TODO()` with `context.Background()`:

     ```go
     func main() {
         ctx := context.Background()
         doSomething(ctx)
     }
     ```

## 2. `Difference Between `context.TODO()` and `context.Background()`:`
   - Both functions create an empty context, but `context.Background()` is typically used when you intend to start a known context. The practical result is the same - an empty context that can be used in your program.

## 3. `Signaling Intent to Other Developers:`
   - The choice between `context.TODO()` and `context.Background()` is often about signaling your intent to other developers reading your code. If you are uncertain about which one to use, `context.Background()` is a good default.

## 4. `Running the Program:`
   - Execute the program again using the `go run` command:

     ```bash
     go run main.go
     ```

## 5. `Output:`
   - The output remains the same: "Doing something!" This is because the functionality of the code didn't change; only the way you create an empty context was updated.

## 6. `Empty Context and Adding Information:`
   - While an empty context doesn't provide much functionality on its own, you can enhance it by adding data. This allows you to pass information to other functions that use the context.

In summary, you updated your program to use `context.Background()` instead of `context.TODO()` to create an empty context. The practical effect is the same, but it may signal your intent better. Empty contexts become more useful when you add data to them, allowing you to pass information between functions.