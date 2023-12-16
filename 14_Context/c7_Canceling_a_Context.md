# `Canceling a Context`

Canceling a context is the most straightforward and controllable way to end a context. Similar to including a value in a context with `context.WithValue`, it’s possible to associate a `cancel function` with a context using the `context.WithCancel function`. 

`context.WithCancel` function takes a parent context as a parameter and returns a new context along with a cancel function.

When we call the `cancel function`, it only cancels the context associated with it and any contexts derived from it.
Importantly, it does not cancel the parent context. The new context is created based on a parent context, and the cancellation of the new context doesn't affect the parent context directly.

**`Excption or special case: Preventing Parent Context Cancellation`**
While canceling a child context, it's essential to note that it doesn't prevent the cancellation of the parent context. The parent context can still be canceled by other means.

## `1. Code explanation`:

```go
package main

import (
	"context"
	"fmt"
	"time"
)

// simulateChildTask simulates an ongoing task in a child context.
func simulateChildTask(ctx context.Context) {
	for {
		select {
        /*
            1.	The ctx.Done() channel is a part of the Go context package,
            2.	<-ctx.Done() syntax is used to read a value from the ctx.Done() channel.
            3.	If the context 'ctx' is canceled, then its channel 'ctx.Done()' will be closed as well.
                At that point, any attempt to read from the closed channel using <-ctx.Done() will immediately return a nil value.
                This behavior indicates that the channel is closed due to the cancellation of the context.
            4.	So, to summarize:
                If ctx.Done() returns a non-nil value, it means the context is not canceled.
                If ctx.Done() returns a nil value, it means the context is canceled, and the associated channel is closed.
        */
		case <-ctx.Done():
			// If the child context is canceled, stop the task
			fmt.Println("Child Task canceled")
			return
		default:
			// Simulate ongoing work in the child task
			fmt.Println("Child Working...")
			time.Sleep(1000 * time.Millisecond) // 1 Second = 1000 Millisecond
		}
	}
}

// checkParentContext checks if the parent context is canceled or not.
func checkParentContext(ctx context.Context) {
	select {
	case <-ctx.Done():
		// Print a message indicating that the parent context is canceled
		fmt.Println("Parent context is canceled.")
	default:
		// Print a message indicating that the parent context is not canceled
		fmt.Println("Parent context is not canceled.")
	}
}

func main() {
	// Create a parent context
	parentContext := context.Background()

	// Use context.WithCancel to create a child context and a cancel function
	childContext, cancel := context.WithCancel(parentContext)

	// Start a goroutine to simulate some task in the child context
	go simulateChildTask(childContext)

	// Simulate main program doing something in the parent context
	fmt.Println("Main program is doing some work in the parent context...")

	// Wait for a while
	time.Sleep(5 * time.Second)

	// Now, let's cancel only the child context to stop its task
	cancel()

	// Wait for a moment to see the effect
	time.Sleep(1 * time.Second)

	// Print a message after canceling the child context
	fmt.Println("Main program completed")

	// Now, let's demonstrate that the parent context is not canceled
	checkParentContext(parentContext)
}
```

### `Output:`
```go
Main program is doing some work in the parent context...
Child Working...
Child Working...
Child Working...
Child Working...
Child Working...
Child Task canceled
Main program completed
Parent context is not canceled.
```

### `Code explanation:`

In the provided Go code, a parent context (`parentContext`) is created using `context.Background()`. Then, a child context (`childContext`) is derived from the parent context using `context.WithCancel(parentContext)`.

The child context is used to simulate an ongoing task in a goroutine (`simulateChildTask`). This goroutine periodically checks if the child context is canceled using the `select` statement with `case <-ctx.Done():` and terminates if it is canceled.

Meanwhile, in the main function, the main program simulates doing some work in the parent context and then waits for a while. Afterward, the `cancel` function is called to cancel the child context, effectively stopping the ongoing task in the goroutine.

