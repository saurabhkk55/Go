# Go Program with Goroutines and Channels

In this Go program, we use goroutines and channels to communicate between the `generateNumbers` and `printNumbers` functions. The `generateNumbers` function generates numbers and sends them to a channel, while the `printNumbers` function reads the numbers from the channel and prints them to the screen. We also use a `sync.WaitGroup` to coordinate the execution of these goroutines.


```go
package main

import (
	"fmt"
	"sync"
)

func generateNumbers(total int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for idx := 1; idx <= total; idx++ {
		fmt.Printf("sending %d to channel\n", idx)
		ch <- idx
	}
}

func printNumbers(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("read %d from channel\n", num)
	}
}

func main() {
	var wg sync.WaitGroup
	numberChan := make(chan int)

	wg.Add(2)
	go printNumbers(numberChan, &wg)

	generateNumbers(3, numberChan, &wg)

	close(numberChan)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}
```

######################################################################################################################################################################

Here, we have two goroutines: one for the `printNumbers` function and one for the `generateNumbers` function. Let's analyze when and why these goroutines may block and provide examples:

1. `printNumbers` Goroutine:
   The `printNumbers` goroutine reads from the channel and prints the numbers received. This goroutine will block when it tries to read from the channel and there is nothing to read. It will remain blocked until data is written to the channel or the channel is closed.

   Example:
   Suppose the `generateNumbers` function has not started yet or has not written any data to the channel, and the `printNumbers` goroutine is trying to read from it. In this case, the `printNumbers` goroutine will be blocked until data is available in the channel.

2. `generateNumbers` Function:
   The `generateNumbers` function writes numbers to the channel. It will block when it tries to write to the channel, and the channel is full. Since the channel is unbuffered in your code, it can only accept one value at a time. Therefore, if the `printNumbers` goroutine is slow in reading from the channel, the `generateNumbers` function will block when trying to write the next number.

Example:
If the `printNumbers` goroutine is slow to read from the channel, and the `generateNumbers` function is generating numbers faster than they are being read, the `generateNumbers` function will be blocked until the `printNumbers` goroutine reads from the channel to free up space for another value.

In your code, you've used a `sync.WaitGroup` to ensure that both the `printNumbers` and `generateNumbers` functions complete before closing the channel and exiting the program. This coordination is important to avoid premature termination of the program and potential data race conditions. Once both functions have completed their work (as indicated by `wg.Done()` calls), the `close(numberChan)` statement is executed, signaling to the `printNumbers` goroutine that no more data will be sent through the channel.

Here's an overview of the execution flow:

1. Both `generateNumbers` and `printNumbers` functions are started as goroutines.
2. `generateNumbers` generates and sends numbers to the channel.
3. `printNumbers` reads from the channel and prints the numbers.
4. The `close(numberChan)` statement is executed once `generateNumbers` completes, indicating that no more data will be sent through the channel.
5. The `Wait` method on the `WaitGroup` ensures that both goroutines have finished.
6. The program prints "Done!" and exits.

This synchronization ensures that the program doesn't exit before both goroutines have completed their tasks and that there are no data races or data sent on a closed channel.

######################################################################################################################################################################

Yes, that's correct. In your provided code, the `close(numberChan)` statement will execute only if the `generateNumbers` function has been executed completely. This is because you've placed the `close(numberChan)` statement after the call to `generateNumbers` and before waiting for the goroutines to complete using the `sync.WaitGroup`.

The purpose of the `sync.WaitGroup` is to ensure that both the `generateNumbers` and `printNumbers` functions have completed their tasks before proceeding to the `close(numberChan)` statement and exiting the program. This ensures that the channel is closed only when both functions have finished their work and the channel is no longer needed.

######################################################################################################################################################################
######################################################################################################################################################################
######################################################################################################################################################################

In the provided code, the goroutine `generateNumbers` will block when it attempts to send a number to the channel `numberChan` if the channel is full. This means that the `generateNumbers` goroutine will be suspended and will not be able to continue executing until there is space available in the channel.

The reason for this is that channels in Go are buffered by default. This means that they can store a limited number of values before they become full. If a goroutine attempts to send a value to a full channel, the goroutine will block until there is space available in the channel. This is done to prevent data loss and to ensure that the data is consumed in the correct order.

