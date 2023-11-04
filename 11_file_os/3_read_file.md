# `Read a file in Go`

There are two main ways to read a file in Go:

1. **Using the `os.ReadFile()` function:** This function reads the entire content of the file into a byte slice. It is the simplest way to read a file, but it can be inefficient for large files.

2. **Using the `os.Open()` function and the `bufio.NewReader()` function:** This method allows you to read the file in chunks, which is more efficient for large files.

## Example 1: Using the `os.ReadFile()` function:

```go
package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    // Read the entire file contents into a byte slice.
    data, err := os.ReadFile("3_example.txt") // Use os.ReadFile to read the file "3_example.txt"
    if err != nil {
        fmt.Println("Error reading file:", err) // If there's an error, print an error message with the details
        return
    }

    // Print the file contents.
    fmt.Println(string(data)) // Convert the byte slice to a string and print the file content
}
```

## `Learning Points`:

1. Reading the contents of a file named "3_example.txt" into a byte slice using `os.ReadFile`. The file's content is stored in the `data` variable, and any error that occurs is stored in the `err` variable.
2. We use an if statement to check if an error occurred during the file read operation. If there's an error (i.e., err is not nil), we print an error message to the console using fmt.Println. The error message will provide details about what went wrong.
3. If there is no error, we proceed to print the file contents. We do this by converting the byte slice (data) to a string using string(data). The fmt.Println function is then used to print the file content to the console.


## Example 2: Using the `os.Open()` function and the `bufio.NewReader()` function:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file for reading
	file, err := os.Open("3_example.txt")
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
```

## `Learning Points`:
1. Use `os.Open` to open a file for reading. Ensure that you check for errors when opening files, and use `defer` to close the file to release resources when done.
2. Utilize `bufio.NewReader` from `bufio` package to create a buffered reader, which can efficiently read the file line by line.
3. Loop through the file with a `for` loop, reading and processing each line until the end of the file is reached.
4. Use `reader.ReadString` to read a line up to the newline character '\n'.
5. Handle potential errors, and check if the error signifies the end of the file (EOF).
6. Remove the trailing newline character from each line to display the content neatly.
7. Print the content of each line using `fmt.Println`.