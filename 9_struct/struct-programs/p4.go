package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name  string
	Class int
}

func main() {
	var input string

	file := play_with_file()
	fmt.Print("\n----------------------------------\n\n")

	for {
		kid_info := encode_decode()
		fmt.Print("\n----------------------------------\n\n")
		store_struct_to_file(kid_info, file)
		fmt.Print("\n----------------------------------\n\n")

		fmt.Print("Do you want to enter information for another student? (y/n): ")
		fmt.Scan(&input)

		if input != "y" {
			break
		}
	}
}

func store_struct_to_file(student_info Student, file_for_student_records *os.File) {
	// Create a JSON encoder
	encoder := json.NewEncoder(file_for_student_records)

	err := encoder.Encode(student_info)
	if err != nil {
		fmt.Println("Error encoding struct:", err)
		return
	}

	fmt.Println("Struct successfully written to file.")
}

func play_with_file() *os.File {
	current_working_directory, err := os.Getwd()
	if err != nil {
		fmt.Printf("ERROR : %v", err)
		panic(err)
	}
	fmt.Println("My pwd: ", current_working_directory)

	newfile := current_working_directory + "/hello.txt"
	fmt.Println("newfile: ", newfile)

	file, err := os.Create(newfile)
	if err != nil {
		panic(err)
	}
	// defer file.Close()

	fmt.Println("File is ready.")

	return file

}

func encode_decode() Student {
	var kid1 Student

	fmt.Print("Enter student info:")
	fmt.Print("\nEnter name: ")
	fmt.Scan(&kid1.Name)
	fmt.Print("Enter class: ")
	fmt.Scan(&kid1.Class)

	return kid1
}

/*
DRABACK:
Content of file will be delete on every execution of program (go run p4.go)
*/
