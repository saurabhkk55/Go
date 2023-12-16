# `Giving a Context a Deadline`

Think of it like setting a deadline for yourself when working on a task. You decide how much time you want to allocate to the task, and if you exceed that time, you consider the task canceled or incomplete.

## `1. Using context.WithDeadline`:
To set a deadline in Go, you use the `context.WithDeadline function`.
`context.WithDeadline function` accepts two things as parameters:
1. The **`parent context`**: the existing context to which you want to add the deadline.
2. A **`time.Time value`**: representing when the context should be automatically canceled if not finished by that time.

## `2. Result of context.WithDeadline`:
- The function returns a new context and a cancellation function.
- The new context is a modified version of the parent context, now with the added deadline.
- The cancellation function can be used to manually cancel the context before the deadline if needed.

## `3. Behavior Similar to context.WithCancel`:
- The mechanism is somewhat similar to context.WithCancel, but with the added aspect of a time deadline.
- If the deadline is exceeded, the context is automatically canceled, similar to manually calling the cancel function in context.WithCancel.

## `4. Scope of Deadline Impact`:
- The deadline affects only the new context created by `context.WithDeadline` and `any other contexts that use it as a parent`.
- It doesn't impact other unrelated contexts.

## `5. Manual Cancellation`:
Just like with `context.WithCancel`, you can manually cancel the context at any time before the deadline by calling the cancellation function.

## `6. Use Case`:
This is useful in scenarios where you want to enforce a time limit for an operation, such as waiting for data or completing a task. If the operation takes longer than the specified deadline, it's automatically considered canceled.

## `1. Code example`:

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a parent context as the starting point for managing contexts.
	parentContext := context.Background()

	// Record the starting time to measure the elapsed time later.
	start := time.Now()

	// Set a deadline 6 seconds from now.
	deadline := time.Now().Add(6 * time.Second)
	// deadline := time.Now().Add(4 * time.Second)
	// deadline := time.Now().Add(5 * time.Second)
	// Try all these 3 one-by-one, will help to understand the use of deadline with context

	// Create a child context with a deadline using context.WithDeadline.
	// This child context will automatically be canceled when the deadline is reached.
	childContext, cancel := context.WithDeadline(parentContext, deadline)

	// Defer the cancellation function to ensure it gets called even if the main function returns early.
	defer cancel()

	// Simulate a long-running operation (e.g., fetching data) in a goroutine.
	go performTask(childContext)

	// Wait for 10 seconds to allow the goroutine to perform its task before the main function exits.
	time.Sleep(10 * time.Second)

	// Measure the total time taken by the code.
	elapsed := time.Since(start)
	fmt.Println("Code execution time:", elapsed)
}

// performTask simulates a task that may take some time to complete.
func performTask(ctx context.Context) {
	for {
		select {
		// Check if the context is done, indicating whether the task completed or the deadline expired.
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				// If the context deadline was exceeded, print a message and cancel the task.
				fmt.Println("Task deadline exceeded. Canceling the task.")
			} else {
				// If the context is done for other reasons, print a message indicating successful completion.
				fmt.Println("Task completed within the deadline.")
			}
			return

		// Use time.After to create a channel that receives the current time after 5 seconds.
		// This case simulates a successful completion of the task after 5 seconds.
		case <-time.After(5 * time.Second):
			fmt.Println("Task completed successfully.")
		}
	}
}
```

### `Output: When, deadline := time.Now().Add(6 * time.Second)`:
```go
Task completed successfully.
Task deadline exceeded. Canceling the task.
Code execution time: 10.0133371s