In the example code, the `generateNumbers` goroutine is trying to send three numbers to the `numberChan` channel. However, there is only one slot available in the channel. Therefore, the `generateNumbers` goroutine will block after sending the first number to the channel. The goroutine will remain blocked until the `printNumbers` goroutine reads the number from the channel.

Once the `printNumbers` goroutine reads the number from the channel, there will be space available in the channel again. This will allow the `generateNumbers` goroutine to continue executing and send the next number to the channel. The process will repeat until all three numbers have been sent to the channel.

After the `generateNumbers` goroutine has sent all three numbers to the channel, the `close()` function is called on the channel. This indicates that no more values will be sent to the channel. The `close()` function will also cause the `printNumbers` goroutine to stop reading from the channel. This is because the `range` loop that the `printNumbers` goroutine is using will exit when the channel is closed.

Once both goroutines have finished executing, the `main` function will print the message "Done!". This indicates that the program has finished running.

######################################################################################################################################################################
######################################################################################################################################################################
######################################################################################################################################################################

In the parameters for `generateNumbers` and `printNumbers`, you’ll see that the chan types are using the `read-only` and `write-only` types. Since `generateNumbers` only needs to be able to write numbers to the channel, it’s a `write-only` type with the `<-` arrow pointing into the channel. `printNumbers` only needs to be able to read numbers from the channel, so it’s a `read-only` type with the `<-` arrow pointing away from the channel.

Even though these types could be a `chan int`, which would allow both reading and writing, it can be helpful to constrain them to only what the function needs to avoid accidentally causing your program to stop running from something known as a `deadlock`. A deadlock can happen when one part of a program is waiting for another part of the program to do something, but that other part of the program is also waiting for the first part of the program to finish. Since both parts of the program are waiting on each other, the program will never continue running, almost like when two gears seize.

The deadlock can happen due to the way channel communication works in Go. When part of a program is writing to a channel, it will wait until another part of the program reads from that channel before continuing on. Similarly, if a program is reading from a channel it will wait until another part of the program writes to that channel before it continues. A part of a program waiting on something else to happen is said to be blocking because it’s blocked from continuing until something else happens. Channels block when they are being written to or read from. So if you have a function where you’re expecting to write to a channel but accidentally read from the channel instead, your program may enter a deadlock because the channel will never be written to. Ensuring this never happens is one reason to use a `chan<- int` or a `<-chan int` instead of just a `chan int`.

One other important aspect of the updated code is using `close()` to close the channel once it’s done being written to by `generateNumbers`. In this program, `close()` causes the `for ... range loop in printNumbers to exit`. Since using range to read from a channel continues until the channel it’s reading from is closed, if close isn’t called on `numberChan` then `printNumbers` will never finish. If `printNumbers` never finishes, the WaitGroup’s Done method will never be called by the `defer` when `printNumbers` exits. If the Done method is never called from `printNumbers`, the program itself will never exit because the WaitGroup’s Wait method in the main function will never continue. This is another example of a deadlock because the main function is waiting on something that will never happen.

######################################################################################################################################################################

It's important to use the read-only (`<-chan int`) and write-only (`chan<- int`) channel types in your function parameters and the significance of using `close()` on channels to avoid deadlocks. Let's summarize the key points:

1. **Using Read-Only and Write-Only Channel Types:**
   - `chan<- int` denotes a write-only channel. It means the function can only send data into the channel but cannot read from it.
   - `<-chan int` denotes a read-only channel. It means the function can only receive data from the channel but cannot send data into it.
   - By constraining functions to the necessary channel operations (reading or writing), you prevent accidental misuse of the channel and help avoid potential deadlocks.

2. **Deadlocks and Channel Communication:**
   - In Go, channels are a fundamental tool for communication between goroutines.
   - When one part of the program writes to a channel, it blocks until another part of the program reads from that channel, and vice versa.
   - A deadlock can occur when two parts of the program are waiting for each other, causing the program to hang indefinitely.
   - Using the correct channel types (`chan<- int` or `<-chan int`) in function parameters helps prevent such deadlocks.

3. **Using `close()` to Terminate Range Loops:**
   - In your code, you use `close(numberChan)` to signal the end of data transmission through the channel.
   - This is crucial, especially when using a `for ... range` loop to read from a channel. The loop will continue until the channel is closed.
   - If you don't close the channel, and the reading goroutine never exits, it can lead to a deadlock, as the main function is waiting for the `WaitGroup` to complete.

Source: https://www.digitalocean.com/community/tutorials/how-to-run-multiple-functions-concurrently-in-go