The key point here is that canceling the child context does not automatically cancel the parent context. The cancellation of the child context only affects the goroutine running in that context. The parent context remains unaffected.

This behavior is demonstrated in the `checkParentContext` function, which checks whether the parent context is canceled or not. After canceling the child context, the `checkParentContext(parentContext)` function is called, and it prints a message indicating that the parent context is not canceled. This is because the cancellation of the child context doesn't propagate up to the parent context.

In summary, the cancellation of a child context does not automatically propagate to its parent context in Go. Each context operates independently, and canceling one context does not affect its parent or siblings. This is a fundamental feature of the context package in Go, allowing for fine-grained control over the cancellation of tasks and goroutines.

## `2. Code explanation`:

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a background context
	parentCtx := context.Background()

	// Call the doSomething function with the parent context
	doSomething(parentCtx)
}

func doSomething(ctx context.Context) {
	// Create a child context with cancel function
	ctx, cancelCtx := context.WithCancel(ctx)

	// Create a channel for communication between doSomething and doAnother
	// Or, It just create a channel to send and receive integers
	printCh := make(chan int)

	// Start a goroutine to execute the doAnother function with the child context 'ctx' and the 'printCh' channel
	go doAnother(ctx, printCh)

	// Iterate through numbers 1 to 3 and send them to the 'printCh' channel
	for num := 1; num <= 3; num++ {
		printCh <- num
	}

	// Cancel the child context to signal the doAnother goroutine to finish
	cancelCtx()

	// Wait for 100 milliseconds, allowing some time for the doAnother goroutine to finish processing
	time.Sleep(100 * time.Millisecond) // comment this line to see the behaviour of goroutine

	// Print a message indicating that doSomething has finished
	fmt.Printf("doSomething: finished\n")
}

func doAnother(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			// If the child context is canceled, print the error (if any) and finish
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			fmt.Printf("doAnother: finished\n")
			return
		// Receive an integer from the 'printCh' channel
		case num := <-printCh:
			fmt.Printf("doAnother: %d\n", num)
		}
	}
}
```

### `Output:`

```go
doAnother: 1
doAnother: 2
doAnother: 3
doAnother err: context canceled
doAnother: finished
doSomething: finished
```

### `Code explanation:`

In this code, `doSomething function` acts like a function that sends work to one or more goroutines reading from a work channel. In this case, `doAnother` is the worker and printing the numbers is the work it’s doing. Once the doAnother goroutine is started, doSomething begins sending numbers to be printed. Inside the `doAnother function`, the select statement is waiting for either the `ctx.Done channel` to close or for a number to be received on the `printCh channel`. 

The doSomething function sends three numbers on the channel, triggering the `fmt.Printf` for each number, and then calls the `cancelCtx function to cancel the context`. After the doAnother function reads the three numbers from the channel, it will wait for the next channel operation. Since the next thing to happen is `doSomething calling cancelCtx`, the `ctx.Done branch is called`.

When the `ctx.Done` branch is called, the code uses the `Err function` provided by `context.Context` to determine how the context ended. Since your program is using the `cancelCtx function` to cancel the context, `the error you see in the output is context canceled`.

Note: If you’ve run Go programs before and looked at the logging output, you may have seen the context canceled error in the past. When using the Go http package, this is a common error to see when a client disconnects from the server before the server handles the full response.

Once the doSomething function has canceled the context, it uses the `time.Sleep function` to wait a short amount of time to give doAnother time to process the canceled context and finish running. After that, it exits the function. In many cases the `time.Sleep` would not be required, but it’s needed since the example code finishes executing so quickly. If `time.Sleep` isn’t included, the program may end before you see the rest of the program’s output on the screen.

The `context.WithCancel` function and the cancel function it returns are most useful when you want to control exactly when a context ends, but sometimes you don’t want or need that amount of control. The next function available to end contexts in the context package is context.WithDeadline and is the first function that will automatically end a context for you.
