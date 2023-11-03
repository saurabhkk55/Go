package main

import "fmt"

func main() {
	var num int = 2

	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}

	var day string = "Sunday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday")
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend")
	default:
		fmt.Println("Not a valid day")
	}

	num = 1
	switch num {
	case 1:
		fmt.Println("This case is executed.")
		fallthrough // This leads to the next case being executed as well.
	case 2:
		fmt.Println("This case is also executed because of fallthrough.")
	case 3:
		fmt.Println("This case is not executed.")
	}
}
