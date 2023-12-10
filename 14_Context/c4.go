package main

import (
	"context"
	"fmt"
)

// doSomething modifies the context by adding a key-value pair and calls doAnother
func doSomething(ctx context.Context) {
	// Print the value associated with "myKey" in the original context (parentContext)
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

/*
Output:
doSomething: myKey's value is originalValue
doAnother: myKey's value is anotherValue
doSomething: myKey's value is originalValue
*/
