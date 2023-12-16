# `Ending a Context`

## `Theory-1: Ending a Context`
Another powerful tool `context.Context` provides is a way to signal to any functions using it that the context has ended and should be considered complete. By signaling to these functions that the context is done, they know to stop processing any work related to the context that they may still be working on. Using this feature of a context allows your programs to be more efficient because instead of fully completing every function, even if the result will be thrown out, that processing time can be used for something else. For example, if a web page request comes to your Go web server, the user may end up hitting the “Stop” button or closing their browser before the page finishes loading. If the page they requested requires running a few database queries, the server may run the queries even though the data won’t be used. However, if your functions are using a `context.Context`, your functions will know the context is done because Go’s web server will cancel it, and that they can skip running any other database queries they haven’t already run. This will free up web server and database server processing time so it can be used on a different request instead.


## `Explanation-1:`
Imagine you're in charge of a busy restaurant. Customers (requests) come in, and you have chefs (functions) preparing their orders. Now, let's say a customer changes their mind and decides to leave before their food is ready. In a traditional setup, your chefs might continue cooking the food even though the customer won't eat it.

Now, think of a "context" like a signal that tells the chefs when a customer has left, and they can stop cooking for that customer immediately. In programming, this is similar to using a tool called `context.Context`.

So, in the programming world, let's say your web server (like a restaurant) gets a request to load a web page. However, the user might close their browser or click the "Stop" button before the page fully loads. Without a context, your server might keep doing unnecessary work, like fetching data from a database even if it won't be used.

But, if your functions use a `context.Context`, they are smart enough to know when the user has "left" (closed the browser or canceled the request). So, they stop any extra work immediately, like skipping unnecessary database queries. This way, your server becomes more efficient because it doesn't waste time on things that won't be useful.

In simpler terms, using a context helps your program be more responsive and not waste time on tasks that aren't needed, just like chefs stopping cooking when a customer leaves the restaurant.

## `Theory-2: Determining if a Context is Done`
No matter the cause of a context ending, determining if a context is done happens the same way. The `context.Context` type provides a method called `Done` that can be checked to see whether a context has ended or not. This method returns a channel that is closed when the context is done, and any functions watching for it to be closed will know they should consider their execution context completed and should stop any processing related to the context. The Done method works because no values are ever written to its channel, and when a channel is closed that channel will start to return nil values for every read attempt. By periodically checking whether the Done channel has closed and doing processing work in-between, you’re able to implement a function that can do work but also knows if it should stop processing early. Combining this processing work, the periodic check of the Done channel, and the select statement goes even further by allowing you to send data to or receive data from other channels simultaneously.

The select statement in Go is used to allow a program to try reading from or writing to a number of channels all at the same time. Only one channel operation happens per select statement, but when performed in a loop, the program can do a number of channel operations when one becomes available. A select statement is created by using the keyword `select`, followed by a code block enclosed in curly braces ({}), with one or more case statements inside the code block. Each case statement can be either a channel read or write operation, and the select statement will block until one of the case statements can be executed. Suppose you don’t want the select statement to block, though. In that case, you can also add a default statement that will be executed immediately if none of the other case statements can be executed. It looks and works similarly to a switch statement, but for channels.

## `Explanation-2:`
Certainly! Let's break down the information point by point:

1. **`Determining if a Context is Done`:**
   - In Go programming language, the `context.Context` type provides a method called `Done`.
   - This method returns a channel that is closed when the context is done.
   - When the channel is closed, any functions watching it should interpret it as a signal to consider their execution context completed and stop any related processing.
   - The `Done` method is useful because no values are written to its channel. When a channel is closed, it starts returning nil values for every read attempt.
   - OR, The Done channel only returns nil when it's closed, which means it's done sending any values.
   - When a Done channel is closed in Go, it indicates that no more values will be sent on it. Reading from a closed channel will immediately return a nil value. It's important to note that only the sender (the goroutine or part of the code that is sending values to the channel) should close a channel. The receiver (the part of the code that is reading from the channel) should not close it. Closing the channel by the receiver could lead to unintended consequences.

2. **`Working with the Done Channel`:**
   - By periodically checking whether the Done channel has closed, a program can determine if the associated context has ended.
   - The periodic checks enable the implementation of functions that can perform work but also know when to stop processing early based on the context's status.
   - This mechanism is crucial for handling scenarios where a context needs to be canceled or timed out.

