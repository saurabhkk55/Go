package main

import (
	"encoding/json"
	"fmt"
)

// Define a struct type called Company
type Company struct {
	Name     string `json:"naam"`
	Location string `json:"jagah"`
	MobNo    int    `json:"ghanti"`
}

func main() {
	var comp Company    // Declare variable of Company type or Declare Company struct instance
	var jsonData []byte // Declare variable of byte type
	var err error       // Declare variable of error type

	comp = Company{"ABC", "Moon", 9099}
	fmt.Printf("\n'comp' info: %v", comp)
	fmt.Printf("\n'comp' type: %T", comp)

	// Encode the Company struct instance to JSON bytes
	jsonData, err = json.Marshal(comp)
	if err != nil {
		fmt.Printf("Error: While encoding to JSON: %v", err)
		return
	}

	fmt.Printf("\n\nData type of 'jsonData': %T", jsonData)
	fmt.Println("\nOriginal content of 'jsonData': ", jsonData)
	fmt.Println("Encoded JSON data: ", string(jsonData))

	// Decode the JSON encoded bytes back to a Company struct instance
	var decodedComp Company
	err = json.Unmarshal(jsonData, &decodedComp)
	if err != nil {
		fmt.Printf("ERROR : While decoding to struct : %v", err)
		return
	}

	// Print the decoded Company struct instance
	fmt.Printf("\n\n'decodedComp' info: %v", decodedComp)
	fmt.Printf("\n'decodedComp' type: %T", comp)
	fmt.Println("\nExtracting fields from 'decodedComp':")
	fmt.Println("NAME :", decodedComp.Name)
	fmt.Println("LOCATION :", decodedComp.Location)
	fmt.Println("MOBNO :", decodedComp.MobNo)
}

/*
saura@DESKTOP-GC3SDTN MINGW64 ~/OneDrive/Desktop/GO (main)
$ go run ./9_struct/struct-programs/p2.go

'comp' info: {ABC Moon 9099}
'comp' type: main.Company

Data type of 'jsonData': []uint8
Original content of 'jsonData':  [123 34 110 97 97 109 34 58 34 65 66 67 34 44 34 106 97 103 97 104 34 58 34 77 111 111 110 34 44 34 103 104 97 110 116 105 34 58 57 48 57 57 125]
Encoded JSON data:  {"naam":"ABC","jagah":"Moon","ghanti":9099}


'decodedComp' info: {ABC Moon 9099}
'decodedComp' type: main.Company
Extracting fields from 'decodedComp':
NAME : ABC
LOCATION : Moon
MOBNO : 9099
*/
