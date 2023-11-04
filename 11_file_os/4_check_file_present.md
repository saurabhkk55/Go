# `Check a file is present (available) or not`

In Go, you can check if a file is present or available using the `os` package. 

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Specify the file path you want to check
    filePath := "3_example.txt"

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

        // You can access fileInfo to get file details if needed
        fmt.Printf("File Size: %d bytes\n", fileInfo.Size())
        fmt.Printf("File Mode: %s\n", fileInfo.Mode())
    }
}
```

1. You import the necessary packages, including "`os`" for file operations and "`fmt`" for printing messages.
2. You specify the `filePath` variable with the path to the file you want to check.
3. You use the `os.Stat` function to check the file's status. If the file exists, `os.Stat` returns information about the file, and `err` is nil. If the file does not exist, `os.IsNotExist` can be used to check if the error indicates that the file is not present.<br>
- When you attempt to access a file or directory using `os.Stat` and the specified path does not exist, `os.Stat` returns an error. The error type is typically of type `*os.PathError`.<br>
- The `os.IsNotExist` function is used to examine the error returned by `os.Stat` (or other file-related functions) and determine whether the error indicates that the file or directory does not exist. It checks if the error message contains the string "`no such file or directory.`"<br>
- If `os.IsNotExist` returns `true`, it means that the error indeed indicates that the file or directory does not exist, and you can handle this specific case in your code.
