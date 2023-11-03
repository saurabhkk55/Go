## `Case 8: JSON Serialization`
Structs can be used to map to JSON objects, allowing you to encode and decode data easily.

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Product struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    // Create a Product instance
    var product Product
    product.ID = 1
    product.Name = "Example Product"

    // Encoding the Product struct to JSON
    var jsonData []byte
    var err error
    jsonData, err = json.Marshal(product)
    if err != nil {
        fmt.Printf("Error encoding JSON: %v\n", err)
        return
    }

    // Display the JSON data as a string
    fmt.Printf("Encoded JSON data:\n%s\n", string(jsonData))

    // Decoding JSON data back to a Product struct
    var decodedProduct Product
    err = json.Unmarshal(jsonData, &decodedProduct)
    if err != nil {
        fmt.Printf("Error decoding JSON: %v\n", err)
        return
    }

    // Display the decoded Product struct
    fmt.Printf("Decoded Product:\nID: %d\nName: %s\n", decodedProduct.ID, decodedProduct.Name)
}
```

**`Output:`**

```go
Encoded JSON data:
{"id":1,"name":"Example Product"}

Decoded Product:
ID: 1
Name: Example Product
```
