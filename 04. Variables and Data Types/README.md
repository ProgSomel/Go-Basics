### ✅ What is a Variable?

**A variable in programming is like a container where you can store data. Each variable has:**

- A name (to identify it).

- A type (which determines what kind of data it holds).

- A value (the actual data).

**Syntax:**

```go

var variableName variableType = value

```

### 🎨 Declaring Variables in Go

1. Basic Declaration
   We declare a variable in Go using the var keyword. Here’s a simple example:

```go

var name string = "John"


```

- var: The keyword used to declare a variable.

- name: The name of the variable.

- string: The type of the variable (a string, in this case).

- "John": The value assigned to the variable.

**This means you created a variable name of type string, and its value is "John".**

```go

package main

import "fmt"

func main(){
	var a int = 10

	fmt.Println(a)

}

```

```go
package main

import "fmt"

func main(){
	var a = 10

	fmt.Println(a)

}
```

### 2. Type Inference (Implicit Declaration)

**Go can infer the type of a variable based on the value you assign to it. This way, you don't need to explicitly mention the type.**

Example:

```go

var age = 30

```

**Here, Go infers that age is an int (since 30 is an integer).**

```go
package main

import "fmt"

func main(){
	var a = 10

	fmt.Println(a)

}
```

### 3. Shorthand Declaration (:=)

**For shorter code, Go provides shorthand variable declaration using :=. This allows us to declare and initialize a variable in one go, inside functions.**

Example:

```go

name := "Alice"

```

**Here, Go automatically infers the type as string because "Alice" is a string.**

```go
package main

import "fmt"

func main(){
	a:= 10
	fmt.Println(a)
}
```

### 4. Multiple Variable Declaration

**You can declare multiple variables at once in Go:**

```go

var name, age = "John", 30

```

**This declares two variables, name (type string) and age (type int), and initializes them with values "John" and 30, respectively.**

### 🔍 Go Variable Types

**Go supports various data types. Let's explore each one.**

_1. Integer Types (Whole Numbers)_

- int: Can be either 32 or 64 bits, depending on the system architecture.

- int8, int16, int32, int64: Fixed-width integers for specific sizes.

- uint, uint8, uint16, uint32, uint64: Unsigned integers (only positive numbers).

Example:

```go
var age int = 25 //? Default int (32/64 bit based on your system)
var maxValue int64 = 9223372036854775807 //? 64 - bit integer
```

_2. Floating-Point Numbers (Decimal Numbers)_

- float32: A floating-point number with 32 bits of precision.

- float64: A floating-point number with 64 bits of precision (default).

Example:

```go
var weight float64 = 70.4 //? Floating-point number (default precision is float64)
```

_3. Strings (Text Data)_
string: Represents text or a sequence of characters.

Example:

```go
var name string = "John"
```

_5. Arrays and Slices_
Go supports both arrays and slices, which are used for holding collections of data.

**Array: Fixed-size collection of elements.**

Example:

```go
var scores [3]int = [3]int{90, 80, 70}
```

**Slice: A dynamically sized collection.**

```go
var scores = []int{90, 80, 70}
```

### 🧩 Zero Values in Go

_In Go, variables are automatically assigned a zero value if they’re declared but not initialized. This is very useful to avoid undefined behavior._

Here are some zero values in Go:

- int: 0

- float64: 0.0

- string: "" (empty string)

- bool: false

- Arrays/Slices/Maps: nil

Example:

```go

var number int       // zero value: 0
var name string      // zero value: ""
var isActive bool    // zero value: false

```

### 🧠 Constants vs Variables

_In Go, constants are like variables, but once set, they cannot change their value._

```go

const Pi = 3.14
```

**const: A keyword used to define a constant value that can’t be modified.**

### 🔥 Example Program Using Variables

````go

package main

import "fmt"

func main(){
    //? Declaring and initializing Variables
    var name string = "John"
    var age = 30 //? Go infers the type
    var weight float64 = 72.4
    isActive:= true //? Short declaration with inferred type

    //? Print Variables
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Weight:", weight)
    fmt.Println("Is Active:", isActive)

    //? Zero value demonstration
    var unintialized int
    fmt.Println("unintialized Variable: ", unintialized)

}

