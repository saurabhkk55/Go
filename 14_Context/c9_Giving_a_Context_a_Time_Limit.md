# `Giving a Context a Time Limit`

## `context.WithTimeout function` takes two parameters:
1. The **`parent context`**:  The existing context to which you want to attach a timeout.
2. A **`time.Duration value`**: The duration for which the context should be automatically canceled if not finished by that time.

## `Result of context.WithTimeout`:
- The function returns a new context and a cancellation function.
- A `new context`: This is a modified version of the parent context, with the added deadline calculated from the specified `time.Duration`.
- A `cancellation function`: This function can be used to manually cancel the new context before the deadline is reached. If you call this function, the context will be canceled even if the operation hasn't finished within the specified time duration.

In simple terms, `context.WithTimeout` creates a new context with a specified timeout, and it automatically internally calculates the deadline by adding the specified duration to the current time using `time.Now()` and `time.Add()`. This function is convenient and efficient, as it allows you to easily set a time limit for a task without having to manually calculate the deadline using `time.Now()` and `time.Add()` that we actually do when we use `context.WithDeadline` function

With `context.WithTimeout`, you specify a duration (e.g., 5 seconds) instead of an exact deadline time. The function internally calculates the deadline by adding the specified duration to the current time using `time.Now()` and `time.Add()`.<br>
On the other hand, with `context.WithDeadline`, you need to explicitly specify the deadline by adding the specified duration to the current time using `time.Now()` and `time.Add()`.

## `In summary`:
`context.WithTimeout`: Takes a duration and internally calculates the deadline by adding it to the current time using `time.Now()` and `time.Add()`.
`context.WithDeadline`: Takes a time.Time value as the deadline and requires you to explicitly specify the deadline using `time.Now()` and `time.Add()`.

## `1. Code example`:

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Example using context.WithTimeout
	// Set the timeout duration for the context
	timeoutDuration := 3 * time.Second
	// Create a new context with timeout and a cancellation function
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), timeoutDuration)
	// Ensure that the cancellation function is called when the main function exits
	defer cancelTimeout()

	// Start a goroutine to simulate a time-consuming task
	go simulateWork(ctxTimeout)

	// Wait for the simulated work to complete or for the timeout to occur
	select {
	case <-ctxTimeout.Done():
		// If the context is canceled due to timeout
		fmt.Println("Timeout occurred. Work did not complete in time.")
	case <-time.After(5 * time.Second):
		// Simulating additional processing in the main function
		fmt.Println("Main function continues after work completion.")
	}

	// Wait for some second, allowing goroutine to perform its task before the main function exits.
	// time.Sleep(1 * time.Second)

	// // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // // //

	// Example using context.WithDeadline
	// Set a deadline 5 seconds from the current time
	deadline := time.Now().Add(5 * time.Second)
	// Create a new context with deadline and a cancellation function
	ctxDeadline, cancelDeadline := context.WithDeadline(context.Background(), deadline)
	// Ensure that the cancellation function is called when the main function exits
	defer cancelDeadline()

	// Start a goroutine to simulate a time-consuming task
	go simulateWork(ctxDeadline)

	// Wait for the simulated work to complete or for the deadline to occur
	select {
	case <-ctxDeadline.Done():
		// If the context is canceled due to the deadline
		fmt.Println("Deadline occurred. Work did not complete in time.")
	case <-time.After(3 * time.Second):
		// Simulating additional processing in the main function
		fmt.Println("Main function continues after work completion.")
	}

	// Sleep to allow some time for the goroutine to print its messages
	time.Sleep(3 * time.Second)
}

// simulateWork simulates a time-consuming task
func simulateWork(ctx context.Context) {
	for {
		select {
		case <-time.After(1 * time.Second):
			// Simulate ongoing work
			fmt.Println("Work in progress...")
		case <-ctx.Done():
			// If the context is canceled, stop the work
			fmt.Println("Work canceled due to context completion.")
			return
		}
	}
}
```

### `Output`: 
As we're using goroutine (thread) output will may change on each execution of the program. Just for refernce below we have one of the outputs. 

```go
Work in progress...
Work in progress...
Work canceled due to context completion.
Timeout occurred. Work did not complete in time.
Work in progress...
Work in progress...
Main function continues after work completion.
Work in progress...
Work in progress...
Work canceled due to context completion.
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
	// Create a new context 'ctx' with a timeout of 1500 milliseconds (1.5 seconds).
	ctx, cancelCtx := context.WithTimeout(ctx, 1500*time.Millisecond)

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

	// context.WithTimeout func will automatically cancel the context once the defined timeout is reached.
	// But if we want to cancel it manually for some good reason, we can do that as well.
	// Cancel the child context explicitly.
	cancelCtx()

	// Wait for some time, allowing goroutine to perform its task before the main function exits.
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

### `Output`:

```go
doAnother: 1
doAnother: 2
doAnother err: context deadline exceeded
doAnother: finished
doSomething: finished
```

### `Code explanation`:

When running the program this time, you should see the same output as when you used `context.WithDeadline`. The error message is even the same, showing you that `context.WithTimeout` is really just a wrapper to do the time.Now math for you.

# `Conclusion`
In this tutorial, you created a program to interact with `Go’s context package` in various ways. `First, you created a function that accepts a context.Context value as a parameter and used the context.TODO and context.Background functions to create empty contexts`. After that you used `context.WithValue to add values to new contexts and used the Value method to retrieve them in other functions`. Lastly, you used a `context’s Done method to determine when a context was done running. When paired with the functions context.WithCancel, context.WithDeadline, and context.WithTimeout`, you implemented your own context management to set limits on how long code using those contexts should run.

If you’d like to learn more about how contexts work with more examples, the [Go context](https://pkg.go.dev/context) package documentation includes additional information.
