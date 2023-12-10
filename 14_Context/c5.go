package main

import (
	"context"
	"fmt"
)

// Function to mix ingredients and bake a cake
func bakeCake(ctx context.Context, flour int, sugar int) {
	// Check if we have enough flour and sugar
	if flour > 0 && sugar > 0 {
		fmt.Println("Mixing flour and sugar... 🍰")
		fmt.Println("Baking the cake! 🎂")
	} else {
		fmt.Println("Not enough ingredients to bake a cake! 😢")
	}
}

// Function to log information, using username from context
func logInformation(ctx context.Context) {
	// Retrieve username from context
	username, ok := ctx.Value("username").(string)
	if !ok {
		username = "UnknownUser"
	}

	fmt.Printf("Logging information for user: %s\n", username)
}

func main() {
	// Create a context with a username
	ctx := context.WithValue(context.Background(), "username", "JohnDoe")

	// Attempt to bake a cake
	// Note: We're passing flour and sugar as parameters, even though they are available in the context
	// This comment suggests that, even though the `flour` and `sugar` values are available in the context (because they could potentially be stored there), the code is intentionally passing them as explicit parameters to the `bakeCake` function.
	// The idea behind this comment is to emphasize that, for the sake of clarity and maintainability, it's often a good practice to pass the specific data needed for a function as parameters, even if similar data is available in the context. In the provided code, the `bakeCake` function explicitly specifies its required ingredients (flour and sugar) as parameters, making it clear what the function depends on.
	// It's a matter of making the code more readable and reducing ambiguity about where the function is getting its necessary data. If a function relies on certain data, it's generally better to see that information as part of its parameter list rather than having to check the context to understand its dependencies.
	bakeCake(ctx, 2, 1)

	// Log information using the context
	logInformation(ctx)
}

/*
Output:
$ go run 14_Context/c5.go
Mixing flour and sugar... 🍰
Baking the cake! 🎂
Logging information for user: JohnDoe
*/
