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

*1. Integer Types (Whole Numbers)*
- int: Can be either 32 or 64 bits, depending on the system architecture.

- int8, int16, int32, int64: Fixed-width integers for specific sizes.

- uint, uint8, uint16, uint32, uint64: Unsigned integers (only positive numbers).

Example:
```go
var age int = 25 //? Default int (32/64 bit based on your system)
var maxValue int64 = 9223372036854775807 //? 64 - bit integer
```

*2. Floating-Point Numbers (Decimal Numbers)*
- float32: A floating-point number with 32 bits of precision.

- float64: A floating-point number with 64 bits of precision (default).
  
Example:
```go
var weight float64 = 70.4 //? Floating-point number (default precision is float64)
```

*3. Strings (Text Data)*
string: Represents text or a sequence of characters.

Example:
```go
var name string = "John"
```

*5. Arrays and Slices*
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
*In Go, variables are automatically assigned a zero value if they’re declared but not initialized. This is very useful to avoid undefined behavior.*

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
*In Go, constants are like variables, but once set, they cannot change their value.*

```go

const Pi = 3.14
```
**const: A keyword used to define a constant value that can’t be modified.**


### 🔥 Example Program Using Variables

```go

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

```