Output:

```bash

Name: John
Age: 30
Weight: 72.5
Is Active: true
Uninitialized Variable: 0

````

### 🧱 Go Data Types – In-Depth Guide

**In Go, data types define what kind of value a variable can hold — such as numbers, text, truth values, collections, etc.**

**Everything in Go has a type. Let’s break them down by category:**

### 📦 1. Basic Data Types

_🔢 Integer Types_
Used to store whole numbers (positive, negative, or zero).

| Typpe     | Size                            | Value Range                     |
| --------- | ------------------------------- | ------------------------------- |
| **int**   | 32 or 64-bit (system dependent) | Platform dependent              |
| **int8**  | 8 bits                          | -128 to 127                     |
| **int16** | 16 bits                         | -32,768 to 32,767               |
| **int32** | 32 bits                         | -2,147,483,648 to 2,147,483,647 |
| **int64** | 64 bits                         | Huge range...                   |

✅ Example:

```go
var  age int = 25
var smallNummber int8 = -100
```

_🧮 Unsigned Integers (No Negatives)_
Same as above, but only positive numbers (and zero).

| Typpe      | Size                            | Value Range        |
| ---------- | ------------------------------- | ------------------ |
| **uint**   | 32 or 64-bit (system dependent) | 0 to platform max  |
| **uint8**  | 8 bits                          | 0 to 127           |
| **uint16** | 16 bits                         | 0 to 32,767        |
| **uint32** | 32 bits                         | 0 to 2,147,483,647 |
| **uint64** | 64 bits                         | Huge range...      |

✅ Example:

```go
var score uint16 = 500;
```

_🔬 Floating Point Types_
Used to store decimal numbers.

| Typpe       | Size    | Precision  |
| ----------- | ------- | ---------- |
| **float32** | 32 bits | 6-7 digits |
| **uint64**  | 64 bits | 0 to 127   |

✅ Example:

```go
var weight float32 = 65.7
var pi float64 = 3.1415926535
```

**👉 Default is float64 if not specified.**

_🧵 String Type_
Used to store text — anything inside double quotes "...".

✅ Example:

```go
var name string = "GoLang"
```

**Strings are immutable (you can’t change individual characters directly).**

_🔘 Boolean Type_
Stores only two values: true or false.

✅ Example:

```go
var isOnline bool = true
```

**Used in conditions, comparisons, logic checks, etc.**

### 🧊 2. Composite Data Types

_🔗 Arrays_
A fixed-size collection of the same type.

✅ Example:

```go
var nums [3]int = [3]int{10, 20, 30}
```

**Size must be declared or inferred.**
**Cannot grow or shrink.**

_🧩 Slices_
A dynamic-size array (much more flexible).

✅ Example:

```go
numbers:= []int{10, 20, 30}
```

**Unlike arrays, slices can grow or shrink.**
**You’ll use slices way more than arrays.**

_📓 Maps_
A key-value pair data structure (like dictionaries in Python or objects in JS).

✅ Example:

```go
person:= map[string]string{
    "name": "Somel",
    "city": "Dhaka"
}
```

**Key and value types must be declared.**
**Very useful for lookup data!**

_🧱 Structs_
Custom data types that group multiple fields.

✅ Example:

```go
type Person struct {
    name string
    age int
}

var p1 = Person{name: "Somel", age: 24};
```

**Think of structs as your own blueprint for objects.**

### 🌀 3. Special Types

_☁️ interface{} (Empty Interface)_
Can hold any type — like a generic container.
Useful for generic functions or unknown data types.

✅ Example:

```go
var x interface{} = "Hello"
x = 42
```

**⚠️ You’ll need type assertion to use it properly.**

_⚠️ nil_
Used to represent “nothing” or “no value”.
Applies to pointers, maps, slices, channels, functions, interfaces, etc.

✅ Example:

```go
var m map[string]int
fmt.Println(m == nil) //true
```

### ✅ Final Notes

- 🧠 Go is statically typed → Every variable must have a known type.

- ⛔ You can’t assign a float to an int without type conversion.

- 🎯 Pick the simplest and most accurate data type for your values.
