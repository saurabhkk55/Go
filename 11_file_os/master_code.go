package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var data string

	// Create a file and get its path
	createdFile, filePath := createFile()
	fmt.Println("File created:", filePath)
	fmt.Println("########################################################")

	data = "This is some data that we are writing to a file in Go.\n"
	writeToFile(createdFile, data)
	fmt.Println("Data written to the file.")
	fmt.Println("########################################################")

	fmt.Println("Reading the file content:")
	readFile(filePath)
	fmt.Println("########################################################")

	data = "1 2 3 4 5 6 7 8 9 0.\n"
	writeToFile(createdFile, data)
	fmt.Println("Data written to the file.")
	fmt.Println("########################################################")

	fmt.Println("Reading the updated file content:")
	readFile(filePath)
	fmt.Println("########################################################")

	// Sleep for 2 seconds
	duration := 2 * time.Second // 2 * 1 => 2 seconds
	time.Sleep(duration)

	// Delete the file
	deleteFile(createdFile, filePath)
	fmt.Println("########################################################")
}

func createFile() (*os.File, string) {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		// Handle the error if getting the current working directory fails
		panic(err)
	}

	fmt.Println("currentDir: ", currentDir)

	// Specify the path of the file you want to create within the current working directory
	filePath := currentDir + "/learn_File.txt"

	fmt.Println("filePath: ", filePath)

	// Create the empty file
	myFile, err := os.Create(filePath)
	if err != nil {
		// If there's an error while creating the file, we use panic to
		// immediately stop the program and display an error message.
		panic(err)
	}

	fmt.Println("File is now created in the current working directory and ready for use.")

	return myFile, filePath
}

func writeToFile(myEmptyFile *os.File, data string) *os.File {
	// Create a buffered writer (myWriter) to write to the file efficiently
	myWriter := bufio.NewWriter(myEmptyFile)

	// Data to be written to the file
	myData := data

	// Write the data to the file using the buffered writer
	_, err := myWriter.WriteString(myData)
	if err != nil {
		panic(err)
	}

	// Flush the writer to ensure data is written to the file
	err = myWriter.Flush()
	if err != nil {
		panic(err)
	}

	return myEmptyFile
}

func readFile(my_file_Path string) {
	var file *os.File
	var err error

	file, err = os.Open(my_file_Path)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when we're done

	// Create a buffered reader to read the file line by line
	reader := bufio.NewReader(file)

	// Loop to read and process each line of the file
	for {
		// Read a line (up to the '\n' character)
		line, err := reader.ReadString('\n')
		if err != nil {
			// Check for the end of the file
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading file:", err)
			break
		}

		// Remove the trailing newline character
		line = line[:len(line)-1]

		// Print the content of each line
		fmt.Println(line)
	}
}

func deleteFile(myEmptyFile *os.File, my_filePath string) {

	// Close the file before deleting
	myEmptyFile.Close()

	// Delete the file.
	err := os.Remove(my_filePath)

	if err != nil {
		// Handle the error.
		fmt.Println(err)
	} else {
		fmt.Println("File deleted successfully.")
	}
}
