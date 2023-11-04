# `Delete a file in Go`

To delete a file in Go, you can use the `os.Remove()` function. This function takes the path of the file to be deleted as its only argument, and returns an error if the file does not exist or cannot be deleted.

To use `os.Remove()`, you must first import the `os` package. Once you have imported the `os` package, you can use the following code to delete a file:

```go
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
```

Output - 1:

```go
File deleted successfully.
```
Output - 2:

```go
remove 5_garbage.txt: The system cannot find the file specified.
```
