## **Predeclared Types**
### **Understanding Go’s Built-in Types**
Go has several built-in types that programmers use daily, such as booleans, integers, floats, and strings. While these might seem familiar if you've worked with other languages, Go has some unique rules and behaviors that you should know.
### **Zero Value Concept**
In Go, when you declare a variable but don’t explicitly assign it a value, it gets a default value called the **zero value**. This prevents unexpected behavior caused by uninitialized variables, which is a common issue in languages like C and C++.
- **bool** → `false`
- **int, float** → `0`
- **string** → `""` (empty string)
- **pointer, slice, map, channel, function, interface** → `nil`
## **Literals in Go**
Literals are hardcoded values that appear directly in your code, like numbers or strings.
### **Integer Literals**
- Written as simple numbers: `123`
- Can use underscores for readability: `1_000_000`
- Support different bases:
    - **Decimal (Base 10)**: `42`
    - **Binary (Base 2)**: `0b1010`
    - **Octal (Base 8)**: `0o755`
    - **Hexadecimal (Base 16)**: `0x2A`
### **Floating-Point Literals**
- Contain a decimal point: `3.14`
- Use scientific notation with `e`: `6.02e23`
- Can be written in hexadecimal: `0x1.2p3` (which means `1.2 × 2^3`)
### **Rune Literals**
A **rune** represents a Unicode character (similar to a `char` in other languages). It is written inside single quotes:
- `'a'` (normal character)
- `'\n'` (new line)
- `'\u03A9'` (Ω, Unicode representation)
### **String Literals**
- **Interpreted strings**: Use **double quotes** `"Hello\nWorld"` (supports escape sequences like `\n` for newline)
- **Raw strings**: Use **backticks** ``Hello\nWorld`` (keeps everything as-is, including newlines)
## **Boolean Type**
- Can be `true` or `false`
- The zero value is `false`
- Used in conditionals (`if`, `for`, etc.)
Example:
```go
var flag bool // Defaults to false
flag = true
```
## **Numeric Types**
Go has a variety of **integer** and **floating-point** types.
### **Integer Types**
Go provides both **signed** (can be negative) and **unsigned** (only positive) integers.
| Type | Size | Range |
| --- | --- | --- |
| `int8` | 1 byte | -128 to 127 |
| `int16` | 2 bytes | -32,768 to 32,767 |
| `int32` | 4 bytes | -2,147,483,648 to 2,147,483,647 |
| `int64` | 8 bytes | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |
| `uint8` | 1 byte | 0 to 255 |
| `uint16` | 2 bytes | 0 to 65,535 |
| `uint32` | 4 bytes | 0 to 4,294,967,295 |
| `uint64` | 8 bytes | 0 to 18,446,744,073,709,551,615 |
### **Special Integer Types**
- **`int`** and **`uint`**: Default integer types, size depends on the system (32-bit or 64-bit).
- **`byte`**: Alias for `uint8`, often used for raw data.
- **`rune`**: Alias for `int32`, represents Unicode characters.
### **Choosing the Right Integer Type**
1. If you're working with a file format or network protocol, match the required size.
2. If you're writing generic functions, use **Go's generics**.
3. Otherwise, just use `int`—it's the most natural and efficient choice.
## **Integer Operations**
You can use standard arithmetic operators:
- `+` (addition)
- `-` (subtraction)
- `*` (multiplication)
- `/` (integer division, discards remainder)
- `%` (modulus)
Example:
```go
var x int = 10
x *= 2  // x is now 20
```
Bitwise operations:
- `&` (AND)
- `|` (OR)
- `^` (XOR)
- `<<` (left shift)
- `>>` (right shift)
Example:
```go
var y uint8 = 0b1100
y = y << 1  // y is now 0b11000 (24 in decimal)
```
## **Floating-Point Types**
Go has two floating-point types:
| Type | Precision | Min Value | Max Value |
| --- | --- | --- | --- |
| `float32` | ~6-7 digits | `1.4e-45` | `3.4e+38` |
| `float64` | ~15-16 digits | `5.0e-324` | `1.8e+308` |
- Use **`float64` by default** because it has higher precision.
- **Floating-point arithmetic is inexact** (e.g., `0.1 + 0.2` might not be exactly `0.3`).
### **Floating-Point Division**
- **`x / 0`** results in **`+Inf` or `Inf`** (positive/negative infinity).
- **`0 / 0`** results in **`NaN`** (Not a Number).
Example:
```go
var a float64 = 1.0 / 0  // +Inf
var b float64 = 0.0 / 0  // NaN
```
### **Floating-Point Comparisons**
Never compare floating-point numbers using `==` due to rounding errors.
Instead, check if the difference is small enough:
```go
const epsilon = 1e-9
if math.Abs(a - b) < epsilon {
    fmt.Println("a and b are approximately equal")
}
```
### **When NOT to Use Floating-Point Numbers**
- **Do NOT use floating-point numbers for money** (e.g., dollars, euros). Instead, use an **integer representation** (e.g., store cents as an `int64`).
- Use `math/big` or a third-party package for precise decimal calculations.
## **Understanding Strings and Runes in Go**
### **1. Strings**
- Go has a built-in `string` type.
- The **zero value** of a string is an **empty string** (`""`).
- Go strings are **immutable**—you cannot modify them after they are created.
- Strings support **Unicode**, so they can store any character from any language.
- Strings can be compared using `==`, `!=`, `<`, `>`, `<=`, and `>=`.
- You can concatenate strings using the `+` operator.
### **2. Runes**
- The `rune` type represents a **single Unicode code point** (a single character).
- `rune` is an alias for `int32`.
- **Use `rune` instead of `int32`** when referring to a single character to improve code readability.
- Example:
    
    ```go
    var myFirstInitial rune = 'J' // ✅ Correct
    var myLastInitial int32 = 'B' // ❌ Legal but confusing
    
    ```
    
## **Type Conversions in Go**
Unlike some languages, Go **does not allow automatic type conversion** between different numeric types (e.g., `int` and `float64`). Instead, you must **explicitly convert** types using type conversion functions.
### **Example 1: Converting Between int and float64**
```go
var x int = 10
var y float64 = 30.2
var sum1 float64 = float64(x) + y // Convert x to float64
var sum2 int = x + int(y)         // Convert y to int
fmt.Println(sum1, sum2) // Output: 40.2 40
```
Here:
- `x` is converted to `float64` so it can be added to `y`.
- `y` is converted to `int`, which removes the decimal part (`30.2` → `30`).
### **Example 2: Converting Between Different Integer Types**
```go
var x int = 10
var b byte = 100
var sum3 int = x + int(b)  // Convert b to int
var sum4 byte = byte(x) + b // Convert x to byte
fmt.Println(sum3, sum4)
```
- Without explicit conversion, Go will **not allow operations between different numeric types**.
### **Boolean Type in Go**
- Unlike Python and JavaScript, Go **does not allow implicit conversion** of numbers or strings to `bool`.
- You **must** use a comparison operator (`==`, `!=`, `>`, `<`, etc.).
- Example:
    
    ```go
    x := 10
    if x {     // ❌ ERROR: x is not a boolean
        fmt.Println("This won't work!")
    }
    if x != 0 { // ✅ Correct
        fmt.Println("x is not zero")
    }
    
    ```
    
## **Declaring Literals in Go**
- **Literals are untyped**, meaning they can adapt to different types as long as the type is compatible.
- Example:
    
    ```go
    var x float64 = 10   // 10 is treated as float64
    var y float64 = 200.3 * 5 // The result is float64
    
    ```
    
- You **cannot** assign a literal of one type to a variable of a different type (e.g., a string to an `int`).
## **Variable Declarations in Go**
There are **two main ways** to declare variables in Go:
1. **Using `var` (explicit declaration)**
2. **Using `:=` (short declaration, only inside functions)**
### **1. Using `var`**
```go
var x int = 10   // Explicit type
var y = 20       // Implicit type (int)
var z int        // Defaults to zero value (0 for int)
```
### **2. Using `:=` (Short Declaration)**
```go
x := 10          // Equivalent to var x = 10
x, y := 30, "hello"  // Multiple variables
```
- `:=` **can only be used inside functions**.
### **When to Use `var` vs `:=`?**
| Scenario | Use `var` | Use `:=` |
| --- | --- | --- |
| Declaring a variable inside a function | ❌ | ✅ |
| Declaring a package-level variable | ✅ | ❌ |
| Assigning the zero value explicitly | ✅ | ❌ |
| Assigning a literal to a specific type | ✅ | ❌ |
| Assigning multiple values at once | ✅ | ✅ (only inside functions) |
### **Potential Pitfall: Shadowing Variables**
- `:=` creates **new** variables even if a variable with the same name exists.
- Example:
    
    ```go
    x := 10
    x, y := 30, "hello"  // x is redefined, y is new
    
    ```
    
    - This might **accidentally create a new variable** instead of updating an existing one.
## **Constants in Go**
Go uses `const` to define **immutable** values.
### **Declaring Constants**
```go
const x int64 = 10
const (
    idKey   = "id"
    nameKey = "name"
)
const z = 20 * 10
```
- Constants **must** be values known at **compile time**.
### **Constants vs Variables**
| Feature | Constants (`const`) | Variables (`var`) |
| --- | --- | --- |
| Can change after assignment? | ❌ No | ✅ Yes |
| Can be computed at runtime? | ❌ No | ✅ Yes |
| Can hold complex types (maps, slices, structs)? | ❌ No | ✅ Yes |
### **Example: Illegal Constant Assignment**
```go
x := 5
y := 10
const z = x + y  // ❌ ERROR: x + y is not constant
```
- You **cannot** use runtime values to define a constant.
### **Typed vs. Untyped Constants**
#### **Untyped Constants (More Flexible)**
```go
const x = 10 // x has no fixed type
var y int = x  // ✅ OK: x is treated as int
var z float64 = x  // ✅ OK: x is treated as float64
```
### **Typed Constants (More Restrictive)**
```go
const typedX int = 10
var y float64 = typedX // ❌ ERROR: cannot use int as float64
```
- Use **untyped constants** unless you **need** a specific type.
## **Unused Variables and Constants**
- Go **does not allow unused variables**—if you declare a variable and don’t use it, you get a **compile-time error**.
- **Example (Error)**:
    
    ```go
    func main() {
        x := 10 // ❌ ERROR: x is declared but not used
    }
    
    ```
    
- Constants, however, **can be unused** because they **don’t affect the final compiled program**.
## Composite Types
### Arrays in Go – Why They Are Rarely Used Directly
Go has built-in support for arrays, but unlike many other programming languages, they are not commonly used directly. Instead, most Go developers rely on **slices**, which are more flexible. Before we dive into slices, let's first understand how arrays work in Go and why they have limitations.
### **Declaring and Using Arrays in Go**
An array in Go is a fixed-size collection of elements of the **same type**. There are multiple ways to declare and initialize arrays:
### **Basic Declaration**
```go
var x [3]int
```
This creates an array of size **3**, with all elements initialized to the **zero value** of `int` (which is `0`).
### **Declaring and Initializing at the Same Time**
```go
var x = [3]int{10, 20, 30}
```
Here, the array `x` is created with the values `10, 20, 30`.
### **Sparse Arrays (Specifying Only Some Indices)**
```go
var x = [12]int{1, 5: 4, 6, 6, 10: 100, 15}
```
This creates a **12-element array**, with nonzero values at specific indices:
```
[1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
```
### **Using `...` to Let the Compiler Infer the Size**
```go
var x = [...]int{10, 20, 30}
```
Here, the compiler automatically determines that the array has **3 elements**.
### **Comparing Arrays**
Go allows comparing arrays using `==` and `!=`. Two arrays are **equal** if they:
1. Have the **same length**
2. Have **identical values** at each index
Example:
```go
var x = [...]int{1, 2, 3}
var y = [3]int{1, 2, 3}
fmt.Println(x == y) // prints true
```
### **Multi-Dimensional Arrays**
Go does **not** have built-in support for multi-dimensional arrays like Fortran or Julia, but you can **simulate them**:
```go
var x [2][3]int
```
This is an array with **2 elements**, where each element is an array of **3 integers**.
### **The Biggest Limitation of Arrays in Go**
Unlike most languages where arrays are more flexible, **Go treats the array size as part of its type**. This leads to several restrictions:
- You **cannot** assign an array of size `[3]int` to a variable of type `[4]int`.
- You **cannot** pass an array of one size to a function expecting a different size.
- You **cannot** use a variable to set the size of an array (e.g., `var size = 5; var x [size]int` is invalid).
Because of these restrictions, **Go arrays are not used for general-purpose programming**. Instead, Go provides **slices**, which solve all of these problems.
## **Slices – The Preferred Way to Work with Collections**
A **slice** is a more powerful, flexible data structure built on top of arrays. Unlike arrays, slices:
1. **Do not have a fixed size** (they can grow and shrink).
2. **Can be passed to functions** without worrying about size.
3. **Support powerful operations like appending elements**.
### **Declaring and Initializing Slices**
A slice is declared **without a size**:
```go
var x = []int{10, 20, 30}
```
This creates a slice with **3 elements**.
### **Sparse Slices**
Just like arrays, slices can be initialized with specific indices:
```go
var x = []int{1, 5: 4, 6, 10: 100, 15}
```
This creates a slice of **12 elements**, with values:
```
[1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
```
### **Slices of Slices (Nested Slices)**
You can create slices of slices (similar to multi-dimensional arrays):
```go
var x [][]int
```
### **Comparing Slices**
Unlike arrays, slices **cannot be compared** directly using `==` or `!=`. The only valid comparison is checking whether a slice is `nil`:
```go
var x []int
fmt.Println(x == nil) // true
```
For comparing slices, Go provides the **`slices` package** (since Go 1.21):
```go
import "slices"
x := []int{1, 2, 3, 4, 5}
y := []int{1, 2, 3, 4, 5}
z := []int{1, 2, 3, 4, 5, 6}
fmt.Println(slices.Equal(x, y)) // prints true
fmt.Println(slices.Equal(x, z)) // prints false
```
> Warning: The reflect package has `DeepEqual`, which was previously used to compare slices, but it is slower and not recommended for new code.
> 
### **Built-in Functions for Slices**
#### **1. `len()` – Get Length of a Slice**
```go
x := []int{10, 20, 30}
fmt.Println(len(x)) // prints 3
```
### **2. `append()` – Grow a Slice**
Unlike arrays, slices can **grow** dynamically:
```go
var x []int
x = append(x, 10) // x now contains [10]
```
Appending multiple elements:
```go
x = append(x, 20, 30, 40)
```
Appending another slice:
```go
y := []int{50, 60, 70}
x = append(x, y...) // Note the `...` to expand `y` into separate elements
```
> Important: append() returns a new slice, so you must reassign it to the original variable.
> 
---
### **Why Are Arrays Still in Go?**
You might be wondering: If arrays are so limited, why does Go include them at all?
The answer is that **arrays serve as the underlying data structure for slices**. Go arrays are efficient and provide the **backing store** for slices, allowing slices to be used without unnecessary memory allocation.
**When should you use arrays instead of slices?**
- When you **know the exact size** ahead of time.
- When dealing with **fixed-size data** (e.g., cryptographic hashes).
- When optimizing for **performance** and avoiding dynamic memory allocation.
## How Slices Work Under the Hood in Go
### 1. Concept Overview
A **slice** in Go is a **descriptor**, not the actual data.
It’s a lightweight structure that describes a contiguous segment of an underlying **array**.
Think of it as a **view** over an array rather than a data container.
### 2. Slice Data Structure
Internally, a slice is represented by the Go runtime as:
```go
type slice struct {
    array unsafe.Pointer // pointer to the underlying array
    len   int            // number of elements accessible through the slice
    cap   int            // total capacity (from start of slice to end of array)
}
```
- `array`: pointer to the **first element** of the slice’s view into the array.
- `len`: how many elements the slice currently exposes.
- `cap`: how many elements can be accessed before a new allocation is required.
### 3. Memory Layout Example
```go
arr := [5]int{10, 20, 30, 40, 50}
s := arr[1:3]
```
| Name | Points to | Length (`len`) | Capacity (`cap`) | Data Visible |
| --- | --- | --- | --- | --- |
| `arr` | full array | 5 | 5 | [10, 20, 30, 40, 50] |
| `s` | element 1 of `arr` | 2 | 4 | [20, 30] |
So:
- `s[0]` → 20
- `s[1]` → 30
- `cap=4` because from index `1` to `4` there are 4 elements in `arr`.
### 4. What Happens When You Append
```go
s = append(s, 99)
```
### Two Cases:
1. **Capacity is not exceeded:**
    
    The append just writes into the existing array.
    
2. **Capacity exceeded:**
    
    Go allocates a **new array**, usually doubling the previous capacity.
    
    The old data is copied to the new array.
    
After that, the slice’s internal pointer (`array`) changes to the new array.
### 5. Copy-on-Write Behavior
Slices **share memory** until one of them appends beyond capacity.
```go
a := []int{1, 2, 3}
b := a[:2]
b[0] = 9
fmt.Println(a) // [9, 2, 3]
```
Both `a` and `b` point to the same backing array.
Changes in one are reflected in the other — until a reallocation occurs.
### 6. Reslicing and Capacity Rules
You can reslice within capacity:
```go
s2 := s[:cap(s)]
```
But if you exceed capacity, Go panics:
```go
s3 := s[:cap(s)+1] // panic: slice bounds out of range
```
### 7. Built-in `make` and `append`
`make([]T, len, cap)`:
- allocates an **array** of length `cap`
- returns a **slice** with length `len` referencing it
Example:
```go
s := make([]int, 2, 5)
fmt.Println(len(s), cap(s)) // 2 5
```
### 8. Important Performance Notes
- Slices are **passed by value**, but they **reference the same underlying array**.
- Copying a slice only copies the 24-byte descriptor (on 64-bit systems).
- Appending can trigger a **reallocation**, breaking the sharing link.
- Using `copy(dst, src)` performs an element-wise copy, not reallocation.
### 9. Visual Summary
```
+-------------+       +---------------------------------------------+
| slice value |       | underlying array                            |
|-------------|       |---------------------------------------------|
| ptr ------->|-----> | [10] [20] [30] [40] [50]                    |
| len: 2      |        ^ start at index 1                           |
| cap: 4      |                                               ^ end |
+-------------+                                               |
                                                             |
                 slice covers elements [20, 30]               |
                                                             |
                 capacity allows appending up to [40, 50] ----+
```
### 10. TL;DR Summary
| Concept | Description |
| --- | --- |
| Slice ≠ Array | A slice *refers* to an array; it does not contain data itself. |
| `len` | Number of accessible elements. |
| `cap` | How many elements fit before reallocation. |
| Append | May reuse or reallocate the underlying array. |
| Copy | Copies the slice header, not data (use `copy()` for data). |
## **Capacity and Growth**
Each time you append elements to a slice:
- The **length** increases.
- If the length reaches the capacity, Go allocates a **new backing array** with a larger capacity, copies the old data, and adds the new values.
### **Growth Strategy (Go 1.18+)**
When appending to a slice:
- If the capacity is **less than 256**, it **doubles** in size.
- If the capacity is **256 or more**, it grows by **(current_capacity + 768) / 4**, gradually converging to a **25% growth rate**.
This strategy balances memory usage and performance.
### **Checking Capacity**
Use the built-in `cap` function to check a slice’s capacity:
```go
x := []int{1, 2, 3}
fmt.Println(cap(x)) // 3
```
For arrays, `cap` always returns the same value as `len`.
### **Using `make` to Create Slices**
The `make` function creates a slice with a specified length and capacity:
```go
x := make([]int, 5)     // Length: 5, Capacity: 5
y := make([]int, 5, 10) // Length: 5, Capacity: 10
z := make([]int, 0, 10) // Length: 0, Capacity: 10
```
- `make([]int, 5)`: Initializes a slice with **5 zero-values**.
- `make([]int, 5, 10)`: Pre-allocates memory for up to 10 elements.
- `make([]int, 0, 10)`: Allows appending without frequent reallocation.
### **Common Mistake**
Appending to a non-empty slice increases the length:
```go
x := make([]int, 5)
x = append(x, 10)
fmt.Println(x) // [0 0 0 0 0 10] (length: 6, capacity: 10)
```
Instead of replacing an element, `append` adds a new one.
### **Clearing a Slice (`Go 1.21`)**
The `clear` function sets all elements to their zero value but **keeps the length unchanged**:
```go
s := []string{"a", "b", "c"}
clear(s)
fmt.Println(s) // ["", "", ""]
```
This is useful for resetting data without reallocating memory.
### **Slicing a Slice**
You can extract parts of a slice using the slicing syntax:
```go
x := []string{"a", "b", "c", "d"}
y := x[:2]    // ["a", "b"]
z := x[1:3]   // ["b", "c"]
w := x[:]     // Full copy: ["a", "b", "c", "d"]
```
- **Start index is included**, end index is **excluded**.
- Leaving indices empty defaults to **start=0** and **end=len(slice)**.
### **Memory Sharing**
Sliced slices share memory:
```go
x := []string{"a", "b", "c"}
y := x[:2]
y[0] = "z"
fmt.Println(x) // ["z", "b", "c"]
```
Changing `y` also changes `x` because they refer to the same memory.
### **Avoiding Shared Memory Issues**
Appending to a slice that shares memory can lead to **unexpected modifications**:
```go
x := []string{"a", "b", "c", "d"}
y := x[:2]
y = append(y, "z")
fmt.Println(x) // ["a", "b", "z", "d"]
fmt.Println(y) // ["a", "b", "z"]
```
Here, `y`'s append **modifies** `x` because it still has shared capacity.
### **Preventing This Issue**
Use a **full slice expression** to limit the capacity:
```go
y := x[:2:2] // Capacity limited to length (2)
z := x[2:4:4] // Capacity limited to 2
```
Now, appending to `y` won’t affect `x`.
### **Copying Slices (`copy` Function)**
To create an independent slice:
```go
x := []int{1, 2, 3}
y := make([]int, len(x))
copy(y, x)
fmt.Println(y) // [1 2 3]
```
- `copy(dest, src)` copies the **minimum of their lengths**.
- Ensures `y` doesn’t share memory with `x`.
### **Best Practices**
1. **Use `make([]T, 0, cap)`** when expecting dynamic growth.
2. **Prefer `append`** over manual indexing when adding elements.
3. **Be cautious when slicing**—memory is shared.
4. **Use `copy`** if you need an independent slice.
5. **Use a full slice expression (`[:len:cap]`)** when modifying subslices.
## Strings in Go: A Sequence of Bytes
In Go, strings are sequences of **bytes**, not runes. A string is an array of bytes, where each byte represents a character. While this may seem confusing at first because we're used to thinking of strings as sequences of characters (or runes), the bytes in Go represent characters encoded in a specific character encoding—UTF-8 in this case.
### UTF-8 Encoding
UTF-8 is the most common way to represent Unicode characters as a sequence of bytes. It’s efficient because characters that fit in one byte (like most ASCII characters) only take one byte, while characters that need more space (like emojis or characters from non-Latin alphabets) take more bytes, up to four. This means that Go’s strings are flexible and efficient, especially for languages like English, where most characters fit in one byte. But it also allows handling more complex characters from other languages or emojis when needed.
### How Go Strings Are Represented
Internally, a Go string is simply a sequence of bytes. For example:
```go
var s string = "Hello"
```
In memory, "Hello" might look like this in UTF-8 encoding:
- **H**: 72 (in UTF-8 byte representation)
- **e**: 101
- **l**: 108
- **l**: 108
- **o**: 111
Each character is one byte long, which makes this string simple to manage in memory.
However, when you deal with non-ASCII characters, like emojis or accented letters, you get a string that contains multiple bytes for each character. For example, a sun emoji ☀️ is encoded in UTF-8 as multiple bytes. So a string like this:
```go
var s string = "Hello ☀️"
```
might be represented in memory as:
- **H**: 72
- **e**: 101
- **l**: 108
- **l**: 108
- **o**: 111
- **Space**: 32
- **☀**: 240, 159, 140, 158 (UTF-8 encoding of the sun emoji)
- **️**: 239, 191, 189 (UTF-8 encoding of the variation selector for emojis)
Notice how some characters, like the sun emoji, use multiple bytes. This can lead to issues when slicing strings directly by byte index, especially when dealing with multi-byte characters.
### Indexing Strings in Go
You can index a string in Go just like you would with an array or slice. But remember, you are indexing by **bytes**, not characters. Here's an example:
```go
var s string = "Hello"
var b byte = s[1] // b is assigned 101, which corresponds to the letter 'e'
```
This works because each character in "Hello" is a single byte in UTF-8. But if you try to index into a string that contains multi-byte characters, you'll be working with the byte representation, which can lead to unexpected results. For example:
```go
var s string = "Hello ☀️"
fmt.Println(s[6])  // Prints the byte value of the first byte of the sun emoji (240)
```
This can give you the first byte of a multi-byte character, which isn’t a valid character by itself and doesn’t make sense as a complete string.
### String Slicing
Slicing strings is also done by bytes. Here's an example:
```go
var s string = "Hello ☀️"
var s2 string = s[4:7]
fmt.Println(s2) // Prints "o "
```
In this case, `s2` prints "o " because we sliced the string by byte index. But as soon as you deal with characters that take multiple bytes (like the emoji), the slicing can become unreliable if you don't consider how many bytes each character takes.
### Converting Between Strings, Bytes, and Runes
To make things easier when dealing with strings that contain multi-byte characters, Go provides type conversions to go from strings to slices of bytes or runes (a rune represents a Unicode code point).
Here's how you convert between strings and slices:
```go
var s string = "Hello ☀️"
var bs []byte = []byte(s)  // Convert string to slice of bytes
var rs []rune = []rune(s)  // Convert string to slice of runes
```
- The byte slice (`bs`) contains the UTF-8 encoded bytes of each character.
- The rune slice (`rs`) contains the Unicode code points of each character. In this case, the sun emoji will be represented as a single rune in the slice, not as multiple bytes.
If you print these slices:
```go
fmt.Println(bs)  // Prints the byte values: [72 101 108 108 111 32 240 159 140 158]
fmt.Println(rs)  // Prints the rune values: [72 101 108 108 111 32 127774]
```
Notice that the sun emoji is represented as a single rune in `rs`, but as multiple bytes in `bs`.
### Common Pitfalls
A common mistake when working with strings in Go is to mistakenly convert an integer to a string directly. For example:
```go
var x int = 65
var y = string(x)
fmt.Println(y) // Prints "A", not "65"
```
This works because the integer 65 corresponds to the ASCII value of the letter 'A'. However, you might expect "65" as a string, which isn’t the case. To get the string "65", you need to explicitly convert it to a string:
```go
var x int = 65
var y = strconv.Itoa(x) // Correctly converts the int to the string "65"
fmt.Println(y) // Prints "65"
```
## How Strings Work Under the Hood in Go
### 1. Concept Overview
In Go, a **string** is an **immutable sequence of bytes**.
It’s not a character array — it’s just a read-only view of raw bytes that often (but not always) represent UTF-8–encoded text.
This means:
- You can **read** from a string, but never modify it.
- Any operation that appears to “change” a string actually **creates a new one**.
### 2. String Data Structure
Under the hood, a string is defined roughly like this in Go’s runtime:
```go
type stringStruct struct {
    str unsafe.Pointer // pointer to the underlying bytes
    len int            // number of bytes (not runes)
}
```
So a Go string is just:
- a pointer to read-only bytes in memory, and
- a length field describing how many bytes it spans.
### 3. Memory Layout Example
```go
s := "hello"
```
Internally looks like:
| Field | Meaning |
| --- | --- |
| `str` | pointer to memory: `[104, 101, 108, 108, 111]` |
| `len` | 5 |
| Data | "hello" (UTF-8: 68 65 6C 6C 6F) |
The string header itself is tiny — only 16 bytes on a 64-bit system — and it points to the actual data stored elsewhere in memory (often read-only).
### 4. Why Strings Are Immutable
Go’s immutability rule for strings is by design:
- Multiple strings can safely **share the same underlying memory**.
- It allows the compiler and runtime to optimize string reuse and interning.
- Modifying a string would require copying, which is explicit in Go (`[]byte` conversion).
Example:
```go
a := "hello world"
b := a[:5] // "hello"
```
Here, `b` **shares memory** with `a`.
No copying happens — both point to the same bytes.
### 5. Conversion Between Strings and Byte Slices
Converting a string to a byte slice:
```go
bs := []byte("abc")
```
- Allocates a **new backing array** and copies the bytes.
- The result is mutable (`[]byte`).
Converting back:
```go
s := string(bs)
```
- Allocates new string memory.
- Copies bytes from the slice into a new read-only region.
So both conversions are **O(n)** operations (linear time).
### 6. Accessing Elements
Indexing a string gives you **bytes**, not runes:
```go
s := "Go💙"
fmt.Println(len(s))    // 6 (because 💙 is 4 bytes in UTF-8)
fmt.Println(s[2])      // 240 (first byte of the emoji)
```
To properly iterate characters (runes):
```go
for i, r := range s {
    fmt.Println(i, r)
}
```
This decodes UTF-8 runes on the fly using the `range` construct.
### 7. Substrings and Slicing
When you slice a string:
```go
s := "abcdef"
sub := s[2:4] // "cd"
```
The new string `sub` points to the same memory as `s` but with a different length and pointer offset.
No new allocation or copy happens unless you explicitly convert it to a new byte slice.
### 8. Comparing Strings
String comparison (`==`, `<`, etc.) works **byte-by-byte**.
- If lengths differ, Go immediately returns false.
- If equal length, Go compares each byte.
- Go does not normalize Unicode — `"é"` (U+00E9) ≠ `"e\u0301"`.
### 9. Empty Strings and Zero Values
A zero-value string (`var s string`) is represented as:
```go
stringStruct{str: nil, len: 0}
```
It’s valid and behaves like an empty string `""`.
### 10. Performance Notes
- Strings are **cheap to copy** — only the 16-byte header is duplicated.
- Operations that produce new strings (concatenation, slicing beyond bounds, conversions) can **allocate**.
- Use `strings.Builder` when constructing strings repeatedly — it avoids multiple allocations.
### 11. Visual Summary
```
string "hello"
+--------------------+       +-------------------------------------+
| stringStruct       |       | memory (read-only)                  |
|--------------------|       |-------------------------------------|
| str ---->----------|-----> | [104][101][108][108][111]           |
| len: 5             |       | "hello"                             |
+--------------------+       +-------------------------------------+
```
### 12. TL;DR Summary
| Concept | Description |
| --- | --- |
| Representation | Pointer + length (immutable view over bytes) |
| Encoding | Usually UTF-8 but not enforced |
| Immutable | You can’t modify it; any change creates a new one |
| Conversion | `string → []byte` or `[]byte → string` copies data |
| Indexing | Returns bytes, not runes |
| Slicing | Reuses same memory without copying |
| Performance | Copying string headers is cheap; creating new data is not |
## **Maps in Go**
### **Declaring Maps**
There are multiple ways to declare and initialize a map in Go.
### **1. Declaring a nil map**
```go
var nilMap map[string]int
```
- This declares a **nil map**, meaning it has no allocated memory.
- Its length is `0`, and reading from it will return the zero value of the value type (`int` in this case, so `0`).
- **Trying to write to a nil map causes a runtime panic.**
### **2. Using a map literal**
```go
totalWins := map[string]int{}
```
- This initializes an **empty but usable map**. Unlike a nil map, you **can write** to it.
- It still has a length of `0`, but you can add key-value pairs.
### **3. Using a map literal with values**
```go
teams := map[string][]string{
    "Orcas":   {"Fred", "Ralph", "Bijou"},
    "Lions":   {"Sarah", "Peter", "Billie"},
    "Kittens": {"Waldo", "Raul", "Ze"},
}
```
- This is a **fully initialized map with predefined keys and values**.
- The keys are strings (`"Orcas"`, `"Lions"`, `"Kittens"`).
- The values are slices of strings (`[]string`), meaning each team has a list of players.
### **4. Using `make` to create a map with a predefined size**
```go
ages := make(map[int][]string, 10)
```
- This pre-allocates memory for `10` key-value pairs, but the map still starts with a length of `0`.
- Maps automatically grow as needed, so the initial size is just a hint for performance optimization.
### **Maps vs. Slices**
Go provides both **maps** and **slices**, but they serve different purposes:
| Feature | Map | Slice |
| --- | --- | --- |
| **Key Type** | Any comparable type (e.g., `string`, `int`) | Always an integer index |
| **Use Case** | Quick lookups using custom keys | Ordered data, sequential processing |
| **Automatic Growth** | Yes | Yes |
| **Zero Value** | `nil` | `nil` |
**When to use a slice vs. a map?**
- **Use a slice** when order matters and you process data sequentially.
- **Use a map** when you need **fast lookups** by a key instead of scanning through data.
### **How Go Implements Maps**
Go uses a **hash map (hash table)** under the hood. This means:
- Keys are hashed, and the hash determines their storage location.
- Lookups, inserts, and deletes are **very fast** (on average, O(1) time complexity).
- Go automatically handles **hashing and equality checks** for key types.
You don’t need to implement your own hash functions—Go takes care of that.
### **Reading and Writing Maps**
Let's see how to use a map in practice.
### **Creating and Modifying a Map**
```go
totalWins := map[string]int{}
totalWins["Orcas"] = 1
totalWins["Lions"] = 2
fmt.Println(totalWins["Orcas"])   // Prints 1
fmt.Println(totalWins["Kittens"]) // Prints 0 (not found, so returns zero value)
```
- Assign values using `map[key] = value`.
- If you **read a missing key**, Go **returns the zero value** of the value type (`0` for `int` in this case).
- **You CANNOT use `:=` when setting a map key**, only when declaring the map.
### **Incrementing Values**
```go
totalWins["Kittens"]++
fmt.Println(totalWins["Kittens"]) // Prints 1
```
Since missing keys return `0`, `++` works even if `"Kittens"` wasn’t originally in the map.
### **Checking if a Key Exists (Comma-OK Idiom)**
Sometimes, you need to check **whether a key exists** rather than assuming it might return a zero value.
```go
m := map[string]int{
    "hello": 5,
    "world": 0,
}
v, ok := m["hello"]
fmt.Println(v, ok)  // 5, true
v, ok = m["world"]
fmt.Println(v, ok)  // 0, true (key exists, but value is 0)
v, ok = m["goodbye"]
fmt.Println(v, ok)  // 0, false (key does not exist)
```
- The **comma-ok idiom** allows us to differentiate between:
    - A key that exists but has a zero value.
    - A key that doesn’t exist at all.
- `ok` is `true` if the key exists and `false` otherwise.
### **Deleting a Key from a Map**
```go
m := map[string]int{
    "hello": 5,
    "world": 10,
}
delete(m, "hello") // Removes "hello" from the map
```
- `delete(map, key)` removes a key-value pair.
- If the key isn’t found, **nothing happens** (no error or panic).
- **If you delete all keys, the map is empty but not nil.**
### **Emptying a Map (`clear`)**
```go
m := map[string]int{
    "hello": 5,
    "world": 10,
}
fmt.Println(m, len(m)) // map[hello:5 world:10] 2
clear(m)
fmt.Println(m, len(m)) // map[] 0
```
- `clear(m)` removes **all keys** from a map.
- The map still exists but has a length of `0`.
### **Comparing Maps**
Go **does not** allow direct comparison of two maps using `==`, but Go 1.21 introduced the `maps` package.
```go
import "maps"
m := map[string]int{
    "hello": 5,
    "world": 10,
}
n := map[string]int{
    "world": 10,
    "hello": 5,
}
fmt.Println(maps.Equal(m, n)) // true
```
- `maps.Equal(m, n)` checks if two maps have identical keys and values.
- You can also use `maps.EqualFunc` for custom comparison logic.
---
### **Using Maps as Sets**
A **set** is a collection of unique values. Go doesn’t have a built-in set type, but you can simulate one using a map.
```go
intSet := map[int]bool{}
vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
for _, v := range vals {
    intSet[v] = true
}
fmt.Println(len(vals), len(intSet)) // 11, 8 (duplicates removed)
fmt.Println(intSet[5])   // true
fmt.Println(intSet[500]) // false
```
- The **keys** store unique values, and the **values** (always `true`) indicate presence.
- **Checking membership (`intSet[v]`) is O(1)**, unlike slices which require scanning.
### **Alternative: Using `struct{}` instead of `bool`**
```go
intSet := map[int]struct{}{}
vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
for _, v := range vals {
    intSet[v] = struct{}{}
}
if _, ok := intSet[5]; ok {
    fmt.Println("5 is in the set")
}
```
- `struct{}` uses **zero bytes of memory**, whereas `bool` takes **1 byte per entry**.
- It requires using the **comma-ok idiom** for checking membership.
---
### **Conclusion**
- **Use a map** when you need **fast lookups** using keys.
- **Always initialize a map** before writing to it (`make` or `{}`).
- **Use the comma-ok idiom** when you need to check if a key exists.
- **Use `delete` and `clear`** to remove keys.
- **Maps can be used as sets** by storing `bool` or `struct{}` as values.
## How Maps Work Under the Hood in Go
### 1. Concept Overview
A **map** in Go is an **unordered collection of key–value pairs** implemented as a **hash table** with built-in concurrency safety for reads and writes at the runtime level (though not thread-safe between goroutines).
Key features:
- Keys are **hashed** into buckets.
- Each bucket can store multiple key-value pairs.
- The runtime dynamically **resizes and redistributes** buckets as needed.
### 2. Map Data Structure
Internally (simplified from `runtime/map.go`), a Go map looks like this:
```go
type hmap struct {
    count     int            // number of key-value pairs
    flags     uint8
    B         uint8          // log2 of number of buckets (2^B buckets)
    noverflow uint16
    hash0     uint32         // hash seed to prevent collision attacks
    buckets    unsafe.Pointer // pointer to array of buckets
    oldbuckets unsafe.Pointer // for growing (during rehash)
    nevacuate  uintptr        // progress counter for bucket evacuation
    extra      *mapextra      // optional fields
}
```
Each **bucket** is a fixed-size structure:
```go
type bmap struct {
    tophash [8]uint8
    // followed by up to 8 key/value pairs
}
```
- Each bucket holds up to 8 entries.
- `tophash` helps quickly identify possible matches before comparing full keys.
### 3. How Buckets Work
When you insert a key:
1. The key’s **hash** is computed.
2. The **lower bits** decide the **bucket index** (`hash & (2^B - 1)`).
3. The **upper bits** go into the `tophash` array for quick matching.
4. If the bucket is full, a new **overflow bucket** is allocated and linked.
Example:
```
hash(key) = 0b1101_0110_0010
B = 3 → 8 buckets
bucket index = lower 3 bits = 010 → bucket #2
tophash = upper bits (for quick match)
```
### 4. Map Lookup
To look up `m[k]`:
1. Compute hash of `k`.
2. Select target bucket using `hash & (2^B - 1)`.
3. Check `tophash` values in that bucket for a match.
4. If no match, follow overflow buckets if present.
5. Compare actual keys when `tophash` matches.
6. Return the value or zero value if not found.
### 5. Map Insertion
When inserting:
- Go finds the correct bucket.
- If there’s an empty slot, it’s filled.
- If not, a new overflow bucket is linked.
- The `count` increases, and if the **load factor** exceeds ~6.5, a **resize** begins.
### 6. Growing and Rehashing
When a map grows:
- The number of buckets doubles (`B` increases by 1).
- Go doesn’t rehash all keys immediately — it performs **incremental rehashing** (evacuation).
Each map operation (lookup, insert, delete) may **migrate a few buckets** from `oldbuckets` to `buckets`.
This keeps rehashing cost amortized and avoids long pauses.
### 7. Deletion
When you `delete(m, k)`:
- Go locates the bucket.
- Marks the slot as empty (zeroes key and value).
- Does not shrink the map (shrinking never happens automatically).
### 8. Iteration
When you use:
```go
for k, v := range m { ... }
```
- Go picks a **random start bucket and offset** each iteration to avoid bias.
- Since maps grow incrementally, iteration order is **non-deterministic**.
- If the map grows or shrinks during iteration, Go ensures consistency but not predictability.
### 9. Hash Collisions
- If multiple keys hash to the same bucket index, Go uses **open addressing with overflow buckets**.
- Only when `tophash` matches does Go perform a full key equality check.
- Go uses a **per-map random hash seed (`hash0`)** to prevent DoS-style collision attacks.
### 10. Type Constraints on Keys
Valid key types must be **comparable** — meaning they can be used with `==` and `!=`.
Examples:
- Valid: `string`, `int`, `float64`, `struct` with comparable fields.
- Invalid: `slice`, `map`, `function` (since they’re not comparable).
### 11. Zero Value of a Map
A zero-value map (`var m map[string]int`) is `nil`.
It behaves as follows:
- Reading: returns zero value.
- Writing: panics (`assignment to entry in nil map`).
    
    Use `make` to allocate:
    
```go
m := make(map[string]int)
```
### 12. Memory Layout Visualization
```
hmap
+-------------------+
| count = 3         |
| B = 3 (8 buckets) |
| buckets ----------|-----------------------+
| oldbuckets (nil)  |                       |
+-------------------+                       |
                                            v
                               +---------------------------+
                               | bucket 0                 |
                               | keys: a, b, c            |
                               | values: 1, 2, 3          |
                               | overflow -> (nil)         |
                               +---------------------------+
```
### 13. Performance Notes
- **Average O(1)** for lookups, inserts, and deletes.
- **Worse case O(n)** if all keys land in one bucket (very rare).
- **Iteration order is randomized** for safety and testing consistency.
- **No automatic shrinking** — use `make` with known size if possible.
- **Amortized rehashing** ensures smooth performance over time.
### 14. TL;DR Summary
| Concept | Description |
| --- | --- |
| Data structure | Hash table with buckets (8 entries each) |
| Lookup | Hash → bucket → tophash → key comparison |
| Insert | Place in bucket or overflow; triggers growth if full |
| Delete | Clears entry but doesn’t shrink map |
| Resize | Incremental, doubles bucket count |
| Iteration | Randomized, non-deterministic order |
| Complexity | O(1) average, O(n) worst case |
Would you like me to continue with a section explaining **Go’s hash function design** (how different key types are hashed and why a per-map random seed is used)?
## **Structs in Go**
A **struct** in Go is a collection of related data fields that are grouped together under a single type. Unlike maps, which allow arbitrary key-value pairs, structs **define a fixed structure** for data. This makes them more type-safe and suitable for passing structured data between functions.
### **Why Use Structs Instead of Maps?**
- **Fixed structure:** A struct defines a specific set of fields, whereas a map allows any key.
- **Different field types:** A map can only store values of the same type, but a struct can mix different types.
- **Better API design:** Structs clearly define what data a function expects.
## **Defining a Struct**
A struct is declared using the `type` keyword followed by the struct name, the `struct` keyword, and a set of fields inside curly braces `{}`.
### **Example: Defining a `person` Struct**
```go
type person struct {
    name string
    age  int
    pet  string
}
```
- The `person` struct has three fields: `name`, `age`, and `pet`, each with a specific type.
- Field names start with lowercase letters, making them **unexported** (private to the package).
---
## **Creating Struct Variables**
Once a struct type is defined, you can create variables of that type.
### **1. Using `var` (Zero Value Initialization)**
```go
var fred person
```
- Since no values are assigned, `fred` gets the **zero value** of each field (`""` for `string`, `0` for `int`).
### **2. Using an Empty Struct Literal**
```go
bob := person{}
fmt.Println(bob)
```
- Equivalent to `var bob person`—all fields are initialized to their zero values.
### **3. Initializing with a Struct Literal**
#### **Positional Style**
```go
julia := person{"Julia", 40, "cat"}
```
- The values must be **provided in the exact order** of the struct definition.
- If the struct definition changes (e.g., a new field is added), this code will break.
### **Named Fields Style**
```go
beth := person{
    age:  30,
    name: "Beth",
}
```
- **Field names are specified**, so the order doesn’t matter.
- Any missing fields will be initialized to their zero values.
🚨 **Important Rule:**
You **cannot mix** the two struct literal styles in a single declaration. This won’t compile:
```go
sam := person{"Sam", age: 25} // ❌ Syntax error
```
## **Accessing and Modifying Struct Fields**
Use **dot notation** to read and modify struct fields.
```go
bob.name = "Bob"
fmt.Println(bob.name) // Outputs: Bob
```
- Just like accessing a map with `map[key]`, you use `struct.field`.
## **Anonymous Structs**
Instead of defining a named struct type, you can create a **one-time-use struct**.
### **Example: Anonymous Struct**
```go
var person struct {
    name string
    age  int
    pet  string
}
person.name = "Bob"
person.age = 50
person.pet = "dog"
```
- This struct exists only in this **variable scope**.
- You cannot reuse the struct type elsewhere.
### **Using an Anonymous Struct with a Literal**
```go
pet := struct {
    name string
    kind string
}{
    name: "Fido",
    kind: "dog",
}
```
- The struct type is defined **inline**.
- Useful when dealing with **temporary data** or **JSON unmarshaling**.
## **Comparing and Converting Structs**
Structs in Go **can be compared** if all their fields are comparable types (i.e., no slices, maps, functions, or channels).
### **Struct Comparison**
```go
type firstPerson struct {
    name string
    age  int
}
type secondPerson struct {
    name string
    age  int
}
f := firstPerson{"Bob", 50}
g := secondPerson{"Bob", 50}
// fmt.Println(08_file_servers_cdn == g) // ❌ ERROR: different struct types
```
- **Structs of different types cannot be compared**, even if they have identical fields.
- However, if one struct is **anonymous**, you can compare them:
```go
var g struct {
    name string
    age  int
}
g = f // ✅ Allowed
fmt.Println(f == g) // ✅ Outputs: true
```
## **Struct Type Conversion**
Go allows converting between struct types **only if**:
1. The structs have the **same field names**.
2. The fields are in the **same order**.
3. The field types are **identical**.
### **Example: Valid Conversion**
```go
type firstPerson struct {
    name string
    age  int
}
type secondPerson struct {
    name string
    age  int
}
f := firstPerson{"Bob", 50}
s := secondPerson(f) // ✅ Allowed
```
### **Examples: Invalid Conversions**
#### **Field Order Mismatch**
```go
type thirdPerson struct {
    age  int
    name string
}
f := firstPerson{"Bob", 50}
t := thirdPerson(f) // ❌ ERROR: different field order
```
### **Field Name Mismatch**
```go
type fourthPerson struct {
    firstName string
    age       int
}
f := firstPerson{"Bob", 50}
fp := fourthPerson(f) // ❌ ERROR: different field names
```
### **Additional Fields**
```go
type fifthPerson struct {
    name          string
    age           int
    favoriteColor string
}
f := firstPerson{"Bob", 50}
fp := fifthPerson(f) // ❌ ERROR: extra field in fifthPerson
```
---
## **When to Use Structs**
- When you need **fixed** sets of related data.
- When passing structured data between functions.
- When you need **different types of values** in one object.
- When you want **better API documentation** than just using maps.
## **Key Takeaways**
✅ **Structs** define a **fixed** set of related fields, unlike dynamic maps.
✅ **Struct literals** come in **two styles**: positional and named fields (named is preferred for readability).
✅ **Anonymous structs** are useful for temporary data but can’t be reused.
✅ **Structs can be compared** if all their fields are comparable types.
✅ **Struct type conversion** is allowed **only if field names, order, and types match exactly**.
## Blocks, Shadows, and Control Structures
### **Understanding Blocks and Scope in Go**
A **block** in Go is simply a section of code enclosed in curly braces `{}`. These blocks determine where variables, constants, and functions exist and how they can be accessed.
There are different types of blocks in Go:
- **Package block** – Contains everything at the package level (e.g., functions, global variables).
- **File block** – Defines names for imported packages, valid only within that file.
- **Function block** – Contains variables and parameters inside a function.
- **Inner blocks** – Every `{}` inside a function creates a new block, including if statements, loops, etc.
- **Universe block** – A special block that contains built-in identifiers like `int`, `true`, `false`, and `nil`.
### **Shadowing Variables in Go**
**Shadowing** happens when a new variable is declared with the same name as an existing variable in an outer block. Once shadowed, the original variable becomes **inaccessible** until the shadowing variable goes out of scope.
### **Example of Shadowing**
```go
package main
import "fmt"
func main() {
    x := 10
    if x > 5 {
        fmt.Println(x) // Prints 10
        x := 5         // Shadows the outer x
        fmt.Println(x) // Prints 5 (new x)
    }
    fmt.Println(x)     // Prints 10 (original x is back)
}
```
### **Why?**
1. The first `fmt.Println(x)` prints `10` because the outer `x` (defined before `if`) is still accessible.
2. Then, `x := 5` inside the `if` block **creates a new x** that exists only inside that block.
3. When inside `if`, `fmt.Println(x)` prints `5` because it refers to the shadowed `x`.
4. Once outside the `if` block, the shadowing `x` is gone, so `fmt.Println(x)` prints `10` again.
### **Shadowing with Multiple Assignments**
Go allows multiple variables to be assigned at once using `:=`. However, `:=` **only reuses variables that are already in the current block**, meaning it will shadow outer variables if needed.
### **Example**
```go
func main() {
    x := 10
    if x > 5 {
        x, y := 5, 20 // x is shadowed, y is new
        fmt.Println(x, y) // Prints 5 20
    }
    fmt.Println(x) // Prints 10 (original x)
}
```
Here, `x := 5` shadows the outer `x`, while `y := 20` is newly declared.
### **Shadowing Built-in Identifiers (Dangerous!)**
Because Go allows shadowing even for built-in names, you can accidentally shadow package names or even built-in types. This can lead to confusing errors.
### **Example: Shadowing the fmt Package**
```go
func main() {
    x := 10
    fmt.Println(x) // Works fine
    fmt := "oops"  // Shadows the fmt package!
    fmt.Println(fmt) // ERROR: fmt is now a string, not a package
}
```
**What happens here?**
- When `fmt := "oops"` is declared, it **shadows** the `fmt` package.
- Now, `fmt.Println()` tries to call `.Println()` on a **string**, which doesn’t exist.
- Result? **Compilation error.**
---
### **The Universe Block: Shadowing True**
Go’s built-in identifiers (like `true`, `false`, `nil`, `int`) live in a special **universe block**. You can technically shadow them, but it’s a terrible idea.
### **Example**
```go
fmt.Println(true) // Prints: true
true := 10
fmt.Println(true) // Prints: 10
```
After redefining `true`, it no longer represents the boolean value `true`, which can lead to weird bugs.
### **Key Takeaways**
✅ **Blocks** define where variables exist (package, file, function, inner `{}` blocks).
✅ **Shadowing** happens when an inner block defines a variable with the same name as an outer block’s variable.
✅ **Be careful with :=** – it shadows outer variables if at least one new variable is introduced.
✅ **Never shadow package imports (e.g., fmt)** – it can break your program.
✅ **Never shadow built-in types (e.g., true, nil, int)** – it leads to confusing errors.
## **if Statements and Variable Scope**
Go allows **scoping a variable to an if statement**, meaning that variable exists **only inside the if/else block**.
### **Example**
```go
if n := rand.Intn(10); n == 0 {
    fmt.Println("That's too low")
} else if n > 5 {
    fmt.Println("That's too big:", n)
} else {
    fmt.Println("That's a good number:", n)
}
```
Here’s what happens:
1. `n := rand.Intn(10)` **declares and initializes** `n`, which exists **only in this if statement**.
2. `n` is used in all `if`, `else if`, and `else` blocks.
3. **Outside the if statement, `n` does not exist!**
If you try:
```go
fmt.Println(n) // ERROR: undefined: n
```
You’ll get a **compilation error** because `n` was scoped to the `if` block.
## **The Four Types of `for` Loops in Go**
Unlike other languages that have multiple looping constructs (`while`, `do-while`, `for`), Go keeps things simple—there’s only **one** keyword for looping: `for`. However, Go provides flexibility by allowing four different ways to use it:
1. **The Classic C-Style `for`**
2. **The Condition-Only `for` (like `while`)**
3. **The Infinite `for` (runs forever unless stopped)**
4. **The `for-range` loop (for iterating over collections)**
Let’s go through each of them with explanations and examples.
## **1. The Classic C-Style `for` Loop**
This is the `for` loop you’re probably familiar with if you've worked with C, Java, or JavaScript:
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```
### **Breaking it down:**
- The loop consists of **three parts**, separated by semicolons (`;`):
    1. **Initialization**: `i := 0` (this sets up the loop variable)
    2. **Condition**: `i < 10` (as long as this is `true`, the loop continues)
    3. **Increment**: `i++` (after each iteration, `i` increases by 1)
- No parentheses `()` are needed around the loop declaration.
- Curly braces `{}` are required for the loop body.
### **Variations of this loop**
You can omit parts of the loop if they are set up elsewhere.
### **(a) Without initialization:**
If `i` is declared before the loop, you can skip the initialization step:
```go
i := 0
for ; i < 10; i++ {
    fmt.Println(i)
}
```
**Notice:** The first semicolon is still needed.
### **(b) Without increment inside the `for` statement:**
Sometimes, you want to modify `i` inside the loop body instead of the `for` header.
```go
for i := 0; i < 10; {
    fmt.Println(i)
    if i % 2 == 0 {
        i++
    } else {
        i += 2
    }
}
```
Here, the loop’s increment logic is inside the loop instead of the `for` statement.
## **2. The Condition-Only `for` Loop (like `while`)**
If you remove both the initialization and the increment, you get a loop that looks like a `while` loop in other languages.
```go
i := 1
for i < 100 {
    fmt.Println(i)
    i = i * 2
}
```
**How it works:**
- The loop runs **while `i < 100` is true**.
- There’s no initialization or increment in the `for` header. Instead, the increment happens inside the loop.
## **3. The Infinite `for` Loop**
If you remove the condition entirely, Go assumes the loop should run **forever**.
```go
for {
    fmt.Println("Hello")
}
```
**Why is this useful?**
- Sometimes, you want a loop that runs indefinitely, waiting for external events (like server requests).
- You can **exit manually** using `break`.
### **Stopping an Infinite Loop**
To avoid getting stuck, you can use `break`:
```go
for {
    fmt.Println("Running...")
    if someCondition {
        break // Exit the loop
    }
}
```
## **4. The `for-range` Loop (for Iterating Over Collections)**
This is a special type of `for` loop for **iterating over arrays, slices, maps, and strings**.
### **Example with a slice:**
```go
evenVals := []int{2, 4, 6, 8, 10, 12}
for i, v := range evenVals {
    fmt.Println(i, v)
}
```
**Output:**
```
0 2
1 4
2 6
3 8
4 10
5 12
```
### **How this works:**
- `range` goes through each element of `evenVals`.
- The first variable (`i`) holds the **index**.
- The second variable (`v`) holds the **value at that index**.
### **Ignoring the Index**
If you don’t need the index, use an underscore `_`:
```go
for _, v := range evenVals {
    fmt.Println(v)
}
```
**Output:**
```
2
4
6
8
10
12
```
This is useful when you **only care about values** and don’t need the index.
### **Ignoring the Value**
If you only need the index, just leave out the second variable:
```go
uniqueNames := map[string]bool{"Alice": true, "Bob": true, "Charlie": true}
for k := range uniqueNames {
    fmt.Println(k)
}
```
Here, `range` iterates over the **keys** of the map.
## **Using `break` and `continue` in Loops**
Like other languages, Go allows you to control loops using `break` and `continue`.
### **`break` (Exit the loop immediately)**
```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // Exit loop when i reaches 5
    }
    fmt.Println(i)
}
```
**Output:**
```
0
1
2
3
4
```
The loop stops as soon as `i == 5`.
### **`continue` (Skip to the next iteration)**
```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // Skip even numbers
    }
    fmt.Println(i)
}
```
**Output:**
```
1
3
5
7
9
```
The `continue` statement **skips even numbers** and moves to the next iteration.
### **Using `continue` to Improve Readability**
Consider this **FizzBuzz** example:
```go
for i := 1; i <= 100; i++ {
    if i%3 == 0 && i%5 == 0 {
        fmt.Println("FizzBuzz")
        continue
    }
    if i%3 == 0 {
        fmt.Println("Fizz")
        continue
    }
    if i%5 == 0 {
        fmt.Println("Buzz")
        continue
    }
    fmt.Println(i)
}
```
Each `if` block **handles one case and exits early** using `continue`, making the logic clear.
## **Iterating Over Maps**
Go’s `for-range` loop works with maps, but **the iteration order is random**. This means that even if you insert keys into a map in a specific order, looping through the map multiple times might produce a different sequence each time.
### **Why Is the Order Random?**
- In earlier versions of Go, maps often - but not always - iterated in a predictable order.
- This led developers to **incorrectly assume** that the iteration order was stable, which caused unexpected bugs.
- The Go team changed the map implementation to **intentionally randomize iteration order** to:
    - Prevent incorrect assumptions about order.
    - Defend against **Hash DoS attacks**, where attackers craft keys that hash to the same bucket, slowing down the server.
### **Debugging Exception**
Even though iteration order is random, **when printing a map using `fmt.Println` or similar functions, the keys are always shown in sorted order** to make debugging easier.
## **Iterating Over Strings**
Go’s `for-range` loop doesn’t iterate over **bytes**, but over **runes** (Unicode code points). This is important because:
- **ASCII characters (English letters, numbers, symbols)** fit into 1 byte.
- **Unicode characters (like `π`)** use **multiple bytes** in UTF-8.
### **Example**
```go
samples := []string{"hello", "apple_π!"}
for _, sample := range samples {
    for i, r := range sample {
        fmt.Println(i, r, string(r))
    }
    fmt.Println()
}
```
### **Expected Output**
```
0 104 h
1 101 e
2 108 l
3 108 l
4 111 o
0 97 a
1 112 p
2 112 p
3 108 l
4 101 e
5 95 _
6 960 π
8 33 !
```
### **Key Observations**
1. The `π` character appears at index `6`, but the next character (`!`) is at index `8`. The number **7 is skipped** because `π` takes **two bytes** in UTF-8.
2. The rune `π` has a **numeric value of 960**, which is larger than what fits in a single byte.
3. If you manually indexed a string using `s[i]`, you’d be working with **bytes**, not runes.
### **Moral of the Story**
- **Use `for-range` for strings** to correctly handle Unicode characters.
- The first loop variable gives the **byte index**, while the second is the **rune**.
## **For-Range Copies Values**
A **common mistake** when using `for-range` is thinking that modifying the loop variable changes the original collection. It doesn’t.
### **Example**
```go
evenVals := []int{2, 4, 6, 8, 10, 12}
for _, v := range evenVals {
    v *= 2
}
fmt.Println(evenVals)
```
### **Expected Output**
```
[2 4 6 8 10 12]
```
Even though `v *= 2` is inside the loop, it **does not modify `evenVals`**. This happens because `v` is a **copy** of each element, not a reference to it.
### **Fix: Use Indexing**
If you need to modify the elements, use indexing:
```go
for i := range evenVals {
    evenVals[i] *= 2
}
```
### **Go 1.22: Iteration Variable Behavior Change**
Before Go 1.22, Go reused the same loop variable for every iteration. This led to **unexpected behavior with goroutines** inside loops.
- **Before Go 1.22:** The same memory location was reused, so goroutines could all reference the **last** value.
- **After Go 1.22:** A new variable is created **for each iteration**, fixing the bug.
## **Using Labels with Loops**
In Go, you can label `for` loops and use `break` or `continue` to control **which loop** they affect.
### **Example**
```go
func main() {
    samples := []string{"hello", "apple_π!"}
outer:
    for _, sample := range samples {
        for i, r := range sample {
            fmt.Println(i, r, string(r))
            if r == 'l' {
                continue outer // Skips the rest of the outer loop and starts next iteration
            }
        }
        fmt.Println()
    }
}
```
### **Expected Output**
```
0 104 h
1 101 e
2 108 l
0 97 a
1 112 p
2 112 p
3 108 l
```
### **Why Use Labels?**
Normally, `continue` and `break` apply to the **closest** loop. But sometimes, you need to skip an **outer** loop. Labels let you do that.
## **Switch Statements in Go**
Unlike languages like C, Java, or JavaScript, Go's `switch` statement:
1. **Doesn’t require `break` statements** – Each case automatically exits after execution.
2. **Allows multiple values per case** – Use commas to group cases.
3. **Allows variable declarations within the switch scope**.
4. **Does not have implicit fall-through** – However, you can explicitly enable it using `fallthrough`.
### **Basic Example**
```go
words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
for _, word := range words {
    switch size := len(word); size {
    case 1, 2, 3, 4:
        fmt.Println(word, "is a short word!")
    case 5:
        wordLen := len(word)
        fmt.Println(word, "is exactly the right length:", wordLen)
    case 6, 7, 8, 9:
    default:
        fmt.Println(word, "is a long word!")
    }
}
```
### **Key Takeaways:**
- **Multiple cases together:** `case 1, 2, 3, 4:` allows multiple values to share logic.
- **Default case:** Executes if no other case matches.
- **No `fallthrough` by default:** Unlike C or JavaScript, execution **does not** move to the next case unless explicitly stated.
- **Scoped variable (`size`)** inside `switch`: It's available in all cases.
## **Empty Cases and `fallthrough`**
If a case has no code, it does nothing.
Example:
```go
case 6, 7, 8, 9:
```
This does nothing when `size` is in `[6, 7, 8, 9]`.
### **Using `fallthrough`**
Although Go discourages fallthrough, you can force it:
```go
switch size {
case 4:
    fmt.Println("Size is 4")
    fallthrough
case 5:
    fmt.Println("Now executing case 5")
}
```
Here, **case 5 runs even if `size == 4`**.
## **Breaking Out of Loops with `switch`**
If a `switch` is inside a loop and you use `break`, it only exits the switch, not the loop.
### **Wrong Example**
```go
for i := 0; i < 10; i++ {
    switch i {
    case 7:
        fmt.Println("exit the loop!")
        break  // Only exits the switch, not the loop
    }
}
```
**Output (unexpected):**
```
0 is even
1 is boring
...
6 is even
exit the loop!
8 is boring
9 is boring
```
### **Fix with Labels**
To break the loop, label it:
```go
loop:
for i := 0; i < 10; i++ {
    switch i {
    case 7:
        fmt.Println("exit the loop!")
        break loop  // Exits the loop, not just the switch
    }
}
```
Now, the loop terminates when `i == 7`.
## **Blank Switch (Expression-less `switch`)**
A **blank switch** doesn’t check a single variable but allows any boolean condition.
### **Example**
```go
words := []string{"hi", "salutations", "hello"}
for _, word := range words {
    switch wordLen := len(word); {
    case wordLen < 5:
        fmt.Println(word, "is a short word!")
    case wordLen > 10:
        fmt.Println(word, "is a long word!")
    default:
        fmt.Println(word, "is exactly the right length.")
    }
}
```
**Why use a blank switch?**
- More readable than multiple `if/else` blocks.
- Allows non-equality conditions (`<`, `>`, `!=`, etc.).
## **Choosing Between `if` and `switch`**
Use `switch` when:
- You have **multiple related conditions**.
- Your conditions use **equality** or **simple comparisons**.
Example (`switch` is better than `if/else`):
```go
for i := 1; i <= 100; i++ {
    switch {
    case i%3 == 0 && i%5 == 0:
        fmt.Println("FizzBuzz")
    case i%3 == 0:
        fmt.Println("Fizz")
    case i%5 == 0:
        fmt.Println("Buzz")
    default:
        fmt.Println(i)
    }
}
```
This is cleaner than:
```go
for i := 1; i <= 100; i++ {
    if i%3 == 0 && i%5 == 0 {
        fmt.Println("FizzBuzz")
    } else if i%3 == 0 {
        fmt.Println("Fizz")
    } else if i%5 == 0 {
        fmt.Println("Buzz")
    } else {
        fmt.Println(i)
    }
}
```
## **The `goto` Statement**
Go includes `goto`, but **you should rarely use it**.
### **Rules:**
- **Cannot jump into a block** (e.g., skipping variable declarations).
- **Cannot skip over variable initialization**.
### **Invalid Example (Won’t compile)**
```go
func main() {
    a := 10
    goto skip
    b := 20  // ERROR: Cannot skip over variable declaration
skip:
    c := 30
    fmt.Println(a, b, c)  // ERROR
}
```
This will throw:
```
goto skip jumps over declaration of b
```
### **Valid Use Case**
When `goto` **improves readability** (rare cases).
```go
func main() {
    a := rand.Intn(10)
    for a < 100 {
        if a%5 == 0 {
            goto done  // Jump out of the loop
        }
        a = a*2 + 1
    }
    fmt.Println("Loop completed normally")
done:
    fmt.Println("This runs regardless of exit reason")
}
```
### **Real-World Example**
Go’s `strconv` package (for string-to-number conversion) uses `goto` to simplify error handling:
```go
overflow:
    mant = 0
    exp = 1<<flt.expbits - 1 + flt.bias
    overflow = true
out:
    bits := mant & (uint64(1)<<flt.mantbits - 1)
    if d.neg {
        bits |= 1 << flt.mantbits << flt.expbits
    }
    return bits, overflow
```
**Why?**
- Some conditions require `overflow` logic.
- Others skip it and jump to `out`.
- Without `goto`, you’d need **duplicate code** or **extra boolean flags**.
### **Key Takeaways on `goto`:**
✅ Use **only** when it **improves readability** (e.g., centralized error handling).
❌ **Avoid whenever possible**—Go’s labeled `break`/`continue` and `return` often work better.
## Functions
### **Declaring and Calling Functions**
- Functions in Go are similar to those in C, Python, or JavaScript.
- Functions consist of four parts:
    1. The `func` keyword
    2. The function name
    3. Input parameters (with types specified)
    4. A return type (if applicable)
Example:
```go
func div(num int, denom int) int {
    if denom == 0 {
        return 0
    }
    return num / denom
}
```
- If a function has no parameters, use `()`, and if it returns nothing, you don’t need a return type.
Example:
```go
func main() {
    result := div(5, 2)
    fmt.Println(result)
}
```
- If multiple parameters share the same type, you can write them like this:
instead of:
    
    ```go
    func div(num, denom int) int
    ```
    
    ```go
    func div(num int, denom int) int
    ```
    
## **Simulating Named and Optional Parameters**
- Go **does not support** named or optional parameters.
- To achieve similar functionality, use a struct.
Example:
```go
type MyFuncOpts struct {
    FirstName string
    LastName  string
    Age       int
}
func MyFunc(opts MyFuncOpts) error {
    // Do something
    return nil
}
func main() {
    MyFunc(MyFuncOpts{
        LastName: "Patel",
        Age:      50,
    })
    MyFunc(MyFuncOpts{
        FirstName: "Joe",
        LastName:  "Smith",
    })
}
```
- This makes parameters flexible without overloading functions.
## **Variadic Parameters**
- Some functions, like `fmt.Println`, accept a variable number of arguments.
- Go achieves this using **variadic parameters**, indicated with `...`.
Example:
```go
func addTo(base int, vals ...int) []int {
    out := make([]int, 0, len(vals))
    for _, v := range vals {
        out = append(out, base+v)
    }
    return out
}
func main() {
    fmt.Println(addTo(3))              // []
    fmt.Println(addTo(3, 2))           // [5]
    fmt.Println(addTo(3, 2, 4, 6, 8))  // [5 7 9 11]
    a := []int{4, 3}
    fmt.Println(addTo(3, a...))        // [7 6]
}
```
- The `...` allows passing a slice instead of separate arguments.
## **Multiple Return Values**
- Unlike many languages, Go **directly supports multiple return values**.
- This is often used for functions that return results **and** errors.
Example:
```go
func divAndRemainder(num, denom int) (int, int, error) {
    if denom == 0 {
        return 0, 0, errors.New("cannot divide by zero")
    }
    return num / denom, num % denom, nil
}
func main() {
    result, remainder, err := divAndRemainder(5, 2)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(result, remainder) // 2 1
}
```
- The `error` is conventionally **the last return value**.
### **Difference from Python**
- In Python, multiple return values are returned as a **tuple**, which can be assigned to a single variable.
    
    ```python
    def div_and_remainder(n, d):
        return n // d, n % d
    
    v = div_and_remainder(5, 2)
    print(v)  # (2, 1)
    
    ```
    
- In Go, you **must** assign all returned values separately.
## **Ignoring Returned Values**
- Go **does not allow** unused variables.
- If you don’t need a return value, use `_`.
Example:
```go
result, _, err := divAndRemainder(5, 2)
```
- You can also ignore **all** return values:
but this is only idiomatic for functions like `fmt.Println`.
    
    ```go
    divAndRemainder(5, 2) // values are dropped, not idiomatic
    ```
    
## **Named Return Values**
- Go allows naming return values inside the function signature.
Example:
```go
func divAndRemainder(num, denom int) (result int, remainder int, err error) {
    if denom == 0 {
        err = errors.New("cannot divide by zero")
        return result, remainder, err
    }
    result, remainder = num/denom, num%denom
    return result, remainder, err
}
```
- The names (`result`, `remainder`, `err`) act like **predeclared variables**.
### **Gotcha: Named Returns Are Ignored If You Explicitly Return Other Values**
Example:
```go
func divAndRemainder(num, denom int) (result int, remainder int, err error) {
    result, remainder = 20, 30  // Assigned, but ignored later
    if denom == 0 {
        return 0, 0, errors.New("cannot divide by zero")
    }
    return num / denom, num % denom, nil // These values are returned instead
}
```
- This prints:
    
    ```
    2 1
    ```
    
- The explicit return **overrides** the named return values.
## **Blank Returns (Naked Returns) – Avoid Using These!**
- If you use named return values, Go allows you to return without specifying them.
Example:
```go
func divAndRemainder(num, denom int) (result int, remainder int, err error) {
    if denom == 0 {
        err = errors.New("cannot divide by zero")
        return // Implicitly returns (result, remainder, err)
    }
    result, remainder = num/denom, num%denom
    return // Implicitly returns (result, remainder, err)
}
```
### **Why is this bad?**
- Makes it harder to see what’s being returned.
- Leads to confusion, especially in large functions.
- Best practice: **always explicitly return values**.
## **Functions as Values**
In Go, functions are first-class citizens, meaning they can be assigned to variables, passed around as arguments, and returned from other functions. The **type of a function** consists of:
- The `func` keyword
- The function’s **parameter types**
- The function’s **return type(s)**
This combination is called the **function’s signature**. If two functions have the same signature, they are interchangeable.
Example:
```go
var myFuncVariable func(string) int
```
This variable can store any function that takes a `string` and returns an `int`.
### **Assigning Functions to Variables**
```go
func f1(a string) int {
    return len(a) // Returns string length
}
func f2(a string) int {
    total := 0
    for _, v := range a {
        total += int(v) // Sum of Unicode values
    }
    return total
}
func main() {
    var myFuncVariable func(string) int
    myFuncVariable = f1
    fmt.Println(myFuncVariable("Hello")) // Output: 5
    myFuncVariable = f2
    fmt.Println(myFuncVariable("Hello")) // Output: 500
}
```
Here, `myFuncVariable` first stores `f1`, then `f2`. When executed, it produces different results.
### **Function Values in Maps (Simple Calculator Example)**
You can store functions in a **map**, associating them with string keys. This is useful for cases like a calculator:
### **Step 1: Define Functions**
```go
func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }
```
### **Step 2: Store Functions in a Map**
```go
var opMap = map[string]func(int, int) int{
    "+": add,
    "-": sub,
    "*": mul,
    "/": div,
}
```
### **Step 3: Use the Map to Perform Operations**
```go
func main() {
    expressions := [][]string{
        {"2", "+", "3"},
        {"2", "-", "3"},
        {"2", "*", "3"},
        {"2", "/", "3"},
        {"2", "%", "3"},   // Unsupported operator
        {"two", "+", "3"}, // Invalid number
        {"5"},             // Invalid format
    }
    for _, expr := range expressions {
        if len(expr) != 3 {
            fmt.Println("invalid expression:", expr)
            continue
        }
        p1, err := strconv.Atoi(expr[0]) // Convert first number
        if err != nil {
            fmt.Println(err)
            continue
        }
        op := expr[1]
        opFunc, ok := opMap[op] // Lookup function in map
        if !ok {
            fmt.Println("unsupported operator:", op)
            continue
        }
        p2, err := strconv.Atoi(expr[2]) // Convert second number
        if err != nil {
            fmt.Println(err)
            continue
        }
        result := opFunc(p1, p2) // Execute function
        fmt.Println(result)
    }
}
```
### **Output:**
```
5
-1
6
0
unsupported operator: %
strconv.Atoi: parsing "two": invalid syntax
invalid expression: [5]
```
This demonstrates error handling and function lookup in maps.
## **Function Type Declarations**
Instead of repeating the function signature, we can **define a function type**:
```go
type opFuncType func(int, int) int
```
Then, rewrite `opMap` as:
```go
var opMap = map[string]opFuncType{
    "+": add,
    "-": sub,
    "*": mul,
    "/": div,
}
```
This improves readability and avoids redundant type declarations.
## **Anonymous Functions**
Anonymous functions are functions **without a name**, which can be assigned to variables or used inline.
### **Assigning an Anonymous Function**
```go
func main() {
    f := func(j int) {
        fmt.Println("printing", j, "from inside of an anonymous function")
    }
    for i := 0; i < 5; i++ {
        f(i) // Calls the anonymous function
    }
}
```
### **Output:**
```
printing 0 from inside of an anonymous function
printing 1 from inside of an anonymous function
printing 2 from inside of an anonymous function
printing 3 from inside of an anonymous function
printing 4 from inside of an anonymous function
```
### **Immediate Execution of an Anonymous Function**
You can define and execute an anonymous function immediately:
```go
func main() {
    for i := 0; i < 5; i++ {
        func(j int) {
            fmt.Println("printing", j, "from inside of an anonymous function")
        }(i) // Call immediately
    }
}
```
This is useful for **defer statements** and **goroutines**.
### **Package-Level Anonymous Functions**
You can define **package-level variables** that store anonymous functions:
```go
var (
    add = func(i, j int) int { return i + j }
    sub = func(i, j int) int { return i - j }
    mul = func(i, j int) int { return i * j }
    div = func(i, j int) int { return i / j }
)
```
### **Usage:**
```go
func main() {
    x := add(2, 3)
    fmt.Println(x) // Output: 5
}
```
### **Modifying a Package-Level Anonymous Function**
Unlike regular functions, package-level function variables **can be reassigned**:
```go
func main() {
    x := add(2, 3)
    fmt.Println(x) // Output: 5
    changeAdd()
    y := add(2, 3)
    fmt.Println(y) // Output: 8
}
func changeAdd() {
    add = func(i, j int) int { return i + j + j } // Modify behavior
}
```
### **Output:**
```
5
8
```
However, changing function behavior dynamically **makes code harder to understand**, so use it carefully.
## **Closures**
A *closure* is a function that **captures and can modify variables from its surrounding scope**. In simple terms, if you declare a function inside another function, the inner function "remembers" the variables from the outer function—even after the outer function has finished executing.
### **Example of a Closure**
```go
func main() {
    a := 20
    f := func() {
        fmt.Println(a) // Accessing 'a' from the outer scope
        a = 30        // Modifying 'a'
    }
    f()
    fmt.Println(a) // Prints the modified value of 'a'
}
```
**Output:**
```
20
30
```
Here, `f` is an *anonymous function* (a function without a name) that can **access and modify** the variable `a`, even though `a` belongs to the outer function `main`. This is the essence of a closure.
### **Variable Shadowing in Closures**
If you **redeclare** a variable inside a closure using `:=`, it creates a *new* variable, instead of modifying the original one.
```go
func main() {
    a := 20
    f := func() {
        fmt.Println(a) // Access outer 'a'
        a := 30        // This 'a' is a new local variable
        fmt.Println(a) // Prints the new 'a', not the outer 'a'
    }
    f()
    fmt.Println(a) // Outer 'a' is still unchanged
}
```
**Output:**
```
20
30
20
```
Since we used `:=` inside the closure, it declared a *new* `a` that only exists inside the closure. The outer `a` remains unchanged.
⚠ **Be careful when using `=` vs `:=` in closures!**
- `a = 30` modifies the **existing** variable from the outer scope.
- `a := 30` creates a **new** variable inside the closure.
### **Why Use Closures?**
Closures can be really useful when:
1. **You want to limit a function’s scope** – Instead of defining a function globally, you keep it local inside another function.
2. **You want to remove repetitive logic** – If you're calling the same logic multiple times inside a function, a closure can simplify things.
3. **You want to "capture" some variables** – Closures allow you to "remember" values from the surrounding function even after the function has exited.
## **Passing Functions as Parameters**
Since functions are *values* in Go, you can pass them around just like integers, strings, or any other data type.
A common use case for this is **sorting slices**.
```go
type Person struct {
    FirstName string
    LastName  string
    Age       int
}
func main() {
    people := []Person{
        {"Pat", "Patterson", 37},
        {"Tracy", "Bobdaughter", 23},
        {"Fred", "Fredson", 18},
    }
    // Sort by last name
    sort.Slice(people, func(i, j int) bool {
        return people[i].LastName < people[j].LastName
    })
    fmt.Println(people)
    // Sort by age
    sort.Slice(people, func(i, j int) bool {
        return people[i].Age < people[j].Age
    })
    fmt.Println(people)
}
```
**How It Works:**
- `sort.Slice` takes a **function** as an argument.
- This function (closure) captures `people` and sorts it based on the specified field.
- The sort function returns `true` if `people[i]` should come before `people[j]`.
## **Returning Functions from Functions**
Closures allow us to return functions from other functions. This is useful when we want to generate functions dynamically.
### **Example: A Multiplier Generator**
```go
func makeMult(base int) func(int) int {
    return func(factor int) int {
        return base * factor
    }
}
func main() {
    twoBase := makeMult(2)
    threeBase := makeMult(3)
    fmt.Println(twoBase(1), threeBase(1)) // 2 3
    fmt.Println(twoBase(2), threeBase(2)) // 4 6
}
```
**Output:**
```
2 3
4 6
```
**How It Works:**
1. `makeMult(2)` returns a closure that always multiplies by `2`.
2. `makeMult(3)` returns a closure that always multiplies by `3`.
3. `twoBase` and `threeBase` store these closure functions.
4. When we call `twoBase(2)`, it returns `2 * 2 = 4`.
5. When we call `threeBase(2)`, it returns `3 * 2 = 6`.
**Real-world Example**: Middleware in web servers often use this pattern to dynamically generate handlers.
## **Higher-Order Functions**
If you've heard the term **higher-order function**, it just means:
- A function **accepts another function as a parameter**.
- A function **returns another function**.
Closures make **higher-order functions** possible in Go.
For example:
```go
func operate(x int, y int, op func(int, int) int) int {
    return op(x, y)
}
func add(a, b int) int { return a + b }
func multiply(a, b int) int { return a * b }
func main() {
    fmt.Println(operate(3, 4, add))      // 7
    fmt.Println(operate(3, 4, multiply)) // 12
}
```
Here, `operate` takes an operation (`op`) as a function and applies it to `x` and `y`.
## **Defer**
The **`defer`** keyword in Go is used to schedule function calls to be executed **after** the surrounding function finishes. This is particularly useful for cleaning up resources like files, network connections, or database transactions, ensuring they are properly closed, even if a function exits early due to an error.
### **Basic Usage of `defer`**
Consider this example, which reads a file and prints its contents:
```go
func main() {
    if len(os.Args) < 2 {
        log.Fatal("no file specified")
    }
    f, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close() // Ensure the file is closed when `main` exits
    data := make([]byte, 2048)
    for {
        count, err := f.Read(data)
        os.Stdout.Write(data[:count])
        if err != nil {
            if err != io.EOF {
                log.Fatal(err)
            }
            break
        }
    }
}
```
**How `defer` Works Here**
1. **Opening a File:** `os.Open` is called to open the file.
2. **Deferring Cleanup:** `defer f.Close()` schedules `f.Close()` to run **only after `main` exits**, ensuring the file is properly closed, even if an error occurs.
3. **Reading the File:** The loop reads and prints the file content.
4. **Error Handling:** If an error occurs (other than `io.EOF`), `log.Fatal` is called, exiting the program. The deferred function (`f.Close()`) **still runs before exiting**.
### **Order of Deferred Calls**
If multiple `defer` statements are used, they execute in **Last-In, First-Out (LIFO) order**.
```go
func deferExample() int {
    a := 10
    defer func(val int) {
        fmt.Println("first:", val)
    }(a)
    a = 20
    defer func(val int) {
        fmt.Println("second:", val)
    }(a)
    a = 30
    fmt.Println("exiting:", a)
    return a
}
```
### **Execution Order**
1. `"exiting: 30"` is printed first (normal execution).
2. **Defer stack unwinds (LIFO order):**
    - The second `defer` captures `a = 20` (since arguments are **evaluated at the time of defer**), so `"second: 20"` is printed.
    - The first `defer` captured `a = 10`, so `"first: 10"` is printed.
### **Output:**
```
exiting: 30
second: 20
first: 10
```
### **Using `defer` with Named Return Values**
A deferred function **can modify the return value** of a function using **named return values**.
```go
func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer func() {
        if err == nil {
            err = tx.Commit()
        }
        if err != nil {
            tx.Rollback()
        }
    }()
    _, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1", value1)
    if err != nil {
        return err
    }
    return nil
}
```
**How This Works**
1. **Start a Database Transaction:** `db.BeginTx(ctx, nil)`.
2. **Defer Cleanup:** The `defer` function determines whether to commit (`tx.Commit()`) or roll back (`tx.Rollback()`) the transaction based on `err`.
3. **Execute SQL Statements:**
    - If an error occurs, `err` is set, and `return err` causes the deferred function to execute before exiting.
    - If no error occurs, the deferred function commits the transaction.
This ensures **all database interactions are cleaned up properly**, reducing the risk of leaving a transaction open.
### **Returning a Cleanup Function**
Instead of handling cleanup inside `defer` within `main`, a **function can return a cleanup function** to ensure proper resource management.
**Example: Returning a Cleanup Function**
```go
func getFile(name string) (*os.File, func(), error) {
    file, err := os.Open(name)
    if err != nil {
        return nil, nil, err
    }
    return file, func() {
        file.Close()
    }, nil
}
```
**Usage in `main`**
```go
f, closer, err := getFile(os.Args[1])
if err != nil {
    log.Fatal(err)
}
defer closer() // Ensure file is closed at the end
```
**Why This is Useful**
1. **Forces Cleanup Usage** – If the caller forgets to call `closer()`, the program **won’t compile** due to an unused variable.
2. **Encapsulates Cleanup Logic** – Keeps file handling and cleanup logic **together** for clarity.
### **Why `defer` is Better than Try-Finally**
Many languages use **try-finally** (Java, Python, JavaScript) to clean up resources:
```python
try:
    f = open("file.txt")
    data = f.read()
finally:
    f.close()  # Cleanup in `finally`
```
### **Advantages of `defer`**
- **No Extra Indentation** – Avoids deep nesting from try-finally blocks.
- **Less Error-Prone** – Ensures cleanup always happens, even if multiple return points exist.
- **More Readable Code** – Keeps cleanup next to resource allocation.
## **Go Is Call by Value**
When people say Go is a **call-by-value** language, they mean that **when you pass a variable to a function, Go makes a copy of the value**. The function operates on this copy, not the original variable. Let's break this down with examples.
### **1. Passing Basic Types (int, string, struct)**
Consider this example where we pass an `int`, a `string`, and a `struct` to a function:
```go
package main
import "fmt"
type person struct {
    age  int
    name string
}
func modifyFails(i int, s string, p person) {
    i = i * 2        // Modify integer
    s = "Goodbye"    // Modify string
    p.name = "Bob"   // Modify struct field
}
func main() {
    p := person{}
    i := 2
    s := "Hello"
    modifyFails(i, s, p)
    fmt.Println(i, s, p) // Output: 2 Hello {0 }
}
```
### **What happens?**
- The function `modifyFails` **tries** to modify `i`, `s`, and `p`, but these modifications **don’t affect** the original values.
- This is because Go passes **copies** of the variables, so the changes inside `modifyFails` are **lost** once the function exits.
- Even though `p` is a struct, it's also passed **by value**, meaning the function only works on a copy of `p`, not the original one in `main`.
### **2. Why Are Maps and Slices Different?**
Maps and slices behave **differently** when passed to a function. Let's look at this example:
```go
package main
import "fmt"
func modMap(m map[int]string) {
    m[2] = "hello"      // Modify an existing entry
    m[3] = "goodbye"    // Add a new entry
    delete(m, 1)        // Remove an entry
}
func modSlice(s []int) {
    for k, v := range s {
        s[k] = v * 2    // Modify existing elements
    }
    s = append(s, 10)   // Append a new element
}
func main() {
    m := map[int]string{
        1: "first",
        2: "second",
    }
    modMap(m)
    fmt.Println(m) // Output: map[2:hello 3:goodbye]
    s := []int{1, 2, 3}
    modSlice(s)
    fmt.Println(s) // Output: [2 4 6] (but no "10" added)
}
```
### **What happens?**
- The **map** `m` is modified inside `modMap`, and those changes persist in `main()`. This is because **maps in Go are implemented with pointers**. So, even though Go is call-by-value, the **value of a map variable is a reference to its underlying data**, meaning modifications affect the original.
- The **slice** `s` behaves differently:
    - The existing values (`1, 2, 3`) are **modified** inside `modSlice`, and those changes persist.
    - However, the **new value (`10`) added via `append` doesn’t persist**. This happens because when you append to a slice, Go might allocate a new underlying array, and the function now works with this new array instead of the original one.
The reason maps and slices behave differently than basic types is that **they are implemented using pointers**.
- **A map variable holds a reference (a pointer) to an internal data structure**. So when you pass a map, you're passing a copy of that reference, which still points to the same underlying data.
- **A slice is a small struct that contains a pointer to an array, along with its length and capacity**. When passed to a function, a copy of this struct is made, but **the copy still points to the same underlying array**—which is why modifications to elements persist. However, if the slice grows and a new array is allocated, the function will be working with the new array, and the caller won’t see the changes.
## Pointers
### **What Is a Pointer?**
A **pointer** is a variable that stores the **memory address** of another variable. Instead of holding a direct value (like an `int` or `string`), a pointer holds a reference to where that value is stored.
Example:
```go
var x int32 = 10
var y bool = true
```
![image.png](media/image.png)
### **Memory Layout and Pointer Storage**
- Every variable is stored in **contiguous memory locations**.
- Different data types take up different amounts of memory (e.g., `int32` takes 4 bytes, `bool` takes 1 byte).
- However, **pointers are always the same size**, regardless of what type they point to (typically 4 or 8 bytes, depending on the system).
For example:
```go
var x int32 = 10
var y bool = true
pointerX := &x
pointerY := &y
var pointerZ *string
```
![image.png](media/image_1.png)
Since `pointerZ` is not assigned any value, it holds `nil` (zero memory address).
### **Pointer Operators**
Go uses two important operators for working with pointers:
1. **Address-of operator (`&`)**: Gets the memory address of a variable.
    
    ```go
    x := "hello"
    pointerToX := &x  // pointerToX stores the address of x
    ```
    
2. **Indirection (dereference) operator ()**: Gets the value stored at a pointer’s address.
    
    ```go
    x := 10
    pointerToX := &x
    fmt.Println(*pointerToX) // prints 10 (dereferencing the pointer)
    ```
    
You can also perform operations using dereferenced pointers:
```go
z := 5 + *pointerToX
fmt.Println(z) // prints 15
```
⚠️ **Dereferencing a nil pointer will cause a panic!**
```go
var x *int
fmt.Println(*x) // This will panic because x is nil
```
Always check that a pointer is **not nil** before dereferencing it.
### **Pointer Types**
A **pointer type** represents a pointer to a specific type. You declare it using `*` before the type name:
```go
var pointerToX *int
```
Here, `pointerToX` is a pointer that can hold the address of an `int` variable.
You can assign a value to a pointer using the address-of operator:
```go
x := 10
pointerToX = &x  // pointerToX now holds the address of x
```
### **Creating Pointers with `new`**
Go provides the `new` function to create a pointer to a zero-initialized variable:
```go
var x = new(int) // x is a pointer to an int initialized to 0
fmt.Println(*x)  // prints 0
```
However, `new` is **rarely used**. The preferred way to create pointers to structs is with the `&` operator:
```go
x := &Foo{}  // Creates a pointer to an instance of Foo
```
---
### **Why Can't You Take the Address of a Constant?**
In Go, you **can’t take the address of a constant** because constants don’t have memory addresses—they exist **only at compile time**.
Example:
```go
type person struct {
    FirstName  string
    MiddleName *string
    LastName   string
}
p := person{
  FirstName:  "Pat",
  MiddleName: "Perry", // This won't compile!
  LastName:   "Peterson",
}
```
You'll get an error:
```
cannot use "Perry" (type string) as type *string in field value
```
Trying to use `&"Perry"` also fails:
```
cannot take the address of "Perry"
```
### **Solution: Using a Helper Function**
A workaround is to **use a helper function** that creates a pointer from a value:
```go
func makePointer[T any](t T) *T {
    return &t
}
```
Now you can do:
```go
p := person{
  FirstName:  "Pat",
  MiddleName: makePointer("Perry"), // This works!
  LastName:   "Peterson",
}
```
Since function parameters are variables (not constants), they **do have memory addresses**, allowing us to return a pointer.
## **Don’t Fear Pointers**
In Java, Python, and JavaScript, when you pass an **object** (class instance) to a function, the function operates on the **same object in memory**. However, if you assign a **new object** to the function parameter, it does **not** change the original variable in the calling function.
This is the same behavior as pointers in Go.
### **Primitive vs. Object Assignment in Other Languages**
#### **Example in Java:**
```java
int x = 10;
int y = x;
y = 20;
System.out.println(x); // prints 10
```
- `y = x` creates a copy.
- Modifying `y` does not affect `x`.
### **Example in Java(With Objects):**
```java
public class Foo {
    private int x;
    public Foo(int x) {
        this.x = x;
    }
    public static void main(String[] args) {
        outer();
    }
    private static void outer() {
        Foo f = new Foo(10);
        inner1(f);
        System.out.println(f.x);
        inner2(f);
        System.out.println(f.x);
        Foo g = null;
        inner2(g);
        System.out.println(g == null);
    }
    private static void inner1(Foo f) {
        f.x = 20;
    }
    private static void inner2(Foo f) {
        f = new Foo(30);
    }
}
```
Running this code prints the following output:
```
20
20
True
```
That’s because the following scenarios are true in Java, Python, JavaScript, and Ruby:
- If you pass an instance of a class to a function and you change the value of a field, the change is reflected in the variable that was passed in.
- If you reassign the parameter, the change is *not* reflected in the variable that was passed in.
- If you pass `nil/null/None` for a parameter value, setting the parameter itself to a new value doesn’t modify the variable in the calling function.
Some people explain this behavior by saying that class instances are passed by reference in these languages. This is untrue. If they were being passed by reference, scenarios two and three would change the variable in the calling function. These languages are always pass-by-value, just as in Go.
What you are seeing is that every instance of a class in these languages is implemented as a pointer. When a class instance is passed to a function or method, the value being copied is the pointer to the instance. Since `outer` and `inner1` are referring to the same memory, changes made to fields in `f` in `inner1` are reflected in the variable in `outer`. When `inner2` reassigns `f` to a new class instance, this creates a separate instance and does not affect the variable in `outer`.
### **Go Equivalent**
```go
package main
import "fmt"
type Foo struct {
	x int
}
func main() {
	outer()
}
func outer() {
	f := &Foo{10} // 08_file_servers_cdn is a pointer to a Foo struct where x = 10
	inner1(f)
	fmt.Println(f.x) // prints 20
	inner2(f)
	fmt.Println(f.x) // still prints 20
	var g *Foo
	inner2(g)
	fmt.Println(g == nil) // prints true
}
func inner1(f *Foo) {
	f.x = 20 // Modifies the object that 08_file_servers_cdn points to
}
func inner2(f *Foo) {
	f = &Foo{30} // Reassigns 08_file_servers_cdn to a new Foo instance, but this does not affect the original 08_file_servers_cdn in outer
}
```
### **Key Takeaways**
1. **Go is still pass-by-value.**
    - When you pass a struct, a copy is made.
    - When you pass a pointer, a copy of the **pointer** is made (but it still points to the same memory).
2. **Objects in Java, JavaScript, Python, and Ruby behave like Go pointers.**
    - Class instances are internally **pointers** in those languages.
    - This means modifying fields of a passed instance **affects the original**.
    - But reassigning a new instance inside a function **does not affect the caller**.
3. **Go gives you control over pointers.**
    - You decide when to use **values** vs. **pointers**.
    - **Use values** for safety and simplicity (e.g., small structs, immutable data).
    - **Use pointers** when you need to modify the original or avoid unnecessary copying.
## **Pointers Indicate Mutable Parameters**
Go does not have a built-in way to declare variables or parameters as immutable, unlike some other languages. However, Go developers use **pointers** to indicate that a parameter is **mutable**—meaning the function is expected to modify the original value.
### **Understanding How Go Passes Values**
Go is always **pass-by-value**, meaning when you pass a variable to a function, Go creates a **copy** of that value.
- **For primitives (e.g., `int`, `string`), arrays and non-pointer structs:** The function gets a copy, so modifications inside the function **do not affect** the original variable.
- **For pointers:** The function gets a copy of the pointer, but since this pointer still references the same memory location as the original, modifications to the underlying value **do affect** the original.
### **Example 1: Why a Nil Pointer Cannot Be Updated**
Let’s look at an example where we try to update a nil pointer:
```go
package main
import "fmt"
func failedUpdate(g *int) {
    x := 10
    g = &x // This only changes g locally, not the original pointer
}
func main() {
    var f *int // 08_file_servers_cdn is nil
    failedUpdate(f)
    fmt.Println(f) // prints nil
}
```
![image.png](media/image_2.png)
### **What’s Happening Here?**
1. In `main()`, `f` is declared as a `int` (a pointer to an int) but is currently `nil`.
2. `failedUpdate(f)` is called, so a **copy** of `f` (which is `nil`) is passed to `g`.
3. Inside `failedUpdate()`, a new variable `x` is created and assigned the value `10`.
4. The pointer `g` is updated to point to `x`, but this change is **local to the function**. The original `f` in `main()` remains unchanged.
5. After `failedUpdate()` finishes, `f` in `main()` is still `nil`, and `x` disappears because it was a local variable in `failedUpdate()`.
### **Example 2: The Correct Way to Modify a Pointer Value**
If you actually want to update the original value, you need to **dereference** the pointer and modify the value in memory, rather than reassigning the pointer itself.
```go
package main
import "fmt"
func failedUpdate(px *int) {
    x2 := 20
    px = &x2 // Only changes px locally, doesn't affect the original
}
func update(px *int) {
    *px = 20 // Dereferencing the pointer modifies the original value
}
func main() {
    x := 10
    failedUpdate(&x)
    fmt.Println(x) // prints 10, because failedUpdate did not modify x
    update(&x)
    fmt.Println(x) // prints 20, because update modified x through the pointer
}
```
![image.png](media/image_3.png)
### **What’s Happening Here?**
1. `x` is created in `main()` and set to `10`.
2. `failedUpdate(&x)` is called, which passes a **copy** of `x`'s memory address to `px`.
    - Inside `failedUpdate()`, a new variable `x2 = 20` is created.
    - `px` is updated to point to `x2`, **but this change is local**—the original `x` in `main()` is not affected.
    - After `failedUpdate()` exits, `x2` is destroyed, and `x` in `main()` is still `10`.
3. `update(&x)` is called, again passing a **copy** of `x`'s memory address to `px`.
    - This time, `px = 20` **modifies the value at the memory location** rather than changing the pointer itself.
    - Since both `px` and `x` in `main()` refer to the same memory location, `x` is successfully updated.
## **Pointers Are a Last Resort**
While pointers in Go allow for mutable parameters, **you should use them sparingly**. Pointers can make code harder to understand and increase memory management overhead. Instead of passing a pointer to a function to modify an existing struct, it's usually better to **return a new struct instance**.
### **Bad Example: Modifying a Struct Using a Pointer**
```go
func MakeFoo(f *Foo) error {
    f.Field1 = "val"
    f.Field2 = 20
    return nil
}
```
Why is this bad?
- It **requires an already allocated `Foo` instance`**, making it unclear who owns the data.
- It **forces mutability**, making debugging and reasoning about the code harder.
### **Good Example: Returning a Struct**
```go
func MakeFoo() (Foo, error) {
    f := Foo{
        Field1: "val",
        Field2: 20,
    }
    return f, nil
}
```
Why is this better?
- It makes ownership **clear**: the function creates and returns a new struct.
- **No pointers** means **no unexpected modifications** elsewhere in the program.
### **Exception: Pointers Are Necessary for Interfaces**
When working with functions like `json.Unmarshal`, pointers **are required** because Go needs to populate an existing variable.
```go
f := struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}{}
err := json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &f)
```
Why do we pass a pointer?
- Without generics, `Unmarshal` **can't create a new struct dynamically**.
- Using a pointer **avoids unnecessary memory allocations** inside a loop.
## **Pointer Passing Performance**
If a struct is large, passing a pointer to it improves performance since pointer size is constant, while passing a value takes longer as size increases.
Returning small structs (e.g., 100 bytes) is faster by value (~10ns) than by pointer (~30ns). But for large structs (e.g., 10MB), returning by value (~1.5ms) is much slower than by pointer (~0.5ms).
In most cases, the difference is negligible, but for large data structures, using pointers—even for immutable data—can improve efficiency.
## **The Zero Value vs. No Value**
Sometimes, you need to **distinguish between a variable that has been set to its zero value and one that hasn’t been set at all**.
### **Using `nil` to Indicate "No Value"**
```go
type User struct {
    Name  *string // nil means "not set"
    Email string  // "" (empty string) means "not set"
}
```
- A `nil` pointer **clearly indicates** the field was **never assigned**.
- A string with `""` might **mean either** "not assigned" or "assigned as empty".
### **Using the `comma ok` idiom instead**
Instead of returning `nil`, use a boolean to indicate if a value is present:
```go
func getUser() (User, bool) {
    return User{Name: "Alice"}, true
}
u, ok := getUser()
if ok {
    fmt.Println("User exists:", u.Name)
} else {
    fmt.Println("No user found")
}
```
## **The Difference Between Maps and Slices**
### **Maps in Go**
Maps in Go are implemented as **pointers** to a struct. This means when you pass a map to a function, you're copying the pointer to the map rather than its contents. This behavior can have some important consequences:
- **Maps are mutable**: Any modification to a map inside a function affects the original map.
- **Unclear API design**: When you pass a map to a function or return a map, there’s no explicit contract regarding the keys and values, making it harder to understand the expected structure without looking through the code.
### **When Should You Use Maps?**
Maps are appropriate when:
- **Keys are not known at compile time**: If you don’t know what the keys will be, a map is an appropriate data structure. For example, in situations where keys are dynamically generated.
However, **maps should not be used in public APIs** when the structure of the data needs to be clear, as they don’t give any clue about the types of data inside.
### **Example of Struct vs Map**
```go
type User struct {
    Name  string
    Age   int
}
userMap := map[string]interface{}{"name": "Alice", "age": 30} // Map is less clear
userStruct := User{Name: "Alice", Age: 30}  // Struct is explicit and clear
```
### **Slices in Go**
Slices in Go are more complex than maps. While they share some similarities, such as being passed around as **pointers**, there are key differences in how their length and capacity work:
- **Slice contents are shared**: When a slice is passed to a function, the contents are **modified** in place (i.e., the function modifies the slice's underlying memory).
- **Changing length**: When you use `append()` on a slice, **the length of the slice can change**. However, if the slice does not have enough capacity, a **new underlying array** is created, which **disconnects the original slice** from the new one.
### **The Memory Layout of Slices**
A slice is implemented as a struct with three fields:
1. **Pointer to the underlying array**
2. **Length**: the number of elements the slice currently holds
3. **Capacity**: the number of elements the slice can hold before it needs to reallocate.
    
    ![image.png](media/image_4.png)
    
Here’s what happens when you pass a slice around:
- **Passing a slice**: When passed to a function, the slice's length, capacity, and pointer are copied. Both the original slice and the copied slice point to the same underlying array.
    
    ![image.png](media/image_5.png)
    
- **Modifying contents**: If you change the contents of the slice (e.g., `s[2] = 4`), both the original and the copy see the change because they share the same underlying array.
    
    ![image.png](media/image_6.png)
    
- If the slice copy is **appended** to and there *is* enough capacity in the slice for the new values, the length changes in the copy, and the new values are stored in the block of memory that’s shared by the copy and the original. However, the length in the original slice remains unchanged. The Go runtime prevents the original slice from seeing those values since they are beyond the length of the original slice.
    
    ![image.png](media/image_7.png)
    
- If the slice copy is appended to and there ***isn’t* enough capacity** in the slice for the new values, a new, bigger block of memory is allocated, values are copied over, and the pointer, length, and capacity fields in the copy are updated. Changes to the pointer, length, and capacity are not reflected in the original, because they are only in the copy
    
    ![image.png](media/image_8.png)
    
- The result is that a slice that’s passed to a function can have its contents modified, but the slice can’t be resized. As the only usable linear data structure, slices are frequently passed around in Go programs. By default, you should assume that a slice is not modified by a function. Your function’s documentation should specify whether it modifies the slice’s contents.
- The reason you can pass a slice of any size to a function is that the data type that’s passed to the function is the same for any size slice: a struct of two `int` values and a pointer. The reason you can’t write a function that takes an array of any size is that the entire array is passed to the function, not just a pointer to the data.
## **Slices as Buffers**
In Go, slices are commonly used as **buffers** for reading from external resources (files, network connections, etc.). Rather than allocating new memory for each chunk of data, you use a preallocated slice (a buffer) and reuse it. This avoids unnecessary memory allocations, reducing the pressure on the garbage collector.
### **Example: Using a Slice as a Buffer**
```go
func loadAndProcess(fileName string, process func([]byte)) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	data := make([]byte, 100)
	for {
		count, err := file.Read(data)
		process(data[:count])
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		if count == 0 {
			return nil
		}
	}
}
```
- **Efficiency**: This approach uses a single slice (`data`) for multiple chunks of data rather than creating a new slice for every chunk. It’s more efficient in terms of both **memory usage** and **garbage collection**.
### **Why Use Slices as Buffers?**
- **Avoid reallocation**: By reusing the same buffer slice, you prevent unnecessary allocations and keep the memory footprint smaller.
- **Memory management**: The slice is allocated once with a fixed capacity, and its length is adjusted as new data is read into it.
- **Cleaner code**: It simplifies your logic by handling chunks of data with one slice, improving both readability and performance.
## **Reducing the Garbage Collector’s Workload**
### **What is Garbage Collection?**
In Go, the term *garbage* refers to data that is no longer being used by the program (i.e., it has no pointers pointing to it). Once data is no longer referenced, it can be safely removed from memory and reused. The *garbage collector* automatically detects unused memory and reclaims it to prevent memory leaks. Although this process is very useful, it’s important to minimize the amount of *garbage* created to reduce the workload on the garbage collector.
### **Stack vs. Heap Memory**
Go uses two primary types of memory allocation: **stack** and **heap**.
- **Stack Memory**:
    - Fast and simple allocation.
    - Each function call gets its own stack frame, and local variables and parameters passed are stored here.
    - The stack grows and shrinks with each function call.
    - **Important**: Stack memory is allocated with known sizes (e.g., primitive values, arrays, structs).
- **Heap Memory**:
    - Memory managed by the garbage collector.
    - Data on the heap remains valid as long as there’s a pointer to it. Once no pointers reference it, the garbage collector reclaims the memory.
### **When Does Data Escape to the Heap?**
Data stored on the stack is faster, but if the data’s size is unknown at compile time or if it’s returned from a function, the data must be moved to the heap. This happens when **escape analysis** determines that data can’t be safely stored on the stack, and it "escapes" to the heap.
### **Why Does Heap Allocation Matter?**
- **Performance**:
    - Heap memory is slower to allocate and access because it’s not laid out sequentially in memory.
    - **Stack-based data** (e.g., structs or slices) is more cache-friendly and can be accessed faster than heap-allocated data.
- **Garbage Collection Overhead**:
    - The garbage collector incurs overhead in tracking which parts of heap memory are in use.
    - A program that generates a lot of garbage will result in slower garbage collection cycles and increased memory usage.
Go's garbage collector is designed to minimize latency, pausing the program for only a small time (under 500 microseconds). However, generating excess garbage increases the time it takes for the garbage collector to clean up.
### **Optimizing Memory Management**
#### **1. Minimizing Pointer Usage**
- Use **value types** (e.g., primitive values, arrays, structs) whenever possible, as these can be allocated on the stack, leading to faster access and less garbage creation.
- Be mindful of using **pointers** since they can force heap allocation, which is more costly.
### **2. Optimizing Data Layout**
- Storing data **sequentially in memory** (e.g., slices of structs) is more efficient than having scattered memory (e.g., slices of pointers), as it reduces the time spent accessing memory.
### **3. Mechanical Sympathy**
- Writing code that is aware of the hardware characteristics (like memory layout) is known as *mechanical sympathy*. This results in more efficient programs by aligning software with how hardware functions.
## **Tuning the Garbage Collector**
Go provides a couple of options to tune garbage collection:
### **GOGC (Garbage Collection Target)**
- **Purpose**: Controls the threshold at which the garbage collector runs based on heap growth.
- **Formula**: The garbage collector will trigger when the heap size reaches `CURRENT_HEAP_SIZE * (1 + GOGC / 100)`.
    - Default value: **100** (heap size doubles before GC runs again).
    - Setting `GOGC` to a smaller value (e.g., 50) reduces heap size before garbage collection, causing more frequent but lighter GC cycles.
    - Setting `GOGC` to a larger value (e.g., 200) reduces the frequency of GC but may result in higher memory usage.
### **GOMEMLIMIT (Memory Limit)**
- **Purpose**: Limits the total memory usage of a Go program, preventing it from consuming more memory than a specified amount.
- **Usage**: `GOMEMLIMIT=3GiB` limits the program to 3 GiB of memory.
- **Behavior**: If memory exceeds the limit, the garbage collector will try to reclaim memory. If it cannot keep up, it will trigger thrashing, where frequent GC cycles occur without freeing enough memory, which can degrade performance.
## Types, Methods, and Interfaces
### **Types in Go**
#### **Structs and User-Defined Types**
In Go, you can define **struct types** using the `type` keyword.
```go
type Person struct {
    FirstName string
    LastName  string
    Age       int
}
```
This declaration creates a **user-defined type** named `Person`, which has an **underlying type** of a struct with three fields.
Go allows you to define new **concrete types** using existing types, such as primitives, functions, and maps:
```go
type Score int
type Converter func(string) Score
type TeamScores map[string]Score
```
Types can be declared at different **block levels**, from the **package block** down to function scopes. However, a type is only accessible within its **scope**, unless it is exported from another package.
### **Abstract vs. Concrete Types**
To better understand types, it's helpful to distinguish between **abstract** and **concrete** types:
- **Abstract Type**: Specifies **what** a type should do but not **how** it does it (e.g., interfaces).
- **Concrete Type**: Specifies both **what** and **how**, meaning it has a defined data structure and associated methods.
Go strictly separates these concepts, unlike languages that allow hybrid types (e.g., abstract classes or interfaces with default methods in Java).
## **Methods in Go**
Like many modern languages, Go allows **methods** to be defined for **user-defined types**.
```go
type Person struct {
    FirstName string
    LastName  string
    Age       int
}
func (p Person) String() string {
    return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}
```
### **Method Syntax**
A **method** in Go is similar to a function but includes a **receiver** before the method name:
```go
func (receiver Type) MethodName() ReturnType { ... }
```
- The **receiver** comes between `func` and the method name.
- The receiver's **name** appears before its **type**.
- By convention, the receiver’s name is a **short abbreviation** of the type's name (e.g., `p` for `Person`).
### **Method Invocation**
Methods are called like in other languages:
```go
p := Person{
    FirstName: "Fred",
    LastName:  "Fredson",
    Age:       52,
}
output := p.String()
```
### **Method Rules in Go**
1. **Methods must be defined at the package level.**
2. **Method overloading is not allowed.**
    - You can have the same method name on different types, but not multiple methods with the same name on the same type.
3. **Methods must be in the same package as the type they belong to.**
    - You cannot add methods to types defined in other packages.
## **Pointer vs. Value Receivers**
Go allows methods to have **pointer receivers** or **value receivers**.
### **When to Use a Pointer Receiver (`Type`)**
Use a **pointer receiver** when:
1. **The method modifies the receiver** (e.g., updating a struct field).
2. **The method needs to handle `nil` instances**.
3. **Other methods on the type use a pointer receiver** (for consistency).
### **When to Use a Value Receiver (`Type`)**
Use a **value receiver** when:
1. **The method does not modify the receiver**.
2. **The receiver type is small and inexpensive to copy** (e.g., `int`, `float64`).
### **Example: Pointer and Value Receivers**
```go
type Counter struct {
    total       int
    lastUpdated time.Time
}
// Pointer receiver (modifies receiver)
func (c *Counter) Increment() {
    c.total++
    c.lastUpdated = time.Now()
}
// Value receiver (does not modify receiver)
func (c Counter) String() string {
    return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}
```
### **Calling Methods with Different Receivers**
```go
var c Counter
fmt.Println(c.String()) // Calls value receiver method
c.Increment()           // Calls pointer receiver method
fmt.Println(c.String())
```
**Output:**
```
total: 0, last updated: 0001-01-01 00:00:00 +0000 UTC
total: 1, last updated: 2009-11-10 23:00:00 +0000 UTC
```
### **Automatic Conversion Between Value and Pointer Receivers**
Go provides **syntactic sugar** for calling pointer receiver methods on **value types** and vice versa.
```go
var c1 Counter
fmt.Println(c1.String())  // Calls c1.String()
c1.Increment()            // Calls (&c1).Increment()
c2 := &Counter{}
fmt.Println(c2.String())  // Calls (*c2).String()
c2.Increment()            // Calls c2.Increment()
```
- When calling a **pointer receiver** method on a **value**, Go **automatically takes the address** of the value.
- When calling a **value receiver** method on a **pointer**, Go **automatically dereferences** the pointer.
### **Beware of Unintended Copies**
When passing a **value type** to a function, it creates a **copy**, meaning modifications inside the function won’t affect the original variable.
```go
func doUpdateWrong(c Counter) {
    c.Increment()  // Updates a copy
    fmt.Println("in doUpdateWrong:", c.String())
}
func doUpdateRight(c *Counter) {
    c.Increment()  // Updates the original
    fmt.Println("in doUpdateRight:", c.String())
}
func main() {
    var c Counter
    doUpdateWrong(c)
    fmt.Println("in main:", c.String())  // Still shows the original value
    doUpdateRight(&c)
    fmt.Println("in main:", c.String())  // Now shows the updated value
}
```
**Output:**
```
in doUpdateWrong: total: 1, last updated: 2009-11-10 23:00:00 +0000 UTC
in main: total: 0, last updated: 0001-01-01 00:00:00 +0000 UTC
in doUpdateRight: total: 1, last updated: 2009-11-10 23:00:00 +0000 UTC
in main: total: 1, last updated: 2009-11-10 23:00:00 +0000 UTC
```
### **Method Sets in Go**
- A **pointer instance** (`T`) can call **both pointer and value receiver methods**.
- A **value instance** (`T`) can only call **value receiver methods**.
This is why `c.Increment()` works even when `c` is a value—it’s converted to `(&c).Increment()` automatically.
### **No Getters and Setters**
Unlike other languages, Go does **not** use **getter** and **setter** methods for struct fields **unless required** (e.g., implementing an interface).
- Direct field access is preferred:
    
    ```go
    p.FirstName = "John"
    ```
    
- Methods should be reserved for **business logic**, such as the `Increment()` method, which updates multiple fields.
## **Code Your Methods for `nil` Instances**
Previously, we discussed **pointer receivers**, which naturally leads to the question: *What happens if you call a method on a `nil` instance?*
In most languages, this would cause an error.
However, **Go does something interesting**—it actually tries to call the method!
- If the method has a **value receiver**, Go panics because it tries to access fields on `nil`.
- If the method has a **pointer receiver**, it can work—**if it is written to handle `nil` properly**.
Sometimes, this behavior makes your code **simpler**. Here’s an example using a **binary tree** that leverages `nil` values:
```go
type IntTree struct {
    val         int
    left, right *IntTree
}
func (it *IntTree) Insert(val int) *IntTree {
    if it == nil {
        return &IntTree{val: val}  // Create a new tree node if nil
    }
    if val < it.val {
        it.left = it.left.Insert(val)  // Recursively insert in the left subtree
    } else if val > it.val {
        it.right = it.right.Insert(val)  // Recursively insert in the right subtree
    }
    return it
}
func (it *IntTree) Contains(val int) bool {
    switch {
    case it == nil:
        return false  // Base case: empty tree doesn’t contain anything
    case val < it.val:
        return it.left.Contains(val)  // Search left
    case val > it.val:
        return it.right.Contains(val) // Search right
    default:
        return true  // Found it!
    }
}
```
In this example:
- The `Insert` method can be called on a `nil` tree because it checks `if it == nil` and initializes a new tree node.
- The `Contains` method is also written to **handle `nil` gracefully**—instead of panicking, it simply returns `false`.
**Why does `Contains` have a pointer receiver?**
Even though `Contains` does not modify the tree, it **needs a pointer receiver** because a value receiver wouldn’t allow it to check for `nil`.
Here’s how you would use the tree in `main`:
```go
func main() {
    var it *IntTree  // `it` is nil
    it = it.Insert(5)
    it = it.Insert(3)
    it = it.Insert(10)
    it = it.Insert(2)
    fmt.Println(it.Contains(2))  // true
    fmt.Println(it.Contains(12)) // false
}
```
### **Key Takeaways**
- Methods on `nil` receivers can work if they are **written to handle `nil` properly**.
- A **pointer receiver** is required if you want to check for `nil`.
- **If your method doesn’t handle `nil`, it will panic**—which might be fine in some cases.
## **Methods Are Just Functions**
In Go, **methods are so similar to functions** that you can use them almost interchangeably in some cases.
Take a look at this struct with a method:
```go
type Adder struct {
    start int
}
func (a Adder) AddTo(val int) int {
    return a.start + val
}
```
### **Method Values**
You can assign a method to a variable. This **binds** the method to the instance it was called from:
```go
myAdder := Adder{start: 10}
f1 := myAdder.AddTo  // f1 is a function that adds to 10
fmt.Println(f1(10))  // prints 20
```
Here, `f1` acts like a **closure**—it remembers `myAdder.start`, so calling `f1(10)` effectively calls `myAdder.AddTo(10)`.
### **Method Expressions**
You can also get a function that takes an explicit receiver as the **first argument**:
```go
f2 := Adder.AddTo
fmt.Println(f2(myAdder, 15))  // prints 25
```
This is different from `f1` because `f2` is a **normal function** that requires `Adder` as an argument.
## **Functions vs. Methods: When to Use Each**
So, when should you use **a method** vs **a function**?
- **Use a function** when it only depends on its **input parameters**.
- **Use a method** anytime your logic depends on values that are configured at startup or changed while your program is running (those values should be stored in a struct)
For example:
```go
// Function: operates only on input parameters
func Multiply(a, b int) int {
    return a * b
}
// Method: depends on the struct’s stored value
type Multiplier struct {
    factor int
}
func (m Multiplier) MultiplyBy(val int) int {
    return m.factor * val
}
```
- `Multiply` doesn’t rely on any state—it just takes inputs and returns an output.
- `MultiplyBy` depends on `m.factor`, so it’s a method.
## **Type Declarations Are Not Inheritance**
Go **does not have inheritance** like Java or Python. But you can **define new types based on existing types**:
```go
type Score int
type HighScore Score
```
At first glance, this looks like inheritance, but it’s **not**.
Go does **not** allow implicit conversion between these types:
```go
var s Score = 100
var hs HighScore = 200
hs = s  // ERROR: Cannot assign Score to HighScore
s = Score(300)  // OK: explicit conversion
hs = HighScore(s) // OK: explicit conversion
```
### **Why Do This?**
Defining new types helps with **clarity and safety**:
```go
type Percentage int
type Milliseconds int
func Delay(ms Milliseconds) {
    fmt.Println("Delaying for", ms, "ms")
}
var p Percentage = 50
Delay(p)  // ERROR: Cannot use Percentage as Milliseconds
```
Even though both `Percentage` and `Milliseconds` are based on `int`, Go **won’t let you mix them up**. This prevents **accidental misuse** in functions.
## **iota Is for Enumerations—Sometimes**
Many programming languages have enumerations, which define a fixed set of named values for a type. Go doesn’t have an explicit `enum` type, but it provides `iota`, a keyword that allows you to define a sequence of related constants with automatically increasing values.
### **How `iota` Works in**
When using `iota`, best practice is to define a new type based on `int` to represent a set of valid values:
```go
type MailCategory int
```
Then, use a `const` block to define the constants for this type:
```go
const (
    Uncategorized MailCategory = iota
    Personal
    Spam
    Social
    Advertisements
)
```
Here’s what happens:
- `iota` starts at **0** and increments for each new constant in the block.
- `Uncategorized = 0`, `Personal = 1`, `Spam = 2`, etc.
- A new `const` block resets `iota` to **0**.
This is an easy way to generate unique identifiers without explicitly assigning values.
### **How `iota` Behaves in a Mixed `const` Block**
When `iota` is used intermittently in a `const` block, it still increments in the background. Consider this example:
```go
const (
    Field1 = 0
    Field2 = 1 + iota
    Field3 = 20
    Field4
    Field5 = iota
)
func main() {
    fmt.Println(Field1, Field2, Field3, Field4, Field5)
}
```
**Output:**
```
0 2 20 20 4
```
### **What’s Happening?**
- `Field1 = 0` (explicitly assigned).
- `Field2 = 1 + iota` (since `iota = 1` here, it becomes `1 + 1 = 2`).
- `Field3 = 20` (explicitly assigned).
- `Field4` inherits the previous value (`20`).
- `Field5 = iota`, which is now **4** (since `iota` was still incrementing).
This shows that `iota` continues increasing, even if it’s not directly used.
### **Best Practices for `iota`**
🔹 **Use `iota` only when the numeric values don’t matter.**
If your constants represent fixed values from an external system (e.g., API responses or database IDs), explicitly assign them instead of relying on `iota`.
🔹 **Be cautious when inserting new values.**
If you add a new constant in the middle of an `iota` list, all subsequent values will shift, which can break your code.
🔹 **For bitmasks, document your intent.**
A common `iota` pattern uses bitwise shifts to define flags:
```go
type BitField int
const (
    Field1 BitField = 1 << iota // 1 (binary: 0001)
    Field2                      // 2 (binary: 0010)
    Field3                      // 4 (binary: 0100)
    Field4                      // 8 (binary: 1000)
)
```
This works well, but if someone inserts a new constant, it shifts the bit values. Always document your approach when using this technique.
🔹 **Avoid `0` if it’s not a valid value.**
By default, `iota` starts at `0`. If `0` doesn’t make sense in your context, use `_` to discard it:
```go
const (
    _ MailCategory = iota // Ignore 0
    Personal
    Spam
    Social
)
```
Now, `Personal` starts at `1`, preventing uninitialized variables from accidentally being `0`.
## **Use Embedding for Composition**
Go doesn’t have **class inheritance**, but it provides **composition** through **embedding**. This allows for **code reuse** while keeping the design **flexible and decoupled**.
```go
type Employee struct {
    Name string
    ID   string
}
func (e Employee) Description() string {
    return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}
type Manager struct {
    Employee
    Reports []Employee
}
func (m Manager) FindNewEmployees() []Employee {
    // business logic
}
```
Here, `Manager` **contains** an `Employee` struct, but **no explicit field name** is assigned to it. This makes `Employee` an **embedded field**.
Because `Employee` is embedded, its **fields and methods** are **promoted** to `Manager`. This allows the following:
```go
m := Manager{
    Employee: Employee{
        Name: "Bob Bobson",
        ID:   "12345",
    },
    Reports: []Employee{},
}
fmt.Println(m.ID)            // prints 12345
fmt.Println(m.Description()) // prints Bob Bobson (12345)
```
**Note:**
You can embed **any type**, not just structs. The methods of the embedded type are promoted to the containing struct.
### **Handling Name Conflicts in Embedded Fields**
If the **containing struct** has a field or method with the same name as the embedded type, you must use the **embedded type’s name** to explicitly reference the shadowed field or method:
```go
type Inner struct {
    X int
}
type Outer struct {
    Inner
    X int
}
```
Here, `Outer` has its **own `X` field**, which hides `Inner.X`:
```go
o := Outer{
    Inner: Inner{
        X: 10,
    },
    X: 20,
}
fmt.Println(o.X)       // prints 20 (Outer's X)
fmt.Println(o.Inner.X) // prints 10 (Inner's X)
```
## **Embedding Is Not Inheritance**
Go’s **embedding** is **not** the same as inheritance. Unlike in object-oriented languages, an embedded field **does not** make the containing struct a subtype of the embedded type.
For example, the following will **not** compile:
```go
var eFail Employee = m        // compilation error!
var eOK Employee = m.Employee // OK!
```
**Error message:**
```
cannot use m (type Manager) as type Employee in assignment
```
`Manager` **contains** an `Employee`, but **it is not an Employee**. You must **explicitly access the embedded field** (`m.Employee`) when assigning to a variable of type `Employee`.
### **No Dynamic Dispatch for Concrete Types**
Go **does not** have **dynamic dispatch** for concrete types. Methods on embedded fields **don’t know** they are embedded.
Consider this example:
```go
type Inner struct {
    A int
}
func (i Inner) IntPrinter(val int) string {
    return fmt.Sprintf("Inner: %d", val)
}
func (i Inner) Double() string {
    return i.IntPrinter(i.A * 2)
}
type Outer struct {
    Inner
    S string
}
func (o Outer) IntPrinter(val int) string {
    return fmt.Sprintf("Outer: %d", val)
}
func main() {
    o := Outer{
        Inner: Inner{
            A: 10,
        },
        S: "Hello",
    }
    fmt.Println(o.Double())
}
```
**Output:**
```
Inner: 20
```
Even though `Outer` has its **own `IntPrinter` method**, `Inner.Double()` still calls `Inner.IntPrinter`, not `Outer.IntPrinter`.
## Defining an Interface
Interfaces in Go are simple. Like other user-defined types, they use the `type` keyword. Here’s an example from the standard library:
```go
type Stringer interface {
    String() string
}
```
This defines an interface called `Stringer` with a single method, `String() string`. Any type that implements this method automatically satisfies the interface—without any explicit declaration.
Let’s look at another example:
```go
type Incrementer interface {
    Increment()
}
```
Now, let’s see how this plays out in practice:
```go
var myStringer fmt.Stringer
var myIncrementer Incrementer
pointerCounter := &Counter{}
valueCounter := Counter{}
myStringer = pointerCounter    // ok
myStringer = valueCounter      // ok
myIncrementer = pointerCounter // ok
myIncrementer = valueCounter   // compile-time error!
```
Why does `valueCounter` fail to implement `Incrementer`? Because `Increment` was defined with a pointer receiver. A value instance only has access to methods defined with value receivers, so it doesn’t satisfy `Incrementer`.
This distinction is important when designing APIs in Go.
## Implicit Interfaces: Type-Safe Duck Typing
In many statically typed languages, you explicitly declare that a type implements an interface. Go does things differently:
If a type has all the methods an interface requires, it **implicitly** implements that interface.
This allows Go to balance the flexibility of dynamically typed languages (like Python or JavaScript) with the safety of statically typed languages (like Java or C#).
Let’s compare Go’s approach to other languages:
### **Dynamically Typed Languages (Duck Typing)**
In Python, you don’t need to declare that a class implements an interface. As long as it provides the expected methods, it works:
```python
class Logic:
    def process(self, data):
        # business logic
def program(logic):
    logic.process(data)
logicToUse = Logic()
program(logicToUse)
```
This is called *duck typing*: “If it walks like a duck and quacks like a duck, it’s a duck.” The downside? There’s no static type checking, which can lead to runtime errors.
### **Statically Typed Languages (Explicit Interfaces)**
In Java, you must explicitly declare that a class implements an interface:
```java
public interface Logic {
    String process(String data);
}
public class LogicImpl implements Logic {
    public String process(String data) {
        // business logic
    }
}
public class Client {
    private final Logic logic;
    public Client(Logic logic) {
        this.logic = logic;
    }
    public void program() {
        this.logic.process(data);
    }
}
public static void main(String[] args) {
    Logic logic = new LogicImpl();
    Client client = new Client(logic);
    client.program();
}
```
This approach ensures type safety but adds boilerplate and makes refactoring harder.
### **Go’s Approach: The Best of Both Worlds**
Go takes a middle ground. You define interfaces based on what the **caller** needs, and any type that satisfies the interface can be used.
```go
type LogicProvider struct {}
func (lp LogicProvider) Process(data string) string {
    // business logic
}
type Logic interface {
    Process(data string) string
}
type Client struct{
    L Logic
}
func(c Client) Program() {
    c.L.Process(data)
}
func main() {
    c := Client{
        L: LogicProvider{},
    }
    c.Program()
}
```
Notice that `LogicProvider` never declares that it implements `Logic`—it just does. This allows for:
1. **Flexibility:** You can swap in new implementations without modifying existing code.
2. **Clarity:** The interface documents exactly what the client needs.
### **Interfaces in the Standard Library**
Go’s standard library has many powerful interfaces, including:
- `io.Reader` (reads from a data source)
- `io.Writer` (writes to a destination)
- `io.Closer` (closes a resource)
Using standard interfaces makes your code more reusable and encourages the *decorator pattern*.
Example: A function that processes an `io.Reader`:
```go
func process(r io.Reader) error
```
Now, you can pass in:
```go
r, err := os.Open(fileName) // File implements io.Reader
if err != nil {
    return err
}
defer r.Close()
return process(r)
```
Or, for a compressed file:
```go
r, err := os.Open(fileName)
if err != nil {
    return err
}
defer r.Close()
gz, err := gzip.NewReader(r) // gzip.Reader also implements io.Reader
if err != nil {
    return err
}
defer gz.Close()
return process(gz)
```
Same function, different input sources—no extra work required.
**Tip:** If the standard library has an interface that matches your needs, use it!
## **Embedding Interfaces**
Just as structs can be embedded, so can interfaces.
For example, `io.ReadCloser` is just a combination of `io.Reader` and `io.Closer`:
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Closer interface {
    Close() error
}
type ReadCloser interface {
    Reader
    Closer
}
```
This allows for flexible, modular design.
## **Accept Interfaces, Return Structs**
A common Go idiom is:
> “Accept interfaces, return structs.”
> 
This means:
1. **Functions should accept interfaces.**
    - This makes the function more flexible.
2. **Functions should return concrete types.**
    - This allows adding new methods later without breaking compatibility.
Example:
```go
func NewLogger(output io.Writer) *Logger {
    return &Logger{output: output}
}
```
Why does this matter?
If you return an interface instead of a struct, adding a new method breaks all existing implementations! But if you return a struct, new methods can be added without issues.
### **Performance Considerations**
While interfaces provide flexibility, they come at a cost:
- Passing an interface parameter requires a heap allocation.
- Returning a struct avoids heap allocations.
This doesn’t mean you should avoid interfaces—it just means you should profile your code before optimizing.
### **When to Break the Rule?**
Returning interfaces is usually a bad idea, but there are exceptions.
For example, in `database/sql/driver`, all database drivers must implement a set of interfaces, so the API uses interface returns.
When Go introduced new features in Go 1.8, they couldn’t modify existing interfaces without breaking backward compatibility. Instead, they created **new** interfaces for the new functionality.
## Understanding Interfaces and `nil`
In Go, interfaces are implemented behind the scenes as structs containing two pointers: one for the value and one for the type of that value. While pointers themselves can be `nil`, an interface only becomes `nil` when **both** the type and the value inside the interface are `nil`. This means that an interface instance can hold a `nil` value, but it will still be **non-nil** if it has a non-`nil` type, even if its value is `nil`.
This might seem a bit tricky, but here’s an example:
```go
func main() {
	var pointerCounter *Counter
	fmt.Println(pointerCounter == nil) // prints true (zero value for pointer is nil)
	var incrementer Incrementer
	fmt.Println(incrementer == nil) // prints true (both type and value are nil)
	incrementer = pointerCounter
	fmt.Println(incrementer == nil) // prints false (now type is Counter, value still nil)
}
```
In this example, `pointerCounter` is `nil`, which makes it the zero value for a pointer. However, when `pointerCounter` is assigned to the `incrementer` interface, even though the pointer inside is `nil`, the interface is still non-`nil` because it has the type `*Counter`.
What this means in practice is that when you check if an interface is `nil`, you’re checking if **both** the type and the value are `nil`. And, if you try to invoke methods on a `nil` interface, it’ll panic because the method lookup fails.
## Interfaces Are Comparable
Go allows comparing interfaces with `==`, but there’s a catch: two interfaces are considered equal only when both their types and their values are equal.
Consider this code:
```go
type Doubler interface {
    Double()
}
type DoubleInt int
func (d *DoubleInt) Double() {
    *d = *d * 2
}
type DoubleIntSlice []int
func (d DoubleIntSlice) Double() {
    for i := range d {
        d[i] = d[i] * 2
    }
}
```
Now, let’s test comparing `Doubler` interfaces with different types:
```go
var di DoubleInt = 10
var di2 DoubleInt = 10
var dis = DoubleIntSlice{1, 2, 3}
var dis2 = DoubleIntSlice{1, 2, 3}
DoublerCompare(&di, &di2)  // prints false
DoublerCompare(&di, dis)   // prints false
DoublerCompare(dis, dis2)  // panics: comparing uncomparable type DoubleIntSlice
```
In this case, we see that while the two `*DoubleInt` pointers have the same value, they point to different memory locations, so they’re not considered equal. Furthermore, when comparing `dis` and `dis2`, Go panics because slices aren't comparable.
This behavior is important to keep in mind when using interfaces in contexts where comparisons might happen, especially with maps, which require comparable keys.
## The Empty Interface Says Nothing
Go’s empty interface, `interface{}`, is a powerful tool for handling data of unknown types. It essentially represents **any** type that implements zero or more methods (which all types do, even primitive types). This is particularly useful when you need to work with data from external sources where the schema might not be defined upfront (e.g., reading from JSON).
Here’s an example:
```go
var i interface{}
i = 20
i = "hello"
i = struct {
    FirstName string
    LastName  string
}{"Fred", "Fredson"}
```
The empty interface is often used with functions like `json.Unmarshal` to store data of any type, but while it offers flexibility, it also introduces some risks. Since the type of the value is unknown, you need to handle it carefully using techniques like type assertions or type switches to safely access the value inside.
## Type Assertions and Type Switches
When working with interfaces, you sometimes need to access the concrete type stored inside. **Type assertions** and **type switches** are the tools Go provides for this. Here’s a breakdown:
### Type Assertion
A **type assertion** allows you to extract the concrete type from an interface. Here’s a simple example:
```go
var i interface{}
var mine MyInt = 20
i = mine
i2 := i.(MyInt)
fmt.Println(i2 + 1)
```
If you assert a type that doesn’t match the value stored in the interface, Go will panic. To avoid this, use the **comma ok** idiom, which allows you to safely check if the assertion is successful:
```go
i2, ok := i.(int)
if !ok {
    return fmt.Errorf("unexpected type: %T", i)
}
fmt.Println(i2 + 1)
```
### Type Switch
A **type switch** allows you to handle multiple possible types for an interface. This is especially useful when you're not sure what concrete type the interface may hold:
```go
func doThings(i interface{}) {
    switch j := i.(type) {
    case int:
        fmt.Println("int:", j)
    case string:
        fmt.Println("string:", j)
    default:
        fmt.Println("unknown type:", j)
    }
}
```
With type switches, you can examine the type of an interface without directly asserting it. This makes it a safer option when you’re dealing with multiple possible types.
## Use Type Assertions and Type Switches Sparingly
While it might seem handy to extract the concrete type behind an interface using type assertions or type switches, these techniques should be used sparingly. The key point is that if you're working with an interface, **treat the parameter or return value as the type you expect it to be** and not what it could possibly be. When you start using type assertions and switches to work with different concrete types, it signals that the function’s API isn't clearly defining the types it expects to work with. If your function requires a different type, that should be explicitly stated in the function's signature.
However, there are valid use cases for type assertions and type switches. Let’s explore those.
### 1. **Optional Interfaces in the Standard Library**
One common use of type assertions is to check if the concrete type behind an interface also implements another interface. This is useful when you want to create a more flexible API, where the user of the API can provide implementations of various interfaces but might only need certain additional capabilities.
For instance, in Go's standard library, the `io.Copy` function efficiently copies data from a reader to a writer. If the `io.Reader` also implements `io.ReaderFrom`, and the `io.Writer` implements `io.WriterTo`, the copy function avoids unnecessary memory allocations and performs the copy directly using those methods. Here’s how it works:
```go
// copyBuffer is the actual implementation of Copy and CopyBuffer.
// if buf is nil, one is allocated.
func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
    // If the reader has a WriteTo method, use it to do the copy.
    // Avoids an allocation and a copy.
    if wt, ok := src.(WriterTo); ok {
        return wt.WriteTo(dst)
    }
    // Similarly, if the writer has a ReadFrom method, use it to do the copy.
    if rt, ok := dst.(ReaderFrom); ok {
        return rt.ReadFrom(src)
    }
    // function continues...
}
```
In the above code, type assertions are used to check if the `src` (reader) implements `WriterTo` and if the `dst` (writer) implements `ReaderFrom`. If they do, the copy operation can be done more efficiently.
### 2. **Evolving APIs with Optional Interfaces**
When evolving an API, you might want to introduce new functionality without breaking backward compatibility. This can be done using optional interfaces. For example, the Go standard library added support for context management in version 1.7. To accommodate this change, they introduced context-aware versions of existing interfaces, like `StmtExecContext` in the `database/sql/driver` package.
```go
func ctxDriverStmtExec(ctx context.Context, si driver.Stmt,
                       nvdargs []driver.NamedValue) (driver.Result, error) {
    if siCtx, is := si.(driver.StmtExecContext); is {
        return siCtx.ExecContext(ctx, nvdargs)
    }
    // fallback code is here
}
```
This allows code written for older database drivers (which don’t support context) to still work while giving newer drivers the option to implement the context-aware method. This is an example of **optional interfaces** that provide a fallback behavior if the interface is not implemented.
### 3. **Decorator Pattern and Type Assertions**
One drawback of using type assertions or type switches with optional interfaces is when you're using the **decorator pattern** to wrap an interface. In Go, interfaces can be wrapped using structs that embed other implementations. If one of these wrapped implementations implements an optional interface, you can’t always detect it with a type assertion or switch because the wrapper might not expose that interface. For example, in the `bufio` package, `bufio.Reader` wraps any `io.Reader`, but if the underlying reader implements `io.ReaderFrom`, wrapping it in `bufio.Reader` may prevent optimizations.
### 4. **Error Wrapping and Type Switches**
Another case where type switches may not work as expected is with error handling. Errors can be wrapped using the `fmt.Errorf` function, and wrapped errors are not directly accessible with type assertions. To access the wrapped error or check its properties, you should use `errors.Is` and `errors.As`, which are specifically designed to handle error wrapping in Go.
```go
if errors.Is(err, someSpecificError) {
    // handle specific error
}
```
### 5. **Type Switches for Differentiating Implementations**
Type switches are most useful when you need to differentiate between multiple concrete implementations of an interface that require different handling. For example, if you're walking through a tree structure where the nodes can hold different types of values (numbers, operators, etc.), a type switch can help you process them accordingly:
```go
func walkTree(t *treeNode) (int, error) {
    switch val := t.val.(type) {
    case nil:
        return 0, errors.New("invalid expression")
    case number:
        return int(val), nil
    case operator:
        left, err := walkTree(t.lchild)
        if err != nil {
            return 0, err
        }
        right, err := walkTree(t.rchild)
        if err != nil {
            return 0, err
        }
        return val.process(left, right), nil
    default:
        return 0, errors.New("unknown node type")
    }
}
```
Here, `walkTree` differentiates between different types that a node can hold (`number`, `operator`, etc.). If a new type is added later, the default case will catch it and help prevent crashes from missing cases in the switch.
### 6. **Protecting Yourself from Unexpected Interfaces**
To avoid issues with unexpected interface implementations, you can make interfaces **unexported** (i.e., private to the package), or make at least one of their methods unexported. This prevents other packages from embedding and implementing the interface unexpectedly. If a package doesn’t export an interface or method, it has better control over what can implement and interact with it.
## Function Types as a Bridge to Interfaces
Go allows you to declare methods on user-defined types, which can even include function types. This concept might sound niche at first, but it’s incredibly useful when working with interfaces. Let's explore how function types are used as a bridge to interfaces, with a specific focus on how HTTP handlers are implemented in Go.
### The HTTP Handler Example
The `http.Handler` interface is a common example of Go’s use of function types with interfaces. This interface defines a method, `ServeHTTP`, which processes HTTP requests:
```go
type Handler interface {
    ServeHTTP(http.ResponseWriter, *http.Request)
}
```
To make a function of type `func(http.ResponseWriter, *http.Request)` implement the `http.Handler` interface, Go provides `http.HandlerFunc`, which is a type alias for `func(http.ResponseWriter, *http.Request)`:
```go
type HandlerFunc func(http.ResponseWriter, *http.Request)
func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    f(w, r)
}
```
With this, you can now pass functions directly as HTTP handlers, which simplifies the handling of HTTP requests. Here's how you can use it:
```go
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
})
http.ListenAndServe(":8080", nil)
```
In this example, `http.HandleFunc` accepts a function with the signature `func(http.ResponseWriter, *http.Request)`, and automatically wraps it into a `http.HandlerFunc`.
### When to Use Function Types vs. Interfaces
Deciding between using a function type or an interface depends on the complexity of the behavior you need to support.
- **Use a function type**: When the behavior you need to implement is a single, simple function that fits neatly into the parameter list of a function.
- **Use an interface**: When the functionality you need involves multiple methods or requires more flexibility. An interface can encapsulate a range of behaviors, making it more suitable for more complex use cases.
## Dependency Injection with Implicit Interfaces
One of the most powerful features of Go's interfaces is their **implicit nature**. Go doesn’t require you to explicitly declare that a type implements an interface; if a type has the required methods, it automatically satisfies the interface. This is especially useful when implementing **dependency injection**, where you pass interfaces to structs rather than concrete implementations.
Let's take a look at how you can leverage dependency injection in Go with implicit interfaces.
### Logger and DataStore Interfaces
Let’s define some basic components for a simple web app, such as a logger and a data store:
```go
type Logger interface {
    Log(message string)
}
type DataStore interface {
    UserNameForID(userID string) (string, bool)
}
type SimpleDataStore struct {
    userData map[string]string
}
func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
    name, ok := sds.userData[userID]
    return name, ok
}
func NewSimpleDataStore() SimpleDataStore {
    return SimpleDataStore{
        userData: map[string]string{
            "1": "Fred",
            "2": "Mary",
            "3": "Pat",
        },
    }
}
type LoggerAdapter func(message string)
func (lg LoggerAdapter) Log(message string) {
    lg(message)
}
```
The `Logger` and `DataStore` interfaces describe the functionality that our business logic will depend on. Notice that `LoggerAdapter` is just a function type that satisfies the `Logger` interface by implementing the `Log` method.
### Business Logic
Now, let’s define the business logic that uses these interfaces. Our logic needs to be able to log messages and look up users in the data store:
```go
type SimpleLogic struct {
    l  Logger
    ds DataStore
}
func (sl SimpleLogic) SayHello(userID string) (string, error) {
    sl.l.Log("in SayHello for " + userID)
    name, ok := sl.ds.UserNameForID(userID)
    if !ok {
        return "", errors.New("unknown user")
    }
    return "Hello, " + name, nil
}
func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
    sl.l.Log("in SayGoodbye for " + userID)
    name, ok := sl.ds.UserNameForID(userID)
    if !ok {
        return "", errors.New("unknown user")
    }
    return "Goodbye, " + name, nil
}
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
    return SimpleLogic{
        l:    l,
        ds: ds,
    }
}
```
The `SimpleLogic` struct has dependencies on `Logger` and `DataStore` interfaces. This allows you to inject any concrete implementation of these interfaces into `SimpleLogic`.
### Controller and API Endpoint
Finally, let’s define the controller that wires everything up and serves HTTP requests:
```go
type Logic interface {
    SayHello(userID string) (string, error)
}
type Controller struct {
    l     Logger
    logic Logic
}
func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
    c.l.Log("In SayHello")
    userID := r.URL.Query().Get("user_id")
    message, err := c.logic.SayHello(userID)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
        return
    }
    w.Write([]byte(message))
}
func NewController(l Logger, logic Logic) Controller {
    return Controller{
        l:     l,
        logic: logic,
    }
}
```
Here, the `Controller` struct takes a `Logger` and a `Logic` interface. This decouples the controller from any specific implementation of these interfaces, allowing for easy swapping of dependencies.
### Wiring Up and Running the Server
Finally, in your `main` function, you can wire up all the components and start the server:
```go
func main() {
    l := LoggerAdapter(LogOutput)
    ds := NewSimpleDataStore()
    logic := SimpleLogic{l: l, ds: ds}
    c := NewController(l, logic)
    http.HandleFunc("/hello", c.SayHello)
    http.ListenAndServe(":8080", nil)
}
```
In this `main` function:
- The logger, data store, and logic are all created and injected into the `Controller`.
- The HTTP handler is set up, and the server starts.
## Generics
### **Generics Reduce Repetitive Code and Increase Type Safety**
Go is a statically typed language, meaning the types of variables and function parameters are checked at compile time. Built-in types like slices, maps, and channels, as well as functions like `len`, `cap`, and `make`, can work with multiple concrete types. However, before Go 1.18, user-defined types and functions did not have this flexibility.
If you come from dynamically typed languages, where type checking happens at runtime, you might not see the immediate benefit of generics. A helpful way to think about them is **“type parameters”**—just as functions take value parameters that are determined at runtime, generics allow functions and types to take **type parameters** that are determined at compile time.
So far, you've seen functions that take specific types. In the example below, `divAndRemainder` takes two `int` values and returns two `int` values along with an `error`:
```go
func divAndRemainder(num, denom int) (int, int, error) {
    if denom == 0 {
        return 0, 0, errors.New("cannot divide by zero")
    }
    return num / denom, num % denom, nil
}
```
Similarly, structs have fields with specific types when they are declared. Here, `Node` has an `int` value and a pointer to another `Node`:
```go
type Node struct {
    val  int
    next *Node
}
```
But what if you need a **binary tree** that can hold not just `int` values, but also `float64`, `string`, or even a custom type? Without generics, you have two options:
1. **Write separate tree implementations for each type**—but that leads to **repetitive code**.
2. **Use interfaces**—but that **removes type safety** and incurs runtime overhead.
### **Using Interfaces as a Workaround**
Before generics, one way to create a tree that works with multiple types was to use an interface:
```go
type Orderable interface {
    Order(any) int
}
```
The `Order` method determines how to compare two values:
- Returns `< 0` if the value is **less than** the other.
- Returns `> 0` if the value is **greater than** the other.
- Returns `0` if the values are **equal**.
Now, you can modify the `Tree` struct to use `Orderable`:
```go
type Tree struct {
    val         Orderable
    left, right *Tree
}
func (t *Tree) Insert(val Orderable) *Tree {
    if t == nil {
        return &Tree{val: val}
    }
    switch comp := val.Order(t.val); {
    case comp < 0:
        t.left = t.left.Insert(val)
    case comp > 0:
        t.right = t.right.Insert(val)
    }
    return t
}
```
To use this tree with `int` values, you define a concrete type:
```go
type OrderableInt int
func (oi OrderableInt) Order(val any) int {
    return int(oi - val.(OrderableInt))
}
```
Now you can create and insert values:
```go
var it *Tree
it = it.Insert(OrderableInt(5))
it = it.Insert(OrderableInt(3))
```
However, this approach has a **major problem**: it allows **mixed types**.
If you create another `Orderable` type:
```go
type OrderableString string
func (os OrderableString) Order(val any) int {
    return strings.Compare(string(os), val.(string))
}
```
And mistakenly insert an `OrderableString` into the `Tree`:
```go
var it *Tree
it = it.Insert(OrderableInt(5))
it = it.Insert(OrderableString("nope"))
```
The compiler **does not catch the error**. Instead, the program panics at runtime:
```
panic: interface conversion: interface {} is main.OrderableInt, not string
```
This defeats Go’s strong type safety. Fortunately, **generics** solve this problem.
## **Introducing Generics in Go**
Go 1.18 introduced **type parameters**, allowing you to write reusable and type-safe code. Let’s look at a **stack** as an example.
A **stack** is a data structure that follows **Last In, First Out (LIFO)** ordering, like a pile of plates: you place new plates on top and remove them from the top.
Here’s how you can define a **generic stack** in Go:
```go
type Stack[T any] struct {
    vals []T
}
func (s *Stack[T]) Push(val T) {
    s.vals = append(s.vals, val)
}
func (s *Stack[T]) Pop() (T, bool) {
    if len(s.vals) == 0 {
        var zero T
        return zero, false
    }
    top := s.vals[len(s.vals)-1]
    s.vals = s.vals[:len(s.vals)-1]
    return top, true
}
```
### **How It Works**
1. `[T any]` defines a **type parameter** named `T` that can be **any type**.
2. `Stack[T]` means the stack stores **values of type `T`**.
3. `Push(val T)` takes a **value of type `T`** and adds it to the stack.
4. `Pop() (T, bool)` removes and returns the last value. If the stack is empty, it returns the **zero value** of `T`.
### **Using the Generic Stack**
```go
func main() {
    var intStack Stack[int]
    intStack.Push(10)
    intStack.Push(20)
    intStack.Push(30)
    v, ok := intStack.Pop()
    fmt.Println(v, ok)  // Output: 30 true
}
```
If you try to push a string into `intStack`:
```go
intStack.Push("nope")
```
The compiler catches the error:
```
cannot use "nope" (untyped string constant) as int value
  in argument to intStack.Push
```
Unlike the `Orderable` interface, **this error happens at compile time**.
### **Generics Constraints**
Let’s say you want to check whether a stack **contains** a certain value:
```go
func (s Stack[T]) Contains(val T) bool {
    for _, v := range s.vals {
        if v == val {
            return true
        }
    }
    return false
}
```
This **does not compile** because `T` could be any type, including types that **don’t support `==`**.
```
invalid operation: v == val (type parameter T is not comparable with ==)
```
To fix this, you use **the built-in `comparable` constraint**:
```go
type Stack[T comparable] struct {
    vals []T
}
```
Now, `T` is restricted to types that support `==` and `!=`, like `int`, `float64`, `string`, etc.
## **Generic Functions: Abstracting Algorithms**
Before generics, writing functions like `map`, `reduce`, and `filter` that work across different types was cumbersome. Now, we can define them using type parameters:
### **1. Map Function**
Transforms a slice of one type into a slice of another type using a mapping function.
```go
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
    r := make([]T2, len(s))
    for i, v := range s {
        r[i] = f(v)
    }
    return r
}
```
### **2. Reduce Function**
Aggregates values in a slice to a single result using an accumulator function.
```go
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
    r := initializer
    for _, v := range s {
        r = f(r, v)
    }
    return r
}
```
### **3. Filter Function**
Returns a new slice containing only the elements that satisfy a given condition.
```go
func Filter[T any](s []T, f func(T) bool) []T {
    var r []T
    for _, v := range s {
        if f(v) {
            r = append(r, v)
        }
    }
    return r
}
```
### **Example Usage**
```go
words := []string{"One", "Potato", "Two", "Potato"}
filtered := Filter(words, func(s string) bool {
    return s != "Potato"
})
fmt.Println(filtered) // Output: [One Two]
lengths := Map(filtered, func(s string) int {
    return len(s)
})
fmt.Println(lengths) // Output: [3 3]
sum := Reduce(lengths, 0, func(acc int, val int) int {
    return acc + val
})
fmt.Println(sum) // Output: 6
```
## **Generics and Interfaces**
In addition to `any`, we can use interfaces as type constraints.
### **Using Interfaces as Type Constraints**
Suppose we need a `Pair` struct that only accepts types implementing `fmt.Stringer`:
```go
type Pair[T fmt.Stringer] struct {
    Val1 T
    Val2 T
}
```
We can also define an interface with a generic type:
```go
type Differ[T any] interface {
    fmt.Stringer
    Diff(T) float64
}
```
Now, we can create a function that finds the closer of two pairs:
```go
func FindCloser[T Differ[T]](pair1, pair2 Pair[T]) Pair[T] {
    d1 := pair1.Val1.Diff(pair1.Val2)
    d2 := pair2.Val1.Diff(pair2.Val2)
    if d1 < d2 {
        return pair1
    }
    return pair2
}
```
### **Defining Structs That Implement `Differ`**
```go
type Point2D struct { X, Y int }
func (p2 Point2D) String() string {
    return fmt.Sprintf("{%d,%d}", p2.X, p2.Y)
}
func (p2 Point2D) Diff(from Point2D) float64 {
    x, y := p2.X-from.X, p2.Y-from.Y
    return math.Sqrt(float64(x*x + y*y))
}
```
```go
type Point3D struct { X, Y, Z int }
func (p3 Point3D) String() string {
    return fmt.Sprintf("{%d,%d,%d}", p3.X, p3.Y, p3.Z)
}
func (p3 Point3D) Diff(from Point3D) float64 {
    x, y, z := p3.X-from.X, p3.Y-from.Y, p3.Z-from.Z
    return math.Sqrt(float64(x*x + y*y + z*z))
}
```
### **Example Usage**
```go
pair2Da := Pair[Point2D]{Point2D{1, 1}, Point2D{5, 5}}
pair2Db := Pair[Point2D]{Point2D{10, 10}, Point2D{15, 5}}
closer := FindCloser(pair2Da, pair2Db)
fmt.Println(closer)
pair3Da := Pair[Point3D]{Point3D{1, 1, 10}, Point3D{5, 5, 0}}
pair3Db := Pair[Point3D]{Point3D{10, 10, 10}, Point3D{11, 5, 0}}
closer2 := FindCloser(pair3Da, pair3Db)
fmt.Println(closer2)
```
## **Using Type Terms for Operators**
Some functions require operations like `/` and `%`, which aren't defined for all types. We can define a constraint for integer types:
```go
type Integer interface {
    int | int8 | int16 | int32 | int64 |
    uint | uint8 | uint16 | uint32 | uint64 | uintptr
}
```
Now we can write a generic `divAndRemainder` function:
```go
func divAndRemainder[T Integer](num, denom T) (T, T, error) {
    if denom == 0 {
        return 0, 0, errors.New("cannot divide by zero")
    }
    return num / denom, num % denom, nil
}
```
### **Example Usage**
```go
var a uint = 18_446_744_073_709_551_615
var b uint = 9_223_372_036_854_775_808
fmt.Println(divAndRemainder(a, b))
```
### **Allowing User-Defined Types**
By default, type constraints match exactly, meaning user-defined types won't work unless explicitly included:
```go
type MyInt int
```
This will fail:
```go
var myA MyInt = 10
var myB MyInt = 20
fmt.Println(divAndRemainder(myA, myB))
```
To fix it, we use `~` to allow types with the same underlying type:
```go
type Integer interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
```
### **Comparison Constraints**
The `Ordered` constraint allows generic comparison functions:
```go
type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
    ~float32 | ~float64 |
    ~string
}
```
Go 1.21 introduced `cmp` with:
- `cmp.Compare(a, b)`: Returns `1`, `0`, or `1` for `<`, `==`, or `>`.
- `cmp.Less(a, b)`: Returns `true` if `a < b`.
### **Mixing Methods and Type Terms**
You can enforce that a type has a method and a specific underlying type:
```go
type PrintableInt interface {
    ~int
    String() string
}
```
This ensures `String()` is implemented while allowing custom types.
## **Type Inference and Generics in Go**
Go simplifies working with generics by supporting type inference, just like it does with the `:=` operator. This can reduce boilerplate code when calling generic functions, making the code cleaner and easier to read. However, there are cases where Go can't infer types, particularly when a type parameter is only used as a return value, and you must explicitly specify the type arguments.
### Example: Type Inference Doesn’t Always Work
```go
type Integer interface {
    int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}
func Convert[T1, T2 Integer](in T1) T2 {
    return T2(in)
}
func main() {
    var a int = 10
    b := Convert[int, int64](a) // Can't infer return type
    fmt.Println(b)
}
```
In this case, Go can't infer the return type, so you must explicitly provide the type arguments (`int, int64`) when calling `Convert`.
## **Type Elements Limit Constants**
Generics in Go can limit the types and constants that are assignable to variables. If a generic type constraint is too broad, such as the `Ordered` interface, it might not be able to handle constants that are valid for all the types in the constraint. Here's an example:
```go
// INVALID: 1,000 can't fit into an 8-bit integer
func PlusOneThousand[T Integer](in T) T {
    return in + 1_000
}
```
However, if the value is smaller, like 100, the code will work because 100 can fit into all the integer types in the `Integer` interface:
```go
// VALID
func PlusOneHundred[T Integer](in T) T {
    return in + 100
}
```
## **Combining Generic Functions with Data Structures**
One powerful use of generics is combining them with data structures. A common example is creating a generic tree, like a binary search tree, that works for any concrete type. To compare the elements in the tree, you’ll need a function that defines the ordering logic.
Here’s how you can combine generics with a tree structure:
```go
type OrderableFunc[T any] func(t1, t2 T) int
type Tree[T any] struct {
    f    OrderableFunc[T]
    root *Node[T]
}
type Node[T any] struct {
    val         T
    left, right *Node[T]
}
func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
    return &Tree[T]{f: f}
}
func (t *Tree[T]) Add(v T) {
    t.root = t.root.Add(t.f, v)
}
func (t *Tree[T]) Contains(v T) bool {
    return t.root.Contains(t.f, v)
}
func (n *Node[T]) Add(f OrderableFunc[T], v T) *Node[T] {
    if n == nil {
        return &Node[T]{val: v}
    }
    switch r := f(v, n.val); {
    case r <= -1:
        n.left = n.left.Add(f, v)
    case r >= 1:
        n.right = n.right.Add(f, v)
    }
    return n
}
func (n *Node[T]) Contains(f OrderableFunc[T], v T) bool {
    if n == nil {
        return false
    }
    switch r := f(v, n.val); {
    case r <= -1:
        return n.left.Contains(f, v)
    case r >= 1:
        return n.right.Contains(f, v)
    }
    return true
}
```
You can create trees with any concrete type and an ordering function. For example, here's how you'd create a tree of integers:
```go
t1 := NewTree(cmp.Compare[int])
t1.Add(10)
t1.Add(30)
t1.Add(15)
fmt.Println(t1.Contains(15)) // true
fmt.Println(t1.Contains(40)) // false
```
For structs, you can use a function that defines how to compare them:
```go
type Person struct {
    Name string
    Age  int
}
func OrderPeople(p1, p2 Person) int {
    out := cmp.Compare(p1.Name, p2.Name)
    if out == 0 {
        out = cmp.Compare(p1.Age, p2.Age)
    }
    return out
}
t2 := NewTree(OrderPeople)
t2.Add(Person{"Bob", 30})
t2.Add(Person{"Maria", 35})
t2.Add(Person{"Bob", 50})
fmt.Println(t2.Contains(Person{"Bob", 30})) // true
fmt.Println(t2.Contains(Person{"Fred", 25})) // false
```
Alternatively, you can define a method on the struct and use it with the tree:
```go
func (p Person) Order(other Person) int {
    out := cmp.Compare(p.Name, other.Name)
    if out == 0 {
        out = cmp.Compare(p.Age, other.Age)
    }
    return out
}
t3 := NewTree(Person.Order)
t3.Add(Person{"Bob", 30})
t3.Add(Person{"Maria", 35})
t3.Add(Person{"Bob", 50})
fmt.Println(t3.Contains(Person{"Bob", 30})) // true
fmt.Println(t3.Contains(Person{"Fred", 25})) // false
```
## **Working with Comparable Types**
In Go, interfaces are comparable types, meaning you can compare them using `==` and `!=`, but only if the underlying types are themselves comparable. Here's how that interacts with generics:
```go
type Thinger interface {
    Thing()
}
type ThingerInt int
func (t ThingerInt) Thing() {
    fmt.Println("ThingInt:", t)
}
type ThingerSlice []int
func (t ThingerSlice) Thing() {
    fmt.Println("ThingSlice:", t)
}
func Comparer[T comparable](t1, t2 T) {
    if t1 == t2 {
        fmt.Println("equal!")
    }
}
var a ThingerInt = 10
var b ThingerInt = 10
Comparer(a, b) // prints "equal!"
```
However, if you try to compare slices, which aren't comparable by default, you'll get a compile-time error:
```go
var a3 ThingerSlice = []int{1, 2, 3}
var b3 ThingerSlice = []int{1, 2, 3}
Comparer(a3, b3) // compile fails: 'ThingerSlice does not satisfy comparable'
```
Interestingly, you can assign `ThingerSlice` to a variable of type `Thinger` because `Thinger` is an interface:
```go
var a4 Thinger = a3
var b4 Thinger = b3
Comparer(a4, b4) // compiles but panics at runtime
```
This compiles but will panic at runtime because slices are not comparable. The program will panic with:
```
panic: runtime error: comparing uncomparable type main.ThingerSlice
```
## **Idiomatic Go and Generics**
With generics in Go, some coding habits will change:
- Instead of using `interface{}`, use `any`.
- Instead of using `float64` as a catch-all numeric type, use generics to support multiple number types.
- Instead of writing separate functions for `[]int` and `[]string`, write one generic function.
But don’t rush to rewrite everything using generics! Generics are **not always faster** and might even slow down performance in some cases.
For example, this **old-style function** using interfaces:
```go
type Ager interface {
    Age() int
}
func doubleAge(a Ager) int {
    return a.Age() * 2
}
```
is **faster** than the equivalent generic version:
```go
func doubleAgeGeneric[T Ager](a T) int {
    return a.Age() * 2
}
```
Why? Because Go **adds runtime checks** to differentiate types when it compiles generics. So, unless generics make your code **clearer and more reusable**, stick to interfaces.
## **Adding Generics to the Standard Library**
When Go first introduced generics in **Go 1.18**, the standard library didn’t change much. By **Go 1.21**, new generic functions were added, such as:
- **`slices.Insert/Delete`** – Simplifies working with slices.
- **`maps.Clone`** – Efficiently copies maps.
- **`sync.OnceValue/OnceValues`** – Ensures a function runs only once.
Expect more generics-based utilities in future Go versions.
## Errors
### Handling Errors: The Basics
Go handles errors by returning an `error` value as the last return value of a function. This is a convention rather than a language feature, but it's such a fundamental convention that it should never be broken. When a function executes successfully, it returns `nil` for the error. If something goes wrong, an error value is returned instead. The calling function then checks this value and either handles the error or propagates it further.
A basic function with error handling looks like this:
```go
func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
    if denominator == 0 {
        return 0, 0, errors.New("denominator is 0")
    }
    return numerator / denominator, numerator % denominator, nil
}
```
The `errors.New` function creates a new error from a string. When writing error messages, avoid capitalization, punctuation, or trailing newlines.
To check for errors, use an `if` statement:
```go
func main() {
    numerator := 20
    denominator := 3
    remainder, mod, err := calcRemainderAndMod(numerator, denominator)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(remainder, mod)
}
```
Here, if an error occurs, the program prints it and exits with a nonzero status.
### Why Go Uses Error Values Instead of Exceptions
Unlike languages with exceptions, Go doesn't have special constructs for detecting errors. Instead, errors are explicit return values. There are two key reasons for this approach:
1. **Control over execution flow**
    
    Exceptions introduce hidden execution paths, which can make it unclear where errors may arise. Without explicit declarations, functions might throw exceptions unexpectedly, leading to crashes or subtle data corruption.
    
2. **Compiler checks enforce error handling**
    
    In Go, all variables must be read at least once. This means that ignoring an error is a conscious choice, often indicated by using `_` as the return variable. This makes unhandled errors explicit.
    
While exception handling can produce shorter code, Go prioritizes clarity over brevity. Error-handling code is indented inside `if` statements, while normal logic remains at the base indentation level, making the “golden path” of execution visually clear.
## Using Strings for Simple Errors
Go provides two ways to create an error from a string:
1. **Using `errors.New`**
    
    This function creates an error with a fixed message:
    
    ```go
    func doubleEven(i int) (int, error) {
        if i%2 != 0 {
            return 0, errors.New("only even numbers are processed")
        }
        return i * 2, nil
    }
    ```
    
2. **Using `fmt.Errorf`**
    
    This function allows for formatted error messages:
    
    ```go
    func doubleEven(i int) (int, error) {
        if i%2 != 0 {
            return 0, fmt.Errorf("%d isn't an even number", i)
        }
        return i * 2, nil
    }
    ```
    
The `fmt.Errorf` approach is useful when including dynamic values in the error message.
## Sentinel Errors
Some errors indicate that an operation cannot proceed due to a specific state. In a well-known blog post, ["Don’t Just Check Errors, Handle Them Gracefully,"](https://dave.cheney.net/2014/12/24/inspecting-errors) Dave Cheney coined the term **sentinel errors** to describe these.
Sentinel errors are declared as package-level variables and usually start with `Err` (except for `io.EOF`):
```go
var ErrNotFound = errors.New("not found")
```
Sentinel errors are used when an operation encounters a known condition that prevents further execution. The standard library provides several examples, such as `zip.ErrFormat`, which indicates that a file isn't a valid ZIP archive:
```go
func main() {
    data := []byte("This is not a zip file")
    notAZipFile := bytes.NewReader(data)
    _, err := zip.NewReader(notAZipFile, int64(len(data)))
    if err == zip.ErrFormat {
        fmt.Println("Told you so")
    }
}
```
Another example is `rsa.ErrMessageTooLong` from `crypto/rsa`, which signals that a message is too large to be encrypted with a given key.
### When to Use Sentinel Errors
Define a sentinel error only if:
- It represents a specific, well-defined state.
- No additional context is needed beyond its presence.
Since a sentinel error becomes part of a package’s public API, it must be maintained for backward compatibility. In many cases, it’s better to define a custom error type instead.
### Checking Sentinel Errors
To check for a sentinel error, use the `==` operator:
```go
if err == zip.ErrFormat {
    fmt.Println("Invalid ZIP file")
}
```
In later sections, we’ll discuss better ways to check for errors, including `errors.Is`.
## Errors Are Values
### Custom Errors with Additional Information
Since `error` is an interface in Go, you can define your own error types that include additional information for logging or error handling. For example, you might want to include a status code to indicate the kind of error being reported, avoiding fragile string comparisons.
### Defining Status Codes
```go
// Define an enumeration for status codes
type Status int
const (
    InvalidLogin Status = iota + 1
    NotFound
)
```
### Creating a Custom Error Type
```go
type StatusErr struct {
    Status  Status
    Message string
}
func (se StatusErr) Error() string {
    return se.Message
}
```
### Using the Custom Error Type
```go
func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
    token, err := login(uid, pwd)
    if err != nil {
        return nil, StatusErr{
            Status:  InvalidLogin,
            Message: fmt.Sprintf("invalid credentials for user %s", uid),
        }
    }
    data, err := getData(token, file)
    if err != nil {
        return nil, StatusErr{
            Status:  NotFound,
            Message: fmt.Sprintf("file %s not found", file),
        }
    }
    return data, nil
}
```
### Best Practices for Returning Errors
Even when using custom error types, always return `error` as the error type to maintain flexibility. This allows different error types to be returned and enables callers to avoid depending on a specific error type.
### Avoid Uninitialized Custom Errors
```go
func GenerateErrorBroken(flag bool) error {
    var genErr StatusErr // Uninitialized instance
    if flag {
        genErr = StatusErr{
            Status: NotFound,
        }
    }
    return genErr // Always non-nil because error is an interface
}
```
Running this function results in unexpected non-nil errors due to Go’s interface nil behavior.
### Correcting the Error Handling
**Approach 1: Explicitly Return nil**
```go
func GenerateErrorOKReturnNil(flag bool) error {
    if flag {
        return StatusErr{
            Status: NotFound,
        }
    }
    return nil
}
```
**Approach 2: Define Error Variable as `error`**
```go
func GenerateErrorUseErrorVar(flag bool) error {
    var genErr error
    if flag {
        genErr = StatusErr{
            Status: NotFound,
        }
    }
    return genErr
}
```
> Warning: Never declare a variable of your custom error type. Use error instead to prevent unintended behavior.
> 
## Wrapping Errors
When propagating errors, it's useful to add additional context while preserving the original error. This technique is called **wrapping errors**.
### Using `fmt.Errorf` for Wrapping
```go
func fileChecker(name string) error {
    f, err := os.Open(name)
    if err != nil {
        return fmt.Errorf("in fileChecker: %w", err) // Wraps the error
    }
    f.Close()
    return nil
}
```
### Unwrapping Errors with `errors.Unwrap`
```go
func main() {
    err := fileChecker("not_here.txt")
    if err != nil {
        fmt.Println(err)
        if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
            fmt.Println(wrappedErr)
        }
    }
}
```
> Note: Instead of calling errors.Unwrap directly, use errors.Is or errors.As to check for specific errors.
> 
### Custom Errors with Wrapping Support
To allow error unwrapping, a custom error type should implement the `Unwrap` method.
```go
type StatusErr struct {
    Status  Status
    Message string
    Err     error
}
func (se StatusErr) Error() string {
    return se.Message
}
func (se StatusErr) Unwrap() error {
    return se.Err
}
```
### Example Usage
```go
func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
    token, err := login(uid, pwd)
    if err != nil {
        return nil, StatusErr{
            Status:  InvalidLogin,
            Message: fmt.Sprintf("invalid credentials for user %s", uid),
            Err:     err,
        }
    }
    data, err := getData(token, file)
    if err != nil {
        return nil, StatusErr{
            Status:  NotFound,
            Message: fmt.Sprintf("file %s not found", file),
            Err:     err,
        }
    }
    return data, nil
}
```
### When to Wrap Errors
Not all errors need to be wrapped. If an error contains unnecessary implementation details, you can create a new error instead.
### Creating a New Error Without Wrapping
```go
err := internalFunction()
if err != nil {
    return fmt.Errorf("internal failure: %v", err) // Uses %v instead of %w
}
```
By using `%v` instead of `%w`, you avoid including the original error for further unwrapping.
## Wrapping Multiple Errors in Go
### Why Wrap Multiple Errors?
When a function encounters multiple errors, you might want to return them together instead of just the first one. For example, validating a struct with multiple fields might result in multiple errors. However, Go’s standard practice is to return a single `error`, not a `[]error`.
To handle this, you can use `errors.Join` to merge multiple errors into a single `error`.
### Using `errors.Join` to Combine Errors
Example: Validating a `User` struct and returning all errors at once.
```go
type User struct {
    Name  string
    Email string
    Age   int
}
func ValidateUser(u User) error {
    var errs []error
    if u.Name == "" {
        errs = append(errs, errors.New("name is required"))
    }
    if u.Email == "" {
        errs = append(errs, errors.New("email is required"))
    }
    if u.Age < 18 {
        errs = append(errs, errors.New("age must be 18 or older"))
    }
    if len(errs) > 0 {
        return errors.Join(errs...)
    }
    return nil
}
```
Now, calling `ValidateUser` will return a **single** error containing all the validation issues.
### Using `fmt.Errorf` for Multiple Errors
Another way to wrap multiple errors is by using `fmt.Errorf` with `%w`:
```go
err1 := errors.New("failed to connect to DB")
err2 := errors.New("timeout occurred")
err3 := errors.New("invalid credentials")
err := fmt.Errorf("db error: %w, network error: %w, auth error: %w", err1, err2, err3)
```
This approach wraps each error separately and allows unwrapping later.
### Custom Error Type for Multiple Errors
Instead of using `errors.Join`, you can create a custom error type that holds multiple errors:
```go
type MultiError struct {
    Errors []error
}
func (m MultiError) Error() string {
    return errors.Join(m.Errors...).Error()
}
func (m MultiError) Unwrap() []error {
    return m.Errors
}
```
**How it works:**
- `Error()` returns all errors as a single string.
- `Unwrap()` allows retrieving the individual errors.
⚠ **Note:** Go’s `errors.Unwrap` function does **not** support `Unwrap() []error`, so calling `errors.Unwrap` directly on `MultiError` won’t work.
### Checking Wrapped Errors
Since Go does not natively support `Unwrap() []error`, you must use a **type switch** to extract wrapped errors:
```go
var err error = someFunction()
switch err := err.(type) {
case interface { Unwrap() error }:
    innerErr := err.Unwrap()
    // Process innerErr
case interface { Unwrap() []error }:
    innerErrs := err.Unwrap()
    for _, e := range innerErrs {
        // Process each error
    }
default:
    // Handle single error
}
```
✅ **Best practice:** Instead of manually unwrapping, use `errors.Is` and `errors.As` for checking errors.
## Is and As
### Using `errors.Is` to Check for Specific Errors
When errors are wrapped, comparing with `==` no longer works. Use `errors.Is` to check if an error exists **anywhere** in the chain.
### Example: Detecting `os.ErrNotExist`
```go
func openFile(name string) error {
    _, err := os.Open(name)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    return nil
}
func main() {
    err := openFile("missing.txt")
    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("File does not exist")
    }
}
```
### Customizing `errors.Is`
By default, `errors.Is` uses `==` for comparison. If your custom error type is non-comparable, define an `Is` method:
```go
type MyError struct {
    Code int
}
func (m MyError) Error() string {
    return fmt.Sprintf("error code: %d", m.Code)
}
func (m MyError) Is(target error) bool {
    if other, ok := target.(MyError); ok {
        return m.Code == other.Code
    }
    return false
}
```
Now `errors.Is` can check for logical equivalence instead of strict identity.
### Matching Errors with `errors.Is`
To make error matching more flexible, implement `Is` to allow partial matches.
### Example: Checking for Specific Resource Errors
```go
type ResourceError struct {
    Resource string
    Code     int
}
func (r ResourceError) Error() string {
    return fmt.Sprintf("%s: %d", r.Resource, r.Code)
}
func (r ResourceError) Is(target error) bool {
    if other, ok := target.(ResourceError); ok {
        matchResource := other.Resource == "" || other.Resource == r.Resource
        matchCode := other.Code == 0 || other.Code == r.Code
        return matchResource && matchCode
    }
    return false
}
```
Now, checking for any database-related error is easy:
```go
if errors.Is(err, ResourceError{Resource: "Database"}) {
    fmt.Println("Database error occurred:", err)
}
```
### Using `errors.As` to Extract Error Types
When checking for a **specific error type**, use `errors.As`.
```go
err := someFunction()
var myErr MyError
if errors.As(err, &myErr) {
    fmt.Println("Error code:", myErr.Code)
}
```
✅ `errors.As` searches the entire error chain and assigns the found error to `myErr`.
### Extracting Errors by Interface
Instead of checking for a concrete type, use an interface:
```go
err := someFunction()
var coder interface {
    GetCode() int
}
if errors.As(err, &coder) {
    fmt.Println("Error code:", coder.GetCode())
}
```
🚨 **Warning:** `errors.As` panics if the second argument is not a pointer.
### `errors.Is` vs `errors.As`
| Function | Purpose |
| --- | --- |
| **errors.Is** | Check if a specific error **instance** exists in the chain |
| **errors.As** | Check if a specific **type** exists and extract it |
✔ **Use `errors.Is` when comparing against known errors (sentinels).**
✔ **Use `errors.As` when extracting custom error types.**
## Wrapping Errors with `defer`
### Why Use `defer` for Error Wrapping?
If you have multiple error checks that all wrap errors with the same message, `defer` can simplify the code.
### Without `defer`: Repeating the Same Error Wrapping
```go
func DoSomeThings(val1 int, val2 string) (string, error) {
    val3, err := doThing1(val1)
    if err != nil {
        return "", fmt.Errorf("in DoSomeThings: %w", err)
    }
    val4, err := doThing2(val2)
    if err != nil {
        return "", fmt.Errorf("in DoSomeThings: %w", err)
    }
    result, err := doThing3(val3, val4)
    if err != nil {
        return "", fmt.Errorf("in DoSomeThings: %w", err)
    }
    return result, nil
}
```
This repeats `fmt.Errorf("in DoSomeThings: %w", err)` multiple times.
### With `defer`: Cleaner and Less Repetitive Code
```go
func DoSomeThings(val1 int, val2 string) (_ string, err error) {
    defer func() {
        if err != nil {
            err = fmt.Errorf("in DoSomeThings: %w", err)
        }
    }()
    val3, err := doThing1(val1)
    if err != nil {
        return "", err
    }
    val4, err := doThing2(val2)
    if err != nil {
        return "", err
    }
    return doThing3(val3, val4)
}
```
### How It Works:
- The return value `err` is named, allowing access to it inside `defer`.
- `defer` only modifies `err` if it is not `nil`.
- This approach avoids redundant wrapping at each error check.
✅ **Best Use Case:** When all errors should be wrapped in the same message.
⚠ **Limitation:** If different errors need different messages, `fmt.Errorf` should still be used individually.
## Panic and Recover
### What is a Panic?
A panic in Go is similar to an `Error` in Java or Python. It occurs when the Go runtime encounters an unrecoverable state, typically due to a programming mistake such as:
- Accessing an out-of-bounds slice index
- Passing an invalid argument (e.g., negative size to `make`)
- Internal runtime issues (though these are rare)
When a panic occurs:
1. The function where the panic happened exits immediately.
2. Any `defer` statements in that function execute.
3. The panic propagates up the call stack, triggering defers along the way.
4. If the panic reaches `main()`, the program crashes with an error message and stack trace.
> Note: If a panic happens in a goroutine other than main, it stops at the function that launched the goroutine. The program exits unless the panic is recovered.
> 
### Creating a Panic
You can manually trigger a panic using the built-in `panic` function. It takes any type of value, though a string is commonly used.
```go
func doPanic(msg string) {
    panic(msg)
}
func main() {
    doPanic("Something went wrong!")
}
```
Running this code produces:
```
panic: Something went wrong!
goroutine 1 [running]:
main.doPanic(...)
    /tmp/sandbox/prog.go:6
main.main()
    /tmp/sandbox/prog.go:10 +0x5f
```
### Recovering from a Panic
Go provides `recover()` to handle panics gracefully. It must be called inside a deferred function.
```go
func div60(i int) {
    defer func() {
        if v := recover(); v != nil {
            fmt.Println("Recovered from panic:", v)
        }
    }()
    fmt.Println(60 / i) // This will panic if i is 0
}
func main() {
    for _, val := range []int{1, 2, 0, 6} {
        div60(val)
    }
}
```
**Output:**
```
60
30
Recovered from panic: runtime error: integer divide by zero
10
```
> Important: recover() only works inside a deferred function. Otherwise, the panic will continue propagating.
> 
### When to Use `panic` and `recover`
- **Use `panic`** for truly fatal errors (e.g., corruption, unrecoverable state).
- **Use `recover`** to log and gracefully handle panics in specific cases, like inside a library.
- **Avoid using panic/recover for error handling.** Instead, return errors explicitly.
> Example: In a library, catch panics at the API boundary and return an error instead.
> 
```go
func SafeOperation() (err error) {
    defer func() {
        if v := recover(); v != nil {
            err = fmt.Errorf("operation failed: %v", v)
        }
    }()
    doSomethingRisky()
    return nil
}
```
## Getting a Stack Trace from an Error
By default, Go does not print a stack trace for errors. However, you can:
> Tip: To avoid exposing full file paths in errors, compile with -trimpath.
> 
- Manually build a call stack using error wrapping.
- Use third-party libraries (e.g., CockroachDB’s `errors` package) to generate stack traces.
- Print errors with `%+v` to see full details.
## **Modules, Packages, and Imports**
### Repositories, Modules, and Packages
A **repository** is the familiar place where the source code for a project is stored. This is typically a version-controlled location, like a GitHub repository. A **module** is a collection of Go source code, versioned and distributed as a single unit, typically stored in a repository. A **module** contains one or more **packages**, which are directories containing Go source code. Packages give a module structure and organization.
### Note on Repository Structure:
While it's possible to store multiple modules in one repository, it’s generally discouraged. Each module is versioned together, and keeping multiple modules in one repository complicates version management. This practice should be avoided to maintain clarity and version control integrity.
Also, it’s helpful to understand the differences in terminology across languages. For instance, in Java, a repository holds multiple artifacts, which is akin to a Go **module**. Meanwhile, Go’s **package** concept is similar to **Node.js modules**, but the terms are swapped between the two languages.
In Go, the **module path** is a globally unique identifier. The module path is typically based on the repository URL. For example, a module path like `github.com/jonbodner/proteus` points to the repository URL. You can create local, non-unique modules for personal projects, but if you intend to share your code, the module path must be globally unique.
## Using go.mod
A directory becomes a Go module when it contains a `go.mod` file. Instead of creating this file manually, you should use the `go mod` command to manage modules. The `go mod init MODULE_PATH` command generates a `go.mod` file and sets the current directory as the module root. The `MODULE_PATH` must be globally unique and case-sensitive, so it’s best to use all lowercase to avoid confusion.
Here’s an example of what a `go.mod` file might look like:
```go
module github.com/learning-go-book-2e/money
go 1.21
require (
    github.com/learning-go-book-2e/formatter v0.0.0-20220918024742-18...
    github.com/shopspring/decimal v1.3.1
)
require (
    github.com/fatih/color v1.13.0 // indirect
    github.com/mattn/go-colorable v0.1.9 // indirect
    github.com/mattn/go-isatty v0.0.14 // indirect
    golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
)
```
The first line of every `go.mod` file is a `module` directive, which declares the module’s unique path. This is followed by a `go` directive specifying the minimum required Go version. The Go compiler will enforce compatibility with this version. For example, if you specify `go 1.12`, you won’t be able to use features introduced in Go 1.13, such as underscores in numeric literals.
### Using the `go` Directive to Manage Go Build Versions
If the `go` directive specifies a newer Go version than what’s installed, the behavior depends on the version of Go you’re using:
- In Go 1.20 and earlier, the specified Go version is ignored, and the installed version is used instead.
- In Go 1.21 and later, Go automatically downloads and uses the newer version to build your code.
To control this behavior, Go 1.21 introduced the `toolchain` directive and the `GOTOOLCHAIN` environment variable. These can take the following values:
- `auto`: Downloads and uses newer Go versions if necessary. (This is the default.)
- `local`: Uses only the installed Go version, restoring the pre-1.21 behavior.
- A specific version, such as `go1.20.4`, which forces the use of that version, downloading it if needed.
For example, you can run:
```
GOTOOLCHAIN=go1.18 go build
```
This builds your code using Go 1.18, even if a newer version is specified in `go.mod`. If both the `toolchain` directive and `GOTOOLCHAIN` are set, `GOTOOLCHAIN` takes precedence.
More details on `go`, `toolchain`, and `GOTOOLCHAIN` can be found in the official Go toolchain documentation.
### Backward Compatibility and `for` Loop Changes in Go 1.22
Go 1.22 introduces the first backward-incompatible change in the language. If a `go.mod` file specifies `go 1.22` or later, a `for-range` loop creates a new variable on each iteration instead of reusing the same one. This behavior is controlled per module—each imported module follows the language level specified in its `go.mod` file.
Consider the following example:
```go
func main() {
    x := []int{1, 2, 3, 4, 5}
    for _, v := range x {
        fmt.Printf("%p\n", &v)
    }
}
```
The `%p` verb in `fmt` prints a pointer’s memory address. If you compile this with `go 1.21`, you’ll see the same memory address printed five times:
```
0x140000140a8
0x140000140a8
0x140000140a8
0x140000140a8
0x140000140a8
```
This means the loop reuses the same `v` variable. However, if you change `go 1.21` to `go 1.22` in `go.mod` and recompile, you’ll see different addresses:
```
0x1400000e0b0
0x1400000e0b8
0x1400000e0d0
0x1400000e0d8
0x1400000e0e0
```
Each iteration now creates a new `v` variable, which prevents common bugs involving loop variable reuse in goroutines.
### The `require` Directive
The `require` directives list the dependencies of your module:
1. The first `require` section lists **direct dependencies**—modules your code imports directly.
2. The second `require` section lists **indirect dependencies**—dependencies of your dependencies. These are marked with `// indirect`.
Functionally, direct and indirect dependencies are the same. However, a module initially added as an indirect dependency might later be promoted to a direct dependency if you import it in your own code.
### Other `go.mod` Directives
While `module`, `go`, and `require` are the most common directives in `go.mod`, three others are worth mentioning:
- `replace`: Used to override a module version or redirect to a local path.
- `exclude`: Prevents a specific module version from being used.
- `retract`: Marks a module version as retracted, signaling that it should not be used.
I’ll cover `replace` and `exclude` in "Overriding Dependencies" and `retract` in "Retracting a Version of Your Module."
## Importing and Exporting in Go
Go's `import` statement allows you to access exported constants, variables, functions, and types from another package. However, Go does not use keywords like `public` or `private` to determine visibility. Instead, it follows a simple convention:
- **Exported identifiers** (accessible from other packages) **must start with an uppercase letter**.
- **Unexported identifiers** (accessible only within the same package) **start with a lowercase letter**.
### Example
```go
package math
// Exported function (visible outside the package)
func Double(a int) int {
    return a * 2
}
// Unexported function (only accessible within the math package)
func half(a int) int {
    return a / 2
}
```
If another package imports `math`, it can only use `Double` but **not** `half`.
### Package API Considerations
- **Anything you export becomes part of your package's API**.
- **Ensure backward compatibility** for exported identifiers to prevent breaking changes.
- **Document all exported identifiers** for clarity.
## Creating and Using a Package
### Directory Structure Example
```
package_example/
│── main.go
│── math/
│   └── math.go
│── do-format/
│   └── formatter.go
```
### `math.go`
```go
package math
func Double(a int) int {
    return a * 2
}
```
### `formatter.go`
```go
package format
import "fmt"
func Number(num int) string {
    return fmt.Sprintf("The number is %d", num)
}
```
### `main.go`
```go
package main
import (
    "fmt"
    "github.com/learning-go-book-2e/package_example/do-format"
    "github.com/learning-go-book-2e/package_example/math"
)
func main() {
    num := math.Double(2)
    output := format.Number(num)
    fmt.Println(output)
}
```
### Key Takeaways
- The **package clause** (e.g., `package math`) **must be the first non-comment line** in a Go file.
- A package's **import path** is composed of the **module path + package directory name**.
- Go **does not allow importing a package without using it**. This prevents unused dependencies.
- **Package names should match their directory names**, unless necessary (e.g., due to invalid characters like ).
### Running the Program
```
go build
./package_example
```
Output:
```
The number is 4
```
## Naming Packages
A well-named package makes code more readable and maintainable.
### Best Practices
- **Use descriptive names** instead of generic names like `util`.
- **Use nouns for package names and verbs for function names**.
- **Avoid redundant prefixes** (e.g., `names.Extract()` is better than `names.ExtractNames()`).
### Example
Instead of:
```go
package util
func ExtractNames(s string) []string { ... }
func FormatNames(s string) string { ... }
```
Use:
```go
package names
func Extract(s string) []string { ... }
func Format(s string) string { ... }
```
## Overriding a Package's Name
If two imported packages have the same name, you can override one of them using an **alias**.
### Example: Resolving `rand` Name Collision
Go has two `rand` packages:
- `crypto/rand` (secure)
- `math/rand` (non-secure)
To use both:
```go
import (
    crand "crypto/rand"
    "encoding/binary"
    "fmt"
    "math/rand"
)
func seedRand() *rand.Rand {
    var b [8]byte
    _, err := crand.Read(b[:])
    if err != nil {
        panic("cannot seed with cryptographic random number generator")
    }
    r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
    return r
}
```
### Alternative Import Strategies
- **Dot Import (`import . "package"`)**
    
    This brings all exported identifiers into the current namespace (not recommended due to potential confusion).
    
- **Blank Import (`import _ "package"`)**
    
    This imports a package solely for its **side effects** (e.g., initializing drivers).
    
## Documenting Your Code with Go Doc Comments
An important part of creating a module for others to use is documenting it properly. Go has its own format for writing comments that are automatically converted into documentation. It’s called **Go Doc format**, and it’s very simple.
### Writing Go Doc Comments
Go Doc comments follow these rules:
- Place the comment **directly before** the item being documented, with no blank lines in between.
- Start each line with `//`  (double slashes followed by a space). While block comments (`/* ... */`) are valid, it’s idiomatic to use `//`.
- The first word in the comment should be **the name of the symbol** being documented, or optionally **"A"** or **"An"** before it for better grammar.
- Separate paragraphs with a **blank comment line** (`//` followed by a newline).
### Formatting Go Doc Comments
Go Doc supports some formatting features:
- **Indented preformatted text**: Add an extra space after `//` to create code blocks or tables.
- **Headers**: Use `#`  (a single `#` after `//`), unlike Markdown, which allows multiple `#` characters.
- **Links to packages**: Wrap the package path in `[ ]`.
- **Links to symbols**: Use `[SymbolName]` or `[pkgName.SymbolName]` for symbols in another package.
- **Raw URLs**: Will be converted into clickable links.
- **Custom text links**: Use `[text]` and define mappings at the end of the comment block (`// [TEXT]: URL`).
### Package-Level Documentation
Comments **before the package declaration** provide documentation for the entire package. If your package requires lengthy documentation, the convention is to place it in a `doc.go` file.
### Example: Package-Level Comment
```go
// Package convert provides various utilities to
// make it easy to convert money from one currency to another.
package convert
```
### Documenting Structs and Functions
#### Example: Struct Documentation
```go
// Money represents an amount of money and its currency.
//
// The value is stored using a [github.com/shopspring/decimal.Decimal].
type Money struct {
    Value    decimal.Decimal
    Currency string
}
```
### Example: Function Documentation
```go
// Convert converts an amount from one currency to another.
//
// It takes a Money instance with the value to convert and a string for the target currency.
// It returns the converted Money instance and an error if the currency is unknown or unconvertible.
//
// If an error occurs, the returned Money instance is set to zero.
//
// Supported currencies:
//   - USD - US Dollar
//   - CAD - Canadian Dollar
//   - EUR - Euro
//   - INR - Indian Rupee
//
// More information on exchange rates can be found at [Investopedia].
//
// [Investopedia]: https://www.investopedia.com/terms/e/exchangerate.asp
func Convert(from Money, to string) (Money, error) {
    // ...
}
```
### Viewing Go Documentation
Use the **`go doc`** command to view documentation:
- `go doc PACKAGE_NAME` – Shows package-level docs.
- `go doc PACKAGE_NAME.IDENTIFIER_NAME` – Shows documentation for a specific identifier.
To preview HTML documentation locally, install `pkgsite`:
```
go install golang.org/x/pkgsite/cmd/pkgsite@latest
pkgsite
```
Visit `http://localhost:8080` to view the docs in your browser.
## Using the `internal` Package
Sometimes you need to share code **within a module** but don’t want it to be part of the public API. Go allows this with the **`internal`** package.
- Identifiers in an `internal` package are accessible **only to sibling packages** or their parent package.
- External packages cannot access `internal` code.
### Example: Internal Package Structure
```
project/
├── foo/
│   ├── foo.go
│   ├── internal/
│   │   ├── internal.go  (can be used by foo.go)
├── sibling/
│   ├── sibling.go  (can use internal.go)
├── bar/
│   ├── bar.go  (CANNOT use internal.go)
```
Trying to access an internal package from an unauthorized package will result in a **compilation error**.
## Avoiding Circular Dependencies
Go **does not allow circular dependencies**. If package A imports package B, package B **cannot import** package A.
### Example: Circular Dependency
#### `pet.go`
```go
import "github.com/example/person"
var owners = map[string]person.Person{
    "Bob":   {"Bob", 30, "Fluffy"},
    "Julia": {"Julia", 40, "Rex"},
}
```
### `person.go`
```go
import "github.com/example/pet"
var pets = map[string]pet.Pet{
    "Fluffy": {"Fluffy", "Cat", "Bob"},
    "Rex":    {"Rex", "Dog", "Julia"},
}
```
**Error:**
```
import cycle not allowed
```
### Solutions
- **Merge** the two packages if they are closely related.
- **Move shared code** to a third package that both can import.
## Organizing Your Go Module
There is no official way to structure Go modules, but over time, best practices have emerged to improve maintainability and readability. Your module's structure should depend on its purpose—whether it is an application or a library.
### When to Organize Your Module
- **Small modules:** Keep everything in a single package if your module is small and has no external dependencies. Over-organizing too early can introduce unnecessary complexity.
- **Growing modules:** As your project expands, structure it logically to improve clarity and maintainability.
### Structuring an Application Module
If your module is meant to be a standalone application, use the `main` package at the root. However, keep the `main` package minimal by placing all logic inside an `internal` directory. This ensures that no external module can depend on your application's internal implementation.
```
myapp/
├── cmd/
│   ├── webserver/
│   │   ├── main.go
│   ├── cli-tool/
│   │   ├── main.go
├── internal/
│   ├── auth/
│   ├── database/
│   ├── handlers/
└── main.go
```
- **`main.go`** should only call functions from `internal` packages.
- **`cmd/`** is used for multiple binaries within the module. Each subdirectory represents an executable program.
- **`internal/`** holds core logic, ensuring it remains inaccessible to external modules.
### Structuring a Library Module
If your module is a library, name the root package after your repository to keep import paths clean. This ensures that the import path matches the package name.
### Guidelines
- Choose a **valid Go package name** for your repository (avoid hyphens).
- If the module contains CLI utilities, place them in a `cmd/` directory.
- Keep package dependencies minimal and group related functionality logically.
Example structure:
```
mylib/
├── cmd/
│   ├── tool/
│   │   ├── main.go
├── mylib/
│   ├── utils.go
│   ├── models.go
│   ├── api.go
├── internal/
│   ├── config/
│   ├── parser/
└── go.mod
```
- **Public API** goes in the root package (`mylib/`).
- **Internal helpers** go in `internal/`, preventing external usage.
- **Command-line tools** belong in `cmd/`.
### Organizing Code by Functionality
Instead of structuring packages based on technical concerns (e.g., separate packages for models, services, and controllers), organize them by functionality.
For example, in an e-commerce application:
```
ecommerce/
├── customer/
│   ├── profile.go
│   ├── orders.go
├── inventory/
│   ├── products.go
│   ├── stock.go
```
This makes it easier to extract functionality into microservices later.
## Gracefully Renaming and Reorganizing Your API
Over time, you may need to rename or move identifiers within your module without breaking existing code. Go provides several ways to do this smoothly.
### Providing Alternate Names
Instead of removing or renaming functions directly, introduce alternate names and gradually phase out the old ones.
### For Functions
```go
func OldFunction() {
    // Implementation
}
func NewFunction() {
    OldFunction()
}
```
The new function calls the old one, ensuring compatibility.
### For Constants
```go
const OldConstant = 10
const NewConstant = OldConstant
```
Both names remain valid until you remove the old one.
### Using Type Aliases
Go allows type aliases to introduce new names without breaking existing references.
### Example
```go
type Foo struct {
    x int
    S string
}
func (f Foo) Hello() string {
    return "hello"
}
// Introduce an alias
type Bar = Foo
func MakeBar() Bar {
    bar := Bar{x: 20, S: "Hello"}
    var f Foo = bar // No type conversion needed
    return bar
}
```
- **Aliases are useful for gradual API changes.**
- **They do not introduce new behavior—modifications must be made to the original type.**
- **Aliases in different packages cannot access unexported fields or methods.**
## Avoiding the `init` Function if Possible
Go prioritizes explicit function calls, but `init` functions allow implicit initialization when a package is imported. They are useful but should be used sparingly.
### What is the `init` Function?
```go
func init() {
    fmt.Println("Initializing package")
}
```
- Runs **once per package** upon the first import.
- **Cannot** take parameters or return values.
- **Executes before `main()`**.
### Why Avoid `init`?
- **Implicit behavior:** It is not clear when or why `init` runs.
- **Uncontrolled side effects:** Modifying global state inside `init` can make debugging difficult.
- **Multiple `init` functions:** Order of execution is hard to track across multiple files.
### When to Use `init`
#### 1. Initializing Package-Level Variables
Use `init` only when a variable cannot be assigned in a single statement.
```go
var config Config
func init() {
    config = LoadConfig()
}
```
Ensure the variable is **effectively immutable** after initialization.
### 2. Legacy Plugin Registration (Obsolete Pattern)
Some libraries register themselves using `init`, but this is discouraged in new code.
```go
import (
    "database/sql"
    _ "github.com/lib/pq"
)
```
This pattern loads the PostgreSQL driver without explicitly referencing it. Instead, prefer explicit registration.
## Importing Third-Party Code
So far, you’ve imported packages from the standard library like `fmt`, `errors`, `os`, and `math`. Go uses the same import system to integrate packages from third parties. Unlike many other compiled languages, Go always builds applications from source code into a single binary file. This includes the source code of your module and the source code of all the modules on which your module depends. (The Go compiler is smart enough to not include unreferenced packages in the binary it produces.) Just as you saw when you imported a package from within your own module, when you import a third-party package, you specify the location in the source code repository where the package is located.
Let’s look at an example. I mentioned back in Chapter 2 that you should never use floating-point numbers when you need an exact representation of a decimal number. If you do need an exact representation, one good option is the `decimal` module from ShopSpring. You are also going to look at a simple formatting module that I’ve written for this book. Both of these modules are used in a small program in the `money` repository for the book. This program calculates the price of an item with the tax included and prints the output in a neat format.
The following code is in `main.go`:
```go
package main
import (
    "fmt"
    "log"
    "os"
    "github.com/learning-go-book-2e/formatter"
    "github.com/shopspring/decimal"
)
func main() {
    if len(os.Args) < 3 {
        fmt.Println("Need two parameters: amount and percent")
        os.Exit(1)
    }
    amount, err := decimal.NewFromString(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    percent, err := decimal.NewFromString(os.Args[2])
    if err != nil {
        log.Fatal(err)
    }
    percent = percent.Div(decimal.NewFromInt(100))
    total := amount.Add(amount.Mul(percent)).Round(2)
    fmt.Println(formatter.Space(80, os.Args[1], os.Args[2], total.StringFixed(2)))
}
```
The two imports `"github.com/learning-go-book-2e/formatter"` and `"github.com/shopspring/decimal"` specify third-party imports. Note that they include the location of the package in the repository. Once they’re imported, you access the exported items in these packages just like any other imported package.
### Updating the `go.mod` File
Before building the application, look at the `go.mod` file. Its contents should be as follows:
```
module github.com/learning-go-book-2e/money
go 1.20
```
If you try to do a build, you get the following message:
```
$ go build
main.go:8:2: no required module provides package
    github.com/learning-go-book-2e/formatter; to add it:
        go get github.com/learning-go-book-2e/formatter
main.go:9:2: no required module provides package
    github.com/shopspring/decimal; to add it:
        go get github.com/shopspring/decimal
```
As the errors indicate, you cannot build the program until you add references to the third-party modules to your `go.mod` file. The `go get` command downloads modules and updates the `go.mod` file.
### Adding Dependencies with `go get`
You have two options when using `go get`. The simplest option is to tell `go get` to scan your module’s source code and add any modules that are found in import statements to `go.mod`:
```
$ go get ./...
go: downloading github.com/shopspring/decimal v1.3.1
go: downloading github.com/learning-go-book-2e/formatter v0.0.0-20220918024742-1835a89362c9
go: downloading github.com/fatih/color v1.13.0
go: downloading github.com/mattn/go-colorable v0.1.9
go: downloading github.com/mattn/go-isatty v0.0.14
go: downloading golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c
go: added github.com/fatih/color v1.13.0
go: added github.com/learning-go-book-2e/formatter v0.0.0-20220918024742-1835a89362c9
go: added github.com/mattn/go-colorable v0.1.9
go: added github.com/mattn/go-isatty v0.0.14
go: added github.com/shopspring/decimal v1.3.1
go: added golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c
```
Because the location of the package is in the source code, `go get` is able to fetch the package’s module and download it.
### Updated `go.mod` File
If you look in the `go.mod` file now, you’ll see this:
```
module github.com/learning-go-book-2e/money
go 1.20
require (
    github.com/learning-go-book-2e/formatter v0.0.0-20220918024742-1835a89362c9
    github.com/shopspring/decimal v1.3.1
)
require (
    github.com/fatih/color v1.13.0 // indirect
    github.com/mattn/go-colorable v0.1.9 // indirect
    github.com/mattn/go-isatty v0.0.14 // indirect
    golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
)
```
The first `require` section lists the modules that you’ve imported into your module. The second `require` section lists modules marked as `// indirect`, meaning they are used only by dependencies of your dependencies.
### Alternative Way to Use `go get`
Instead of scanning your source code, you can pass module paths directly to `go get`. First, roll back changes and remove `go.sum`:
```
$ git restore go.mod
$ rm go.sum
```
Then, add modules one by one:
```
$ go get github.com/learning-go-book-2e/formatter
go: added github.com/learning-go-book-2e/formatter v0.0.0-20200921021027-5abc380940ae
$ go get github.com/shopspring/decimal
go: added github.com/shopspring/decimal v1.3.1
```
This approach doesn’t analyze your source code, so all dependencies are marked as `// indirect`. To fix this, run:
```
$ go mod tidy
```
This command synchronizes `go.mod` and `go.sum` with your module’s source code.
### Why Use `go get` Manually?
Using `go get` with specific module names allows you to update the version of a dependency without affecting other modules.
## **Understanding Go Module Versioning**
Go modules handle dependencies in a structured way, making sure you get the correct version of each package while avoiding unnecessary conflicts. Let’s break it down from the basics to more advanced concepts like **minimal version selection**, **downgrading dependencies**, and **major version upgrades**.
### **1. How Go Resolves Dependencies**
Whenever you import a third-party package in your Go project (without an existing `go.mod` file), Go automatically fetches the latest available version when you run:
```
go get ./...
```
This does two things:
- Downloads the required packages
- Updates `go.mod` and `go.sum` to track dependencies
For example, suppose your `main.go` imports the following:
```go
import (
    "github.com/learning-go-book-2e/simpletax"
    "github.com/shopspring/decimal"
)
```
When you run `go get ./...`, Go fetches the latest versions:
```
go: downloading github.com/learning-go-book-2e/simpletax v1.1.0
go: added github.com/learning-go-book-2e/simpletax v1.1.0
go: added github.com/shopspring/decimal v1.3.1
```
And the `go.mod` file is updated to:
```
module github.com/learning-go-book-2e/region_tax
go 1.20
require (
    github.com/learning-go-book-2e/simpletax v1.1.0
    github.com/shopspring/decimal v1.3.1
)
```
Now, `go.sum` is also updated. It records checksums to ensure package integrity.
### **2. Downgrading a Package Version**
Let’s say `simpletax` v1.1.0 has a bug, and you want to go back to v1.0.0. First, check available versions:
```
go list -m -versions github.com/learning-go-book-2e/simpletax
```
Output:
```
github.com/learning-go-book-2e/simpletax v1.0.0 v1.1.0
```
To downgrade:
```
go get github.com/learning-go-book-2e/simpletax@v1.0.0
```
Now Go updates `go.mod`:
```
require (
    github.com/learning-go-book-2e/simpletax v1.0.0
    github.com/shopspring/decimal v1.3.1
)
```
### **3. Minimal Version Selection (MVS)**
If multiple dependencies in your project require different versions of the same module, Go applies **Minimal Version Selection (MVS)** to choose the highest necessary version.
Example:
| Module | Declared Dependency |
| --- | --- |
| A | `D v1.1.0` |
| B | `D v1.2.0` |
| C | `D v1.2.3` |
Even if `D v1.3.0` exists, Go picks `D v1.2.3`, since it's the **minimum** version that satisfies all requirements.
You can inspect this with:
```
go mod graph
```
### **4. Upgrading a Package**
If the library author releases **patches and minor versions**, you can update using:
- **Patch updates only** (e.g., `v1.1.0 -> v1.1.1`):
    
    ```
    go get -u=patch github.com/learning-go-book-2e/simpletax
    ```
    
- **Update to the latest minor version** (e.g., `v1.0.0 -> v1.2.1`):
    
    ```
    go get -u github.com/learning-go-book-2e/simpletax
    ```
    
### **5. Major Version Upgrades**
Now, let’s say a **breaking change** happens, and `simpletax` releases **v2.0.0**. Go enforces **Semantic Import Versioning**, which means:
- The major version (`v2`) **must** be included in the import path.
- The package must be explicitly updated in `go.mod`.
First, update your import:
```go
import "github.com/learning-go-book-2e/simpletax/v2"
```
Then, run:
```
go get github.com/learning-go-book-2e/simpletax/v2
```
Now, `go.mod` shows:
```
require (
    github.com/learning-go-book-2e/simpletax/v2 v2.0.0
    github.com/shopspring/decimal v1.3.1
)
```
To clean up old, unused versions, run:
```
go mod tidy
```
## **Summary of Commands**
| Task | Command |
| --- | --- |
| Fetch all dependencies | `go get ./...` |
| List available versions | `go list -m -versions <module>` |
| Downgrade a module | `go get <module>@<version>` |
| Upgrade to latest patch | `go get -u=patch <module>` |
| Upgrade to latest minor version | `go get -u <module>` |
| Upgrade to a major version | `go get <module>/v<major>` |
| Remove unused dependencies | `go mod tidy` |
| Show dependency graph | `go mod graph` |
## **Vendoring in Go**
Vendoring ensures that a module always builds with identical dependencies by keeping copies of dependencies inside the module.
### **How it Works**
- Running `go mod vendor` creates a `vendor/` directory that contains all dependencies.
- Go uses these vendored dependencies instead of the module cache on the system.
- If dependencies change (`go get` updates), you must run `go mod vendor` again.
- Forgetting to update the vendor folder leads to build errors.
### **Why (and Why Not) Use Vendoring?**
**Advantages:**
- Makes builds faster in ephemeral CI/CD environments where caches are not preserved.
- Avoids downloading dependencies on every build.
**Disadvantages:**
- Increases repository size.
- Not necessary for most projects due to Go modules and proxy servers.
## **Using `pkg.go.dev`**
- `pkg.go.dev` is an index of Go modules that provides documentation, dependency information, and licensing details.
- You can search for any Go module and see its API, versions, and dependencies.
## **Publishing Your Module**
Publishing a Go module is as simple as pushing it to a version control system (GitHub, GitLab, or a private repository). Unlike ecosystems like npm or Maven, Go modules don’t require uploading to a central registry—Go automatically fetches the source code from the repository.
### **Key Steps to Publish a Module**
1. **Ensure your module has a `go.mod` file**
    - Run `go mod init example.com/yourmodule` if you haven’t already.
2. **Push the module to a repository**
    
    ```
    git init
    git add .
    git commit -m "Initial commit"
    git remote add origin git@github.com:yourname/yourmodule.git
    git push -u origin main
    
    ```
    
3. **Tag a release version**
    - Use semantic versioning (e.g., `v1.0.0`) so Go can recognize versions.
    
    ```
    git tag v1.0.0
    git push origin v1.0.0
    
    ```
    
4. **Ensure a LICENSE file is present**
    - Go prefers permissive licenses like BSD, MIT, or Apache to avoid legal issues when code is compiled into binaries.
## **Versioning Your Module**
### **Semantic Versioning**
Go follows semantic versioning (`vMAJOR.MINOR.PATCH`):
- **MAJOR**: Incompatible changes
- **MINOR**: New features, backward-compatible
- **PATCH**: Bug fixes, backward-compatible
Example:
```
git tag v1.3.4  # Stable version
git tag v1.4.0-beta1  # Pre-release
git tag v2.0.0  # Breaking change (major version bump)
```
### **Breaking Changes (Major Version Updates)**
- If a module reaches `v2.0.0` or later, its **module path must include `/v2`**:
    
    ```go
    module example.com/yourmodule/v2
    
    ```
    
- All imports must be updated:
    
    ```go
    import "example.com/yourmodule/v2"
    
    ```
    
- Two approaches for structuring major versions:
    1. **Subdirectory approach** (`/v2` folder with new code)
    2. **Branch approach** (`v2` branch in version control)
## **Overriding Dependencies**
Sometimes you need to override dependencies—either to fork a module or use a local version.
### **Using `replace`** (for forked modules)
```go
replace example.com/original/module => github.com/yourname/forkedmodule v1.0.0
```
### **Using `replace` for local modules** (not recommended)
```go
replace example.com/original/module => ../local_module
```
- Local `replace` is fragile—use **workspaces** instead.
## **Retracting a Version**
If a version was published by mistake or has security issues, use `retract` in `go.mod`:
```go
retract v1.5.0 // Unstable version
retract [v1.7.0, v1.8.5] // Known issues
```
- This prevents `go get` from fetching retracted versions.
- A new version must be published after adding a `retract` directive.
## **Using Workspaces (`go.work`)**
Workspaces allow working with multiple modules simultaneously.
### **Setup a Workspace**
```
go work init
go work use ./module1
go work use ./module2
```
Creates a `go.work` file:
```go
go 1.20
use (
    ./module1
    ./module2
)
```
### **Why Use Workspaces?**
- Lets multiple modules be developed together without modifying `go.mod`
- Avoids accidental commits of `replace` directives
## Go Tooling
### **Using `go run` to Try Out Small Programs**
- Go is a compiled language, but `go run` allows running code immediately, similar to scripting languages.
- It compiles the code into a temporary binary, executes it, and then deletes the binary.
- Useful for testing small programs without generating unnecessary binaries.
## **Adding Third-Party Tools with `go install`**
- Installs Go programs from source repositories.
- Usage: `go install <module_path>@<version>` (e.g., `go install github.com/rakyll/hey@latest`).
- Always specify `@version` or `@latest` to avoid unintended behaviors.
- Installs binaries in `$HOME/go/bin` by default; can be changed using the `GOBIN` environment variable.
- Ensure `$HOME/go/bin` is in the system `PATH` to run installed tools easily.
## **Improving Import Formatting with `goimports`**
- Enhanced version of `go fmt` that also organizes import statements.
- Automatically removes unused imports and adds missing ones (but may guess incorrectly).
- Install: `go install golang.org/x/tools/cmd/goimports@latest`.
- Run: `goimports -l -w .` to format all files in the current directory.
## **Code Quality Scanners (Linters)**
### **`go vet`**
- Built-in Go tool to detect common programming mistakes.
- Should be part of every automated build process.
### **`staticcheck`**
- More extensive than `go vet`, detecting over 150 issues.
- Catches inefficiencies like unnecessary `fmt.Sprintf` usage.
- Finds unused assignments that the Go compiler may miss.
- Install: `go install honnef.co/go/tools/cmd/staticcheck@latest`.
- Run: `staticcheck ./...`.
### **`revive`**
- Successor to `golint`, allowing configurable linting rules.
- Detects naming issues, missing comments, and style violations.
- Can enforce stricter code quality rules via configuration.
- Install: `go install github.com/mgechev/revive@latest`.
- Example: Configure to detect built-in identifier shadowing using a `built_in.toml` file.
### **`golangci-lint`**
- Meta-linter running multiple tools (`go vet`, `staticcheck`, `revive`, etc.).
- Configurable via `.golangci.yml`.
- Can catch issues like unused variables and identifier shadowing.
- Recommended to maintain a shared config file for team consistency.
- Install via binary release (not `go install`).
- Run: `golangci-lint run`.
## **Scanning for Vulnerable Dependencies with `govulncheck`**
- Detects known security vulnerabilities in standard and third-party Go libraries.
- Uses a public vulnerability database maintained by the Go team.
- Install: `go install golang.org/x/vuln/cmd/govulncheck@latest`.
- Run: `govulncheck ./...`.
- Helps ensure dependencies are updated to patched versions.
## **Embedding Content in Go**
### **Why Use Embedding?**
- Embedding files within a Go binary avoids the need for external files, making distribution easier.
- Useful for including static assets like configuration files, templates, or data files.
- The `//go:embed` directive allows embedding files directly into Go binaries.
### **Embedding a Single File**
#### **Example: Embedding a Password List**
```go
package main
import (
    _ "embed"
    "fmt"
    "os"
    "strings"
)
//go:embed passwords.txt
var passwords string
func main() {
    pwds := strings.Split(passwords, "\n")
    if len(os.Args) > 1 {
        for _, v := range pwds {
            if v == os.Args[1] {
                fmt.Println("true")
                os.Exit(0)
            }
        }
        fmt.Println("false")
    }
}
```
### **Key Points**
- The `embed` package must be imported (even if not used explicitly).
- The `//go:embed` directive must be placed directly above a package-level variable.
- The embedded variable must be of type `string`, `[]byte`, or `embed.FS`.
- Embedding is typically used for immutable data.
### **Embedding a Directory**
#### **Example: Simple Help System**
```go
package main
import (
    "embed"
    "fmt"
    "io/fs"
    "os"
)
//go:embed help
var helpInfo embed.FS
func main() {
    if len(os.Args) == 1 {
        printHelpFiles()
        os.Exit(0)
    }
    data, err := helpInfo.ReadFile("help/" + os.Args[1])
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(string(data))
}
```
### **Key Points**
- `embed.FS` represents an embedded virtual filesystem.
- The directory name (e.g., `"help"`) is included in the embedded structure.
- To read a file, prepend the directory name when calling `ReadFile()`.
### **Listing Embedded Files**
```go
func printHelpFiles() {
    fmt.Println("contents:")
    fs.WalkDir(helpInfo, "help", func(path string, d fs.DirEntry, err error) error {
        if !d.IsDir() {
            fmt.Println(path)
        }
        return nil
    })
}
```
- Uses `fs.WalkDir()` to traverse the embedded directory.
- Checks if an entry is a file before printing its path.
### **Advanced Embedding Techniques**
#### **Embedding Multiple Files or Directories**
- Multiple files can be embedded using space-separated names:
    
    ```go
    //go:embed file1.txt file2.txt dir1/*
    var content embed.FS
    ```
    
- Supports wildcards (`` matches multiple characters, `?` matches a single character).
- If a pattern doesn’t match any file, the compilation fails.
### **Embedding Hidden Files**
- By default, hidden files (`.filename`, `_filename`) are ignored.
- Use `/*` to include hidden files only in the root directory.
    
    ```go
    //go:embed parent_dir/*
    var parentHiddenOnly embed.FS
    ```
    
- Use `all:` to include hidden files in all subdirectories.
    
    ```go
    //go:embed all:parent_dir
    var allHidden embed.FS
    ```
    
## **Using `go generate`**
### **What is `go generate`?**
- The `go generate` tool does not perform any actions by itself.
- It looks for specially formatted comments in source code and executes programs specified within them.
- Commonly used for generating source code automatically (e.g., code for interacting with Protocol Buffers or enumerations).
### **Example: Using `go generate` with Protobufs**
#### **Install Prerequisites**
1. Install `protoc` (Protocol Buffers compiler).
2. Install the Go protobuf plugin:
    
    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    
    ```
    
### **Sample `person.proto` File**
```
syntax = 'proto3';
message Person {
  string name = 1;
  int32 id = 2;
  string email = 3;
}
```
### **Run `go generate`**
In the `main.go` file, include the following magic comment:
```go
//go:generate protoc -I=. --go_out=. --go_opt=module=github.com/learning-go-book-2e/proto_generate --go_opt=Mperson.proto=github.com/learning-go-book-2e/proto_generate/data person.proto
```
Then run:
```bash
go generate ./...
```
### **Generated Output**
- After running `go generate`, a `person.pb.go` file is generated in the `data` directory, containing the necessary Go code for the `Person` struct and methods to marshal/unmarshal data.
### **Usage in Code**
```go
package main
import (
    "fmt"
    "github.com/golang/protobuf/proto"
    "github.com/learning-go-book-2e/proto_generate/data"
)
func main() {
    p := &data.Person{
        Name:  "Bob Bobson",
        Id:    20,
        Email: "bob@bobson.com",
    }
    fmt.Println(p)
    protoBytes, _ := proto.Marshal(p)
    fmt.Println(protoBytes)
    var p2 data.Person
    proto.Unmarshal(protoBytes, &p2)
    fmt.Println(&p2)
}
```
Output:
```bash
name:"Bob Bobson" id:20 email:"bob@bobson.com"
[10 10 66 111 98 32 66 111 98 115 111 110 16 20 26 14 98 111 98 64 98 111 98 115 111 110 46 99 111 109]
name:"Bob Bobson" id:20 email:"bob@bobson.com"
```
### **Using `go generate` with `stringer`**
- `stringer` generates a `String()` method for enumerations to allow printable names.
### **Install `stringer`**
```bash
go install golang.org/x/tools/cmd/stringer@latest
```
### **Example**
```go
type Direction int
const (
    _ Direction = iota
    North
    South
    East
    West
)
//go:generate stringer -type=Direction
func main() {
    fmt.Println(North.String())
}
```
### **Run `go generate`**
```bash
go generate ./...
```
This generates a file `direction_string.go` with the `String()` method for the `Direction` type.
Output when run:
```bash
North
```
### **Best Practices for `go generate`**
- **Commit Generated Code**: It is recommended to commit generated source code to version control. This ensures that others don’t need to have the tools installed to build the code.
- **Automation**: Automate `go generate` to run before the build step to avoid issues with forgotten code generation.
- **Limitations**:
    - Avoid committing generated files with minor differences (e.g., timestamps).
    - Avoid long `go generate` times that impact build speed.
## Reading the Build Info Inside a Go Binary
When you build a Go program, the binary includes valuable build information, such as the version of the Go toolchain used, the versions of dependencies, and details about the repository and revision the code was built from. This data can be essential when deploying software and tracking versions, especially for security or debugging purposes.
The build information is embedded automatically by the Go toolchain and can be viewed using the `go version -m` command, as shown in the example below:
```bash
$ go build
$ go version -m vulnerable
vulnerable: go1.20
    path     github.com/learning-go-book-2e/vulnerable
    mod      github.com/learning-go-book-2e/vulnerable    (devel)
    dep      gopkg.in/yaml.v2  v2.2.7  h1:VUgggvou5XRW9mHwD/yXxIYSMtY0zoKQf/v...
    build    -compiler=gc
    build    CGO_ENABLED=1
    build    GOARCH=arm64
    build    GOOS=darwin
    build    vcs=git
    build    vcs.revision=623a65b94fd02ea6f18df53afaaea3510cd1e611
    build    vcs.time=2022-10-02T03:31:05Z
    build    vcs.modified=false
```
This build info is helpful for tracking the exact code and dependencies used in production. It also aids in vulnerability scanning by providing insight into the exact versions of third-party libraries, as tools like `govulncheck` can scan for known vulnerabilities:
```bash
$ govulncheck -mode binary vulnerable
```
## Building Go Binaries for Other Platforms
Go binaries are compiled for a specific OS and CPU architecture, making cross-compilation a useful feature. You can target different platforms by setting the `GOOS` and `GOARCH` environment variables. For example, to build a binary for Linux on a 64-bit Intel CPU from a macOS environment:
```bash
$ GOOS=linux GOARCH=amd64 go build
$ file vulnerable
vulnerable: ELF 64-bit LSB executable, x86-64
```
This allows you to build and deploy Go applications on different platforms without needing separate build environments.
## Using Build Tags
In cases where you need platform-specific or version-specific code, you can use build tags. These tags are placed in the source file as comments (e.g., `//go:build`), and control when a file is included in the build. You can target specific platforms, architectures, or Go versions, and you can also define custom tags. For example:
```go
//go:build linux && amd64
```
This would ensure that the file is only compiled when targeting Linux and AMD64 platforms. Custom tags can also be defined for specialized builds, such as skipping files during certain builds:
```go
//go:build ignore
```
Build tags offer flexibility in managing platform-specific code and conditional builds.
## Testing Versions of Go
To ensure compatibility with multiple Go versions, you can install and test specific Go versions. This is especially useful for testing backward compatibility. For example, to install Go 1.19.2 and test your code:
```bash
$ go install golang.org/dl/go1.19.2@latest
$ go1.19.2 download
$ go1.19.2 build
```
Once done, you can remove the secondary Go environment:
```bash
$ rm -rf ~/sdk/go.19.2
$ rm ~/go/bin/go1.19.2
```
This allows you to validate your code with different Go versions to ensure it works as expected across versions.
## Concurrency
### When to Use Concurrency
Concurrency is an effective tool, but it’s important to use it when it will actually benefit your program. Many developers, especially when starting with Go, go through a cycle where they try to add concurrency everywhere, believing it will automatically make the program faster. However, this approach often leads to:
1. **Initial excitement:** "I’m going to put everything in goroutines!"
2. **Realization:** "My program isn’t any faster, maybe I should add buffers to my channels."
3. **Frustration:** "My channels are blocking, and I’m getting deadlocks. Let’s try using buffered channels with big buffers."
4. **Confusion:** "My channels are still blocking. I need mutexes."
5. **Giving up:** "Forget it, I’m done with concurrency."
The key to understanding when to use concurrency is recognizing that concurrency isn’t the same as parallelism. Concurrency structures a problem by breaking it into independent tasks, but it doesn’t guarantee that those tasks will run in parallel. Whether or not they run in parallel depends on the hardware and the problem’s characteristics.
### Amdahl’s Law
Amdahl’s Law helps in understanding the potential benefits of parallelism, which is directly tied to the proportion of work that can be parallelized. If most of your work must be done sequentially, adding more threads won’t drastically speed up your program. More concurrency does not always equal more speed.
### When to Use Concurrency
You should consider using concurrency in your program when:
- **Operations can run independently**: For example, when combining results from multiple independent services, as they can work in parallel without waiting on each other.
- **I/O-bound tasks**: Concurrency is very effective when performing tasks like network calls or reading/writing files since these tasks often spend time waiting for external resources.
- **Complex workflows**: If your program involves transforming data in multiple stages that can occur concurrently (e.g., fetching data from different sources and merging them).
A great rule of thumb is: If an operation doesn’t take much time to run, adding concurrency could introduce unnecessary overhead, making the program slower. Before jumping into concurrency, start by writing the program serially, and then benchmark to compare performance with a concurrent approach.
## Goroutines
In Go, the concept of a **goroutine** is at the core of concurrency. A goroutine is a lightweight thread, managed by the Go runtime. Goroutines are not directly mapped to operating system threads. Instead, the Go scheduler handles them in a more efficient way.
### Goroutines vs. Threads
- **Processes**: A process is a program instance with allocated resources, including memory. Each process can have multiple threads.
- **Threads**: A thread is the unit of execution within a process. Threads within a process share memory and resources, and multiple threads can run simultaneously (if the system supports it).
A goroutine is similar to a thread, but it is much more lightweight. The Go runtime schedules these goroutines onto threads managed by the operating system. Goroutines are created with the `go` keyword, which allows functions to run concurrently.
### Advantages of Goroutines
- **Speed**: Goroutines are quicker to create compared to threads because they don’t require operating system-level resources.
- **Memory efficiency**: The initial stack size of a goroutine is smaller than that of a thread and can grow as needed.
- **Context switching**: The Go runtime’s scheduler can switch between goroutines efficiently, avoiding the overhead of thread switching.
- **Scalability**: Go can launch thousands of goroutines simultaneously without significant performance penalties.
### Goroutines in Action
To launch a goroutine, use the `go` keyword before calling a function. Here’s an example:
```go
func process(val int) int {
    // Do something with val
}
func processConcurrently(inVals []int) []int {
    in := make(chan int, 5)
    out := make(chan int, 5)
    // Launch processing goroutines
    for i := 0; i < 5; i++ {
        go func() {
            for val := range in {
                out <- process(val)
            }
        }()
    }
    // Load data into the input channel and read from the output channel
    // Return the results
}
```
In this example:
- We launch a goroutine for each item to process using a closure.
- The closure reads from an input channel, processes the data using the `process` function, and writes the result to an output channel.
- The actual execution of `process` is unaware that it’s running in a goroutine.
This example demonstrates how Go’s concurrency model makes it simple to manage concurrent operations, while keeping the business logic and concurrency concerns separate, leading to modular and testable code.
For further reading, you can explore the complete example on The Go Playground or check out the `sample_code/goroutine` directory in the Chapter 12 repository.
## Channels
Goroutines in Go communicate using **channels**, which provide a way to send and receive values between them safely. Similar to slices and maps, channels are a **reference type** and are created using the `make` function:
```go
ch := make(chan int)
```
A channel has a **zero value** of `nil`, which means it is not usable until initialized.
### Reading, Writing, and Buffering
The `<-` operator is used to read from and write to channels:
```go
a := <-ch  // Read a value from ch and assign it to a
ch <- b    // Write the value in b to ch
```
Each value written to a channel is **consumed only once**—if multiple goroutines are reading from the same channel, a value will be received by only one of them.
### **Direction-Specific Channels**
To improve type safety, Go allows defining **read-only** and **write-only** channels:
```go
var readOnly <-chan int  // Can only receive values
var writeOnly chan<- int // Can only send values
```
By marking a function parameter as read-only or write-only, the Go compiler ensures the correct usage of channels.
### **Unbuffered Channels**
By default, channels are **unbuffered**, meaning every **write** operation blocks until another goroutine reads the value, and vice versa. This ensures synchronization but requires multiple goroutines:
```go
ch := make(chan int) // Unbuffered channel
go func() {
    ch <- 42 // Blocks until another goroutine reads
}()
fmt.Println(<-ch) // Blocks until a value is available
```
### **Buffered Channels**
Buffered channels allow storing a limited number of values without blocking:
```go
ch := make(chan int, 3) // Buffered channel with capacity 3
ch <- 1
ch <- 2
ch <- 3  // This is fine, as the buffer has space
// ch <- 4  // This would block until a read occurs
fmt.Println(<-ch) // Outputs 1, unblocking space for another write
```
The **`len(ch)`** function returns the number of items currently in the buffer, and **`cap(ch)`** gives the buffer’s capacity.
> Note: len and cap return 0 for unbuffered channels, as they do not store values.
> 
### Using `for-range` with Channels
A `for-range` loop can iterate over values received from a channel:
```go
for v := range ch {
    fmt.Println(v)
}
```
This loop continues reading values from the channel until the channel is **closed**.
### Closing a Channel
A channel can be **closed** using the built-in `close` function:
```go
close(ch)
```
Once closed:
- Writing to the channel or trying to close it again **panics**.
- Reading from a **buffered** channel returns remaining values before yielding the **zero value**.
- Reading from an **empty** channel immediately returns the zero value.
Example:
```go
ch := make(chan int, 2)
ch <- 10
ch <- 20
close(ch)
fmt.Println(<-ch) // 10
fmt.Println(<-ch) // 20
fmt.Println(<-ch) // 0 (zero value)
```
### **Detecting Closed Channels (`comma ok` idiom)**
To differentiate between a real zero value and a channel closure, use:
```go
v, ok := <-ch
if !ok {
    fmt.Println("Channel is closed")
}
```
Here, `ok` is `false` when the channel is closed.
> Tip: Always use the comma ok idiom when reading from a potentially closed channel.
> 
### **Who Should Close the Channel?**
The **writing goroutine** should be responsible for closing the channel, ensuring no more writes occur after closing. **Closing is not required** unless another goroutine depends on detecting the channel to close (e.g., using `for-range`).
### Understanding Channel Behavior
Different types of channels behave differently depending on whether they are open, closed, buffered, or unbuffered:
| Channel Type | Read Behavior | Write Behavior | Closing Behavior |
| --- | --- | --- | --- |
| **Unbuffered, Open** | Blocks if empty | Blocks if no reader | Allowed |
| **Unbuffered, Closed** | Returns zero value | Panics | Panics |
| **Buffered, Open** | Blocks if empty | Blocks if full | Allowed, remaining values still there |
| **Buffered, Closed** | Returns remaining values, then zero value (use comma ok) | Panics | Panics |
| **Nil** | Blocks forever | Blocks forever | Panics |
### **Avoiding Channel Panics**
- **Never close a channel twice** (`close(ch)` on an already closed channel panics).
- **Never write to a closed channel** (`ch <- val` on a closed channel panics).
- **Multiple goroutines writing?** Use a `sync.WaitGroup` to ensure all writers are done before closing.
## `select` Statement
The `select` statement is a **concurrency control structure** that allows a goroutine to **wait on multiple channels** and proceed with whichever operation is ready first. Unlike a `switch`, it **randomly picks a ready case**, ensuring **no starvation**.
### **Basic Syntax**
```go
select {
case v := <-ch:
    fmt.Println(v) // Read from ch
case v := <-ch2:
    fmt.Println(v) // Read from ch2
case ch3 <- x:
    fmt.Println("Wrote", x) // Write to ch3
case <-ch4:
    fmt.Println("Received from ch4 (ignored)") // Read and discard
}
```
Each case must be a **read or write** on a channel. If multiple cases are ready, Go **randomly picks one**. If none are ready, `select` **blocks** until a case is ready.
### **Avoiding Deadlocks with `select`**
Deadlocks occur when all goroutines are waiting indefinitely. Consider this **deadlocked** program:
```go
func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    go func() {
        ch1 <- 1       // Blocks until read
        fmt.Println(<-ch2) // Never executes
    }()
    ch2 <- 2       // Blocks because no one is reading yet
    fmt.Println(<-ch1) // Never executes
}
```
**Why does it deadlock?**
- `ch1 <- 1` blocks **until it is read**.
- `ch2 <- 2` blocks **until it is read**.
- Since **both writes block**, neither goroutine can proceed.
Using `select` **prevents deadlocks**:
```go
package main
import (
	"fmt"
)
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v           // Sends v to ch1
		v2 := <-ch2        // Blocks here, waiting for a value from ch2
		fmt.Println(v, v2) // This line is never executed
	}()
	v := 2
	var v2 int
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}
	fmt.Println(v, v2)
	// The main goroutine exits immediately after this,
	// terminating the program and killing the other goroutine.
	// The launched goroutine is stuck waiting to receive from ch2,
	// so it never reaches the fmt.Println statement.
}
```
**Why does this work?**
- If `ch2 <- 2` is **possible**, it executes.
- Otherwise, `fromGoroutine = <-ch1` executes.
- This **avoids waiting on a blocked channel**.
### **`for-select` Loops**
A `for` loop with `select` is common when continuously reading from channels:
```go
for {
    select {
    case v := <-ch:
        fmt.Println(v)
    case <-done:
        return // Graceful exit
    }
}
```
> Note: A for-select loop must have an exit condition (e.g., receiving from done), or it runs indefinitely.
> 
### **Non-blocking Channel Operations**
If `select` has a **default case**, it **does not block** and executes the `default` if no channels are ready:
```go
select {
case v := <-ch:
    fmt.Println("Read:", v)
default:
    fmt.Println("No value available")
}
```
> Warning: Avoid default inside a for-select loop—it executes constantly when no channels are ready, wasting CPU.
> 
## **Keep Your APIs Concurrency-Free**
- Concurrency should be **an implementation detail**, not exposed in API design.
- **Do not expose channels or mutexes** in exported types, functions, or methods.
    - Exposing them **shifts responsibility** to API users, leading to issues like deadlocks, channel management, and unexpected behavior.
- **Exception:** APIs designed as **concurrency helpers** may expose channels (e.g., worker pools).
- **Best Practice:**
    - Channels and mutexes **can be used internally** but should not be exported.
## **Goroutines, `for` Loops, and Captured Variables**
- Before Go **1.22**, `for` loops reused the same **index (`i`) and value (`v`)** variables, leading to unintended behavior in goroutines.
- **Bug in Go ≤ 1.21:**
    
    ```go
    func main() {
        a := []int{2, 4, 6, 8, 10}
        ch := make(chan int, len(a))
        for _, v := range a {
            go func() {
                ch <- v * 2
            }()
        }
        for i := 0; i < len(a); i++ {
            fmt.Println(<-ch)
        }
    }
    ```
    
    **Unexpected output:**
    
    ```
    20
    20
    20
    20
    20
    ```
    
    **Why?**
    
    - The **same `v` is captured** by all goroutines, which all see its last value (`10`).
- **Fix in Go 1.22+:**
    - Each iteration gets **a new index and value variable**, so goroutines capture the expected values.
    - Correct output:
        
        ```
        20
        8
        4
        12
        16
        ```
        
- **Workarounds for Go ≤ 1.21:**
    - **Shadow `v` inside the loop:**
        
        ```go
        for _, v := range a {
            v := v // Creates a new variable
            go func() {
                ch <- v * 2
            }()
        }
        
        ```
        
    - **Pass `v` as a parameter to the goroutine:**
        
        ```go
        for _, v := range a {
            go func(val int) {
                ch <- val * 2
            }(v)
        }
        
        ```
        
- **General Rule:**
    - **Always pass variables to closures** if their values might change before the closure executes.
## **Always Clean Up Your Goroutines**
- **Goroutines do not automatically exit**—they must be explicitly stopped.
- **Leaking goroutines** keep memory allocated and prevent garbage collection.
- **Example of a leaking goroutine:**
    
    ```go
    func countTo(max int) <-chan int {
        ch := make(chan int)
        go func() {
            for i := 0; i < max; i++ {
                ch <- i
            }
            close(ch)
        }()
        return ch
    }
    
    func main() {
        for i := range countTo(10) {
            if i > 5 {
                break // Exits early, leaking the goroutine
            }
            fmt.Println(i)
        }
    }
    
    ```
    
    **Why does it leak?**
    
    - The goroutine **waits indefinitely** to send more values, but the main function **stopped reading**.
## **Use `context.Context` to Stop Goroutines**
- Use `context.Context` to **signal cancellation** and stop goroutines gracefully.
- **Fixed version using `context.Context`:**
    
    ```go
    func countTo(ctx context.Context, max int) <-chan int {
        ch := make(chan int)
        go func() {
            defer close(ch)
            for i := 0; i < max; i++ {
                select {
                case <-ctx.Done(): // Stop if context is cancelled
                    return
                case ch <- i:
                }
            }
        }()
        return ch
    }
    
    func main() {
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel() // Ensures cleanup when `main` exits
    
        ch := countTo(ctx, 10)
        for i := range ch {
            if i > 5 {
                break // Exits early
            }
            fmt.Println(i)
        }
    }
    
    ```
    
    **How it works:**
    
    - `context.WithCancel` creates a **cancellable context**.
    - The goroutine **checks `ctx.Done()`** and exits if the context is cancelled.
    - `defer cancel()` ensures the context **signals all goroutines to exit** when `main` ends.
- **Benefits of using `context.Context`:**
    - Allows **cancellation from any part of the call stack**.
    - Ensures **goroutines terminate cleanly**.
    - A common pattern for **graceful shutdowns and timeouts**.
## Know When to Use Buffered and Unbuffered Channels
### Unbuffered Channels
Unbuffered channels are straightforward: one goroutine writes to the channel and waits for another goroutine to pick up the value. The receiver must be ready to receive when the sender sends a value, which results in synchronization between the two goroutines.
### Buffered Channels
Buffered channels are more complex. They require a size to be set at creation. They are useful when you want to limit the number of goroutines or manage queued work. A buffered channel doesn’t block until the buffer is full. Once full, any goroutine attempting to write will block until there’s space. These channels are perfect when you want to manage a limited pool of goroutines or cap the work that can be queued up.
### Example of Buffered Channels:
In this example, you launch 10 goroutines and use a buffered channel to manage the results:
```go
func processChannel(ch chan int) []int {
    const conc = 10
    results := make(chan int, conc)
    for i := 0; i < conc; i++ {
        go func() {
            v := <-ch
            results <- process(v)
        }()
    }
    var out []int
    for i := 0; i < conc; i++ {
        out = append(out, <-results)
    }
    return out
}
```
Here, the buffered channel ensures that each goroutine writes to the channel without blocking, and the main function can collect the results efficiently.
## Implement Backpressure
Buffered channels can also implement backpressure, which limits the number of concurrent operations in a system. Backpressure prevents overloading by restricting how many tasks can be processed simultaneously.
### Example of Backpressure:
```go
type PressureGauge struct {
	ch chan struct{}
}
func New(limit int) *PressureGauge {
	return &PressureGauge{
		ch: make(chan struct{}, limit),
	}
}
func (pg *PressureGauge) Process(f func()) error {
	select {
	case pg.ch <- struct{}{}:
		f()
		<-pg.ch
		return nil
	default:
		return errors.New("no more capacity")
	}
}
func doThingThatShouldBeLimited() string {
	time.Sleep(2 * time.Second)
	return "done"
}
func main() {
	pg := New(10)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
		}
	})
	http.ListenAndServe(":8080", nil)
}
```
In this example, the `PressureGauge` struct uses a buffered channel to limit how many operations can run concurrently. If the channel is full, a request will be rejected, signaling to the caller that the system is overwhelmed.
## Turn Off a Case in a Select
In `select` statements, if a channel is closed, reading from it always succeeds but returns the zero value. To prevent reading from a closed channel, you can disable the case by setting the channel to `nil` once it is closed.
### Example of Handling Closed Channels:
```go
for count := 0; count < 2; {
    select {
    case v, ok := <-in:
        if !ok {
            in = nil
            count++
            continue
        }
        // process v
    case v, ok := <-in2:
        if !ok {
            in2 = nil
            count++
            continue
        }
        // process v
    }
}
```
This loop processes values from two channels, stopping when both channels are closed. The `nil` assignment disables a closed channel case, preventing unnecessary reads.
## Time Out Code
Timeouts are a crucial part of managing concurrency, especially in systems that require a response within a specific time. You can implement timeouts using a `select` statement combined with a context.
### Example of Time Limiting with Context:
```go
func timeLimit[T any](worker func() T, limit time.Duration) (T, error) {
    out := make(chan T, 1)
    ctx, cancel := context.WithTimeout(context.Background(), limit)
    defer cancel()
    go func() {
        out <- worker()
    }()
    select {
    case result := <-out:
        return result, nil
    case <-ctx.Done():
        var zero T
        return zero, errors.New("work timed out")
    }
}
```
Here, the `timeLimit` function ensures that if the `worker` function doesn’t finish within the specified time, it will return a timeout error. The context is canceled once the time limit is reached, and the select statement waits for either the worker’s result or a timeout.
## Use WaitGroups
When a goroutine needs to wait for multiple goroutines to complete, `sync.WaitGroup` from the standard library is the right tool.
### Basic Usage
A `sync.WaitGroup` has three key methods:
- `Add(n int)`: Increments the counter by `n` (the number of goroutines to wait for).
- `Done()`: Decrements the counter (called inside the goroutine).
- `Wait()`: Blocks execution until the counter reaches zero.
Example:
```go
func main() {
    var wg sync.WaitGroup
    wg.Add(3) // Waiting for 3 goroutines
    go func() {
        defer wg.Done()
        doThing1()
    }()
    go func() {
        defer wg.Done()
        doThing2()
    }()
    go func() {
        defer wg.Done()
        doThing3()
    }()
    wg.Wait() // Blocks until all goroutines call Done()
}
```
### Key Considerations
- The `sync.WaitGroup` should be shared among all goroutines, usually captured in a closure to ensure a single instance is used.
- `defer wg.Done()` ensures `Done()` is always called, even if the goroutine panics.
- `WaitGroups` should primarily be used when there's a need to perform cleanup (like closing a shared channel) after all goroutines finish.
### Using WaitGroups to Close a Channel Once
When multiple goroutines write to the same channel, ensuring the channel is closed only once is critical. A `sync.WaitGroup` helps manage this.
Example:
```go
func processAndGather[T, R any](in <-chan T, processor func(T) R, num int) []R {
    out := make(chan R, num)
    var wg sync.WaitGroup
    wg.Add(num)
    for i := 0; i < num; i++ {
        go func() {
            defer wg.Done()
            for v := range in {
                out <- processor(v)
            }
        }()
    }
    // Close `out` only after all goroutines are done
    go func() {
        wg.Wait()
        close(out)
    }()
    var result []R
    for v := range out {
        result = append(result, v)
    }
    return result
}
```
- This pattern ensures that `out` is closed only when all worker goroutines complete.
- The `for-range` loop on `out` will exit cleanly when the channel is closed.
### `golang.org/x/sync/errgroup`
The `errgroup` package builds on `sync.WaitGroup` and adds error handling. It allows launching goroutines that stop processing when one of them returns an error. Check the [errgroup documentation](https://pkg.go.dev/golang.org/x/sync/errgroup) for more details.
## Run Code Exactly Once
### The `sync.Once` Type
The `sync.Once` type ensures that a function runs only once, even in concurrent scenarios.
Example:
```go
type SlowComplicatedParser interface {
    Parse(string) string
}
func initParser() SlowComplicatedParser {
    // Expensive initialization logic
}
```
```go
var parser SlowComplicatedParser
var once sync.Once
func Parse(data string) string {
    once.Do(func() {
        parser = initParser()
    })
    return parser.Parse(data)
}
```
### Key Points
- `sync.Once` ensures `initParser()` runs only once, even if `Parse()` is called multiple times.
- `sync.Once` should not be copied; use it as a package-level variable.
- Declaring `sync.Once` inside a function would reset its state on every function call, negating its purpose.
### Go 1.21+: `sync.OnceFunc`, `sync.OnceValue`, and `sync.OnceValues`
These helper functions make it easier to execute a function exactly once.
Example using `sync.OnceValue`:
```go
var initParserCached func() SlowComplicatedParser = sync.OnceValue(initParser)
func Parse(data string) string {
    parser := initParserCached()
    return parser.Parse(data)
}
```
- The first time `initParserCached()` is called, `initParser()` runs and its return value is cached.
- Subsequent calls return the cached value, avoiding unnecessary recomputation.
- This approach eliminates the need for a global `parser` variable.
## Put Your Concurrent Tools Together
### Overview
This example demonstrates a structured approach to making concurrent API calls, handling errors, and enforcing a timeout using `context.Context`. It involves three web services:
1. **Service A and Service B** process separate parts of the input concurrently.
2. **Service C** takes the results from A and B as input and produces the final output.
3. The entire operation must complete within **50 milliseconds**, or an error is returned.
### Main Function: `GatherAndProcess`
```go
func GatherAndProcess(ctx context.Context, data Input) (COut, error) {
    ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
    defer cancel()
    ab := newABProcessor()
    ab.start(ctx, data)
    inputC, err := ab.wait(ctx)
    if err != nil {
        return COut{}, err
    }
    c := newCProcessor()
    c.start(ctx, inputC)
    out, err := c.wait(ctx)
    return out, err
}
```
### Key Concepts:
- **Timeout Handling**: A `context.WithTimeout` ensures the operation doesn't exceed 50ms.
- **Cancellation**: `defer cancel()` prevents resource leaks.
- **Parallel Execution**: `abProcessor` calls **Service A** and **Service B** concurrently.
- **Sequential Execution**: The results of A & B are used as input for **Service C**.
### `abProcessor`: Handling Service A and B
#### Struct Definition
```go
type abProcessor struct {
    outA chan aOut
    outB chan bOut
    errs chan error
}
func newABProcessor() *abProcessor {
    return &abProcessor{
        outA: make(chan aOut, 1),
        outB: make(chan bOut, 1),
        errs: make(chan error, 2),
    }
}
```
- Uses **three buffered channels**:
    - `outA`: Stores the result from **Service A**.
    - `outB`: Stores the result from **Service B**.
    - `errs`: Captures errors (buffered to avoid blocking).
### `start` Method: Making Parallel Calls
```go
func (p *abProcessor) start(ctx context.Context, data Input) {
    go func() {
        aOut, err := getResultA(ctx, data.A)
        if err != nil {
            p.errs <- err
            return
        }
        p.outA <- aOut
    }()
    go func() {
        bOut, err := getResultB(ctx, data.B)
        if err != nil {
            p.errs <- err
            return
        }
        p.outB <- bOut
    }()
}
```
- **Goroutines**:
    - Each goroutine calls **Service A or B**.
    - Errors are sent to `errs` if a failure occurs.
    - Results are sent to `outA` or `outB`.
### `wait` Method: Collecting Results
```go
func (p *abProcessor) wait(ctx context.Context) (cIn, error) {
    var cData cIn
    for count := 0; count < 2; count++ {
        select {
        case a := <-p.outA:
            cData.a = a
        case b := <-p.outB:
            cData.b = b
        case err := <-p.errs:
            return cIn{}, err
        case <-ctx.Done():
            return cIn{}, ctx.Err()
        }
    }
    return cData, nil
}
```
- **Loops twice** to collect results from A & B.
- **Handles errors & timeouts** using `select`.
- **Returns** combined results if successful.
---
### `cProcessor`: Handling Service C
#### Struct Definition
```go
type cProcessor struct {
    outC chan COut
    errs chan error
}
func newCProcessor() *cProcessor {
    return &cProcessor{
        outC: make(chan COut, 1),
        errs: make(chan error, 1),
    }
}
```
- **Channels**:
    - `outC`: Stores the result from **Service C**.
    - `errs`: Captures errors.
### `start` Method: Calling Service C
```go
func (p *cProcessor) start(ctx context.Context, inputC cIn) {
    go func() {
        cOut, err := getResultC(ctx, inputC)
        if err != nil {
            p.errs <- err
            return
        }
        p.outC <- cOut
    }()
}
```
- Launches a **goroutine** to call **Service C**.
- Handles **errors & results** similarly to `abProcessor`.
### `wait` Method: Collecting Results
```go
func (p *cProcessor) wait(ctx context.Context) (COut, error) {
    select {
    case out := <-p.outC:
        return out, nil
    case err := <-p.errs:
        return COut{}, err
    case <-ctx.Done():
        return COut{}, ctx.Err()
    }
}
```
- Uses **select** to:
    - Return the first available **result**.
    - Return **errors or timeouts** immediately.
---
### Summary
- **Efficient parallel processing** with goroutines.
- **Proper synchronization** using channels.
- **Error handling** via an `errs` channel.
- **Timeout enforcement** with `context.Context`.
This pattern ensures a **responsive** and **robust** concurrent pipeline.
## When to Use Mutexes Instead of Channels
### Overview
In Go, channels and `select` are the preferred concurrency management tools, but mutexes (`sync.Mutex` and `sync.RWMutex`) are still useful in specific scenarios. Channels clarify data flow by ensuring only one goroutine accesses a value at a time, while mutexes protect shared memory without explicit data transfer. The Go philosophy favors communication over shared memory: *“Share memory by communicating; do not communicate by sharing memory.”*
### When to Use Mutexes
Mutexes are useful when:
- Goroutines need to **read or write a shared value** but don't process it.
- Data is stored **in-memory** rather than in external services (e.g., databases, HTTP servers).
### Example: Scoreboard with Channels
Using channels, a goroutine manages the scoreboard by receiving functions that modify the map:
```go
func scoreboardManager(ctx context.Context, in <-chan func(map[string]int)) {
    scoreboard := map[string]int{}
    for {
        select {
        case <-ctx.Done():
            return
        case f := <-in:
            f(scoreboard)
        }
    }
}
```
A type is defined to interact with the scoreboard manager:
```go
type ChannelScoreboardManager chan func(map[string]int)
func NewChannelScoreboardManager(ctx context.Context) ChannelScoreboardManager {
    ch := make(ChannelScoreboardManager)
    go scoreboardManager(ctx, ch)
    return ch
}
func (csm ChannelScoreboardManager) Update(name string, val int) {
    csm <- func(m map[string]int) {
        m[name] = val
    }
}
```
Reading a value requires sending a function that writes the result to another channel:
```go
func (csm ChannelScoreboardManager) Read(name string) (int, bool) {
    type Result struct {
        out int
        ok  bool
    }
    resultCh := make(chan Result)
    csm <- func(m map[string]int) {
        out, ok := m[name]
        resultCh <- Result{out, ok}
    }
    result := <-resultCh
    return result.out, result.ok
}
```
This approach ensures safe concurrent access but is cumbersome and allows only one reader at a time.
### Example: Scoreboard with a Mutex
A mutex simplifies shared access:
```go
type MutexScoreboardManager struct {
    l          sync.RWMutex
    scoreboard map[string]int
}
func NewMutexScoreboardManager() *MutexScoreboardManager {
    return &MutexScoreboardManager{
        scoreboard: map[string]int{},
    }
}
func (msm *MutexScoreboardManager) Update(name string, val int) {
    msm.l.Lock()
    defer msm.l.Unlock()
    msm.scoreboard[name] = val
}
func (msm *MutexScoreboardManager) Read(name string) (int, bool) {
    msm.l.RLock()
    defer msm.l.RUnlock()
    val, ok := msm.scoreboard[name]
    return val, ok
}
```
Using a `sync.RWMutex`:
- `Lock` and `Unlock` ensure exclusive writes.
- `RLock` and `RUnlock` allow multiple readers at once.
### Guidelines for Choosing Between Channels and Mutexes
**Use Channels if:**
- Coordinating goroutines.
- Tracking a value as it transforms through multiple goroutines.
**Use Mutexes if:**
- Protecting **a field in a struct** that multiple goroutines access.
- Performance issues arise with channels.
### Mutex Considerations
- Always use `defer` to unlock after acquiring a lock.
- **Mutexes in Go are non-reentrant**—a goroutine locking the same mutex twice causes deadlocks.
- **Never copy a mutex**—always use pointers (`sync.Mutex` or `sync.RWMutex`).
- **Use the data race detector** (`go run -race`) to find concurrency issues.
### `sync.Map` – Not a General Replacement for Maps
`sync.Map` is a concurrency-safe map but should only be used when:
- **Keys are written once and read many times.**
- **Goroutines do not access each other’s keys and values.**
Because `sync.Map` predates generics, it uses `any` for keys and values, leading to potential type safety issues. In most cases, a regular `map` with `sync.RWMutex` is a better choice.
## **The Standard Library**
### io and Friends
Go's `io` package is at the heart of how programs read and write data. Two key interfaces, `io.Reader` and `io.Writer`, define the foundation of I/O operations in Go.
### Key Interfaces
#### `io.Reader`
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```
- Reads up to `len(p)` bytes into the provided slice.
- Returns the number of bytes read and an error, if any.
- Uses a preallocated slice to avoid unnecessary memory allocations.
### `io.Writer`
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
- Writes bytes from a slice to an implementation of `io.Writer`.
- Returns the number of bytes written and an error, if any.
### Why `io.Reader` Uses a Buffer Instead of Returning a Slice
A more intuitive definition might be:
```go
type NotHowReaderIsDefined interface {
    Read() (p []byte, err error)
}
```
However, this would require allocating a new slice on every call, leading to excessive heap allocations and increased garbage collection overhead. Instead, Go gives the developer control over memory by passing a preallocated buffer to `Read`.
### Example: Counting Letter Frequencies in a Stream
```go
func countLetters(r io.Reader) (map[string]int, error) {
    buf := make([]byte, 2048)
    out := map[string]int{}
    for {
        n, err := r.Read(buf)
        for _, b := range buf[:n] {
            if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
                out[string(b)]++
            }
        }
        if err == io.EOF {
            return out, nil
        }
        if err != nil {
            return nil, err
        }
    }
}
```
### Key Takeaways
1. **Reuses a buffer** to minimize allocations.
2. **Uses `n` to process only the bytes read**, avoiding unnecessary operations on uninitialized parts of `buf`.
3. **Handles `io.EOF`** as a signal of completion, not an error.
### Using `io.Reader` with a String
```go
s := "The quick brown fox jumped over the lazy dog"
sr := strings.NewReader(s)
counts, err := countLetters(sr)
if err != nil {
    return err
}
fmt.Println(counts)
```
The `strings.NewReader` function creates an `io.Reader` from a string.
### Chaining `io.Reader` Implementations: Processing Gzip Files
Go encourages **decorator patterns** to extend `io.Reader`. This allows processing a gzip-compressed file without modifying `countLetters`.
### Creating a Gzip Reader
```go
func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
    r, err := os.Open(fileName)
    if err != nil {
        return nil, nil, err
    }
    gr, err := gzip.NewReader(r)
    if err != nil {
        return nil, nil, err
    }
    return gr, func() {
        gr.Close()
        r.Close()
    }, nil
}
```
- **Opens a file** and creates a `gzip.Reader` from it.
- **Returns a cleanup function** to close both the gzip reader and the underlying file.
### Using the Gzip Reader
```go
r, closer, err := buildGZipReader("my_data.txt.gz")
if err != nil {
    return err
}
defer closer()
counts, err := countLetters(r)
if err != nil {
    return err
}
fmt.Println(counts)
```
### `io` Utilities for Composing Readers and Writers
- **`io.MultiReader(r1, r2, ...)`** – Reads from multiple `io.Reader` instances sequentially.
- **`io.LimitReader(r, n)`** – Reads up to `n` bytes from `r`.
- **`io.MultiWriter(w1, w2, ...)`** – Writes to multiple `io.Writer` instances simultaneously.
### Other One-Method Interfaces in `io`
#### `io.Closer` (For Cleanup)
```go
type Closer interface {
    Close() error
}
```
- Implemented by types like `os.File`.
- **Best practice**: Use `defer f.Close()` after opening a resource.
- **Exception**: Do not use `defer` inside a loop; instead, close resources before each iteration ends.
### `io.Seeker` (For Random Access)
```go
type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}
```
- `whence` values:
    - `io.SeekStart` (0) – Relative to the beginning.
    - `io.SeekCurrent` (1) – Relative to the current position.
    - `io.SeekEnd` (2) – Relative to the end.
### Combined Interfaces in `io`
Go combines basic interfaces into more useful abstractions:
- `io.ReadCloser` = `io.Reader` + `io.Closer`
- `io.ReadSeeker` = `io.Reader` + `io.Seeker`
- `io.ReadWriteCloser` = `io.Reader` + `io.Writer` + `io.Closer`
- `io.ReadWriteSeeker` = `io.Reader` + `io.Writer` + `io.Seeker`
These are useful when designing functions that require specific behaviors without depending on concrete types like `*os.File`.
### `io.NopCloser`: Adapting an `io.Reader` to an `io.ReadCloser`
If a function requires an `io.ReadCloser`, but you only have an `io.Reader`, use `io.NopCloser`:
```go
func NopCloser(r io.Reader) io.ReadCloser {
    return nopCloser{r}
}
type nopCloser struct {
    io.Reader
}
func (nopCloser) Close() error { return nil }
```
This pattern allows **extending types with additional methods** without modifying the original implementation.
### Reading and Writing Files with `os` Package
- **`os.ReadFile(filename)`** – Reads a file into a byte slice.
- **`os.WriteFile(filename, data, perm)`** – Writes a byte slice to a file.
- **For large files**, use `os.Open()` and `os.Create()`, which return `os.File`, implementing `io.Reader` and `io.Writer`.
### Buffered Reading with `bufio.Scanner`
```go
f, err := os.Open("file.txt")
if err != nil {
    return err
}
defer f.Close()
scanner := bufio.NewScanner(f)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
if err := scanner.Err(); err != nil {
    return err
}
```
- `bufio.Scanner` is **memory-efficient** and reads line by line.
- Handles errors using `scanner.Err()`.
## Time in Go
Go's `time` package provides robust support for working with time and durations. It includes two main types:
- `time.Duration` – Represents a period of time.
- `time.Time` – Represents a specific moment in time, including time zone information.
### Working with Durations
A `time.Duration` represents a span of time, stored as an `int64` counting nanoseconds. Go provides predefined constants to make durations readable and type-safe.
```go
d := 2 * time.Hour + 30 * time.Minute // 2 hours and 30 minutes
```
### Parsing Durations
Go defines a consistent string format for durations, which `time.ParseDuration` can interpret:
```go
d, err := time.ParseDuration("2h45m")
if err != nil {
    return err
}
fmt.Println(d) // 2h45m0s
```
Valid units:
- `"ns"` (nanoseconds)
- `"us"` or `"µs"` (microseconds)
- `"ms"` (milliseconds)
- `"s"` (seconds)
- `"m"` (minutes)
- `"h"` (hours)
### Duration Methods
A `time.Duration` provides methods to extract values in different units:
```go
fmt.Println(d.Hours())   // 2.75
fmt.Println(d.Minutes()) // 165
fmt.Println(d.Seconds()) // 9900
```
The `Truncate` and `Round` methods adjust durations:
```go
fmt.Println(d.Truncate(time.Hour)) // 2h0m0s
fmt.Println(d.Round(time.Hour))    // 3h0m0s
```
### Working with Time
A specific moment in time is represented using `time.Time`. The current time is obtained via:
```go
t := time.Now()
fmt.Println(t) // 2025-02-25 14:05:36.123456789 -0700 MST
```
### Comparing `time.Time` Instances
Because `time.Time` includes a time zone, direct equality comparison (`==`) may not work correctly. Instead, use the `Equal` method:
```go
t1 := time.Now()
t2 := t1.Add(10 * time.Second)
fmt.Println(t1.Equal(t2))  // false
fmt.Println(t1.Before(t2)) // true
fmt.Println(t1.After(t2))  // false
```
### Formatting and Parsing Time
Go uses **reference time formatting**, based on:
```
January 2, 2006 at 3:04:05PM MST
```
This corresponds to:
```
01/02 03:04:05PM '06 -0700
```
### Formatting a `time.Time`
```go
fmt.Println(t.Format("2006-01-02 15:04:05 -0700"))
// 2025-02-25 14:05:36 -0700
```
### Parsing a String into `time.Time`
```go
t, err := time.Parse("2006-01-02 15:04:05 -0700", "2023-03-13 00:00:00 +0000")
if err != nil {
    return err
}
fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))
// March 13, 2023 at 12:00:00AM UTC
```
For common formats, `time` provides constants like `time.RFC3339`:
```go
t, _ := time.Parse(time.RFC3339, "2025-02-25T14:05:36Z")
fmt.Println(t.Format(time.RFC3339))
// 2025-02-25T14:05:36Z
```
### Extracting Components
A `time.Time` provides methods to extract details:
```go
fmt.Println(t.Year())     // 2025
fmt.Println(t.Month())    // February
fmt.Println(t.Day())      // 25
fmt.Println(t.Hour())     // 14
fmt.Println(t.Minute())   // 5
fmt.Println(t.Second())   // 36
fmt.Println(t.Weekday())  // Tuesday
```
`Clock()` and `Date()` return multiple components:
```go
hour, min, sec := t.Clock()
fmt.Println(hour, min, sec) // 14 5 36
year, month, day := t.Date()
fmt.Println(year, month, day) // 2025 February 25
```
### Calculating Differences
The `Sub` method finds the difference between two `time.Time` instances, returning a `time.Duration`:
```go
start := time.Now()
time.Sleep(2 * time.Second)
elapsed := time.Since(start)
fmt.Println(elapsed) // ~2s
```
### Adding and Rounding Time
```go
future := t.Add(48 * time.Hour)
fmt.Println(future) // Two days later
nextMonth := t.AddDate(0, 1, 0)
fmt.Println(nextMonth) // One month later
rounded := t.Round(time.Hour)
fmt.Println(rounded) // Rounds to the nearest hour
```
### Monotonic Time
Go internally tracks **monotonic time** to measure elapsed durations accurately, preventing issues caused by daylight saving time changes or clock adjustments.
- **Monotonic timestamps are included in `time.Now()`**.
- **Elapsed durations (`Sub`) use monotonic time if available**.
Example:
```go
t1 := time.Now()
time.Sleep(500 * time.Millisecond)
t2 := time.Now()
fmt.Println(t2.Sub(t1)) // ~500ms
```
Monotonic timestamps are **invisible to the user** but ensure accurate duration calculations.
### Timers and Tickers
Go's `time` package provides channel-based timers for scheduling tasks.
### One-Time Delays with `time.After`
```go
<-time.After(2 * time.Second)
fmt.Println("2 seconds later")
```
### Repeating Tasks with `time.Ticker`
`time.Tick` creates a **leaking goroutine** because it cannot be stopped. Use `time.NewTicker` instead:
```go
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()
for i := 0; i < 5; i++ {
    fmt.Println(<-ticker.C)
}
```
### Running a Function After a Delay
```go
time.AfterFunc(2*time.Second, func() {
    fmt.Println("Executed after 2 seconds")
})
time.Sleep(3 * time.Second) // Wait to see output
```
## `encoding/json`
REST APIs have standardized JSON as the primary data format for communication, and Go's `encoding/json` package provides built-in support for converting Go data types to and from JSON. The process of converting Go data to JSON is called **marshaling**, while converting JSON into Go data is called **unmarshaling**.
### Using Struct Tags to Add Metadata
Struct tags define how Go struct fields map to JSON fields. Consider the following JSON structure for an order management system:
```json
{
    "id": "12345",
    "date_ordered": "2020-05-01T13:01:02Z",
    "customer_id": "3",
    "items": [
        { "id": "xyz123", "name": "Thing 1" },
        { "id": "abc789", "name": "Thing 2" }
    ]
}
```
The corresponding Go structs would be:
```go
type Order struct {
    ID          string    `json:"id"`
    DateOrdered time.Time `json:"date_ordered"`
    CustomerID  string    `json:"customer_id"`
    Items       []Item    `json:"items"`
}
type Item struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
```
- The `json` struct tag specifies how a field should be named in the JSON output.
- If no struct tag is provided, Go uses the struct field name with its first letter capitalized.
- JSON field names are matched **case-insensitively** during unmarshaling.
- To exclude a field, use `json:"-"`.
- To omit a field when it has an empty value, use `json:",omitempty"`.
### Important Note on `omitempty`
The definition of "empty" varies:
- **Strings, slices, and maps** are empty if their length is `0`.
- **Pointers and interfaces** are empty if `nil`.
- **Structs** are never empty (zero values are still considered "set").
### Marshaling and Unmarshaling
#### Unmarshaling JSON into a Go Struct
To convert JSON (in a `[]byte` slice) into a Go struct:
```go
var o Order
err := json.Unmarshal([]byte(data), &o)
if err != nil {
    return err
}
```
- `json.Unmarshal` modifies the provided struct, requiring a pointer (`&o`).
- This approach optimizes memory usage by allowing reuse of an existing struct.
### Marshaling a Go Struct to JSON
To convert a struct to JSON:
```go
out, err := json.Marshal(o)
if err != nil {
    return err
}
```
This generates a `[]byte` slice containing JSON data.
### JSON, Readers, and Writers
Since Go heavily uses the `io.Reader` and `io.Writer` interfaces, the `encoding/json` package provides:
- `json.Decoder`: Reads JSON from any `io.Reader`.
- `json.Encoder`: Writes JSON to any `io.Writer`.
### Writing JSON to a File
```go
tmpFile, err := os.CreateTemp(os.TempDir(), "sample-")
if err != nil {
    panic(err)
}
defer os.Remove(tmpFile.Name())
toFile := Person{Name: "Fred", Age: 40}
err = json.NewEncoder(tmpFile).Encode(toFile)
if err != nil {
    panic(err)
}
tmpFile.Close()
```
### Reading JSON from a File
```go
tmpFile2, err := os.Open(tmpFile.Name())
if err != nil {
    panic(err)
}
var fromFile Person
err = json.NewDecoder(tmpFile2).Decode(&fromFile)
if err != nil {
    panic(err)
}
tmpFile2.Close()
fmt.Printf("%+v\n", fromFile)
```
### Encoding and Decoding JSON Streams
If multiple JSON objects are stored sequentially (not in an array), `json.Decoder` can process them efficiently:
```go
var t struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
dec := json.NewDecoder(strings.NewReader(streamData))
for {
    err := dec.Decode(&t)
    if err != nil {
        if errors.Is(err, io.EOF) {
            break
        }
        panic(err)
    }
    // Process t
}
```
Similarly, `json.Encoder` can be used to write multiple JSON objects:
```go
var b bytes.Buffer
enc := json.NewEncoder(&b)
for _, input := range allInputs {
    t := process(input)
    err := enc.Encode(t)
    if err != nil {
        panic(err)
    }
}
out := b.String()
```
### Streaming JSON from Arrays
Instead of loading an entire JSON array into memory, `json.Decoder` can process elements one at a time:
```go
dec := json.NewDecoder(reader)
_, err := dec.Token() // Read the opening `[`
if err != nil {
    panic(err)
}
for dec.More() {
    var elem MyStruct
    err := dec.Decode(&elem)
    if err != nil {
        panic(err)
    }
    process(elem)
}
_, err = dec.Token() // Read the closing `]`
if err != nil {
    panic(err)
}
```
This approach reduces memory usage and improves performance.
## **Custom JSON Parsing in Go**
### **Implementing Custom JSON Marshalling and Unmarshalling**
The `encoding/json` package provides standard ways to convert Go data types to and from JSON. However, in some cases, you need to override the default behavior—such as when dealing with non-standard time formats. This can be done by implementing the `json.Marshaler` and `json.Unmarshaler` interfaces for a custom type.
### **Creating a Custom Type for Time Formatting**
Go’s `time.Time` natively supports RFC 3339 for JSON encoding and decoding, but if you need a different format (e.g., RFC 822Z), you can create a custom wrapper type:
```go
type RFC822ZTime struct {
    time.Time
}
func (rt RFC822ZTime) MarshalJSON() ([]byte, error) {
    out := rt.Time.Format(time.RFC822Z)
    return []byte(`"` + out + `"`), nil
}
func (rt *RFC822ZTime) UnmarshalJSON(b []byte) error {
    if string(b) == "null" {
        return nil
    }
    t, err := time.Parse(`"`+time.RFC822Z+`"`, string(b))
    if err != nil {
        return err
    }
    *rt = RFC822ZTime{t}
    return nil
}
```
This struct embeds `time.Time` so that it retains all `time.Time` methods. The `MarshalJSON` function ensures that time is formatted in RFC 822Z, while `UnmarshalJSON` parses incoming JSON time strings into this format.
Now, you can use this type in your struct:
```go
type Order struct {
    ID          string      `json:"id"`
    DateOrdered RFC822ZTime `json:"date_ordered"`
    CustomerID  string      `json:"customer_id"`
    Items       []Item      `json:"items"`
}
```
This ensures that `DateOrdered` fields are always marshaled and unmarshaled in RFC 822Z format.
### **Alternative Approach: Overriding JSON Methods in Struct**
Instead of defining a separate type, you can override JSON methods in the struct itself. This approach uses an embedded duplicate struct to handle JSON processing while maintaining the original field types:
```go
type Order struct {
    ID          string    `json:"id"`
    Items       []Item    `json:"items"`
    DateOrdered time.Time `json:"date_ordered"`
    CustomerID  string    `json:"customer_id"`
}
func (o Order) MarshalJSON() ([]byte, error) {
    type Alias Order
    tmp := struct {
        DateOrdered string `json:"date_ordered"`
        Alias
    }{
        Alias:      Alias(o),
        DateOrdered: o.DateOrdered.Format(time.RFC822Z),
    }
    return json.Marshal(tmp)
}
func (o *Order) UnmarshalJSON(b []byte) error {
    type Alias Order
    tmp := struct {
        DateOrdered string `json:"date_ordered"`
        *Alias
    }{
        Alias: (*Alias)(o),
    }
    err := json.Unmarshal(b, &tmp)
    if err != nil {
        return err
    }
    o.DateOrdered, err = time.Parse(time.RFC822Z, tmp.DateOrdered)
    return err
}
```
This method avoids creating a separate `time.Time` wrapper but still allows for custom JSON parsing. However, the struct remains tied to a specific JSON format.
### **Keeping JSON Parsing Separate from Business Logic**
A better approach is to use separate structs for JSON handling and business logic. You define one struct for JSON serialization and another for internal data processing:
```go
type OrderJSON struct {
    ID          string `json:"id"`
    DateOrdered string `json:"date_ordered"`
    CustomerID  string `json:"customer_id"`
    Items       []Item `json:"items"`
}
type Order struct {
    ID          string
    DateOrdered time.Time
    CustomerID  string
    Items       []Item
}
func (o *Order) ToJSON() ([]byte, error) {
    return json.Marshal(OrderJSON{
        ID:          o.ID,
        DateOrdered: o.DateOrdered.Format(time.RFC822Z),
        CustomerID:  o.CustomerID,
        Items:       o.Items,
    })
}
func (o *Order) FromJSON(data []byte) error {
    var tmp OrderJSON
    err := json.Unmarshal(data, &tmp)
    if err != nil {
        return err
    }
    o.ID = tmp.ID
    o.CustomerID = tmp.CustomerID
    o.Items = tmp.Items
    o.DateOrdered, err = time.Parse(time.RFC822Z, tmp.DateOrdered)
    return err
}
```
This approach ensures that JSON formatting concerns are handled separately, making the business logic independent of any specific serialization format.
### **Avoiding `map[string]any` for JSON Parsing**
While Go allows parsing JSON into `map[string]any`, this should be used only during initial prototyping. A concrete struct provides:
- **Stronger type safety**: Avoids runtime errors due to unexpected types.
- **Better documentation**: Makes expected data structures clear.
- **Easier debugging**: Reduces ambiguity in how data is processed.
If you must use a map for flexibility, always convert it to a typed struct once you understand the data format.
### **Using Other Encoding Formats**
Besides JSON, Go’s standard library includes encoders for XML and Base64. If no existing package supports your format, you can implement a custom encoder using Go’s `encoding` interfaces.
### **Avoiding `encoding/gob` for RPC**
Go’s `encoding/gob` is a binary serialization format primarily for Go-specific RPC via `net/rpc`. However, it’s **not recommended** for general RPC because:
- It’s **Go-specific**, making it incompatible with other languages.
- It lacks **versioning support**, making backward compatibility difficult.
Instead, use a **language-neutral** protocol like **gRPC** (which uses Protocol Buffers) for cross-service communication.
## HTTP Client
The `net/http` package provides an `http.Client` type for making requests and handling responses. Instead of using the default client (`http.DefaultClient`), create your own to ensure timeouts are set:
```go
client := &http.Client{
    Timeout: 30 * time.Second,
}
```
### Making Requests
Create an `*http.Request` using `http.NewRequestWithContext`, passing a context, the HTTP method, and the URL. For `PUT`, `POST`, or `PATCH` requests, provide a request body; otherwise, use `nil`:
```go
req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
if err != nil {
    panic(err)
}
```
### Setting Headers
Modify request headers using the `Header` field:
```go
req.Header.Add("X-My-Client", "Learning Go")
```
### Executing the Request
Use the `Do` method to send the request and receive a response:
```go
res, err := client.Do(req)
if err != nil {
    panic(err)
}
```
### Processing the Response
The `http.Response` object contains:
- `StatusCode` – numeric HTTP status
- `Status` – status text
- `Header` – response headers
- `Body` – response body (`io.ReadCloser`)
To properly handle the response, ensure the body is closed and check the status:
```go
defer res.Body.Close()
if res.StatusCode != http.StatusOK {
    panic(fmt.Sprintf("unexpected status: got %v", res.Status))
}
```
For JSON responses, use `json.Decoder`:
```go
fmt.Println(res.Header.Get("Content-Type"))
var data struct {
    UserID    int    `json:"userId"`
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}
err = json.NewDecoder(res.Body).Decode(&data)
if err != nil {
    panic(err)
}
fmt.Printf("%+v\n", data)
```
**Warning:** Avoid using package-level functions like `http.Get`, `http.Post`, and `http.Head` as they use `http.DefaultClient`, which lacks timeouts.
## HTTP Server
A request to a server is handled by an implementation of the `http.Handler` interface that’s assigned to the `Handler` field. This interface defines a single method:
```go
type Handler interface {
    ServeHTTP(http.ResponseWriter, *http.Request)
}
```
The `*http.Request` should look familiar, as it’s the exact same type that’s used to send a request to an HTTP server. The `http.ResponseWriter` is an interface with three methods:
```go
type ResponseWriter interface {
        Header() http.Header
        Write([]byte) (int, error)
        WriteHeader(statusCode int)
}
```
### Order of Calls
1. **Set Headers** (if needed) using `Header()`:
    
    ```go
    w.Header().Set("Content-Type", "application/json")
    ```
    
2. **Write Status Code** using `WriteHeader()`. This is optional if the status is `200 OK`:
    
    ```go
    w.WriteHeader(http.StatusCreated)
    ```
    
3. **Write the Response Body** using `Write()`:
    
    ```go
    w.Write([]byte(`{"message": "Hello, world!"}`))
    ```
    
### Example Handler
```go
type HelloHandler struct{}
func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello!\n"))
}
```
If the status code is `200 OK`, calling `WriteHeader(http.StatusOK)` is unnecessary, as `Write()` implicitly sets it.
### Creating an HTTP Server
Instantiate an `http.Server`:
```go
s := http.Server{
    Addr:         ":8080",
    ReadTimeout:  30 * time.Second,
    WriteTimeout: 90 * time.Second,
    IdleTimeout:  120 * time.Second,
    Handler:      HelloHandler{},
}
err := s.ListenAndServe()
if err != nil && err != http.ErrServerClosed {
    panic(err)
}
```
Key configurations:
- `Addr` – host and port to listen on
- `ReadTimeout`, `WriteTimeout`, `IdleTimeout` – prevent slow client attacks
- `Handler` – the request handler
### Using `http.ServeMux`
For routing multiple endpoints, use `http.ServeMux`:
```go
mux := http.NewServeMux()
mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello!\n"))
})
```
Set the `Handler` field of `http.Server` to the `ServeMux`:
```go
s.Handler = mux
```
### Path Variables and HTTP Methods (Go 1.22+)
Go 1.22 introduces path-based routing with variables:
```go
mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
    name := r.PathValue("name")
    w.Write([]byte(fmt.Sprintf("Hello, %s!\n", name)))
})
```
### Nested Routing
Multiple `ServeMux` instances can be composed for better organization:
```go
person := http.NewServeMux()
person.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("greetings!\n"))
})
dog := http.NewServeMux()
dog.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("good puppy!\n"))
})
mux := http.NewServeMux()
mux.Handle("/person/", http.StripPrefix("/person", person))
mux.Handle("/dog/", http.StripPrefix("/dog", dog))
```
A request to `/person/greet` is handled by `person`, while `/dog/greet` is handled by `dog`.
**Warning:** Avoid package-level functions (`http.Handle`, `http.HandleFunc`, `http.ListenAndServe`) as they use `http.DefaultServeMux`, which can be modified by third-party libraries, leading to unpredictable behavior. Always instantiate your own `http.Server` and `http.ServeMux`.
## Middleware
Middleware is a common pattern in Go's `net/http` package for performing cross-cutting concerns like authentication, logging, request timing, and more. It works by wrapping `http.Handler` instances to modify request handling before or after calling the next handler.
### **Creating Middleware in Go**
Middleware in Go is implemented as a function that:
1. Accepts an `http.Handler`.
2. Returns a new `http.Handler` that performs extra logic before or after calling the original handler.
Example of middleware that logs request duration:
```go
func RequestTimer(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        h.ServeHTTP(w, r)
        dur := time.Since(start)
        slog.Info("request time", "path", r.URL.Path, "duration", dur)
    })
}
```
**Simple Authentication Middleware**
This middleware checks for a secret password in the request header before allowing access.
```go
var securityMsg = []byte("You didn't give the secret password\n")
func TerribleSecurityProvider(password string) func(http.Handler) http.Handler {
    return func(h http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if r.Header.Get("X-Secret-Password") != password {
                w.WriteHeader(http.StatusUnauthorized)
                w.Write(securityMsg)
                return
            }
            h.ServeHTTP(w, r)
        })
    }
}
```
This middleware:
1. Reads the `X-Secret-Password` header.
2. If it doesn't match the expected value, returns `401 Unauthorized`.
3. Otherwise, calls the next handler.
### **Chaining Middleware**
To use multiple middleware layers, wrap handlers progressively:
```go
terribleSecurity := TerribleSecurityProvider("GOPHER")
mux.Handle("/hello", terribleSecurity(RequestTimer(
    http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello!\n"))
    }),
)))
```
This applies:
1. `terribleSecurity`, which checks the password.
2. `RequestTimer`, which logs the request duration.
3. The final handler, which returns `"Hello!\n"`.
### **Applying Middleware to All Routes**
Since `*http.ServeMux` implements `http.Handler`, you can wrap all registered handlers:
```go
terribleSecurity := TerribleSecurityProvider("GOPHER")
wrappedMux := terribleSecurity(RequestTimer(mux))
s := http.Server{
    Addr:    ":8080",
    Handler: wrappedMux,
}
```
This applies middleware globally.
### **Third-Party Middleware Solutions**
While Go’s function-based middleware pattern is powerful, function chaining can be verbose. The `alice` package simplifies this:
```go
import "github.com/justinas/alice"
helloHandler := func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello!\n"))
}
chain := alice.New(terribleSecurity, RequestTimer).ThenFunc(helloHandler)
mux.Handle("/hello", chain)
```
The `alice.New()` function accepts multiple middleware functions and returns a composed chain.
### **Advanced Routing with Third-Party Packages**
While `http.ServeMux` improved in Go 1.22, it still has limited features. Popular third-party routers like `gorilla/mux` and `chi` provide:
- Route variables.
- Advanced path matching (regex, headers, etc.).
- Middleware support.
Example using `chi`:
```go
import "github.com/go-chi/chi/v5"
r := chi.NewRouter()
r.Use(RequestTimer)
r.Get("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
    name := chi.URLParam(r, "name")
    w.Write([]byte(fmt.Sprintf("Hello, %s!\n", name)))
})
http.ListenAndServe(":8080", r)
```
This allows:
- Named route parameters (`/hello/{name}`).
- Middleware (`r.Use(RequestTimer)`).
For frameworks with even more built-in features, **Echo** and **Gin** provide:
- Automatic request binding (JSON, form data).
- Response JSON rendering.
- Integrated middleware support.
## **ResponseController in Go 1.20+**
Go’s `http.ResponseWriter` interface cannot be changed without breaking compatibility. Instead, Go introduces optional interfaces like `http.Flusher` and `http.Hijacker`.
To improve discoverability and extendability, Go 1.20 introduced `http.ResponseController`:
```go
func handler(rw http.ResponseWriter, req *http.Request) {
    rc := http.NewResponseController(rw)
    for i := 0; i < 10; i++ {
        result := doStuff(i)
        _, err := rw.Write([]byte(result))
        if err != nil {
            slog.Error("error writing", "msg", err)
            return
        }
        err = rc.Flush()
        if err != nil && !errors.Is(err, http.ErrNotSupported) {
            slog.Error("error flushing", "msg", err)
            return
        }
    }
}
```
### **Why Use `ResponseController`?**
1. Provides methods for optional behaviors (e.g., `Flush`).
2. Avoids modifying `http.ResponseWriter`, maintaining Go’s compatibility guarantee.
3. New methods can be added over time without breaking existing code.
### When to Use `http.ResponseController` Instead of `http.ResponseWriter`
#### **Understanding `http.ResponseWriter`**
`http.ResponseWriter` is the primary interface used to send HTTP responses in Go. It provides basic methods for writing headers (`Header()`), setting status codes (`WriteHeader()`), and writing response bodies (`Write()`).
However, `http.ResponseWriter` has limitations:
- It was designed as an interface, which means adding new methods would break backward compatibility.
- Some advanced functionalities, like flushing data immediately or setting deadlines, require additional interfaces (`http.Flusher`, `http.Hijacker`, `http.Pusher`), which are not always implemented by all `http.ResponseWriter` types.
### **Why `http.ResponseController` Exists**
Go 1.20 introduced `http.ResponseController` as a concrete type that acts as a wrapper around `http.ResponseWriter`. It provides a stable way to access optional features without relying on type assertions or checking for interface implementations.
### **When to Use `http.ResponseController`**
Use `http.ResponseController` when you need to:
1. **Flush Data Immediately**
    
    If you need to send partial responses to the client as they are generated (e.g., streaming responses), you can use:
    
    ```go
    rc := http.NewResponseController(w)
    err := rc.Flush()
    if err != nil && !errors.Is(err, http.ErrNotSupported) {
        log.Println("Flushing not supported:", err)
    }
    
    ```
    
    - This avoids having to check if `w` implements `http.Flusher`.
2. **Set Read/Write Deadlines**
    
    If you need to prevent slow clients from hanging your server, `http.ResponseController` provides:
    
    ```go
    rc.SetWriteDeadline(time.Now().Add(2 * time.Second))
    rc.SetReadDeadline(time.Now().Add(2 * time.Second))
    
    ```
    
    - This is useful for time-sensitive applications like WebSockets or long polling.
3. **Work with Optional Features Cleanly**
    
    Instead of checking if `http.ResponseWriter` implements `http.Flusher`, `http.Hijacker`, etc., `http.ResponseController` centralizes these features into a single, forward-compatible API.
    
### **When to Stick with `http.ResponseWriter`**
- If you only need basic response writing (setting headers, writing body, setting status codes), `http.ResponseWriter` is sufficient.
- If your application doesn’t require flushing, hijacking, or setting deadlines, `http.ResponseController` adds no real benefit.
## Structured Logging
### Overview
- The Go standard library initially included the `log` package, which is simple but lacks structured logging.
- Modern applications require structured logs for better processing and analysis.
- The `log/slog` package, introduced in Go 1.21, provides structured logging support.
### Benefits of Structured Logging
- Uses a documented format for each log entry.
- Easier for log processing tools to analyze and detect patterns.
- JSON is commonly used, but even whitespace-separated key-value pairs improve readability and parsing.
### Why `log/slog` Was Added to the Standard Library
1. **Prevents Logging Fragmentation**
    - Many third-party structured loggers exist (`zap`, `logrus`, `go-kit log`), leading to inconsistencies.
    - `log/slog` standardizes structured logging, making modules work together more seamlessly.
2. **Separate from the `log` Package**
    - Structured and unstructured logging have different design philosophies.
    - Keeping them separate avoids API confusion.
3. **Scalable API**
    - Provides default logging functions for different levels:
        
        ```go
        slog.Debug("debug log message")
        slog.Info("info log message")
        slog.Warn("warning log message")
        slog.Error("error log message")
        ```
        
    - Default logger suppresses debug messages.
    - Uses structured formatting (whitespace-separated fields).
### Adding Custom Fields to Logs
- Logs can include additional key-value pairs:
**Output:**
    
    ```go
    userID := "fred"
    loginCount := 20
    slog.Info("user login",
        "id", userID,
        "login_count", loginCount)
    ```
    
    ```
    2023/04/20 23:36:38 INFO user login id=fred login_count=20
    ```
    
### Using JSON for Logs
- To output logs as JSON and control the logging level:
    
    ```go
    options := &slog.HandlerOptions{Level: slog.LevelDebug}
    handler := slog.NewJSONHandler(os.Stderr, options)
    mySlog := slog.New(handler)
    
    lastLogin := time.Date(2023, 01, 01, 11, 50, 00, 00, time.UTC)
    mySlog.Debug("debug message",
        "id", userID,
        "last_login", lastLogin)
    ```
    
    ```json
    {"time":"2023-04-22T23:30:01.170243-04:00","level":"DEBUG",
     "msg":"debug message","id":"fred","last_login":"2023-01-01T11:50:00Z"}
    ```
    
### Custom Logging Handlers
- If JSON and text formats are insufficient, implement a custom `slog.Handler` and pass it to `slog.New`.
### Performance Considerations
- Logging can impact performance, especially with high-frequency logs.
- Use `LogAttrs` for more efficient logging:
    
    ```go
    mySlog.LogAttrs(ctx, slog.LevelInfo, "faster logging",
                    slog.String("id", userID),
                    slog.Time("last_login", lastLogin))
    ```
    
    - Uses `slog.Attr` instances to reduce memory allocations.
### Compatibility with `log`
- The `log` package remains available due to Go’s compatibility promise.
- Bridge `log.Logger` with `slog.Handler` using:
    
    ```go
    myLog := slog.NewLogLogger(mySlog.Handler(), slog.LevelDebug)
    myLog.Println("using the mySlog Handler")
    ```
    
    ```json
    {"time":"2023-04-22T23:30:01.170269-04:00","level":"DEBUG",
     "msg":"using the mySlog Handler"}
    ```
    
## Context in Go
### Overview
- Servers need a way to handle metadata in requests.
- Metadata falls into two categories:
    - **Required metadata** (e.g., tracking IDs).
    - **Metadata for stopping execution** (e.g., timeouts).
- Many languages use **thread-local storage**, but Go doesn’t have thread identity due to goroutines.
- Instead, Go provides **context** via the `context` package.
### What Is `context.Context`?
- `context.Context` is an interface, not a language feature.
- Encourages **explicit data passing** as a function parameter.
- Convention:
    - The **first** parameter in functions that use context should be `ctx context.Context`.
Example:
```go
func logic(ctx context.Context, info string) (string, error) {
    // process the request using context
    return "", nil
}
```
### Creating Contexts
- When no context exists (e.g., in `main`), use:
    
    ```go
    ctx := context.Background()
    ```
    
- During development, you can use `context.TODO()` as a placeholder.
### Context in HTTP Servers
- `http.Request` includes **two context-related methods**:
    - `req.Context()` → Extracts the `context.Context` from the request.
    - `req.WithContext(ctx)` → Creates a new request with a modified context.
### Using Context in Middleware
- Middleware extracts the existing context, modifies it, and passes it down.
Example:
```go
func Middleware(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        ctx := req.Context() // Extract context
        req = req.WithContext(ctx) // Modify and attach new context
        handler.ServeHTTP(rw, req)
    })
}
```
### Using Context in HTTP Handlers
- Extract the context from `http.Request` and pass it to business logic.
Example:
```go
func handler(rw http.ResponseWriter, req *http.Request) {
    ctx := req.Context()
    err := req.ParseForm()
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    data := req.FormValue("data")
    result, err := logic(ctx, data)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }
    rw.Write([]byte(result))
}
```
### Using Context for HTTP Requests
- When making requests to another service, use `http.NewRequestWithContext()`.
Example:
```go
type ServiceCaller struct {
    client *http.Client
}
func (sc ServiceCaller) callAnotherService(ctx context.Context, data string) (string, error) {
    req, err := http.NewRequestWithContext(ctx, http.MethodGet,
                "http://example.com?data="+data, nil)
    if err != nil {
        return "", err
    }
    resp, err := sc.client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("Unexpected status code %d", resp.StatusCode)
    }
    id, err := processResponse(resp.Body)
    return id, err
}
```
## Context and Values in Go
### Explicit Data Passing
- Go prefers **explicit data passing** over implicit approaches.
- If a function depends on certain data, it should be **clear** where the data comes from.
- **Exception:** When explicit passing isn't possible (e.g., HTTP handlers and middleware).
### Using Context for Value Storage
- Middleware often needs to store data in **context** (e.g., user extracted from a JWT, request GUID).
- Use `context.WithValue` to store values:
    
    ```go
    ctx := context.WithValue(parentCtx, key, value)
    ```
    
    - Returns a **new child context** that wraps the parent.
    - Contexts are **immutable**; values are added by wrapping existing contexts.
- Retrieving values from context:
    
    ```go
    val, ok := ctx.Value(key).(int)
    if !ok {
        fmt.Println("no value")
    } else {
        fmt.Println("value:", val)
    }
    ```
    
    - If key is missing, `ctx.Value` returns `nil`.
    - Searching the context chain is a **linear** operation (avoid excessive values).
### Choosing a Context Key
- **Keys must be comparable** (like map keys).
- Avoid using simple strings like `"id"` (risk of key collision across packages).
- **Use an unexported type for keys** to prevent collisions:
    
    ```go
    type userKey int
    const (
        _ userKey = iota
        key
    )
    ```
    
    - Ensures only your package can set/retrieve the value.
- **Alternative:** Use an empty struct as the key:
    
    ```go
    type userKey struct{}
    ```
    
    - Works well when only one key is needed.
### Creating Helper Functions
- Use **`ContextWith` prefix** for setting values and **`FromContext` suffix** for retrieval.
    
    ```go
    func ContextWithUser(ctx context.Context, user string) context.Context {
        return context.WithValue(ctx, key, user)
    }
    
    func UserFromContext(ctx context.Context) (string, bool) {
        user, ok := ctx.Value(key).(string)
        return user, ok
    }
    ```
    
### Middleware Example (Extracting User)
- Middleware extracts user ID from a cookie and stores it in the request's context.
    
    ```go
    // a real implementation would be signed to make sure
    // the user didn't spoof their identity
    func extractUser(req *http.Request) (string, error) {
        userCookie, err := req.Cookie('identity')
        if err != nil {
            return '', err
        }
        return userCookie.Value, nil
    }
    
    func Middleware(h http.Handler) http.Handler {
        return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
            user, err := extractUser(req)
            if err != nil {
                rw.WriteHeader(http.StatusUnauthorized)
                rw.Write([]byte("unauthorized"))
                return
            }
            ctx := ContextWithUser(req.Context(), user)
            req = req.WithContext(ctx)
            h.ServeHTTP(rw, req)
        })
    }
    ```
    
    - Calls `extractUser(req)`.
    - Stores user in the request's **context**.
    - Wraps request with new context (`req.WithContext(ctx)`) before passing to the handler.
### Retrieving Context Data in Handlers
- **Extract user from context** before passing to business logic:
    
    ```go
    func (c Controller) DoLogic(rw http.ResponseWriter, req *http.Request) {
        ctx := req.Context()
        user, ok := identity.UserFromContext(ctx)
        if !ok {
            rw.WriteHeader(http.StatusInternalServerError)
            return
        }
        data := req.URL.Query().Get("data")
        result, err := c.Logic.BusinessLogic(ctx, user, data)
        if err != nil {
            rw.WriteHeader(http.StatusInternalServerError)
            rw.Write([]byte(err.Error()))
            return
        }
        rw.Write([]byte(result))
    }
    ```
    
    - Extracts `user` from context.
    - Passes it **explicitly** to business logic.
### When to Keep Values in Context
- **Pass values explicitly** to business logic.
- **Keep system metadata in context** (e.g., tracking GUIDs).
- Tracking GUIDs:
    - Avoid cluttering function signatures.
    - Allow seamless propagation through third-party libraries.
## Example: Tracking GUID in Context
### Middleware to Attach a GUID
- Uses a GUID **from request headers** or **generates a new one**.
    
    ```go
    func Middleware(h http.Handler) http.Handler {
        return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
            ctx := req.Context()
            if guid := req.Header.Get("X-GUID"); guid != "" {
                ctx = contextWithGUID(ctx, guid)
            } else {
                ctx = contextWithGUID(ctx, uuid.New().String())
            }
            req = req.WithContext(ctx)
            h.ServeHTTP(rw, req)
        })
    }
    ```
    
### Logging with GUID
- Logs messages with the request’s GUID (if available).
    
    ```go
    type Logger struct{}
    
    func (Logger) Log(ctx context.Context, message string) {
        if guid, ok := guidFromContext(ctx); ok {
            message = fmt.Sprintf("GUID: %s - %s", guid, message)
        }
        fmt.Println(message)
    }
    ```
    
### Propagating GUID to External Requests
- Attach GUID to outbound requests:
    
    ```go
    func Request(req *http.Request) *http.Request {
        ctx := req.Context()
        if guid, ok := guidFromContext(ctx); ok {
            req.Header.Add("X-GUID", guid)
        }
        return req
    }
    ```
    
### Business Logic with Context Logging
- Business logic **logs events** and **propagates GUID** to downstream services.
    
    ```go
    type LogicImpl struct {
        RequestDecorator RequestDecorator
        Logger           Logger
        Remote           string
    }
    
    func (l LogicImpl) Process(ctx context.Context, data string) (string, error) {
        l.Logger.Log(ctx, "starting Process with "+data)
        req, err := http.NewRequestWithContext(ctx, http.MethodGet, l.Remote+"/second?query="+data, nil)
        if err != nil {
            l.Logger.Log(ctx, "error building remote request:"+err.Error())
            return "", err
        }
        req = l.RequestDecorator(req)
        resp, err := http.DefaultClient.Do(req)
        // process the response...
    }
    ```
    
### Wiring Up Dependencies in `main`
```go
controller := Controller{
    Logic: LogicImpl{
        RequestDecorator: tracker.Request,
        Logger:           tracker.Logger{},
        Remote:           "http://localhost:4000",
    },
}
```
### Best Practices
- **Use context for passing values through standard APIs**.
- **Copy values from context into explicit parameters** for business logic.
- **Keep system metadata in context** (e.g., tracking GUIDs).
## Cancellation in Go with `context`
### Overview
Go’s `context` package allows for:
- Passing metadata and working around HTTP API limitations.
- Controlling application responsiveness and coordinating concurrent goroutines.
- Cancelling operations when a condition is met.
### Creating a Cancellable Context
To create a cancellable context, use `context.WithCancel`:
```go
ctx, cancelFunc := context.WithCancel(context.Background())
defer cancelFunc() // Ensures cleanup
```
- `context.WithCancel` returns:
    - A child `context.Context`
    - A `context.CancelFunc` that, when called, signals cancellation.
### Detecting Cancellation
- `ctx.Done()` returns a channel that is closed when cancellation occurs.
- Reading from a nil channel (when using `Done()` on a non-cancellable context) results in a deadlock. → do it in case inside a select statement
### Example: Cancelling Goroutines
Consider a program that calls two HTTP services:
1. A service returning random status codes (e.g., `200, 500`).
2. A service with a delayed response.
### Setup
```go
ctx, cancelFunc := context.WithCancel(context.Background())
defer cancelFunc()
ch := make(chan string)
var wg sync.WaitGroup
wg.Add(2)
```
### Goroutine: Random Status Check
```go
go func() {
    defer wg.Done()
    for {
        resp, err := makeRequest(ctx, "http://httpbin.org/status/200,200,200,500")
        if err != nil {
            fmt.Println("error in status goroutine:", err)
            cancelFunc()
            return
        }
        if resp.StatusCode == http.StatusInternalServerError {
            fmt.Println("bad status, exiting")
            cancelFunc()
            return
        }
        select {
        case ch <- "success from status":
        case <-ctx.Done(): // Stop if cancelled
        }
        time.Sleep(1 * time.Second)
    }
}()
```
### Goroutine: Delayed Response
```go
go func() {
    defer wg.Done()
    for {
        resp, err := makeRequest(ctx, "http://httpbin.org/delay/1")
        if err != nil {
            fmt.Println("error in delay goroutine:", err)
            cancelFunc()
            return
        }
        select {
        case ch <- "success from delay: " + resp.Header.Get("date"):
        case <-ctx.Done():
        }
    }
}()
```
### Main Execution Loop
```go
loop:
    for {
        select {
        case s := <-ch:
            fmt.Println("in main:", s)
        case <-ctx.Done():
            fmt.Println("in main: cancelled!")
            break loop
        }
    }
    wg.Wait()
```
### Sample Output
```
in main: success from status
in main: success from delay: Thu, 16 Feb 2023 03:53:57 GMT
bad status, exiting
in main: cancelled!
error in delay goroutine: Get "http://httpbin.org/delay/1": context canceled
```
- Calling `cancelFunc()` multiple times is safe.
- The Go HTTP client respects cancellation, stopping in-progress requests.
## Using `WithCancelCause` to Capture Errors
Introduced in Go 1.20, `WithCancelCause` allows passing an error when cancelling:
```go
ctx, cancelFunc := context.WithCancelCause(context.Background())
defer cancelFunc(nil) // Cleanup
```
### Modifying the Status Goroutine
```go
resp, err := makeRequest(ctx, "http://httpbin.org/status/200,200,200,500")
if err != nil {
    cancelFunc(fmt.Errorf("in status goroutine: %w", err))
    return
}
if resp.StatusCode == http.StatusInternalServerError {
    cancelFunc(errors.New("bad status"))
    return
}
ch <- "success from status"
time.Sleep(1 * time.Second)
```
### Capturing the Cancellation Cause
Modify the main loop:
```go
loop:
    for {
        select {
        case s := <-ch:
            fmt.Println("in main:", s)
        case <-ctx.Done():
            fmt.Println("in main: cancelled with error", context.Cause(ctx))
            break loop
        }
    }
    wg.Wait()
    fmt.Println("context cause:", context.Cause(ctx))
```
### Sample Output
```
in main: success from status
in main: success from delay: Thu, 16 Feb 2023 04:11:49 GMT
in main: cancelled with error bad status
in delay goroutine: Get "http://httpbin.org/delay/1": context canceled
context cause: bad status
```
- `context.Cause(ctx)` returns the first error passed to `cancelFunc()`.
- Later calls to `cancelFunc(err)` do not overwrite the original error.
## Contexts with Deadlines in Go
### Why Manage Request Timeouts?
A server must manage requests efficiently to ensure fair resource allocation. Without limits, users may monopolize resources, causing performance degradation. Servers can manage load by:
- **Limiting simultaneous requests** (controlled via goroutines).
- **Limiting queued requests** (controlled via buffered channels).
- **Limiting request execution time** (handled via `context`).
- **Limiting resource usage** (manual implementation required).
### Using Context for Timeouts
Go provides two functions to enforce time limits:
- **`context.WithTimeout(parentCtx, duration)`**
    - Cancels the context automatically after `duration`.
- **`context.WithDeadline(parentCtx, time)`**
    - Cancels the context automatically at `time`.
Both return:
1. A new context that enforces the timeout.
2. A cancellation function that should be called to release resources.
**Note:** If you pass a past deadline, the context is created as already canceled.
### Checking Context Expiration
- **`ctx.Deadline()`** → Returns `(time.Time, bool)` indicating timeout settings.
- **`ctx.Err()`** → Returns:
    - `nil` if still active.
    - `context.Canceled` if explicitly canceled.
    - `context.DeadlineExceeded` if it timed out.
### Parent-Child Context Relationship
Child contexts inherit constraints from their parent:
```go
ctx := context.Background()
parent, cancel := context.WithTimeout(ctx, 2*time.Second)
defer cancel()
child, cancel2 := context.WithTimeout(parent, 3*time.Second)
defer cancel2()
<-child.Done()
fmt.Println("Elapsed:", time.Since(start).Truncate(time.Second)) // Outputs 2s
```
Here, `child` has a 3s timeout, but since `parent` cancels after 2s, `child` cancels at 2s as well.
### Applying Timeouts in HTTP Requests
For HTTP requests:
```go
ctx, cancelFuncParent := context.WithTimeout(context.Background(), 3*time.Second)
defer cancelFuncParent()
ctx, cancelFunc := context.WithCancelCause(ctx)
defer cancelFunc(nil)
```
If the request takes longer than 3 seconds, it automatically cancels.
If you want the option of returning the error for the cancellation cause, you need to wrap a context created by `WithTimeout` or `WithDeadline` in a context created by `WithCancelCause`. You must `defer` both cancellation functions to keep resources from being leaked. If you want to return a custom sentinel error when a context times out, use the `context.WithTimeoutCause` or `context.WithDeadlineCause` functions instead.
### Handling Cancellation in Your Code
**Use `select` with `ctx.Done()`** when working with channels:
```go
select {
case msg := <-ch:
    process(msg)
case <-ctx.Done():
    return
}
```
**Check `context.Cause(ctx)` periodically** for long-running operations:
```go
for {
    if err := context.Cause(ctx); err != nil {
        return "", err
    }
    // Continue processing
}
```
### Example: Computing π with Cancellation
```go
i := 0
for {
    if err := context.Cause(ctx); err != nil {
        fmt.Println('cancelled after', i, 'iterations')
        return sum.Text('g', 100), err
    }
    var diff big.Float
    diff.SetInt64(4)
    diff.Quo(&diff, &d)
    if i%2 == 0 {
        sum.Add(&sum, &diff)
    } else {
        sum.Sub(&sum, &diff)
    }
    d.Add(&d, two)
    i++
}
```
This ensures the computation stops if the context is canceled.
## Writing Tests
### Understanding the Basics of Testing
Go’s testing system has **two key parts**:
- **Libraries**: The `testing` package provides types and functions for writing tests.
- **Tooling**: The `go test` command runs tests and generates reports.
Unlike many languages, **Go places tests in the same directory and package** as the production code. This allows tests to access unexported functions and variables.
## Writing a Simple Test
### Production Code (`adder.go`)
```go
func addNumbers(x, y int) int {
    return x + x // Bug: should be x + y
}
```
### Test Code (`adder_test.go`)
```go
func Test_addNumbers(t *testing.T) {
    result := addNumbers(2, 3)
    if result != 5 {
        t.Error("incorrect result: expected 5, got", result)
    }
}
```
**Key points:**
- Test file names **must end with** `_test.go` (e.g., `adder_test.go` for `adder.go`).
- Test function names **must start with** `Test` and take a `testing.T` parameter.
- Use `t.Error()` to report failures.
### Running the Test
```
$ go test
--- FAIL: Test_addNumbers (0.00s)
    adder_test.go:8: incorrect result: expected 5, got 4
FAIL
```
Fix the bug (`return x + y`), rerun `go test`, and it should pass.
## Reporting Test Failures
### Methods for Reporting Failures
| Method | Behavior |
| --- | --- |
| `t.Error(args...)` | Prints an error message but **continues execution**. |
| `t.Errorf(format, args...)` | Similar to `t.Error`, but with `Printf`-style formatting. |
| `t.Fatal(args...)` | Prints an error message and **stops execution** of the test function. |
| `t.Fatalf(format, args...)` | Similar to `t.Fatal`, but with formatting. |
**Guidelines:**
- Use `t.Fatal()` when a failure **makes further tests meaningless**.
- Use `t.Error()` when you want to **report multiple failures in one run**.
## Setting Up and Tearing Down
### Using `TestMain`
`TestMain` is used for **setting up state before tests** and **cleaning up afterward**.
```go
var testTime time.Time
func TestMain(m *testing.M) {
    fmt.Println("Setting up tests")
    testTime = time.Now()
    exitVal := m.Run() // Run all tests
    fmt.Println("Cleaning up")
    os.Exit(exitVal)
}
func TestFirst(t *testing.T) {
    fmt.Println('TestFirst uses stuff set up in TestMain', testTime)
}
func TestSecond(t *testing.T) {
    fmt.Println('TestSecond also uses stuff set up in TestMain', testTime)
}
```
**Key points:**
- `TestMain` is **called once per package**, not before each test.
- **Use cases**:
    - Setting up external dependencies (e.g., databases).
    - Initializing package-level variables.
- **Avoid relying on `TestMain`** for package-level state; consider refactoring your code so it doesn’t use package level variables
## Cleaning Up After Individual Tests
### Using `t.Cleanup()`
`t.Cleanup()` **registers a function to run after the test completes**.
```go
// createFile is a helper function called from multiple tests
func createFile(t *testing.T) (_ string, err error) {
    f, err := os.Create('tempFile')
    if err != nil {
        return '', err
    }
    defer func() {
        err = errors.Join(err, f.Close())
    }()
    // write some data to 08_file_servers_cdn
    t.Cleanup(func() {
        os.Remove(f.Name())
    })
    return f.Name(), nil
}
func TestFileProcessing(t *testing.T) {
    fName, err := createFile(t)
    if err != nil {
        t.Fatal(err)
    }
    // do testing, don't worry about cleanup
}
```
**Alternative:** `t.TempDir()`
```go
// createFile is a helper function called from multiple tests
func createFile(tempDir string) (_ string, err error) {
    f, err := os.CreateTemp(tempDir, 'tempFile')
    if err != nil {
        return '', err
    }
    defer func() {
        err = errors.Join(err, f.Close())
    }()
    // write some data to 08_file_servers_cdn
    return f.Name(), nil
}
func TestFileProcessing(t *testing.T) {
    tempDir := t.TempDir()
    fName, err := createFile(tempDir)
    if err != nil {
        t.Fatal(err)
    }
    // do testing, don't worry about cleanup
}
```
## Testing with Environment Variables
Use `t.Setenv()` to set environment variables **that automatically reset after the test**.
```go
func TestEnvVarProcess(t *testing.T) {
    t.Setenv("OUTPUT_FORMAT", "JSON")
    cfg := ProcessEnvVars()
    if cfg.OutputFormat != "JSON" {
        t.Error("OutputFormat not set correctly")
    }
}
```
### Best Practices:
- **Store env vars in a config struct** before running application logic.
- **Use third-party libraries** like [Viper](https://github.com/spf13/viper) or [GoDotEnv](https://github.com/joho/godotenv) for better configuration management.
## Storing Sample Test Data
- Go **reserves the directory name** `testdata` for sample files.
- **Always use relative paths** to access files inside `testdata`.
Example:
```go
data, err := os.ReadFile("testdata/sample.txt")
if err != nil {
    t.Fatal(err)
}
```
## Caching Test Results
- Go **caches test results** if files haven't changed.
- Use `count=1` to **force tests to rerun**:
    
    ```
    $ go test -count=1
    ```
    
## Testing Your Public API
By default, tests in Go are written in the same package as the production code. This allows testing both exported and unexported functions. However, if you want to test only the **public API**, Go provides a convention for this.
Instead of using the same package name, the test file should use `packagename_test`. The test files remain in the same directory as the source code but must import the package explicitly.
### Example
If the `pubadder` package has the following function:
```go
package pubadder
func AddNumbers(x, y int) int {
    return x + y
}
```
You can test it as a public API using `pubadder_test` as the package name:
```go
package pubadder_test
import (
    "testing"
    "github.com/learning-go-book-2e/ch15/sample_code/pubadder"
)
func TestAddNumbers(t *testing.T) {
    result := pubadder.AddNumbers(2, 3)
    if result != 5 {
        t.Errorf("incorrect result: expected 5, got %d", result)
    }
}
```
### Key Points
- The test file uses `pubadder_test` instead of `pubadder`.
- The package must be imported explicitly (`pubadder`).
- This method forces testing the package **as a black box**, interacting only through exported functions.
## Using `go-cmp` to Compare Test Results
When comparing structs, maps, or slices, Go provides `reflect.DeepEqual`. However, **Google's `go-cmp` package** provides a more detailed comparison.
### Example
Given this struct and function:
```go
type Person struct {
    Name      string
    Age       int
    DateAdded time.Time
}
func CreatePerson(name string, age int) Person {
    return Person{
        Name:      name,
        Age:       age,
        DateAdded: time.Now(),
    }
}
```
A test using `go-cmp` looks like this:
```go
import (
    "testing"
    "github.com/google/go-cmp/cmp"
)
func TestCreatePerson(t *testing.T) {
    expected := Person{
        Name: "Dennis",
        Age:  37,
    }
    result := CreatePerson("Dennis", 37)
    if diff := cmp.Diff(expected, result); diff != "" {
        t.Error(diff)
    }
}
```
### Output if `DateAdded` differs:
```
--- FAIL: TestCreatePerson (0.00s)
    ch13_cmp_test.go:16:   ch13_cmp.Person{
              Name:      "Dennis",
              Age:       37,
        -     DateAdded: "0001-01-01 00:00:00 +0000 UTC",
        +     DateAdded: "2020-03-01 22:53:58.087229 -0500 EST",
    }
```
The `-` and `+` indicate differences. Since `DateAdded` is time-dependent, it should be ignored in comparisons.
### Ignoring Fields with a Custom Comparator
Use `cmp.Comparer` to define custom comparison logic:
```go
comparer := cmp.Comparer(func(x, y Person) bool {
    return x.Name == y.Name && x.Age == y.Age
})
if diff := cmp.Diff(expected, result, comparer); diff != "" {
    t.Error(diff)
}
```
- This ignores `DateAdded`, only comparing `Name` and `Age`.
- The function must be **symmetric, deterministic, and pure**.
## Running Table Tests
For functions with multiple test cases, a **table-driven test** prevents code repetition.
### Example
Consider this function in the `table` package:
```go
func DoMath(num1, num2 int, op string) (int, error) {
    switch op {
    case "+":
        return num1 + num2, nil
    case "-":
        return num1 - num2, nil
    case "*":
        return num1 * num2, nil
    case "/":
        if num2 == 0 {
            return 0, errors.New("division by zero")
        }
        return num1 / num2, nil
    default:
        return 0, fmt.Errorf("unknown operator %s", op)
    }
}
```
A **repetitive test** would look like this:
```go
func TestDoMath(t *testing.T) {
    result, err := DoMath(2, 2, "+")
    if result != 4 {
        t.Error("Should have been 4, got", result)
    }
    if err != nil {
        t.Error("Should have been nil error, got", err)
    }
    result2, err2 := DoMath(2, 2, "-")
    if result2 != 0 {
        t.Error("Should have been 0, got", result2)
    }
    if err2 != nil {
        t.Error("Should have been nil error, got", err2)
    }
    // and so on...
}
```
This can be **simplified with a table-driven test**:
```go
func TestDoMath(t *testing.T) {
    data := []struct {
        name     string
        num1     int
        num2     int
        op       string
        expected int
        errMsg   string
    }{
        {"addition", 2, 2, "+", 4, ""},
        {"subtraction", 2, 2, "-", 0, ""},
        {"multiplication", 2, 2, "*", 4, ""},
        {"division", 2, 2, "/", 1, ""},
        {"bad_division", 2, 0, "/", 0, "division by zero"},
    }
    for _, d := range data {
        t.Run(d.name, func(t *testing.T) {
            result, err := DoMath(d.num1, d.num2, d.op)
            if result != d.expected {
                t.Errorf("Expected %d, got %d", d.expected, result)
            }
            var errMsg string
            if err != nil {
                errMsg = err.Error()
            }
            if errMsg != d.errMsg {
                t.Errorf("Expected error message `%s`, got `%s`", d.errMsg, errMsg)
            }
        })
    }
}
```
### Key Points
- Uses a **slice of structs** to store test cases.
- The `t.Run` function runs each test case separately.
- Running tests with `v` shows named subtests.
### Comparing Error Messages
Error messages might change in future versions, so comparing strings can be fragile.
### Better Error Checking
If an error has a **custom type** or **sentinel error**, use:
- `errors.Is(err, targetErr)`: Checks if `err` matches `targetErr`.
- `errors.As(err, &targetType)`: Extracts the underlying error type.
```go
if errors.Is(err, ErrDivisionByZero) {
    // Handle division by zero
}
```
## Running Tests Concurrently
By default, unit tests in Go run sequentially. Since each test is meant to be **independent**, they are good candidates for concurrency. To run a test concurrently, call `t.Parallel()` at the start of the test:
```go
func TestMyCode(t *testing.T) {
    t.Parallel()
    // rest of test goes here
}
```
- Parallel tests run **concurrently** with other parallel tests.
- This can **speed up long-running test suites**.
- However, tests relying on **shared mutable state** should **not** be parallel, as they may produce inconsistent results.
### Caution: Using `t.Parallel()` with `Setenv()`
Tests that call `t.Parallel()` **must not** use `Setenv()`, or they will **panic**.
## Running Table Tests in Parallel
When running **table-driven tests** in parallel, be aware of how **loop variables** are captured.
### Example Issue in Go 1.21 and Earlier
In Go **1.21 or earlier**, the following test will behave unpredictably:
```go
func TestParallelTable(t *testing.T) {
    data := []struct {
        name   string
        input  int
        output int
    }{
        {"a", 10, 20},
        {"b", 30, 40},
        {"c", 50, 60},
    }
    for _, d := range data {
        t.Run(d.name, func(t *testing.T) {
            t.Parallel()
            fmt.Println(d.input, d.output)
            out := toTest(d.input)
            if out != d.output {
                t.Error("didn't match", out, d.output)
            }
        })
    }
}
```
### Unexpected Output
```
=== CONT  TestParallelTable/a
50 60
=== CONT  TestParallelTable/c
50 60
=== CONT  TestParallelTable/b
50 60
```
**Why?**
- The loop variable `d` is **shared** across all parallel tests.
- Each goroutine sees the **last** value of `d` (50, 60).
### Solution: Shadow the Loop Variable
If using **Go 1.21 or earlier**, shadow the loop variable before `t.Run`:
```go
for _, d := range data {
    d := d // Shadowing the loop variable
    t.Run(d.name, func(t *testing.T) {
        t.Parallel()
        fmt.Println(d.input, d.output)
        out := toTest(d.input)
        if out != d.output {
            t.Error("didn't match", out, d.output)
        }
    })
}
```
This ensures each test gets its **own copy** of `d`.
### Fix in Go 1.22
In **Go 1.22 and later**, loop variables are scoped correctly, so this issue no longer occurs.
## Checking Your Code Coverage
Code coverage helps ensure that all parts of your code are tested. However, **100% code coverage does not guarantee bug-free code**.
### Running Tests with Coverage
Use the `-cover` flag to display **code coverage**:
```
go test -v -cover
```
To save coverage data to a file:
```
go test -v -cover -coverprofile=c.out
```
### Visualizing Coverage
Use `go tool cover` to generate an **HTML report**:
```
go tool cover -html=c.out
```
This opens a browser showing:
- **Gray**: Non-testable code (e.g., comments, type declarations).
- **Green**: Covered by tests.
- **Red**: Not covered by tests
![image.png](media/image_9.png)
### Example: Missing a Test Case
Consider this function:
```go
func DoMath(num1, num2 int, op string) (int, error) {
    switch op {
    case "+":
        return num1 + num2, nil
    case "-":
        return num1 - num2, nil
    case "*":
        return num1 * num2, nil
    case "/":
        if num2 == 0 {
            return 0, errors.New("division by zero")
        }
        return num1 / num2, nil
    default:
        return 0, fmt.Errorf("unknown operator %s", op)
    }
}
```
### Missing Test Case
Initially, the test cases might be:
```go
data := []struct {
    name     string
    num1     int
    num2     int
    op       string
    expected int
    errMsg   string
}{
    {"addition", 2, 2, "+", 4, ""},
    {"subtraction", 2, 2, "-", 0, ""},
    {"multiplication", 2, 2, "*", 4, ""},
    {"division", 2, 2, "/", 1, ""},
    {"bad_division", 2, 0, "/", 0, "division by zero"},
}
```
- Running `go tool cover -html=c.out` **reveals untested lines**.
- The **default case** (`"?"` operator) is **not covered**.
### Coverage ≠ Bug-Free Code
Even with **100% coverage**, your code may still have **bugs**.
### Example: A Logic Error in Multiplication
```go
case "*":
    return num1 + num2, nil // Oops! Should be multiplication
```
- This **passes coverage tests** but **fails logically**.
- Adding another test case reveals the issue:
```go
{"another_mult", 2, 3, "*", 6, ""},
```
## Fuzzing in Go
### Importance of Fuzzing
- **All input data is suspect** – Even if a program works correctly for expected input, real-world data can be corrupted, malformed, or malicious.
- **Unit tests cannot cover everything** – Even with 100% test coverage, unexpected edge cases can still break a program.
- **Fuzzing automates input testing** – By generating random inputs, fuzzing helps discover issues that might not be obvious in manually written tests.
- **Ensures robustness** – By running functions against a variety of unexpected inputs, fuzzing can reveal crashes, panics, excessive resource usage, or incorrect behavior.
### How Fuzzing Works
1. **Generates test cases** – The fuzzer creates random input data to feed into a function.
2. **Uses a seed corpus** – It starts with known good data and mutates it to find edge cases.
3. **Detects failures** – If the function panics, crashes, or returns unexpected output, the fuzzer records the failing input.
4. **Finds minimal failure cases** – When a bug is found, the fuzzer simplifies the input to the smallest failing example.
5. **Saves test cases** – The fuzzer stores failing inputs as new test cases for regression testing.
## Example: Fuzzing `ParseData` Function
### The `ParseData` Function
This function reads structured data from an `io.Reader`. The input format:
- The **first line** contains an integer, indicating the number of strings to read.
- The function then reads exactly that many lines and returns them as a slice.
- If the format is invalid (wrong number, missing lines, etc.), it returns an error.
### Implementation:
```go
func ParseData(r io.Reader) ([]string, error) {
    s := bufio.NewScanner(r)
    // Read first line
    if !s.Scan() {
        return nil, errors.New("empty")
    }
    countStr := s.Text()
    // Convert first line to an integer
    count, err := strconv.Atoi(countStr)
    if err != nil {
        return nil, err
    }
    // Read the specified number of lines
    out := make([]string, 0, count)
    for i := 0; i < count; i++ {
        hasLine := s.Scan()
        if !hasLine {
            return nil, errors.New("too few lines")
        }
        line := s.Text()
        out = append(out, line)
    }
    return out, nil
}
```
## Unit Testing `ParseData`
Before fuzzing, we write **unit tests** to check expected inputs.
```go
func TestParseData(t *testing.T) {
    data := []struct {
        name   string
        in     []byte
        out    []string
        errMsg string
    }{
        {"simple", []byte("3\nhello\ngoodbye\ngreetings\n"), []string{"hello", "goodbye", "greetings"}, ""},
        {"empty_error", []byte(""), nil, "empty"},
        {"zero", []byte("0\n"), []string{}, ""},
        {"number_error", []byte("asdf\nhello\ngoodbye\ngreetings\n"), nil, `strconv.Atoi: parsing "asdf": invalid syntax`},
        {"line_count_error", []byte("4\nhello\ngoodbye\ngreetings\n"), nil, "too few lines"},
    }
    for _, d := range data {
        t.Run(d.name, func(t *testing.T) {
            r := bytes.NewReader(d.in)
            out, err := ParseData(r)
            var errMsg string
            if err != nil {
                errMsg = err.Error()
            }
            if diff := cmp.Diff(d.out, out); diff != "" {
                t.Error(diff)
            }
            if diff := cmp.Diff(d.errMsg, errMsg); diff != "" {
                t.Error(diff)
            }
        })
    }
}
```
## Adding a Fuzz Test
### Why Use Fuzzing?
- While unit tests cover expected cases, fuzzing helps find **unexpected** inputs that might break the function.
- It can discover:
    - **Panics** (e.g., negative slice capacities).
    - **Excessive memory usage** (e.g., allocating massive slices).
    - **Logical errors** (e.g., blank lines causing incorrect parsing).
### Writing the Fuzz Test
- We create **random inputs** and check if `ParseData` behaves correctly.
- If the function returns valid data, we **round-trip test** it:
    1. Convert the output back into an input format.
    2. Parse the re-generated data.
    3. Ensure both outputs match.
```go
func FuzzParseData(f *testing.F) {
    testcases := [][]byte{
        []byte("3\nhello\ngoodbye\ngreetings\n"),
        []byte("0\n"),
    }
    for _, tc := range testcases {
        f.Add(tc)
    }
    f.Fuzz(func(t *testing.T, in []byte) {
        r := bytes.NewReader(in)
        out, err := ParseData(r)
        if err != nil {
            t.Skip("handled error")
        }
        roundTrip := ToData(out)
        rtr := bytes.NewReader(roundTrip)
        out2, err := ParseData(rtr)
        if diff := cmp.Diff(out, out2); diff != "" {
            t.Error(diff)
        }
    })
}
```
### Running the Fuzzer
```
go test -fuzz=FuzzParseData
```
- Runs indefinitely, generating new test cases.
- Stops when it finds an error or when manually terminated (`Ctrl+C`).
## Debugging Fuzzer Failures
### Crash 1: Excessive Memory Allocation
#### Failing Input:
```
go test fuzz v1
[]byte("300000000000")
```
### Issue:
- The function tries to allocate a slice of **300 billion strings**, which exhausts memory.
### Fix:
Limit the number of expected text elements.
```go
if count > 1000 {
    return nil, errors.New("too many")
}
```
---
### Crash 2: Negative Slice Capacity
#### Failing Input:
```
go test fuzz v1
[]byte("-1")
```
### Issue:
- `make([]string, 0, count)` **panics** when `count` is negative.
### Fix:
```go
if count < 0 {
    return nil, errors.New("no negative numbers")
}
```
### Crash 3: Blank Lines
#### Failing Input:
```
go test fuzz v1
[]byte("\r")
```
### Issue:
- Blank lines **cause unexpected behavior** (e.g., an empty count string).
### Fix:
```go
line = strings.TrimSpace(line)
if len(line) == 0 {
    return nil, errors.New("blank line")
}
```
## Final Fuzzing Results
```
go test -fuzz=FuzzParseData
```
- No more failures found.
- **Ctrl+C** to stop fuzzing after a few minutes.
## Key Takeaways
- **Fuzzing discovers real-world issues that unit tests miss.**
- **Bugs found include memory exhaustion, panics, and unexpected input handling.**
- **Fuzzer automatically saves failing test cases for regression testing.**
- **Always validate inputs before allocating memory or processing data.**
## Using Benchmarks in Go
### Importance of Benchmarking
Benchmarking is crucial in determining how fast (or slow) code runs, which is often difficult to assess by intuition alone. Go's built-in benchmarking tools in the testing framework make this easier, allowing you to evaluate and optimize your code more effectively.
### Example Function to Benchmark
Consider the `FileLen` function, which counts the number of characters in a file using a specific buffer size:
```go
func FileLen(f string, bufsize int) (int, error) {
    file, err := os.Open(f)
    if err != nil {
        return 0, err
    }
    defer file.Close()
    count := 0
    for {
        buf := make([]byte, bufsize)
        num, err := file.Read(buf)
        count += num
        if err != nil {
            break
        }
    }
    return count, nil
}
```
### Test for `FileLen`
Before benchmarking, it’s essential to ensure the function works correctly:
```go
func TestFileLen(t *testing.T) {
    result, err := FileLen("testdata/data.txt", 1)
    if err != nil {
        t.Fatal(err)
    }
    if result != 65204 {
        t.Error("Expected 65204, got", result)
    }
}
```
### Benchmarking the Function
To benchmark the `FileLen` function, Go uses special benchmark functions. These start with `Benchmark` and accept a `*testing.B` parameter, which provides the necessary benchmarking functionality.
### Basic Benchmark with 1-byte Buffer
```go
var blackhole int
func BenchmarkFileLen1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        result, err := FileLen("testdata/data.txt", 1)
        if err != nil {
            b.Fatal(err)
        }
        blackhole = result
    }
}
```
- **Blackhole**: The `blackhole` variable ensures the result of the benchmark is used so the compiler doesn't optimize the function call away.
- **Benchmark Loop**: The benchmark function must run in a loop with `b.N` iterations. Go's testing framework increases `b.N` over multiple iterations to ensure accurate timing results.
### Running the Benchmark
To run the benchmark, use the `-bench` flag:
```
go test -bench=. -benchmem
```
This command runs all benchmarks and includes memory allocation data.
### Sample Benchmark Output
```
BenchmarkFileLen1-12   25  47201025 ns/op  65342 B/op  65208 allocs/op
```
- **BenchmarkFileLen1-12**: The benchmark name with the number of CPUs used (GOMAXPROCS).
- **25**: The number of times the test was run to produce stable results.
- **47201025 ns/op**: The time taken to run one iteration of the benchmark in nanoseconds.
- **65342 B/op**: The number of bytes allocated during a single pass.
- **65208 allocs/op**: The number of heap allocations made during a single benchmark pass.
### Benchmarking with Different Buffer Sizes
To evaluate performance for various buffer sizes, you can create a benchmark function that runs the `FileLen` function with different buffer sizes:
```go
func BenchmarkFileLen(b *testing.B) {
    for _, v := range []int{1, 10, 100, 1000, 10000, 100000} {
        b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                result, err := FileLen("testdata/data.txt", v)
                if err != nil {
                    b.Fatal(err)
                }
                blackhole = result
            }
        })
    }
}
```
### Sample Benchmark Results
```
BenchmarkFileLen/FileLen-1-12          25  47828842 ns/op   65342 B/op  65208 allocs/op
BenchmarkFileLen/FileLen-10-12        230   5136839 ns/op  104488 B/op   6525 allocs/op
BenchmarkFileLen/FileLen-100-12      2246    509619 ns/op   73384 B/op    657 allocs/op
BenchmarkFileLen/FileLen-1000-12    16491     71281 ns/op   68744 B/op     70 allocs/op
BenchmarkFileLen/FileLen-10000-12   42468     26600 ns/op   82056 B/op     11 allocs/op
BenchmarkFileLen/FileLen-100000-12  36700     30473 ns/op  213128 B/op      5 allocs/op
```
These results show:
- As the buffer size increases, the number of allocations decreases, and the time per operation (`ns/op`) decreases as well.
- Larger buffers improve performance until the buffer exceeds the file size, causing unnecessary allocations.
### Optimizing the Function
In the original `FileLen` function, a new buffer is allocated every time the file is read. This can be optimized by allocating the buffer once before the loop:
```go
func FileLen(f string, bufsize int) (int, error) {
    file, err := os.Open(f)
    if err != nil {
        return 0, err
    }
    defer file.Close()
    buf := make([]byte, bufsize)
    count := 0
    for {
        num, err := file.Read(buf)
        count += num
        if err != nil {
            break
        }
    }
    return count, nil
}
```
### Optimized Benchmark Results
After the optimization, benchmarks show fewer allocations (only 4 per test run):
```
BenchmarkFileLen/FileLen-1-12          25  46167597 ns/op     137 B/op  4 allocs/op
BenchmarkFileLen/FileLen-10-12        261   4592019 ns/op     152 B/op  4 allocs/op
BenchmarkFileLen/FileLen-100-12      2518    478838 ns/op     248 B/op  4 allocs/op
BenchmarkFileLen/FileLen-1000-12    20059     60150 ns/op    1160 B/op  4 allocs/op
BenchmarkFileLen/FileLen-10000-12   62992     19000 ns/op   10376 B/op  4 allocs/op
BenchmarkFileLen/FileLen-100000-12  51928     21275 ns/op  106632 B/op  4 allocs/op
```
### Conclusion on Benchmarking
- **Benchmarking** helps identify the optimal buffer size for performance.
- **Optimizations** (like reusing buffers) can significantly reduce memory allocations and improve speed.
- Benchmarks allow you to measure trade-offs between **memory usage** and **performance**. For example, smaller buffers reduce memory usage but might slow down the program.
## Profiling Your Go Code
Once benchmarks reveal performance bottlenecks, **profiling** helps pinpoint the exact issue. Go provides profiling tools that collect CPU and memory usage data from running programs.
- Profiling data can be visualized with tools like `pprof`, helping identify where most CPU time or memory is spent.
- Remote profiling can be exposed through web endpoints, providing insight into live Go services.
For more on profiling Go programs, check out the blog post "Profiling Go Programs with pprof" by Julia Evans.
## Using Stubs in Go
### Introduction
In Go, writing unit tests for functions that depend on other code requires abstraction. There are two primary ways to abstract function calls:
- **Function types**
- **Interfaces**
These abstractions enable modular production code and simplify unit testing.
> Tip: Code that depends on abstractions is easier to test!
> 
### Example: Stubbing a Dependency
Consider the following example, where `Processor` depends on an interface called `MathSolver`:
### **Processor and MathSolver Interface**
```go
type Processor struct {
    Solver MathSolver
}
type MathSolver interface {
    Resolve(ctx context.Context, expression string) (float64, error)
}
```
The `Processor` struct has a method that reads an expression from an `io.Reader` and calculates its value using `MathSolver`:
```go
func (p Processor) ProcessExpression(ctx context.Context, r io.Reader) (float64, error) {
    curExpression, err := readToNewLine(r)
    if err != nil {
        return 0, err
    }
    if len(curExpression) == 0 {
        return 0, errors.New("no expression to read")
    }
    answer, err := p.Solver.Resolve(ctx, curExpression)
    return answer, err
}
```
### **Writing a Stub for Testing**
A **stub** provides predefined responses for method calls. Here’s a stub implementation of `MathSolver`:
```go
type MathSolverStub struct{}
func (ms MathSolverStub) Resolve(ctx context.Context, expr string) (float64, error) {
    switch expr {
    case "2 + 2 * 10":
        return 22, nil
    case "( 2 + 2 ) * 10":
        return 40, nil
    case "( 2 + 2 * 10":
        return 0, errors.New("invalid expression: ( 2 + 2 * 10")
    }
    return 0, nil
}
```
### **Unit Test Using the Stub**
A test case for `ProcessExpression` can be written as follows:
```go
func TestProcessorProcessExpression(t *testing.T) {
    p := Processor{MathSolverStub{}}
    in := strings.NewReader(`2 + 2 * 10
( 2 + 2 ) * 10
( 2 + 2 * 10`)
    data := []float64{22, 40, 0}
    hasErr := []bool{false, false, true}
    for i, d := range data {
        result, err := p.ProcessExpression(context.Background(), in)
        if err != nil && !hasErr[i] {
            t.Error(err)
        }
        if result != d {
            t.Errorf("Expected result %08_file_servers_cdn, got %08_file_servers_cdn", d, result)
        }
    }
}
```
This test ensures that:
- Valid expressions return correct results.
- Invalid expressions return errors.
## **Stubbing Large Interfaces**
Interfaces with multiple methods can be challenging to stub. Consider the `Entities` interface:
```go
type Entities interface {
    GetUser(id string) (User, error)
    GetPets(userID string) ([]Pet, error)
    GetChildren(userID string) ([]Person, error)
    GetFriends(userID string) ([]Person, error)
    SaveUser(user User) error
}
```
The `Logic` struct depends on `Entities`:
```go
type Logic struct {
    Entities Entities
}
```
A method in `Logic` fetches pet names for a user:
```go
func (l Logic) GetPetNames(userId string) ([]string, error) {
	pets, err := l.Entities.GetPets(userId)
	if err != nil {
		return nil, err
	}
	out := make([]string, len(pets))
	for i, p := range pets {
		out[i] = p.Name
	}
	return out, nil
}
```
### **Creating a Stub for a Single Method**
Instead of implementing all `Entities` methods, stub only `GetPets`:
```go
type GetPetNamesStub struct {
    Entities
}
func (ps GetPetNamesStub) GetPets(userID string) ([]Pet, error) {
    switch userID {
    case "1":
        return []Pet{{Name: "Bubbles"}}, nil
    case "2":
        return []Pet{{Name: "Stampy"}, {Name: "Snowball II"}}, nil
    default:
        return nil, fmt.Errorf("invalid id: %s", userID)
    }
}
```
### **Unit Test Using the Stub**
```go
func TestLogicGetPetNames(t *testing.T) {
    data := []struct {
        name     string
        userID   string
        petNames []string
    }{
        {"case1", "1", []string{"Bubbles"}},
        {"case2", "2", []string{"Stampy", "Snowball II"}},
        {"case3", "3", nil},
    }
    l := Logic{GetPetNamesStub{}}
    for _, d := range data {
        t.Run(d.name, func(t *testing.T) {
            petNames, err := l.GetPetNames(d.userID)
            if err != nil {
                t.Error(err)
            }
            if diff := cmp.Diff(d.petNames, petNames); diff != "" {
                t.Error(diff)
            }
        })
    }
}
```
> Warning: If a test calls an unimplemented method, it will panic. Ensure that all required methods have implementations.
> 
## **Using Function Fields for More Flexible Stubs**
A more flexible approach is using **function fields** instead of implementing the interface directly.
### **Defining a Stub with Function Fields**
```go
type EntitiesStub struct {
    getUser     func(id string) (User, error)
    getPets     func(userID string) ([]Pet, error)
    getChildren func(userID string) ([]Person, error)
    getFriends  func(userID string) ([]Person, error)
    saveUser    func(user User) error
}
```
Each method in `Entities` is implemented by calling the corresponding function field:
```go
func (es EntitiesStub) GetUser(id string) (User, error) {
    return es.getUser(id)
}
func (es EntitiesStub) GetPets(userID string) ([]Pet, error) {
    return es.getPets(userID)
}
```
### **Unit Test with Function Fields**
```go
func TestLogicGetPetNames(t *testing.T) {
    data := []struct {
        name     string
        getPets  func(userID string) ([]Pet, error)
        userID   string
        petNames []string
        errMsg   string
    }{
        {"case1", func(userID string) ([]Pet, error) {
            return []Pet{{Name: "Bubbles"}}, nil
        }, "1", []string{"Bubbles"}, ""},
        {"case2", func(userID string) ([]Pet, error) {
            return nil, errors.New("invalid id: 3")
        }, "3", nil, "invalid id: 3"},
    }
    l := Logic{}
    for _, d := range data {
        t.Run(d.name, func(t *testing.T) {
            l.Entities = EntitiesStub{getPets: d.getPets}
            petNames, err := l.GetPetNames(d.userID)
            if diff := cmp.Diff(d.petNames, petNames); diff != "" {
                t.Error(diff)
            }
            var errMsg string
            if err != nil {
                errMsg = err.Error()
            }
            if errMsg != d.errMsg {
                t.Errorf("Expected error `%s`, got `%s`", d.errMsg, errMsg)
            }
        })
    }
}
```
This approach makes it easy to define different behaviors for different tests.
## **Mocks vs. Stubs**
- **Stubs**: Return predefined values based on input.
- **Mocks**: Validate that calls happen in the expected order with expected inputs.
To generate mocks, you can use:
- [**gomock**](https://github.com/golang/mock) (by Google)
- [**testify/mock**](https://github.com/stretchr/testify) (by Stretchr)
## **Testing HTTP Clients in Go with `httptest`**
### **Introduction**
Writing tests for functions that make HTTP requests can be challenging. Traditionally, integration tests require running a real instance of the service. However, Go's `net/http/httptest` package provides a way to stub HTTP services, allowing unit testing without needing an actual server.
## **Using `httptest` to Stub an HTTP Service**
### **Implementing `MathSolver` as an HTTP Client**
Consider a `RemoteSolver` struct that calls an HTTP service to evaluate expressions:
```go
type RemoteSolver struct {
    MathServerURL string
    Client        *http.Client
}
func (rs RemoteSolver) Resolve(ctx context.Context, expression string) (float64, error) {
    req, err := http.NewRequestWithContext(ctx, http.MethodGet,
        rs.MathServerURL+"?expression="+url.QueryEscape(expression),
        nil)
    if err != nil {
        return 0, err
    }
    resp, err := rs.Client.Do(req)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()
    contents, err := io.ReadAll(resp.Body)
    if err != nil {
        return 0, err
    }
    if resp.StatusCode != http.StatusOK {
        return 0, errors.New(string(contents))
    }
    result, err := strconv.ParseFloat(string(contents), 64)
    if err != nil {
        return 0, err
    }
    return result, nil
}
```
This implementation sends a request to a math-solving HTTP service and parses the response.
## **Testing `RemoteSolver` with `httptest`**
Instead of requiring a real HTTP server, use `httptest.NewServer` to create a stubbed server.
### **Defining a Stub Server**
A `httptest.NewServer` instance will act as the HTTP service:
```go
server := httptest.NewServer(
    http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        expression := req.URL.Query().Get("expression")
        if expression != io.expression {
            rw.WriteHeader(http.StatusBadRequest)
            fmt.Fprintf(rw, "expected expression '%s', got '%s'", io.expression, expression)
            return
        }
        rw.WriteHeader(io.code)
        rw.Write([]byte(io.body))
    }))
defer server.Close()
```
This stub server:
- Reads the `expression` query parameter.
- Checks if it matches the expected input.
- Responds with a predefined HTTP status and body.
### **Creating an Instance of `RemoteSolver`**
```go
rs := RemoteSolver{
    MathServerURL: server.URL,
    Client:        server.Client(),
}
```
This instance communicates with the stubbed server, allowing tests to verify expected behavior.
## **Writing a Table-Driven Test**
### **Defining Test Cases**
Each test case specifies an input expression, expected HTTP response, and expected output:
```go
type info struct {
    expression string
    code       int
    body       string
}
var io info
data := []struct {
    name   string
    io     info
    result float64
    errMsg string
}{
    {"case1", info{"2 + 2 * 10", http.StatusOK, "22"}, 22, ""},
    {"case2", info{"( 2 + 2 ) * 10", http.StatusOK, "40"}, 40, ""},
    {"case3", info{"( 2 + 2 * 10", http.StatusBadRequest, "invalid expression"}, 0, "invalid expression"},
}
```
### **Executing Test Cases**
```go
for _, d := range data {
    t.Run(d.name, func(t *testing.T) {
        io = d.io
        result, err := rs.Resolve(context.Background(), d.io.expression)
        if result != d.result {
            t.Errorf("expected `%08_file_servers_cdn`, got `%08_file_servers_cdn`", d.result, result)
        }
        var errMsg string
        if err != nil {
            errMsg = err.Error()
        }
        if errMsg != d.errMsg {
            t.Errorf("expected error `%s`, got `%s`", d.errMsg, errMsg)
        }
    })
}
```
### **Key Takeaways**
- `io` is a shared variable captured by both the test cases and the stub server.
- This pattern works in unit tests but should be avoided in production.
- The test verifies correct responses and error handling.
## **Integration Tests and Build Tags**
Even though `httptest` allows unit testing without a real server, integration tests are still necessary to validate real API interactions.
### **Running a Math Server Locally**
To test against a real server, use Docker:
```
docker pull jonbodner/math-server
docker run -p 8080:8080 jonbodner/math-server
```
Alternatively, build the server from its [GitHub repository](https://github.com/jonbodner/math-server).
### **Using Build Tags for Integration Tests**
Create a test file with the `//go:build integration` directive:
```go
//go:build integration
package solver
import (
    "context"
    "testing"
)
func TestRemoteSolver_ResolveIntegration(t *testing.T) {
    rs := RemoteSolver{
        MathServerURL: "http://localhost:8080",
        Client:        http.DefaultClient,
    }
    result, err := rs.Resolve(context.Background(), "2 + 2 * 10")
    if err != nil {
        t.Fatal(err)
    }
    if result != 22 {
        t.Errorf("expected 22, got %08_file_servers_cdn", result)
    }
}
```
### **Running Integration Tests**
Use the `-tags` flag to run integration tests:
```
go test -tags integration -v ./...
```
> Note: Some developers prefer environment variables over build tags.
> 
Example using an environment variable:
```go
func TestRemoteSolver_ResolveIntegration(t *testing.T) {
    if os.Getenv("RUN_INTEGRATION_TESTS") != "1" {
        t.Skip("Skipping integration test. Set RUN_INTEGRATION_TESTS=1 to enable.")
    }
}
```
Run the test:
```
RUN_INTEGRATION_TESTS=1 go test -v ./...
```
### **Comparison: Build Tags vs. Environment Variables**
| **Method** | **Pros** | **Cons** |
| --- | --- | --- |
| **Build Tags** | Groups tests at compile time | Hard to discover required tags |
| **Env Variables** | Easily configurable at runtime | Requires explicit check in each test function |
## **Using `short` for Skipping Slow Tests**
For another approach, use the `-short` flag:
```go
if testing.Short() {
    t.Skip("skipping test in short mode.")
}
```
Run only short tests:
```
go test -short -v ./...
```
### **Limitations of `short`**
- Only distinguishes between "short" and "all" tests.
- Cannot group tests based on external dependencies.
- Philosophically, `short` is meant for skipping slow tests, not indicating dependencies.
## **Finding Concurrency Problems with the Data Race Detector**
### **Introduction**
Go provides built-in concurrency support, but improper synchronization can still lead to **data races**, where multiple goroutines access the same variable concurrently without proper locking. Go includes a **race detector** that helps identify such issues.
## **Example of a Data Race**
Consider the following function, which launches multiple goroutines to increment a shared counter:
```go
func getCounter() int {
    var counter int
    var wg sync.WaitGroup
    wg.Add(5)
    for i := 0; i < 5; i++ {
        go func() {
            for i := 0; i < 1000; i++ {
                counter++
            }
            wg.Done()
        }()
    }
    wg.Wait()
    return counter
}
```
### **Explanation:**
- Five goroutines are launched.
- Each increments `counter` 1,000 times.
- The expected result is `5,000`, but due to concurrent writes, updates can be lost.
## **Detecting the Issue with a Test**
The following unit test checks whether `getCounter()` returns the expected value:
```go
func TestGetCounter(t *testing.T) {
    counter := getCounter()
    if counter != 5000 {
        t.Error("unexpected counter:", counter)
    }
}
```
Running `go test` multiple times may produce inconsistent results, such as:
```
unexpected counter: 3673
```
This confirms a data race is occurring.
## **Using the Data Race Detector**
### **Running Tests with `race`**
Enable the race detector when running tests:
```
go test -race
```
This produces output similar to:
```
==================
WARNING: DATA RACE
Read at 0x00c000128070 by goroutine 10:
  test_examples/race.getCounter.func1()
      test_examples/race/race.go:12 +0x45
Previous write at 0x00c000128070 by goroutine 8:
  test_examples/race.getCounter.func1()
      test_examples/race/race.go:12 +0x5b
==================
```
### **Key Takeaways**
- The race detector pinpoints `counter++` as the issue.
- It identifies the exact **memory address**, **goroutines involved**, and **source file/line numbers.**
## **Reflect**
### Reflection Lets You Work with Types at Runtime
Go is a statically typed language, meaning types are defined at compile time. However, there are cases where runtime type inspection is necessary. Reflection allows Go programs to examine and manipulate types, values, and functions at runtime.
### When Is Reflection Needed?
Reflection is commonly used in the Go standard library for:
- **Database interactions**: The `database/sql` package maps records to structs using reflection.
- **Templating**: `text/template` and `html/template` use reflection to process values.
- **Formatting**: The `fmt` package relies on reflection for dynamic type handling.
- **Error handling**: The `errors` package uses reflection in `errors.Is` and `errors.As`.
- **Sorting**: The `sort` package supports generic slice sorting with reflection.
- **Serialization**: JSON, XML, and other encoders use reflection to read struct tags and process fields.
Reflection is useful when dealing with data that comes from outside the program, such as user input, network responses, or database records.
### Downsides of Reflection
Reflection has trade-offs:
- **Performance cost**: Reflection is significantly slower than direct type operations.
- **Complexity**: Reflective code is harder to understand and maintain.
- **Panics on misuse**: Incorrect usage of reflection often results in runtime panics.
## Types, Kinds, and Values
Now that you understand what reflection is and when it’s useful, let’s explore how it works. The `reflect` package in Go provides the types and functions that implement reflection. Reflection revolves around three core concepts: **types**, **kinds**, and **values**.
### Types and Kinds
A **type** defines the properties of a variable—what it can hold and how you can interact with it. Reflection allows you to inspect a type dynamically.
You can obtain the reflection representation of a variable's type using the `reflect.TypeOf` function:
```go
vType := reflect.TypeOf(v)
```
The `reflect.TypeOf` function returns a value of type `reflect.Type`, which provides methods for obtaining information about a variable’s type.
### Getting the Name of a Type
The `Name()` method returns the name of a type:
```go
var x int
xt := reflect.TypeOf(x)
fmt.Println(xt.Name()) // "int"
type Foo struct{}
f := Foo{}
ft := reflect.TypeOf(f)
fmt.Println(ft.Name()) // "Foo"
xpt := reflect.TypeOf(&x)
fmt.Println(xpt.Name()) // ""
```
- **For primitive types** (e.g., `int`), `Name()` returns the type name (`"int"`).
- **For structs**, it returns the struct name.
- **For unnamed types** (e.g., pointers, slices), `Name()` returns an empty string.
### Getting the Kind of a Type
The `Kind()` method returns a value of type `reflect.Kind`, which categorizes a type into broad groups:
```go
fmt.Println(reflect.TypeOf(42).Kind())   // reflect.Int
fmt.Println(reflect.TypeOf("hello").Kind()) // reflect.String
fmt.Println(reflect.TypeOf([]int{}).Kind()) // reflect.Slice
fmt.Println(reflect.TypeOf(map[string]int{}).Kind()) // reflect.Map
fmt.Println(reflect.TypeOf(&x).Kind()) // reflect.Pointer
```
- **Type** represents a named type (e.g., `Foo`, `int`, `string`).
- **Kind** categorizes the type into broader categories (e.g., `reflect.Struct`, `reflect.Int`, `reflect.Pointer`).
### Difference Between Type and Kind
If you define a struct named `Foo`:
- The **type** is `"Foo"`.
- The **kind** is `reflect.Struct`.
### Using Kind to Avoid Panics
Many `reflect.Type` methods only apply to specific kinds. For example:
- `NumIn()` returns the number of input parameters for a function type.
- Calling `NumIn()` on a non-function type **will panic**.
**Warning:** Always check the kind before calling methods that may not apply.
### Inspecting Pointer Types with `Elem()`
The `Elem()` method retrieves the underlying type for pointers, slices, maps, channels, and arrays:
```go
var x int
xpt := reflect.TypeOf(&x)
fmt.Println(xpt.Name())        // ""
fmt.Println(xpt.Kind())        // reflect.Pointer
fmt.Println(xpt.Elem().Name()) // "int"
fmt.Println(xpt.Elem().Kind()) // reflect.Int
```
- `xpt.Kind()` returns `reflect.Pointer`, indicating a pointer type.
- `xpt.Elem()` returns the `reflect.Type` of the pointed-to value (`int`).
### Inspecting Struct Fields
The `NumField()` and `Field()` methods allow you to inspect struct fields:
```go
type Foo struct {
    A int    `myTag:"value"`
    B string `myTag:"value2"`
}
var f Foo
ft := reflect.TypeOf(f)
for i := 0; i < ft.NumField(); i++ {
    curField := ft.Field(i)
    fmt.Println(curField.Name, curField.Type.Name(), curField.Tag.Get("myTag"))
}
```
**Output:**
```
A int value
B string value2
```
- `NumField()` returns the number of struct fields.
- `Field(i)` retrieves field information, including:
    - **Name**
    - **Type**
    - **Struct tag (`Tag.Get()`)** for metadata.
## Values
Reflection can also inspect and manipulate **values** at runtime.
### Getting a Reflective Value
Use `reflect.ValueOf()` to create a `reflect.Value` instance:
```go
vValue := reflect.ValueOf(v)
```
Every `reflect.Value` has a `Type()` method:
```go
vType := vValue.Type()
```
And a `Kind()` method:
```go
vKind := vValue.Kind()
```
### Reading Values from `reflect.Value`
To extract the original value from `reflect.Value`, use `Interface()`:
```go
s := []string{"a", "b", "c"}
sv := reflect.ValueOf(s)
s2 := sv.Interface().([]string) // Convert back to []string
```
### Specialized Methods for Primitive Types
For primitive types, `reflect.Value` provides specific getters:
```go
vInt := reflect.ValueOf(42)
fmt.Println(vInt.Int()) // 42
vStr := reflect.ValueOf("hello")
fmt.Println(vStr.String()) // "hello"
```
Using an incorrect method will panic:
```go
fmt.Println(vInt.String()) // Panics! `vInt` is an int, not a string
```
### Checking Convertibility
Use `CanInt()`, `CanFloat()`, `CanConvert()`, etc., to verify a value's compatibility before calling getters:
```go
if vInt.CanInt() {
    fmt.Println(vInt.Int()) // Safe
}
```
## Modifying Values with Reflection
You can modify values using reflection, but only if you pass a **pointer**.
### Steps to Modify a Value
1. Pass a **pointer** to `reflect.ValueOf()`:
    
    ```go
    i := 10
    iv := reflect.ValueOf(&i)
    ```
    
2. Use `Elem()` to access the actual value:
    
    ```go
    ivv := iv.Elem()
    ```
    
3. Modify the value:
    
    ```go
    ivv.SetInt(20)
    fmt.Println(i)
    ```
    
### Using `Set()`
For non-primitive types, use `Set()`:
```go
v1 := reflect.ValueOf("hello")
v2 := reflect.ValueOf("world")
v1.Set(v2) // Panics! v1 is not settable
```
### Why a Pointer Is Required
Reflection follows Go’s standard rules for modifying function parameters. Just like a regular function must receive a pointer to modify a value:
```go
func changeInt(i *int) {
    *i = 20
}
```
The reflection equivalent is:
```go
func changeIntReflect(i *int) {
    iv := reflect.ValueOf(i)
    iv.Elem().SetInt(20)
}
```
### Checking if a Value Is Settable
Calling `Set()` on an unmodifiable `reflect.Value` will **panic**. Use `CanSet()` to check first:
```go
iv := reflect.ValueOf(10)
if iv.CanSet() {
    iv.SetInt(20)
} else {
    fmt.Println("Cannot modify value")
}
```
Since `iv` was created from a **non-pointer**, `CanSet()` returns `false`.
## Make New Values
The `reflect.New` function is the reflection analogue of the `new` function. It takes in a `reflect.Type` and returns a `reflect.Value` that’s a **pointer** to a `reflect.Value` of the specified type. Since it’s a pointer, you can modify it and then assign the modified value to a variable by using the `Interface` method.
### Creating Channels, Maps, and Slices with Reflection
Just as `reflect.New` creates a pointer to a scalar type, you can also use reflection to do the same thing as the `make` keyword with the following functions:
```go
func MakeChan(typ Type, buffer int) Value
func MakeMap(typ Type) Value
func MakeMapWithSize(typ Type, n int) Value
func MakeSlice(typ Type, len, cap int) Value
```
Each of these functions takes in a `reflect.Type` that represents the compound type, **not** the contained type.
### Constructing `reflect.Type` Without a Value
You must always start from a value when constructing a `reflect.Type`. However, a trick lets you create a variable to represent a `reflect.Type` if you don’t have a value handy:
```go
var stringType = reflect.TypeOf((*string)(nil)).Elem()
var stringSliceType = reflect.TypeOf([]string(nil))
```
- The variable `stringType` contains a `reflect.Type` that represents a `string`.
- The variable `stringSliceType` contains a `reflect.Type` that represents a `[]string`.
The first line is tricky:
1. Convert `nil` to a pointer to a `string`: `(*string)(nil)`.
2. Use `reflect.TypeOf` to get the `reflect.Type` of that pointer.
3. Call `Elem` on the pointer’s `reflect.Type` to get the underlying type.
Parentheses around `*string` are necessary because of Go’s order of operations. Without them, the compiler interprets it as converting `nil` to `string`, which is illegal.
For `stringSliceType`, the process is simpler because `nil` is a valid value for a slice. You just convert `nil` to `[]string` and pass that to `reflect.TypeOf`.
### Using `reflect.New` and `reflect.MakeSlice`
Now that you have these types, you can see how to use `reflect.New` and `reflect.MakeSlice`:
```go
ssv := reflect.MakeSlice(stringSliceType, 0, 10) // Create an empty []string with capacity 10
sv := reflect.New(stringType).Elem() // Create a new string value
sv.SetString("hello")                // Set the string value
ssv = reflect.Append(ssv, sv)        // Append the new value to the slice
ss := ssv.Interface().([]string)     // Convert reflect.Value back to []string
fmt.Println(ss) // prints [hello]
```
## Use Reflection to Check If an Interface’s Value Is `nil`
If a `nil` variable of a concrete type is assigned to an interface variable, the **interface variable itself is not `nil`**. This is because a **type** is still associated with the interface variable.
To check whether the **value** associated with an interface is `nil`, you can use reflection with the `IsValid` and `IsNil` methods:
```go
func hasNoValue(i any) bool {
    iv := reflect.ValueOf(i)
    if !iv.IsValid() {
        return true
    }
    switch iv.Kind() {
    case reflect.Pointer, reflect.Slice, reflect.Map, reflect.Func,
         reflect.Interface:
        return iv.IsNil()
    default:
        return false
    }
}
```
### Explanation:
1. `IsValid` returns `true` if `reflect.Value` holds anything **other than** a nil interface.
    - This check is **essential** because calling any method on an **invalid** `reflect.Value` **panics**.
2. `IsNil` checks whether the `reflect.Value` **is nil**, but **only** if the kind is:
    - `reflect.Pointer`
    - `reflect.Slice`
    - `reflect.Map`
    - `reflect.Func`
    - `reflect.Interface`
    - If `IsNil` is called on any other type, it **panics**.
3. The function returns `true` if the interface contains **no value**, otherwise it returns `false`.
### Best Practices
Even though it is **possible** to detect an interface with a nil value, you should **strive to write your code so that it works correctly even when an interface’s value is nil**. Reserve this technique for situations where **you have no other options**.
## Use Reflection to Write a Data Marshaler
Reflection is what the Go standard library uses to implement marshaling and unmarshaling. However, while Go provides `csv.NewReader` and `csv.NewWriter` to read and write CSV files as `[][]string`, it does not include a built-in way to **map CSV data to struct fields**. This section covers how to implement a CSV marshaler and unmarshaler using reflection.
### Defining the Struct with Tags
Like other Go marshaling libraries (`json`, `xml`), you define **struct tags** to specify how CSV column names map to struct fields:
```go
type MyData struct {
    Name   string `csv:"name"`
    Age    int    `csv:"age"`
    HasPet bool   `csv:"has_pet"`
}
```
The marshaler and unmarshaler functions will use these tags to **map CSV headers to struct fields**.
## Marshaling: Converting Structs to CSV
### Public API for Marshaling
```go
// Marshal converts a slice of structs into a CSV-compatible [][]string format.
// The first row contains the headers based on struct tags.
func Marshal(v any) ([][]string, error)
```
This function takes a slice of structs, extracts field values using reflection, and converts them into a `[][]string` format, suitable for writing to a CSV file.
### Implementation of `Marshal`
```go
func Marshal(v any) ([][]string, error) {
    sliceVal := reflect.ValueOf(v)
    if sliceVal.Kind() != reflect.Slice {
        return nil, errors.New("must be a slice of structs")
    }
    structType := sliceVal.Type().Elem()
    if structType.Kind() != reflect.Struct {
        return nil, errors.New("must be a slice of structs")
    }
    var out [][]string
    header := marshalHeader(structType) // Get CSV headers from struct tags
    out = append(out, header)
    for i := 0; i < sliceVal.Len(); i++ {
        row, err := marshalOne(sliceVal.Index(i))
        if err != nil {
            return nil, err
        }
        out = append(out, row)
    }
    return out, nil
}
```
### Explanation:
1. Ensure the input is a **slice of structs** (not a pointer or other type).
2. Extract the **struct type** from the slice.
3. Generate **CSV headers** from the struct field tags.
4. Iterate over the slice:
    - Extract each struct.
    - Convert each struct's fields into a `[]string` using `marshalOne`.
    - Append the result to the output.
### Helper Function: `marshalHeader`
Extracts the CSV headers from struct field tags:
```go
func marshalHeader(vt reflect.Type) []string {
    var row []string
    for i := 0; i < vt.NumField(); i++ {
        field := vt.Field(i)
        if curTag, ok := field.Tag.Lookup("csv"); ok {
            row = append(row, curTag)
        }
    }
    return row
}
```
### Explanation:
1. Iterate over struct fields.
2. Read the `csv` tag.
3. Append field names to a string slice.
4. Return the headers.
### Helper Function: `marshalOne`
Converts a struct instance to a CSV row:
```go
func marshalOne(vv reflect.Value) ([]string, error) {
    var row []string
    vt := vv.Type()
    for i := 0; i < vv.NumField(); i++ {
        fieldVal := vv.Field(i)
        if _, ok := vt.Field(i).Tag.Lookup("csv"); !ok {
            continue
        }
        switch fieldVal.Kind() {
        case reflect.Int:
            row = append(row, strconv.FormatInt(fieldVal.Int(), 10))
        case reflect.String:
            row = append(row, fieldVal.String())
        case reflect.Bool:
            row = append(row, strconv.FormatBool(fieldVal.Bool()))
        default:
            return nil, fmt.Errorf("cannot handle field of kind %v", fieldVal.Kind())
        }
    }
    return row, nil
}
```
### Explanation:
1. Iterate over struct fields.
2. Check if the field has a `csv` tag.
3. Convert each field to a **string** based on its type.
4. Return the row as a `[]string`.
## Unmarshaling: Converting CSV to Structs
### Public API for Unmarshaling
```go
// Unmarshal converts CSV data ([][]string) into a slice of structs.
// The first row is assumed to be the CSV header.
func Unmarshal(data [][]string, v any) error
```
### Implementation of `Unmarshal`
```go
func Unmarshal(data [][]string, v any) error {
    sliceValPointer := reflect.ValueOf(v)
    if sliceValPointer.Kind() != reflect.Pointer {
        return errors.New("must be a pointer to a slice of structs")
    }
    sliceVal := sliceValPointer.Elem()
    if sliceVal.Kind() != reflect.Slice {
        return errors.New("must be a pointer to a slice of structs")
    }
    structType := sliceVal.Type().Elem()
    if structType.Kind() != reflect.Struct {
        return errors.New("must be a pointer to a slice of structs")
    }
    // Map CSV headers to struct field positions
    header := data[0]
    namePos := make(map[string]int, len(header))
    for i, name := range header {
        namePos[name] = i
    }
    // Convert CSV rows into struct instances
    for _, row := range data[1:] {
        newVal := reflect.New(structType).Elem()
        err := unmarshalOne(row, namePos, newVal)
        if err != nil {
            return err
        }
        sliceVal.Set(reflect.Append(sliceVal, newVal))
    }
    return nil
}
```
### Explanation:
1. Validate that `v` is a **pointer to a slice of structs**.
2. Extract the **slice element type** (struct type).
3. Assume the **first row is the CSV header** and map field names to indices.
4. Iterate over the remaining rows:
    - Create a **new struct instance**.
    - Populate it using `unmarshalOne`.
    - Append it to the slice.
### Helper Function: `unmarshalOne`
```go
func unmarshalOne(row []string, namePos map[string]int, vv reflect.Value) error {
    vt := vv.Type()
    for i := 0; i < vv.NumField(); i++ {
        typeField := vt.Field(i)
        pos, ok := namePos[typeField.Tag.Get("csv")]
        if !ok {
            continue
        }
        val := row[pos]
        field := vv.Field(i)
        switch field.Kind() {
        case reflect.Int:
            i, err := strconv.ParseInt(val, 10, 64)
            if err != nil {
                return err
            }
            field.SetInt(i)
        case reflect.String:
            field.SetString(val)
        case reflect.Bool:
            b, err := strconv.ParseBool(val)
            if err != nil {
                return err
            }
            field.SetBool(b)
        default:
            return fmt.Errorf("cannot handle field of kind %v", field.Kind())
        }
    }
    return nil
}
```
### Explanation:
1. Iterate over struct fields.
2. Use the `csv` tag to **map field names to CSV column indices**.
3. Convert CSV values from `string` to the correct type.
4. Set field values accordingly.
## Using the CSV Marshaler & Unmarshaler
### Example CSV Data
```go
data := `name,age,has_pet
Jon,100,true
"Fred 'The Hammer' Smith",42,false
Martha,37,true
`
```
### Unmarshaling CSV Data
```go
r := csv.NewReader(strings.NewReader(data))
allData, err := r.ReadAll()
if err != nil {
    panic(err)
}
var entries []MyData
Unmarshal(allData, &entries)
fmt.Println(entries) // Output: [{Jon 100 true} {Fred 'The Hammer' Smith 42 false} {Martha 37 true}]
```
### Marshaling Back to CSV
```go
out, err := Marshal(entries)
if err != nil {
    panic(err)
}
sb := &strings.Builder{}
w := csv.NewWriter(sb)
w.WriteAll(out)
fmt.Println(sb.String()) // Outputs the reconstructed CSV
```
## Summary
- **Marshaling**: Convert a slice of structs to `[][]string` using struct tags.
- **Unmarshaling**: Convert `[][]string` into a slice of structs.
- **Reflection** allows this to work dynamically for any struct.
This implementation extends Go's CSV capabilities by **automatically mapping CSV data to struct fields**.
## Build Functions with Reflection to Automate Repetitive Tasks
Go allows you to create functions dynamically using reflection. This technique can help wrap existing functions with additional behavior without writing repetitive code.
### Creating a Timed Function Wrapper
The following factory function adds timing to any function that’s passed into it:
```go
func MakeTimedFunction(f any) any {
    ft := reflect.TypeOf(f)
    fv := reflect.ValueOf(f)
    wrapperF := reflect.MakeFunc(ft, func(in []reflect.Value) []reflect.Value {
        start := time.Now()
        out := fv.Call(in)
        end := time.Now()
        fmt.Println(end.Sub(start))
        return out
    })
    return wrapperF.Interface()
}
```
- This function takes in any function (`any` type).
- It uses `reflect.TypeOf` and `reflect.ValueOf` to inspect and call the function.
- `reflect.MakeFunc` creates a wrapper that:
    - Captures the start time.
    - Calls the original function.
    - Captures the end time.
    - Prints the execution duration.
    - Returns the original function's output.
### Usage Example
```go
func timeMe(a int) int {
    time.Sleep(time.Duration(a) * time.Second)
    return a * 2
}
func main() {
    timed := MakeTimedFunction(timeMe).(func(int) int)
    fmt.Println(timed(2))
}
```
This technique is useful but should be used carefully to maintain clarity in your code.
## You Can Build Structs with Reflection, but Don’t
Go’s `reflect.StructOf` allows you to create struct types dynamically:
- `reflect.StructOf([]reflect.StructField)` returns a `reflect.Type` for a new struct.
- These structs can only be assigned to `any` variables.
- Their fields must be accessed via reflection.
This feature is mostly of academic interest. An example of its use is in a memoization function that dynamically generates struct keys.
## Reflection Can’t Make Methods
While reflection allows you to create functions and struct types dynamically, **it cannot be used to add methods to a type**. This means:
- You **cannot** use reflection to create a type that implements an interface.
- Methods must be explicitly defined in Go.
## Use Reflection Only if It’s Worthwhile
Reflection is necessary for some cases, like handling unknown data types at runtime, but it has drawbacks:
- **Performance impact**: Reflection is significantly slower than generics or direct implementations.
- **Type safety issues**: The compiler cannot enforce type correctness when using reflection.
### Implementing a Filter Function with Reflection
```go
func Filter(slice any, filter any) any {
    sv := reflect.ValueOf(slice)
    fv := reflect.ValueOf(filter)
    sliceLen := sv.Len()
    out := reflect.MakeSlice(sv.Type(), 0, sliceLen)
    for i := 0; i < sliceLen; i++ {
        curVal := sv.Index(i)
        values := fv.Call([]reflect.Value{curVal})
        if values[0].Bool() {
            out = reflect.Append(out, curVal)
        }
    }
    return out.Interface()
}
```
### Usage Example
```go
names := []string{"Andrew", "Bob", "Clara", "Hortense"}
longNames := Filter(names, func(s string) bool {
    return len(s) > 3
}).([]string)
fmt.Println(longNames)
ages := []int{20, 50, 13}
adults := Filter(ages, func(age int) bool {
    return age >= 18
}).([]int)
fmt.Println(adults)
```
**Output:**
```
[Andrew Clara Hortense]
[20 50]
```
### Performance Comparison
Benchmarking a reflection-based filter vs. generic and direct implementations:
| Function | Operations | Time per op (ns) | Memory (B) | Allocations |
| --- | --- | --- | --- | --- |
| Filter (Reflection, String) | 5,870 | 203,962 ns | 46,616 B | 2,219 |
| Filter (Generic, String) | 294,355 | 3,920 ns | 16,384 B | 1 |
| Filter (Direct, String) | 302,636 | 3,885 ns | 16,384 B | 1 |
| Filter (Reflection, Int) | 5,756 | 204,530 ns | 45,240 B | 2,503 |
| Filter (Generic, Int) | 439,100 | 2,698 ns | 8,192 B | 1 |
| Filter (Direct, Int) | 443,745 | 2,677 ns | 8,192 B | 1 |
### Key Takeaways
- **Reflection is ~50–75x slower** than generics or direct implementations.
- **Memory usage and allocations increase significantly**, adding strain to the garbage collector.
- **Generics offer type safety and performance without code duplication**.
### When to Use Reflection
Reflection should be reserved for cases where generics are not feasible, such as:
- **Working with unknown data structures** (e.g., CSV marshaling/unmarshaling).
- **Memoization functions that require dynamic struct generation**.
Before using reflection, ensure that **its benefits outweigh the performance and maintainability costs**.
## Unsafe
### unsafe Is Unsafe
The `unsafe` package in Go allows you to manipulate memory directly. It is a small but powerful package, which can perform operations that aren’t available through standard Go types and functions.
### Why Does `unsafe` Exist?
Given Go's focus on memory safety, you might wonder why `unsafe` is necessary. Just as `reflect` is used for translating text data, `unsafe` helps with translating binary data. A 2020 paper by Diego Elias Costa et al. called “Breaking Type-Safety in Go: An Empirical Study on the Usage of the unsafe Package” revealed:
- 24% of Go projects use `unsafe` at least once.
- Most `unsafe` usages were for integration with operating systems and C code.
- Developers often use `unsafe` for more efficient Go code.
The Go standard library uses `unsafe` to interact with the operating system, as seen in the `syscall` package and higher-level `sys` package.
### unsafe.Pointer
The `unsafe.Pointer` type allows for conversions between pointers of any type and `unsafe.Pointer`. It can also be converted to and from `uintptr`, a special integer type. This is particularly useful for:
- Performing pointer arithmetic.
- Manipulating the bytes in a variable by converting it to `unsafe.Pointer` and then back to the type’s pointer.
These operations are similar to those you might perform in C or C++ but with Go’s stricter safety measures in place.
### Common Patterns in `unsafe`
There are two main patterns in `unsafe` code:
1. **Type Conversion:** Converting between types that are usually not directly convertible, using `unsafe.Pointer` as an intermediary.
2. **Byte Manipulation:** Converting a variable to an `unsafe.Pointer`, and then modifying the underlying bytes. This requires knowing the size of the data you're working with.
### Using `Sizeof` and `Offsetof`
The `unsafe` package provides functions like `Sizeof` and `Offsetof` to understand how data is laid out in memory.
### Sizeof
The `Sizeof` function returns the size of a value in bytes. Some important points:
- For numeric types, the size is straightforward (e.g., `int16` is 2 bytes, `byte` is 1 byte).
- For pointers, it returns the size of the pointer itself, not the data the pointer points to.
- A slice is 24 bytes on a 64-bit system (2 `int` fields for lenghth and capacity + a pointer to data).
- A string is 16 bytes (an `int` for the length and a pointer to the data).
- Arrays’ sizes are calculated by multiplying the number of elements by the size of the type.
- Structs can be larger than the sum of their fields due to padding for alignment.
### Offsetof
The `Offsetof` function tells you the byte offset of a field within a struct. The layout of struct fields can affect their size due to padding added for alignment.
### Example: Field Order in Structs
Consider the following two structs:
```go
type BoolInt struct {
    b bool
    i int64
}
type IntBool struct {
    i int64
    b bool
}
```
Running this code:
```go
fmt.Println(unsafe.Sizeof(BoolInt{}), unsafe.Offsetof(BoolInt{}.b), unsafe.Offsetof(BoolInt{}.i))
fmt.Println(unsafe.Sizeof(IntBool{}), unsafe.Offsetof(IntBool{}.i), unsafe.Offsetof(IntBool{}.b))
```
Produces the following output:
```
16 0 8
16 0 8
```
Both structs are 16 bytes in size. The field order affects padding, as seen in this example:
- In `BoolInt`, 7 bytes of padding are added between `b` and `i` to ensure proper alignment.
- In `IntBool`, padding is added after `b` to maintain alignment.
### More Complex Example
With three fields:
```go
type BoolIntBool struct {
    b  bool
    i  int64
    b2 bool
}
type BoolBoolInt struct {
    b  bool
    b2 bool
    i  int64
}
type IntBoolBool struct {
    i  int64
    b  bool
    b2 bool
}
```
The output for these structs is:
```
24 0 8 16
16 0 1 8
16 0 8 9
```
Here, reordering the fields minimizes the amount of padding:
- `BoolIntBool` is 24 bytes long.
- `BoolBoolInt` and `IntBoolBool` are 16 bytes long, as their bool fields are grouped together, minimizing padding.
### Use Cases for `unsafe`
Although primarily academic, `unsafe` can be valuable in specific situations:
1. **Memory Optimization:** For programs managing large data sets, reordering struct fields to minimize padding can save memory.
2. **Mapping Bytes to Structs:** If you need to map a sequence of bytes directly into a struct (e.g., reading binary data), `unsafe` is necessary.
## Using Unsafe to Convert External Binary Data
As mentioned earlier, one of the main reasons people use the `unsafe` package is for performance, especially when reading data from a network. If you want to map the data into or out of a Go data structure, `unsafe.Pointer` gives you a very fast way to do so. Here's an example based on a wire protocol.
### Example Protocol
Imagine a wire protocol with the following structure:
- **Value**: 4 bytes, representing an unsigned, big-endian 32-bit integer.
- **Label**: 10 bytes, ASCII name for the value.
- **Active**: 1 byte, boolean flag to indicate whether the field is active.
- **Padding**: 1 byte to ensure everything fits into 16 bytes.
**Note**: Data sent over a network is usually in big-endian format (most significant byte first). Since most CPUs are little-endian (or bi-endian), you need to be careful when reading or writing data to the network.
### Define the Data Structure
You define a data structure in Go that matches the wire protocol:
```go
type Data struct {
    Value  uint32   // 4 bytes
    Label  [10]byte // 10 bytes
    Active bool     // 1 byte
    // Go pads this with 1 byte to make it align
}
```
### Calculate Size Using `unsafe.Sizeof`
You can use `unsafe.Sizeof` to define a constant that represents the size of the structure:
```go
const dataSize = unsafe.Sizeof(Data{}) // sets dataSize to 16
```
### Mapping Binary Data to the Struct
When reading bytes from the network, you might receive the following data:
```go
[0 132 95 237 80 104 111 110 101 0 0 0 0 0 1 0]
```
You can read these bytes into an array and convert it into the struct as follows:
### Safe Go Code
```go
func DataFromBytes(b [dataSize]byte) Data {
    d := Data{}
    d.Value = binary.BigEndian.Uint32(b[:4])
    copy(d.Label[:], b[4:14])
    d.Active = b[14] != 0
    return d
}
```
### Unsafe Go Code
```go
func DataFromBytesUnsafe(b [dataSize]byte) Data {
    data := *(*Data)(unsafe.Pointer(&b))
    if isLE {
        data.Value = bits.ReverseBytes32(data.Value)
    }
    return data
}
```
The first line is a little confusing, but you can take it apart and understand what’s going on. First, you take the address of the byte array and convert it to an `unsafe.Pointer`. Then you convert the `unsafe.Pointer` to a `(*Data)` (you have to put `(*Data)` in parentheses because of Go’s order of operations). You want to return the struct, not a pointer to it, so you dereference the pointer. Next, you check your flag to see if you are on a little-endian platform. If so, you reverse the bytes in the `Value` field. Finally, you return the value.
### Little-Endian Detection
To detect if the platform is little-endian:
```go
var isLE bool
func init() {
    var x uint16 = 0xFF00
    xb := *(*[2]byte)(unsafe.Pointer(&x))
    isLE = (xb[0] == 0x00)
}
```
In this example, `x` is stored as `[00 FF]` on little-endian systems and `[FF 00]` on big-endian systems. By checking the first byte, you can determine whether you're on a little-endian platform.
### Writing Data Back to the Network
To write the data back, you could use the following functions:
### Safe Go Code
```go
func BytesFromData(d Data) [dataSize]byte {
    out := [dataSize]byte{}
    binary.BigEndian.PutUint32(out[:4], d.Value)
    copy(out[4:14], d.Label[:])
    if d.Active {
        out[14] = 1
    }
    return out
}
```
### Unsafe Go Code
```go
func BytesFromDataUnsafe(d Data) [dataSize]byte {
    if isLE {
        d.Value = bits.ReverseBytes32(d.Value)
    }
    b := *(*[dataSize]byte)(unsafe.Pointer(&d))
    return b
}
```
### Using `unsafe.Slice` for Slice Conversion
If you need to work with slices instead of arrays:
```go
func BytesFromDataUnsafeSlice(d Data) []byte {
    if isLE {
        d.Value = bits.ReverseBytes32(d.Value)
    }
    bs := unsafe.Slice((*byte)(unsafe.Pointer(&d)), dataSize)
    return bs
}
func DataFromBytesUnsafeSlice(b []byte) Data {
    data := *(*Data)((unsafe.Pointer)(unsafe.SliceData(b)))
    if isLE {
        data.Value = bits.ReverseBytes32(data.Value)
    }
    return data
}
```
### Performance Comparison
Benchmark results on an Apple Silicon M1 (little-endian) show the following timings:
| Benchmark | Iterations | ns/op | Allocations |
| --- | --- | --- | --- |
| `BenchmarkBytesFromData` | 548,607,271 | 2.185 | 0 B/op |
| `BenchmarkBytesFromDataUnsafe` | 1,000,000,000 | 0.8418 | 0 B/op |
| `BenchmarkBytesFromDataUnsafeSlice` | 91,179,056 | 13.14 | 16 B/op |
| `BenchmarkDataFromBytes` | 538,443,861 | 2.186 | 0 B/op |
| `BenchmarkDataFromBytesUnsafe` | 1,000,000,000 | 1.160 | 0 B/op |
| `BenchmarkDataFromBytesUnsafeSlice` | 1,000,000,000 | 0.9617 | 0 B/op |
### Key Observations
1. **Conversion to Slice is Slow**: Converting a struct to a slice is the slowest operation and is the only one that allocates memory. Allocating memory on the heap is typically slower than using memory on the stack.
2. **Efficiency**: Using `unsafe` for conversions between arrays and structs is roughly 2–2.5 times faster than the safe approach. If your program involves many such conversions, using `unsafe` can provide significant performance benefits. However, for most programs, sticking to safe code is recommended.
### Accessing Unexported Fields
You can also use `unsafe` to access unexported fields in structs by combining it with reflection. Here's an example:
```go
type HasUnexportedField struct {
    A int
    b bool
    C string
}
```
Normally, code outside this package cannot access the `b` field. However, by using reflection and `unsafe`, you can modify it:
```go
func SetBUnsafe(huf *one_package.HasUnexportedField) {
    sf, _ := reflect.TypeOf(huf).Elem().FieldByName('b')
    offset := sf.Offset
    start := unsafe.Pointer(huf)
    pos := unsafe.Add(start, offset)
    b := (*bool)(pos)
    fmt.Println(*b) // read the value
    *b = true       // write the value
}
```
### Using Unsafe Tools
Go provides a compiler flag (`-gcflags=-d=checkptr`) to help identify misuse of `unsafe.Pointer`. This is similar to the race checker and adds additional checks at runtime. However, it does slow down your program and isn’t guaranteed to find every issue, so it’s primarily useful during testing.
## Cgo Is for Integration, Not Performance
Just like reflection and unsafe, cgo is most useful at the border between Go programs and the outside world. Reflection helps integrate with external textual data, unsafe is best used with operating system and network data, and cgo is best for integrating with C libraries.
Despite being nearly 50 years old, C is still the lingua franca of programming languages. All the major operating systems are primarily written in either C or C++, which means that they are bundled with libraries written in C. It also means that nearly every programming language provides a way to integrate with C libraries. Go calls its foreign function interface (FFI) to C cgo.
As you have seen many times, Go is a language that favors explicit specification. Go developers sometimes deride automatic behaviors in other languages as “magic.” However, using cgo feels a bit like spending time with Merlin. Let’s take a look at this magical glue code.
### Simple C Integration Example
Start with a very simple program that calls C code to do some math.
```go
package main
import "fmt"
/*
    #cgo LDFLAGS: -lm
    #include <stdio.h>
    #include <math.h>
    #include "mylib.h"
    int add(int a, int b) {
        int sum = a + b;
        printf("a: %d, b: %d, sum %d\n", a, b, sum);
        return sum;
    }
*/
import "C"
func main() {
    sum := C.add(3, 2)
    fmt.Println(sum)
    fmt.Println(C.sqrt(100))
    fmt.Println(C.multiply(10, 20))
}
```
The `mylib.h` header is in the same directory as your `main.go`:
```c
int multiply(int a, int b);
```
There is also `mylib.c`:
```c
#include "mylib.h"
int multiply(int a, int b) {
    return a * b;
}
```
Assuming you have a C compiler installed on your computer, all you need to do is compile your program with `go build`:
```bash
$ go build
$ ./call_c_from_go
a: 3, b: 2, sum 5
5
10
200
```
### Explanation of Cgo Integration
The standard library doesn’t have a real package named `C`. Instead, `C` is an automatically generated package whose identifiers mostly come from the C code embedded in the comments that immediately precede it. In this example, you declare a C function called `add`, and cgo makes it available to your Go program as `C.add`. You can also invoke functions or global variables that are imported into the comment block from libraries via header files, as you can see when you call `C.sqrt` from `main` (imported from `math.h`) or `C.multiply` (imported from `mylib.h`).
In addition to the identifier names that appear in the comment block (or are imported into the comment block), the C pseudopackage also defines types like `C.int` and `C.char` to represent the built-in C types and functions, such as `C.CString` to convert a Go string to a C string.
### Calling Go Functions from C
You can use more magic to call Go functions from C functions. Go functions can be exposed to C code by putting an `//export` comment before the function. You can see this in use in the `sample_code/call_go_from_c` directory in the Chapter 16 repository. In `main.go`, you export the `doubler` function:
```go
//export doubler
func doubler(i int) int {
    return i * 2
}
```
When you export a Go function, you can no longer write C code directly in the comment before the import `C` statement. You can only list function headers:
```c
/*
    extern int add(int a, int b);
*/
import "C"
```
Put your C code into a `.c` file in the same directory as your Go code and include the magic header `_cgo_export.h`. You can see this in the `example.c` file:
```c
#include "_cgo_export.h"
int add(int a, int b) {
    int doubleA = doubler(a);
    int sum = doubleA + b;
    return sum;
}
```
Running this program with `go build` gives the following output:
```bash
$ go build
$ ./call_go_from_c
8
```
### Cgo and Memory Management Issues
One major stumbling block when using cgo is that Go is a garbage-collected language, and C is not. This makes it difficult to integrate nontrivial Go code with C. While you can pass a pointer into C code, you cannot directly pass something that contains a pointer. This is very limiting, as things like strings, slices, and functions are implemented with pointers and therefore cannot be contained in a struct passed into a C function.
Additionally, a C function cannot store a copy of a Go pointer that lasts after the function returns. If you break these rules, your program will compile and run, but it may crash or behave incorrectly at runtime when the memory pointed to by the pointer is garbage collected.
### Using `cgo.Handle` for Go Objects in C
If you need to pass an instance of a type that contains a pointer from Go to C and then back into Go, you can use a `cgo.Handle` to wrap the instance. Here’s a short example:
```go
package main
/*
    #include <stdint.h>
    extern void in_c(uintptr_t handle);
*/
import "C"
import (
    "fmt"
    "runtime/cgo"
)
type Person struct {
    Name string
    Age  int
}
func main() {
    p := Person{
        Name: "Jon",
        Age:  21,
    }
    C.in_c(C.uintptr_t(cgo.NewHandle(p)))
}
//export processor
func processor(handle C.uintptr_t) {
    h := cgo.Handle(handle)
    p := h.Value().(Person)
    fmt.Println(p.Name, p.Age)
    h.Delete()
}
```
And this is the C code:
```c
#include <stdint.h>
#include "_cgo_export.h"
void in_c(uintptr_t handle) {
    processor(handle);
}
```
In this Go code, you’re passing a `Person` instance into the C function `in_c`. This function in turn calls the Go function `processor`. You can’t safely pass a `Person` into C via cgo, because one of its fields is a string, and every string contains a pointer. To make this work, use the `cgo.NewHandle` function to convert `p` to a `cgo.Handle`. You then cast the Handle to a `C.uintptr_t` so you can pass it to the C function `in_c`.
### Cgo Limitations
Other cgo limitations include:
- You cannot use cgo to call a variadic C function (such as `printf`).
- Union types in C are converted into byte arrays.
- You cannot invoke a C function pointer, but you can assign it to a Go variable and pass it back into a C function.
### When to Use Cgo
In most cases, Go code is many times faster than Python or Ruby, so the need to rewrite algorithms in a lower-level language is greatly reduced. However, if you find yourself needing to use a C library and there is no suitable Go replacement, cgo can be a useful tool. For example, if you want to embed SQLite in a Go application, check out GitHub.
Since cgo isn’t fast, and it isn’t easy to use for nontrivial programs, the only reason to use cgo is if you must use a C library and there is no suitable Go replacement. Rather than writing cgo yourself, see if a third-party module already provides the wrapper.
If you find yourself needing to use an internal C library or third-party library that doesn’t have a wrapper, you can find additional details on how to write your integration in the Go documentation.