3. **`Select Statement in Go`:**
   - The `select` statement in Go allows a program to interact with multiple channels concurrently.
   - It is used to try reading from or writing to several channels simultaneously.
   - The `select` statement is placed within a loop to perform multiple channel operations as they become available.
   - It consists of a code block enclosed in curly braces (`{}`) and contains one or more `case` statements representing channel operations.
   - Each `case` statement can be either a channel read or write operation.

4. **`Blocking Behavior of Select`:**
   - The `select` statement blocks until one of its `case` statements can be executed.
   - If none of the `case` statements can be executed immediately, the `select` statement can include a `default` statement that is executed immediately.
   - This behavior is similar to a `switch` statement, but it is specifically designed for working with channels.

5. **`Simultaneous Channel Operations`:**
   - The `select` statement allows a program to perform channel operations concurrently.
   - When one of the channels becomes available, the associated `case` statement is executed.
   - This enables efficient handling of multiple communication channels, such as in the context of Go routines and channels.

6. **`Additional Information`:**
   - The combination of processing work, periodic checks on the `Done` channel, and the `select` statement allows for a flexible and responsive approach in handling contexts and channels.
   - Go's concurrency model, including goroutines and channels, facilitates the development of concurrent and parallel programs.

Overall, the integration of `context.Context`, the `Done` method, and the `select` statement provides a robust mechanism for managing concurrent processes, handling context completion, and efficiently interacting with multiple channels in Go.

### `Code example`
The following code example shows how a select statement could potentially be used in a long-running function that receives results from a channel, but also watches for when a context’s Done channel is closed:

```go
ctx := context.Background()
resultsCh := make(chan *WorkResult)

for {
	select {
	case <- ctx.Done():
		// The context is over, stop processing results
		return
	case result := <- resultsCh:
		// Process the results received
	}
}
```

Certainly! Let's break down the provided code example and explanation point by point:

1. **`Initialization of Context and Channel`:**
   ```go
   ctx := context.Background()
   resultsCh := make(chan *WorkResult)
   ```
   - Creates a background context (`context.Background()`) and a channel (`resultsCh`) for receiving `WorkResult` values.

2. **`Select Statement in a Loop`:**
   ```go
   for {
      select {
      case <-ctx.Done():
         // The context is over, stop processing results
         return
      case result := <-resultsCh:
         // Process the results received
      }
   }
   ```
   - Utilizes an `infinite for loop` with a `select` statement to continuously monitor two cases:
      - **Case 1:** If the context (`ctx`) is done (closed), the loop stops processing results and returns.
      - **Case 2:** If there is a new result available in the `resultsCh` channel, it processes the result.

3. **`Handling Context Cancellation`:**
   - The `ctx.Done()` case is responsible for checking whether the context is done. If the context is canceled or times out, this case is triggered.
   - The loop exits, and subsequently, the function returns, stopping any further processing.

4. **`Order of Case Execution`:**
   - The order of execution for the `select` cases is not guaranteed. The select statement will execute the first case that is ready, and if multiple cases are ready simultaneously, it may seem random which one is executed.

5. **`Busy Loop and Default Case`:**
   - The provided example does not include a default case. If a default case were added without any code, it would not change the behavior. A default case could be useful if you want to perform an action when none of the other cases are immediately ready.
   - The absence of a default case causes the select statement to be a busy loop, as it repeatedly executes without waiting for any channel operations to be ready.

6. **`Context Cancellation and Channel Operations`:**
   - The only way to exit the loop is by closing the channel returned by `Done`, and the only way to close the `Done` channel is by ending the context.
   - The context can be canceled explicitly by calling a function to "cancel" it, using the `context.WithCancel` function, for example.

7. **`Busy Loop and CPU Consumption`:**
   - The provided loop, without any waiting mechanism, can be considered a busy loop. It repeatedly runs without waiting for something to happen, potentially consuming a significant amount of CPU.
   - A busy loop might be useful when you want to check if a channel is ready before performing a non-channel operation.

8. **`Ending the Context`:**
   - The Go context package provides various ways to end a context. The most direct option is to call a function to "cancel" the context, such as using `context.WithCancel`.

In summary, the code example demonstrates a pattern where a function continuously processes results from a channel while also monitoring the closure of a context's `Done` channel to gracefully stop processing when the context is canceled or times out. The select statement, in conjunction with the context, enables flexible handling of multiple conditions concurrently. The absence of waiting mechanisms, such as sleep or timeouts, makes the loop a busy loop that might consume CPU resources continuously.
