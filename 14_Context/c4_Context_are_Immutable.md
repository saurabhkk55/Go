# `context.Context` are immutable

When using contexts, it’s important to know that the values stored in a specific context.Context are immutable, meaning they can’t be changed. When you called the context.WithValue, you passed in the parent context and you also received a context back. You received a context back because the context.WithValue function didn’t modify the context you provided. Instead, it wrapped your parent context inside another one with the new value.

## `Explanation:`

Imagine you have a locked box (context) with some contents inside (values). You can't change what's inside directly. Instead, if you want to add something new, you need to create a new, bigger box (another context) that contains the original box (parent context) and the new item.

The `context.WithValue` function acts like adding a new item to a box. It doesn't modify the original box, but creates a new one with the original contents and the new item.

Here's how it works in the example:

```go
package main

import (
	"context"
	"fmt"
)

// doSomething modifies the context by adding a key-value pair and calls doAnother
func doSomething(ctx context.Context) {
	// Print the value associated with "myKey" in the original context
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))

	// Creating a new context by adding new/another key-value pair
	anotherCtx := context.WithValue(ctx, "myKey", "anotherValue")

	// Calling another function with the modified context
	doAnother(anotherCtx)

	// Print the value associated with "myKey" in the original context
	// Even though we modified the context, the original context remains unchanged
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))
}

// doAnother prints the value associated with "myKey" in the provided context
func doAnother(ctx context.Context) {
	// Retrieve and print the value associated with "myKey" in the context
	fmt.Printf("doAnother: myKey's value is %s\n", ctx.Value("myKey"))
}

func main() {
	// Creating a parent context with an initial key-value pair
	parentContext := context.WithValue(context.Background(), "myKey", "originalValue")

	// Calling the doSomething function with the parent context
	doSomething(parentContext)
}
```

## `Explanation:`

In `main`:
- The main function creates a parent context with an initial key-value pair "myKey" -> "originalValue".
- The doSomething function is then called with this parent context.

In `doSomething`:
- It prints the value associated with "myKey" in the original context, which is "originalValue".
- It creates a new context (anotherCtx) by adding a new key-value pair "myKey" -> "anotherValue".
- Calls doAnother with the modified context (anotherCtx).
- After returning from doAnother, it prints the value associated with "myKey" in the original context again.<br>
  Note: Even though we modified the context in anotherCtx, the original context (ctx) remains unchanged.

In `doAnother`:
- It prints the value associated with "myKey" in the provided context, which is "anotherValue" when called from doSomething.

The final output shows the sequence of events and the values associated with "myKey" at different points in the program.
