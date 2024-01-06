// Running Multiple Servers at One Time

package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func getAnotherEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /another endpoint request\n")
	io.WriteString(w, "This is another endpoint!\n")
}

func main() {
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", getRoot)
	mux1.HandleFunc("/hello", getHello)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/another", getAnotherEndpoint)

	go func() {
		err := http.ListenAndServe(":3333", mux1)
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server 1 closed\n")
		} else if err != nil {
			fmt.Printf("Error starting server 1: %s\n", err)
			os.Exit(1)
		}
	}()

	go func() {
		err := http.ListenAndServe(":4444", mux2)
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server 2 closed\n")
		} else if err != nil {
			fmt.Printf("Error starting server 2: %s\n", err)
			os.Exit(1)
		}
	}()

	// Wait for some signal to exit, for example, Ctrl+C
	select {}
}
