# `How To Use Contexts in Go`

`Theory 1: Definition/Purpose of having Context`

When developing a large application, especially in server software, sometimes it’s helpful for a function to know more about the environment it’s being executed in aside from the information needed for a function to work on its own. For example, if a web server function is handling an HTTP request for a specific client, the function may only need to know which URL the client is requesting to serve the response. The function might only take that URL as a parameter. However, things can always happen when serving a response, such as the client disconnecting before receiving the response. If the function serving the response doesn’t know the client disconnected, the server software may end up spending more computing time than it needs calculating a response that will never be used.

In this case, being aware of the context of the request, such as the client’s connection status, allows the server to stop processing the request once the client disconnects. This saves valuable compute resources on a busy server and frees them up to handle another client’s request. This type of information can also be helpful in other contexts where functions take time to execute, such as making database calls. To enable ubiquitous access to this type of information, Go has included a context package in its standard library.

`Explanation 1:`

Imagine you're a server at a busy restaurant. You take orders from customers (functions) and deliver their food (responses).

Sometimes, customers leave the restaurant before their food arrives (client disconnects). It would be silly for you to keep cooking their food (wasting resources) if you knew they were gone.

The context package in Go is like a waiter who tells you if a customer has left. This way, you can stop cooking their food and focus on other customers' orders.

This saves valuable resources (compute time) and makes the server more efficient. This is especially helpful for functions that take a long time to execute, like database calls.

So, the context package helps functions "know their place" and avoid wasting resources by being aware of the wider environment they're working in. It's like having extra eyes and ears to make your server run smoothly.

<div style="border: 10px solid white;"></div>

`Theory 2:`

Many functions in Go use the context package to gather additional information about the environment they’re being executed in, and will typically provide that context to the functions they also call. By using the context.Context interface in the context package and passing it from function to function, programs can convey that information from the beginning function of a program, such as main, all the way to the deepest function call in the program. The Context function of an http.Request, for example, will provide a context.Context that includes information about the client making the request and will end if the client disconnects before the request is finished. By passing this context.Context value into a function that then makes a call to the QueryContext function of a sql.DB, the database query will also be stopped if it’s still running when the client disconnects.

`Explanation 2:`

Imagine you're working on a big construction project. You have various teams working on different parts of the building.

* **The architect (main function):** Starts the whole project and provides overall plans and context to everyone.
* **Individual workers (functions):** Each team member has a specific task, like building a wall or installing windows. They need some information from the architect to do their job, like the blueprints and deadlines.
* **The foreman (context.Context):** The foreman acts like a messenger between the architect and the workers. They carry information about the project, deadlines, and any changes to the plan.

In the context of Go, the context package acts like the foreman. It carries information about the program's environment, such as deadlines and cancellations, and passes it along to functions as they are called. This allows functions to be aware of the bigger picture and adjust their behavior accordingly.

Here's an example:

* You have a function that downloads a file from the internet.
* This function uses the context to check if the user has canceled the download.
* If the user cancels, the function stops downloading the file and avoids wasting resources.

Another example:

* You have a function that queries a database for information.
* This function uses the context to check if the client has disconnected.
* If the client disconnects, the function stops the database query and avoids unnecessary processing.

By using the context package, you can:

* **Make your programs more efficient:** By stopping functions that are no longer needed, you can free up resources for other tasks.
* **Make your programs more robust:** By being aware of potential errors and cancellations, you can prevent your programs from crashing.
* **Make your programs more maintainable:** By having a clear way to pass information between functions, you can make your code easier to understand and modify.

So, the context package is like a communication hub that helps all parts of your program work together smoothly and efficiently. It allows functions to be "informed" citizens that contribute to the overall success of the project.
