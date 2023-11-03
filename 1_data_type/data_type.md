# 1. **`Numeric Types:`**
   Go has several numeric data types, including:
   - `int`: This is a signed integer type. Its size depends on the architecture of the machine (32-bit or 64-bit).
   - `int8`, `int16`, `int32`, `int64`: These are signed integer types with specific bit sizes.
   - `uint`: This is an unsigned integer type, which means it cannot store negative values.
   - `uint8`, `uint16`, `uint32`, `uint64`: These are unsigned integer types with specific bit sizes.
   - `float32`, `float64`: These are floating-point types for representing real numbers.

   Example:
   ```go
   var age int = 30
   var temperature float64 = 98.6
   ```

1. **Boolean Type:**
   Go has a `bool` type, which can only have two values: `true` or `false`. It's used for conditions and logic.

   Example:
   ```go
   var isSunny bool = true
   ```

2. **String Type:**
   Go represents text as strings. Strings are made up of a sequence of characters. They are immutable, meaning once created, you can't change their content.

   Example:
   ```go
   var name string = "Alice"
   ```

3. **Composite Types:**
   Go provides composite types that can hold multiple values.
   - **Arrays**: Arrays have a fixed size and contain elements of the same data type.

     Example:
     ```go
     var numbers [3]int = [3]int{1, 2, 3}
     ```

   - **Slices**: Slices are more flexible than arrays. They are like dynamic arrays that can grow or shrink.

     Example:
     ```go
     var scores []int = []int{95, 85, 75}
     ```

   - **Maps**: Maps are key-value pairs, similar to dictionaries in other languages.

     Example:
     ```go
     var person map[string]string = map[string]string{
         "name": "Bob",
         "age": "25",
     }
     ```

   - **Structs**: Structs allow you to create your own composite data types. They are made up of fields with different data types.

     Example:
     ```go
     type Person struct {
         Name string
         Age  int
     }

     var someone Person = Person{"Alice", 30}
     ```

4. **Pointers:**
   Go supports pointers, which are variables that store memory addresses. They are used to refer to the memory location of a value.

   Example:
   ```go
   var x int = 10
   var ptr *int = &x  // ptr is a pointer to the memory location of x
   ```

5. **Other Types:**
   Go also provides more specialized data types such as complex numbers (`complex64` and `complex128`) for advanced mathematical operations and `byte` (an alias for `uint8`) for handling binary data.

6. **Type Conversion:**
   You can convert values from one data type to another using type casting.

   Example:
   ```go
   var number int = 42
   var result float64 = float64(number)  // Convert int to float64
   ```
