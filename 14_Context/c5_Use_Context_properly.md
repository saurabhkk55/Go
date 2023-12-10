# `How to use context in the code properly`

Note: Contexts can be a powerful tool with all the values they can hold, but a balance needs to be struck between data being stored in a context and data being passed to a function as parameters. It may seem tempting to put all of your data in a context and use that data in your functions instead of parameters, but that can lead to code that is hard to read and maintain. A good rule of thumb is that any data required for a function to run should be passed as parameters. Sometimes, for example, it can be useful to keep values such as usernames in context values for use when logging information for later. However, if the username is used to determine if a function should display some specific information, you’d want to include it as a function parameter even if it’s already available from the context. This way when you, or someone else, looks at the function in the future, it’s easier to see which data is actually being used.

## `Explanation`:

Imagine you're baking a cake. In this analogy:

- **`Context`** is like your kitchen. It's where you have all your tools and ingredients.
- **`Function`** is like a specific task, such as mixing ingredients or putting the cake in the oven.

Now, let's talk about parameters and why we need a balance:

1. **`Context (Kitchen)`:** You have all your ingredients and tools here. Putting everything in the kitchen makes it cluttered, just like putting too much data in the context can make your code messy.

2. **`Function (Task)`:** Each task (function) needs specific ingredients (data) to be performed. If you put everything in the kitchen (context) and don't specify which ingredients (parameters) are needed for each task (function), it can get confusing.

3. **`Balance`:** It's like deciding what to keep in the kitchen (context) and what to bring to each specific task (function). Too much in the kitchen makes it chaotic, but too little means you might not have what you need for a task.

4. **`Rule of Thumb`:** If a task (function) needs certain ingredients (data) to be performed, it's better to bring those ingredients with you (pass them as parameters) to that task, even if they're already in the kitchen (context). It makes it clear what each task is using.

5. **`Example`:** If you have a cake recipe (function) that needs flour and sugar, you bring those specific ingredients (parameters) to the task, even though your kitchen (context) already has them.

6. **`Logging Information`:** Sometimes, you might want to keep certain info in the kitchen (context), like a recipe book (usernames), but if a specific task (function) needs that info, it's better to bring it along as a parameter.

So, in coding, it's about organizing your data. Too much in one place can make it messy, but you still need to bring the right data to the right part of your code when it's needed.


Let's consider a simple code. We'll create a basic scenario where a user is trying to bake a cake with functions and contexts.

```go
package main

import (
	"fmt"
	"context"
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
```

```go
Output:
$ go run 14_Context/c5.go
Mixing flour and sugar... 🍰
Baking the cake! 🎂
Logging information for user: JohnDoe
```
**`Code explanation`**:

- We import necessary packages, including the `context` package.
- The `bakeCake` function represents a task (like baking a cake) that requires specific ingredients (`flour` and `sugar`). We pass these ingredients as parameters.
- The `logInformation` function logs information and uses the `username` stored in the context. We retrieve the `username` from the context and use it as a parameter to show how it's being used explicitly.
- In the `main` function, we create a context with a username (`JohnDoe`), and then we use this context to bake a cake and log information. We pass the required ingredients explicitly to the `bakeCake` function, even though they are available in the context.

This example demonstrates the idea of bringing only the necessary data (ingredients) to each task (function) and using the context for shared information (like the username) across functions.
