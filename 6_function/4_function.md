# `Function Recursion`

Recursion is a programming concept where a function calls itself in order to solve a problem. 
It is a powerful technique for solving complex problems by breaking them down into smaller, similar subproblems. Recursion is based on the idea of defining a problem in terms of a smaller instance of the same problem. This smaller instance is usually simpler to solve and leads to the base case, which is a situation where the problem can be solved directly without further recursion.

## `Key Components of a Recursive Function:`

1. **`Base Case`**: A condition that defines when the recursion should stop. It is the simplest form of the problem that can be solved directly.

2. **`Recursive Case`**: The part of the function where it calls itself with a modified input. This leads to breaking down the problem into smaller, more manageable subproblems.

## `Example: Calculating Factorial Using Recursion`
Let's write a Python program to calculate the factorial of a number using recursion. The factorial of a non-negative integer `n` is defined as the product of all positive integers less than or equal to `n`.

```go
package main

import "fmt"

func factorial(n int) int {
	// Base case: If n is 0 or 1, the factorial is 1.
	if n == 0 || n == 1 {
		return 1
	}

	// Recursive case: Calculate n * factorial of (n-1).
	return n * factorial(n-1)
}

func main() {
	var n, result int
	n = 5
	result = factorial(n)
	fmt.Printf("The factorial of %d is %d\n", n, result)
}
```

In this code, the `factorial` function calls itself with a smaller value `(n-1)` until it reaches the base case (n=0 or n=1), at which point it returns 1. The recursive calls are then resolved in reverse order, and the final result is computed.

**`Let's see how it iterates when calling factorial(7):`**

- `factorial(7)` enters the recursive case because `n` is not 0 or 1. It calculates `7 * factorial(6)`.
- Now, it needs to compute `factorial(6)`. It again enters the recursive case and calculates `6 * factorial(5)`.
- This process continues until `factorial(1)` is reached. At this point, the base case is satisfied, and `factorial(1)` returns 1.
- The previous call, `factorial(2)`, can now complete its calculation: `2 * 1`, and it returns 2.
- This chain of calculations continues to unwind until `factorial(7)` gets its final result: `7 * (6 * (5 * (4 * (3 * (2 * 1)))))`
- The recursion is essentially computing the product of numbers from 1 to 7, which is the definition of the factorial.

In summary, the code iterates through the recursive function by breaking the problem into smaller, similar subproblems and gradually simplifying them until the base case is reached, allowing the recursion to unwind and calculate the final result.
