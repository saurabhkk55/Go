package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		// Handle the error if getting the current working directory fails
		panic(err)
	}

	fmt.Println("currentDir: ", currentDir)

	// Specify the path of the file you want to create within the current working directory
	filePath := currentDir + "/emptyFile.txt"

	fmt.Println("filePath: ", filePath)

	// Create the empty file
	file, err := os.Create(filePath)
	if err != nil {
		// If there's an error while creating the file, we use panic to
		// immediately stop the program and display an error message.
		panic(err)
	}

	// Defer the file.Close() function call to ensure the file is properly closed
	// just before the main function returns, even in the presence of errors.
	defer file.Close()

	fmt.Println("File is now created in the current working directory and ready for use.")
}
