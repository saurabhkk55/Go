package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file for reading
	file, err := os.Open("learn_File.txt")
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
