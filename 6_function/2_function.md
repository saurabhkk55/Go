# `Pass-by-value and Pass-by-reference in Go functions`

## 1. `Pass-by-Value`:
   When you pass a variable to a function, you're making a copy of the value. Any changes made to this copy do not affect the original variable outside of the function.

   ```go
   package main

   import "fmt"

   // modifyValue takes an integer by value.
   // It receives a copy of the original value.
   func modifyValue(x int) {
       // Modifying the copy, not the original.
       x = 10
   }

   func main() {
       value := 5
       // Call modifyValue with 'value,' which is passed by value.
       modifyValue(value)
       // 'value' remains unchanged outside of the function.
       fmt.Println(value) // Output: 5
   }
   ```

## 2. `Pass-by-Reference (using Pointers)`:
   In Go, you can use pointers to achieve pass-by-reference behavior. When you pass a pointer to a function, you're still passing a copy of the pointer, but this copy points to the same memory location as the original variable. Therefore, changes made through the pointer affect the original variable.

   ```go
   package main

   import "fmt"

   // modifyValueByReference takes a pointer to an integer.
   // It can modify the value pointed to by the pointer.
   func modifyValueByReference(x *int) {
       // Modifying the value pointed to by the pointer.
       *x = 10
   }

   func main() {
       value := 5
       // Call modifyValueByReference with a pointer to 'value.'
       // Changes made through the pointer affect the original 'value.'
       modifyValueByReference(&value)
       fmt.Println(value) // Output: 10
   }
   ```

In the second example, the function `modifyValueByReference` receives a copy of the pointer to the original variable. It dereferences the pointer using `*x` to modify the value it points to, which affects the original `value` variable.

### `Summary`:
Pass-by-value creates a copy of the value, and pass-by-reference (via pointers) allows you to modify the original variable indirectly.