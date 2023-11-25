package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Define a struct named `Student` to store student information
type Student struct {
	Name  string
	Class int
}

func main() {
	var input string

	// Open a file for student records
	file := play_with_file()
	fmt.Print("\n----------------------------------\n\n")

	// loop to enter and store student information
	for {
		kid_info := encode_decode()
		fmt.Print("\n----------------------------------\n\n")
		store_struct_to_file(kid_info, file)
		fmt.Print("\n----------------------------------\n\n")

		fmt.Print("Do you want to enter information for another student? (y/n): ")
		fmt.Scan(&input)

		// Close the file when not needed anymore
		if input != "y" {
			defer file.Close()
			break
		}
	}
}

// store_struct_to_file encodes a Student struct to JSON and stores it in a file.
func store_struct_to_file(student_info Student, file_for_student_records *os.File) {
	// Create a JSON encoder
	encoder := json.NewEncoder(file_for_student_records)

	// Encode the studentInfo struct to JSON and write it to the file
	err := encoder.Encode(student_info)
	if err != nil {
		fmt.Println("Error encoding struct:", err)
		return
	}

	fmt.Println("Struct successfully written to file.")
}

// play_with_file creates or opens a file for student records.
func play_with_file() *os.File {
	// Get the current working directory
	current_working_directory, err := os.Getwd()
	if err != nil {
		fmt.Printf("ERROR : %v", err)
		panic(err)
	}
	fmt.Println("My pwd: ", current_working_directory)

	// Create a new file path
	newfile := current_working_directory + "/hello.txt"
	fmt.Println("newfile: ", newfile)

	// Open or create the file for writing and appending
	file, err := os.OpenFile(newfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("File is ready.")
	return file

}

// encode_decode prompts the user to enter student information and returns a Student struct.
func encode_decode() Student {
	var kid1 Student

	fmt.Print("Enter student info:")
	fmt.Print("\nEnter name: ")
	fmt.Scan(&kid1.Name)
	fmt.Print("Enter class: ")
	fmt.Scan(&kid1.Class)

	return kid1
}
