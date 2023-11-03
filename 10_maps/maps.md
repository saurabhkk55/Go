# `Map`

- In Go, a map is a built-in data structure that allows you to create collections of key-value pairs.
- A map is a collection of unordered key-value pairs in Go.
- It is a hash table implementation, which means it provides efficient look-up and retrieval of values based on their keys. 
- Maps are often used to implement associative arrays or dictionaries in other programming languages.
- When a map is declared but not initialized, its zero value is nil. This means it's an empty or uninitialized map. You need to make sure to initialize a map before adding key-value pairs to it, or you'll encounter a runtime error.

Here is the syntax of a map in Go:

```go
map[KeyType]ValueType
```

- `map`: The keyword used to declare a map.
- `KeyType`: The data type of the keys in the map.
- `ValueType`: The data type of the values associated with the keys.

For example, you can declare a map of strings to integers like this:

```go
var myMap map[string]int
```

**There are several ways to create and use maps in Go**:

## 1. **`Declaration and Initialization`**:
   You can declare and initialize a map in one step using a map literal.

   ```go
   myMap := map[string]int{"one": 1, "two": 2, "three": 3}
   ```
   
   You can declare and initialize a map without using shorthand by following these steps:

```go
// Declare a map variable
var myMap map[string]int

// Initialize the map with key-value pairs using a map literal
myMap = map[string]int{"one": 1, "two": 2, "three": 3}
```

**`In the above code:`**

1. First, you declare a map variable `myMap` with the desired key and value types.

2. Then, you initialize the map by creating a map literal (a map literal is a key-value pair enclosed in curly braces `{}`). Inside the curly braces, you specify the key-value pairs for the map.

3. This initializes the `myMap` map with the specified key-value pairs.

## 2. **`Make Function`**:
You can use the `make` function to create an empty map with a specific capacity.

```go
myMap := make(map[string]int)
```

**Alternate method: Without short-hand**
```go
var myMap map[string]int
myMap = make(map[string]int)
```
	
You can use the `make` function to create an empty map with a specific capacity.

```go
myMap := make(map[string]int, 10) // This creates a map with a capacity of 10.
```

**Alternate method: Without short-hand**
```go
var myMap map[string]int
myMap = make(map[string]int, 10)
```

## 3. **`Adding Key-Value Pairs`**:
You can add new key-value pairs to a map by using the `square bracket` notation.

```go
myMap["four"] = 4
```

## 4. **`Accessing Values`**:

In Go, you can access values from a map using the square bracket notation with the key, just as you mentioned:

```go
value := myMap["two"]
```

**Alternatively: without short-hand**

```go
var value int
value = myMap["two"]
```

In this example, `myMap` is a map, and we are trying to retrieve the value associated with the key "two" and store it in the variable `value`. If the key "two" exists in the map, the value associated with it will be assigned to `value`.

It's essential to be aware that if the key is not present in the map, accessing it using this notation will return the zero value for the value type. For example, if your map contains integers, accessing a nonexistent key will return 0.

To check whether a key is present in the map before accessing its value, you can use a second return value, typically named `present` or `ok`. Here's how you can do it:

```go
value, present := myMap["nonexistent"]
```

**Alternatively: without short-hand**

```go
var value int
var present bool

value, present = myMap["nonexistent"]
```

In this code, `value` will be set to the value associated with the key "nonexistent" if it exists in the map, and `present` will be a boolean that indicates whether the key was found in the map or not. If the key is not present, `present` will be `false`, and `value` will be the zero value for the value type.

This way, you can safely check if a key exists in the map and handle the case when it doesn't without causing runtime errors.

## 5. **`Deleting Key-Value Pairs`**:
   You can delete a key-value pair from a map using the `delete` function.

   ```go
   delete(myMap, "three")
   ```

## 6. **`Iterating Over a Map`**:
   You can iterate over the keys and values in a map using a `for range` loop.

   ```go
   for key, value := range myMap {
       // Do something with key and value
   }
   ```
## 7. **`Length of Map`**
You can find the number of key-value pairs (the size) in a map by using the `len` function. Here's how you can do it:

```go
size := len(myMap)
```

**Alternatively: without short-hand**

```go
var size int
size = len(myMap)
```
