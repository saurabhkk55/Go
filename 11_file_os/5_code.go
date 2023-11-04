package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the path of the file to be deleted.
	// filename := "/path/to/file.txt"
	filename := "5_garbage.txt"

	// Delete the file.
	err := os.Remove(filename)

	if err != nil {
		// Handle the error.
		fmt.Println(err)
	} else {
		fmt.Println("File deleted successfully.")
	}
}
