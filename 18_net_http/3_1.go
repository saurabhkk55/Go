// Running Multiple Servers at One Time

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// getRoot handles requests to the root ("/") endpoint.
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

// getHello handles requests to the "/hello" endpoint.
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

// getAnotherEndpoint handles requests to the "/another" endpoint.
func getAnotherEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /another endpoint request\n")
	io.WriteString(w, "This is another endpoint!\n")
}

// startServer starts an HTTP server on the specified address with the provided multiplexer.
// It listens for a shutdown signal on the given context.
func startServer(addr string, mux *http.ServeMux, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	// Create an HTTP server with the provided address and multiplexer.
	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// Goroutine to handle graceful shutdown when the context is canceled.
	go func() {
		<-ctx.Done() // Wait for the context to be canceled (e.g., on signal reception).
		fmt.Printf("Shutting down server on %s\n", addr)
		server.Shutdown(context.Background()) // Shutdown the server gracefully.
	}()

	// Start the HTTP server.
	fmt.Printf("Starting server on %s\n", addr)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error starting server on %s: %s\n", addr, err)
	}
}

func main() {
	// Create a context and a cancellation function to manage server shutdown.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure cancellation is called when main exits.

	// Create a multiplexer for the first set of endpoints.
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", getRoot)
	mux1.HandleFunc("/hello", getHello)

	// Create a multiplexer for the second set of endpoints.
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/another", getAnotherEndpoint)

	// Use WaitGroup to wait for all servers to finish before exiting the program.
	var wg sync.WaitGroup
	wg.Add(2)

	// Start the first server in a goroutine.
	go startServer(":3333", mux1, &wg, ctx)
	// Start the second server in another goroutine.
	go startServer(":4444", mux2, &wg, ctx)

	// Handle graceful shutdown on SIGINT and SIGTERM signals.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-sigCh:
		// Received a signal, trigger context cancellation to shut down servers.
		fmt.Printf("Received signal: %v. Shutting down...\n", sig)
		cancel()
	}

	// Wait for all servers to finish before exiting the program.
	wg.Wait()
	fmt.Println("All servers gracefully shut down.")
}

/*
Certainly! Let's go through how the `sync.WaitGroup` (denoted as `wg`) is used in this code:

1. **WaitGroup Creation:**
   ```go
   var wg sync.WaitGroup
   wg.Add(2)
   ```
   - A `sync.WaitGroup` is created to keep track of the number of goroutines that need to finish before the program can exit.
   - The initial count is set to 2 using `wg.Add(2)`, indicating that there are two goroutines (servers) that need to complete.

2. **Goroutine Launching:**
   ```go
   go startServer(":3333", mux1, &wg, ctx)
   go startServer(":4444", mux2, &wg, ctx)
   ```
   - Two goroutines are launched concurrently to start the HTTP servers using the `startServer` function.
   - The `&wg` is a pointer to the `WaitGroup`, allowing the `startServer` function to decrement the count when it completes.

3. **Server Shutdown Handling:**
   ```go
   go func() {
       <-ctx.Done()
       fmt.Printf("Shutting down server on %s\n", addr)
       server.Shutdown(context.Background())
   }()
   ```
   - Within each `startServer` goroutine, a goroutine is spawned to wait for the cancellation of the context (`<-ctx.Done()`).
   - When the context is canceled (e.g., on receiving a signal), it triggers the graceful shutdown of the respective server (`server.Shutdown`).
   - After the server is shut down, the `WaitGroup` count is decremented using `defer wg.Done()`.

4. **Signal Handling in Main:**
   ```go
   select {
   case sig := <-sigCh:
       fmt.Printf("Received signal: %v. Shutting down...\n", sig)
       cancel() // Trigger context cancellation
   }
   ```
   - The main goroutine waits for a signal (SIGINT or SIGTERM) on the `sigCh` channel.
   - When a signal is received, it triggers the cancellation of the context (`cancel()`), which initiates the shutdown process for all servers.

5. **Waiting for Goroutines to Finish:**
   ```go
   wg.Wait()
   ```
   - The `WaitGroup` is used to block the main goroutine until the count becomes zero.
   - Each call to `wg.Done()` (deferred in `startServer`) decrements the count, and when both servers have completed, the `Wait` call unblocks the main goroutine.

6. **Print Statement after Waiting:**
   ```go
   fmt.Println("All servers gracefully shut down.")
   ```
   - Once the `WaitGroup` count becomes zero, the program prints a message indicating that all servers have gracefully shut down.

In summary, the `sync.WaitGroup` ensures that the main goroutine waits for the completion of the two server goroutines before allowing the program to exit. It helps synchronize the concurrent execution and provides a mechanism to wait for the completion of parallel tasks.
*/
