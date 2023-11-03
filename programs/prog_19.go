// Patterns
package main

import "fmt"

func main() {
	var x, y int

	for x = 4; x >= 1; x-- {
		for y = 1; y <= x; y++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
	// Output
	// * * * *
	// * * *
	// * *
	// *

	fmt.Print("\n")

	for x = 1; x <= 4; x++ {
		for y = 1; y <= x; y++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
	// Output
	// *
	// * *
	// * * *
	// * * * *

	fmt.Print("\n")

	for x = 1; x <= 4; x++ {
		for y = 0; y < 4; y++ {
			fmt.Printf("%c ", 'A'+byte(y))
		}
		fmt.Print("\n")
	}
	// Output
	// A B C D
	// A B C D
	// A B C D
	// A B C D

	fmt.Print("\n")

	for x = 1; x <= 4; x++ {
		for y = 0; y < 4; y++ {
			fmt.Printf("%c ", 'D'-byte(y))
		}
		fmt.Print("\n")
	}
	// Output
	// D C B A
	// D C B A
	// D C B A
	// D C B A

	fmt.Print("\n")

	var item byte
	for x = 0; x < 4; x++ {
		item = 'A' + byte(x)
		for y = 0; y < 4; y++ {
			fmt.Printf("%c ", item)
		}
		fmt.Print("\n")
	}
	// Output
	// A A A A
	// B B B B
	// C C C C
	// D D D D

	fmt.Print("\n")

	for x = 0; x < 4; x++ {
		item = 'D' - byte(x)
		for y = 0; y < 4; y++ {
			fmt.Printf("%c ", item)
		}
		fmt.Print("\n")
	}
	// Output
	// D D D D
	// C C C C
	// B B B B
	// A A A A

	fmt.Print("\n")

	for x = 4; x > 0; x-- {
		for y = 0; y < x; y++ {
			fmt.Printf("%c ", 'A'+byte(y))
		}
		fmt.Print("\n")
	}
	// Output
	// A B C D
	// A B C
	// A B
	// A

	fmt.Print("\n")

	for x = 0; x < 4; x++ {
		item = 'A' + byte(x)
		for y = x; y < 4; y++ {
			fmt.Printf("%c ", item)
		}
		fmt.Print("\n")
	}
	// Output
	// A A A A
	// B B B
	// C C
	// D

	fmt.Print("\n")

	for x = 4; x > 0; x-- {
		for y = 0; y < x; y++ {
			fmt.Printf("%c ", 'D'-byte(y))
		}
		fmt.Print("\n")
	}
	// Output
	// D C B A
	// D C B
	// D C
	// D

	fmt.Print("\n")

	for x = 0; x < 4; x++ {
		item = 'D' - byte(x)
		for y = x; y < 4; y++ {
			fmt.Printf("%c ", item)
		}
		fmt.Print("\n")
	}
	// Output
	// D D D D
	// C C C
	// B B
	// A

	fmt.Print("\n")

	for x = 0; x < 4; x++ {
		for y = 0; y <= x; y++ {
			fmt.Printf("%c ", 'A'+byte(y))
		}
		fmt.Print("\n")
	}
	// Output
	// A
	// A B
	// A B C
	// A B C D

	fmt.Print("\n")

	for x = 0; x < 4; x++ {
		item = 'A' + byte(x)
		for y = 0; y <= x; y++ {
			fmt.Printf("%c ", item)
		}
		fmt.Print("\n")
	}
	// Output
	// A
	// B B
	// C C C
	// D D D D

	fmt.Print("\n")

	var counter int
	counter = 0

	for x = 0; x < 4; x++ {
		item = 'A' + byte(counter)
		for y = 0; y < 4; y++ {
			fmt.Printf("%c ", item+byte(y))
			counter++
		}
		fmt.Print("\n")
	}
	// Output
	// A B C D
	// E F G H
	// I J K L
	// M N O P

	fmt.Print("\n")

	for x = 0; x < 4; x++ {
		item = 'D' - byte(x)
		for y = x; y >= 0; y-- {
			fmt.Printf("%c ", item)
		}
		fmt.Print("\n")
	}
	// Output
	// D
	// C C
	// B B B
	// A A A A

	fmt.Print("\n")

	counter = 0
	for x = 4; x > 0; x-- {
		item = 'A' + byte(counter)
		for y = 0; y < x; y++ {
			fmt.Printf("%c ", item+byte(y))
			counter++
		}
		fmt.Print("\n")
	}
}
