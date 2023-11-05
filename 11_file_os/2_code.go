package main

import "os"

func main() {
	// Specify the file path
	filePath := "2_example.txt"

	// Create or truncate the file, and handle any errors
	myFile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer myFile.Close() // Ensure the file is closed when done

	// Define the content to be written as a string
	newContent := "This is the new content that will override the file.\n"

	// Write the content to the file
	_, err = myFile.WriteString(newContent)
	if err != nil {
		panic(err)
	}

	// Add more data to the file
	newContent = "1 2 3 4 5 6 7 8 9 0\n"
	_, err = myFile.WriteString(newContent)
	if err != nil {
		panic(err)
	}
}
