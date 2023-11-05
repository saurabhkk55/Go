package main

import (
	"bufio"
	"os"
)

func main() {
	// Open the file "present.txt" in write, append, and create mode (if it doesn't exist).
	// - os.O_WRONLY: Open the file for write-only access.
	// - os.O_APPEND: Open the file and move the file pointer to the end, allowing us to append data.
	// - os.O_CREATE: Create the file if it doesn't exist.
	// - 0644: Specifies the file permission (in octal notation). 0644 allows read and write for the owner and read-only for others.
	file, err := os.OpenFile("present.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err) // Immediately stop the program and display an error message.
	}
	defer file.Close() // Ensure the file is closed when we are done.

	////////////////// Write to the file as a string //////////////////

	newData := "Data to be written to the file as a string.\n"

	// Write the new data to the file.
	// The WriteString function writes the string to the file and returns the number of bytes written and an error (if any).
	_, err = file.WriteString(newData)
	if err != nil {
		panic(err) // Immediately stop the program and display an error message.
	}

	////////////////// Alternate way to write to the file as a byte slice //////////////////

	// Data to be written to the file as a byte slice
	newContent := []byte("Data to be written to the file as a byte slice.\n")

	// Write the new data to the file.
	// The Write function writes the byte slice to the file and returns the number of bytes written and an error (if any).
	_, err = file.Write(newContent)
	if err != nil {
		panic(err)
	}

	////////////////// Alternate way to write to the file using a buffered writer //////////////////

	// Create a buffered writer to write to the file efficiently
	myWriter := bufio.NewWriter(file)

	// Write the new data to the file.
	// The WriteString function writes the string to the file and returns the number of bytes written and an error (if any).
	myData := "Data to be written to the file using a buffered writer.\n"

	// Write the data to the file using the buffered writer
	_, err = myWriter.WriteString(myData)
	if err != nil {
		panic(err)
	}

	// Flush the writer to ensure data is written to the file
	err = myWriter.Flush()
	if err != nil {
		panic(err)
	}
}
