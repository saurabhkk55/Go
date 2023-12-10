# `Using Data Within a Context`

One benefit of using `context.Context` in a program is the ability to access data stored inside a context. By adding data to a context and passing the context from function to function, each layer of a program can add additional information about what’s happening. For example, the first function may add a username to the context. The next function may add the file path to the content the user is trying to access. Finally, a third function could then read the file from the system’s disk and log whether it was loaded successfully or not as well as which user tried to load it.


To add a new value to a context, use the **`context.WithValue`** function in the context package. The function accepts three parameters: 1. the parent context.Context, 2. the key, and 3. the value. 
The `context.WithValue` function allows you to add a new key-value pair to a `context.Context` in Go. Here's a breakdown of its parameters and functionalities:

**`Parameters`:**

1. **`parent context.Context`:** This is the existing context to which you want to add the new value.
2. **`key`:** This is an identifier for the value being added. It can be any data type.
3. **`value`:** This is the actual data you want to store in the context. It can also be any data type.

**`Functionality`:**

- `context.WithValue` creates a new copy of the parent context with the specified key-value pair added to it.
- It **does not modify the original parent context**. This ensures that the original context remains untouched and can be used elsewhere without the added value.
- The new context returned by `context.WithValue` contains all the information of the parent context plus the newly added key-value pair.

Once you have a context.Context with a value added to it, you can access those values using the Value method of a context.Context. Providing the Value method with a key will return the value stored..

**Example:**

```go

func doSomething(ctx context.Context) {
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))
}

func main() {
	ctx := context.Background() // new-1

	ctx = context.WithValue(ctx, "myKey", "myValue") // new-2

	doSomething(ctx)
}
```

In this code, you’re assigning the new context ( ctx => // new-1 ) back to the ctx ( // new-2 ) variable that is used to hold the parent context ( ctx => // new-1 ). This is a relatively common pattern to use if you don’t have a reason to reference a specific parent context. If you did need access to the parent context as well, you can assign this value to a new variable, as you’ll see shortly.

Alternatively:

```go
package main

import (
	"context"
	"fmt"
)

// doSomething is a function that takes a context as an argument and retrieves
// a value associated with the key "myKey" from the context. It then prints the
// value if the key is found, and prints a message if the key is not found.
func doSomething(ctx context.Context) {
	// Retrieve the value associated with the key "myKey" from the provided context.
	value := ctx.Value("myKey")
	
	// If the key is not found, the value will be nil.
	// Check if the value is not nil.
	if value != nil {
		// Convert the value to a string.
		myValue := value.(string)

		// Print the value associated with the key.
		fmt.Printf("doSomething: myKey's value is %s\n", myValue)
	} else {
		// Print a message indicating that the key "myKey" does not exist in the provided context.
		fmt.Println("The key 'myKey' does not exist in the context.")
	}
}

func main() {
	// Create a new background context.
	backgroundContext := context.Background()

	// This function adds a new key "myKey" with the value "myValue" to the background context.
	contextWithKeyValue := context.WithValue(backgroundContext, "myKey", "myValue")
	// This function returns a new context with the added key "myKey" and value "myValue" to the background context.

	// Pass the context with the key-value pair to the doSomething function.
	doSomething(contextWithKeyValue)
	// This function receives the context containing the key-value pair and uses it.
}

```

To see the program’s output, run it using the go run command:
```go
Output:
doSomething: myKey's value is myValue
```


**`Key Points:`**
- `context.WithValue` adds data to a context while leaving the original context untouched.
- Use it to pass information across functions and layers for better code organization and readability.
- Remember to choose appropriate key names and data types for stored values.
