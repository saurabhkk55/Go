package main

import (
	"fmt"
	"os"
)

func main() {
	// Specify the file path you want to check
	filePath := "4_example.txt"

	// Use os.Stat to get file information and check for errors
	fileInfo, err := os.Stat(filePath)

	// Check if there was an error while accessing the file
	if err != nil {
		// Check if the error indicates that the file does not exist
		if os.IsNotExist(err) {
			fmt.Printf("File %s does not exist.\n", filePath)
		} else {
			// Handle other errors, if any
			fmt.Printf("Error checking file: %v\n", err)
		}
	} else {
		// File exists, and fileInfo contains information about the file
		fmt.Printf("File %s is available.\n", filePath)

		// Access fileInfo to get file details if needed
		fmt.Printf("File Size: %d bytes\n", fileInfo.Size())
		fmt.Printf("File Mode: %s\n", fileInfo.Mode())
	}
}
