package main

import (
	"encoding/json"
	"fmt"
)

// User is a struct representing user data.
type User struct {
	ID       int
	UserName string
	Email    string
}

func main() {
	// Create a User instance
	var u1 User
	u1 = User{101, "Ratan", "ratantata@gmail.com"}

	// Marshal the User instance to JSON
	jsonData, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("Error encoding to JSON: %v\n", err)
		return
	}

	// Print the JSON data (Encoded data)
	/*
		In the context of this Go program, "encoded" refers to the process of converting a Go data structure,
		in this case, a User struct, into a JSON representation.
		Encoding is the process of converting data from one format to another so that it can be easily transmitted or stored and then later decoded (or "decoded") to its original format.
	*/
	fmt.Printf("Data type of 'jsonData': %T\n", jsonData)
	fmt.Print("\n")
	fmt.Println("Original content of 'jsonData': ", jsonData)
	fmt.Print("\n")
	fmt.Printf("Encoded JSON data:\n%s\n", string(jsonData))
	fmt.Print("\n")

	// Display the original User data
	fmt.Println("User u1 Information: ", u1)
	fmt.Print("\n")

	fmt.Print("##################################################################\n\n")

	// Decode JSON data back to a User struct
	var decodedUser User
	err = json.Unmarshal(jsonData, &decodedUser)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	// Display the decoded User struct
	fmt.Printf("Decoded User u1:\nID: %d\nUserName: %s\nEmail: %s\n", decodedUser.ID, decodedUser.UserName, decodedUser.Email)
}
