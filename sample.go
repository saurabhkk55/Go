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