```

### `Output: When, deadline := time.Now().Add(4 * time.Second):`
```go
Task deadline exceeded. Canceling the task.
Code execution time: 10.008484s
```

### `Output: When, deadline := time.Now().Add(5 * time.Second):`
```go
Task completed successfully.
Task deadline exceeded. Canceling the task.
Code execution time: 10.0076197s
```

## `2. Code example`:

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
	// Set a deadline of 1500 milliseconds (1.5 seconds) from now
	deadline := time.Now().Add(1500 * time.Millisecond)

	// Create a child context with the deadline
	ctx, cancelCtx := context.WithDeadline(ctx, deadline)

	// Ensure that cancelCtx is called to release resources when doSomething returns
	// defer cancelCtx() will execute just before doSomething terminates/exits
	defer cancelCtx()

	// Create a channel for communication between goroutines
	printCh := make(chan int)

	// Start a goroutine to execute the doAnother function with the child context and communication channel
	go doAnother(ctx, printCh)

	// Loop to send integers to the printCh channel and handle context cancellation
	for num := 1; num <= 3; num++ {
		select {
		// Send the current integer to the printCh channel if possible
		case printCh <- num:
			// Simulate some work with a sleep of 1 second
			time.Sleep(1 * time.Second)
		// Check if the context is canceled (parent or child)
		case <-ctx.Done():
			break
		}
	}

	// Cancel the child context explicitly
	cancelCtx()

	// Simulate a short delay
	time.Sleep(100 * time.Millisecond)

	// Print a message indicating that doSomething has finished
	fmt.Printf("doSomething: finished\n")
}

func doAnother(ctx context.Context, printCh <-chan int) {
	// Infinite loop to listen for events or context cancellation
	for {
		select {
		// Check if the context is canceled
		case <-ctx.Done():
			// If the child context is canceled, print the error (if any) and finish
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			// Print a message indicating that doAnother has finished
			fmt.Printf("doAnother: finished\n")
			return
		// Receive an integer from the 'printCh' channel
		case num := <-printCh:
			// Print the received integer
			fmt.Printf("doAnother: %d\n", num)
		}
	}
}
```

### `Code explanation`:
**`In this code, you will notice `cancelCtx` is called twice:`**
- The first call to `cancelCtx() function` is within the defer statement of the doSomething function. This ensures that even if the function exits early due to an error or other reason, the context will be canceled and resources cleaned up.

- The purpose of this defer statement is to ensure that `cancelCtx()` is called even if there are return statements or panics in the function 'doSomething' . This is a common practice to ensure proper cleanup and resource release.

- In the provided code, the defer `cancelCtx()` statement is not strictly necessary because cancelCtx() is also explicitly called later in the function. However, as mentioned in the comments, it can serve as a safety measure. If the function 'doSomething' is modified in the future and a return statement is added, the defer statement ensures that the cancellation is still performed, providing a more robust and defensive coding style.

- `cancelCtx() function`  explicitly cancels the context. It means that it triggers the cancellation of the context, and any resources associated with that context can be cleaned up.

**`In the output:`**
- You’ll see that the context was canceled due to a context deadline exceeded error before all three numbers were printed. 
- Since the deadline was set to 1.5 seconds after the doSomething function started running and doSomething is set to wait one second after sending each number, the deadline will be exceeded before the third number is printed. 
- Once the deadline is exceeded, both the doAnother and doSomething functions finish running because they’re both watching for the ctx.Done channel to close. 
- You can also tweak the amount of time added to time.Now to see how various deadlines affect the output. 
- If the deadline ends up at or over 3 seconds, you could even see the error change back to the context canceled error because the deadline is no longer being exceeded.

The `context deadline exceeded error` may also be familiar to you if you’ve used Go applications or looked at their logs in the past. This error is common to see in web servers that take some time to complete sending responses to the client. If a database query or some processing takes a long time, it could cause the web request to take longer than is allowed by the server. Once the request surpasses the limit, the request’s context would be canceled and you’d see the context deadline exceeded error message.

Ending a context using `context.WithDeadline` instead of `context.WithCancel` allows you to specify a specific time when a context needs to end without needing to keep track of that time yourself. If you know the `time.Time` when a context should end, `context.WithDeadline` is likely a good candidate for managing your context’s endpoint. Other times, you don’t care about the specific time a context ends and you just know you want it to end one minute after it’s started. It’s possible to do this using `context.WithDeadline` and other time package functions and methods, but Go also provides the context.WithTimeout function to make that work easier.