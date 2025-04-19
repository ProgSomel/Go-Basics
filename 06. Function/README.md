### 🧠 What is a Function?

In Go, a function is a block of code designed to perform a specific task. It can take some input (called parameters) and return output (called return values). Functions are useful to encapsulate logic, making code more readable and maintainable.

**✅ Basic Structure of a Function in Go**
Here’s the basic syntax to define a function in Go:

```go
func functionName(parameter1 type, parameter2 type) returnType {
    //Function body
    // Some Logic

    return value //optional, if return type exists
}
```

- func: The keyword used to define a function.

- functionName: The name of the function (should follow camelCase).

- parameter1 type, parameter2 type: Function parameters and their types.

- returnType: The type of value the function will return (if any).

- return value: The value returned by the function.

### 🔍 Function Components Breakdown

_1. Function Declaration_
A function is declared using the func keyword.

Syntax:

```go

func functionName(parameters) returnType
```

Example:

```go

func greet(name string) {
    fmt.Println("Hello, " + name)
}
```

- greet is the function name.
- It accepts one parameter name of type string.
- It returns nothing (void).

_2. Calling a Function_
You call a function by using its name followed by the arguments (if any).

Syntax:

```go

functionName(arguments)

```

Example:

```go

greet("Alice")

```

This will output:

```bash
Hello, Alice
```

_3. Functions with Return Values_
A function can also return values. In Go, you use the return keyword to send a value back from the function.

Syntax:

```go

func functionName(parameters) returnType {
    // some logic
    return value
}

```

Example:

```go

func add(a int, b int) int {
    return a + b
}

```

- The function add accepts two integers and returns an integer.
- When called, it adds the two numbers together and returns the result.

Calling the function:

```go

sum := add(5, 3)
fmt.Println(sum)  // Output: 8

```

```go

package main

import "fmt"

func main() {
	a:= 10
	b:=20

	sum:=a+b
	fmt.Println("Sum of a and b is: ", sum)
}

//Sum of a and b is:  30
```

```go
package main

import "fmt"

func add(num1 int, num2 int){
	sum:= num1 + num2

	fmt.Println(sum)
}

func main() {
	a:= 10
	b:=20

	add(a, b)
}

//30
```

```go
package main

import "fmt"

func add(num1 int, num2 int) int {
	sum:=num1 + num2

	return sum
}

func main(){
	a:= 10
	b:=20

	fmt.Println(add(a, b))
}

//30
```

```go
package main

import "fmt"

func add(num1 int, num2 int) int {
	sum:=num1 + num2

	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum:=num1 + num2
	mul:= num1 * num2

	return sum, mul
}

func main(){
	a:= 10
	b:=20

	fmt.Println(getNumbers(a, b))
}

//30 200
```

```go
package main

import "fmt"

func add(num1 int, num2 int) int {
	sum:=num1 + num2

	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum:=num1 + num2
	mul:= num1 * num2

	return sum, mul
}

func main(){
	a:= 10
	b:=20

	p, q:=getNumbers(a, b)
	fmt.Println(p)
	fmt.Println(q)
}

//30
//200
```

```go
package main

import "fmt"

func add(num1 int, num2 int) int {
	sum:=num1 + num2

	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum:=num1 + num2
	mul:= num1 * num2

	return sum, mul
}

func main(){
	a:= 10
	b:=20

	p, q:=getNumbers(a, b)
	fmt.Println(p)
	fmt.Println(q)
}

//30
//200
```

```go
package main

import "fmt"

func add(num1 int, num2 int) int {
	sum:=num1 + num2

	return sum
}

func main(){
	a:= 10
	b:=20

	fmt.Println(add(a, b))
}

//30
```

```go
package main

import "fmt"

func add(num1 int, num2 int) int {
	sum:=num1 + num2

	return sum
}

func main(){
	a:= 10
	b:=20

	fmt.Println(add(a, b))
}

//30
```

```go
package main

import "fmt"

func add(num1 int, num2 int) int {
	sum:=num1 + num2

	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum:=num1 + num2
	mul:= num1 * num2

	return sum, mul
}

func main(){
	a:= 10
	b:=20

	fmt.Println(getNumbers(a, b))
}

//30 200
```

```go
package main

import "fmt"

func add(num1 int, num2 int) int {
	sum:=num1 + num2

	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum:=num1 + num2
	mul:= num1 * num2

	return sum, mul
}

func main(){
	a:= 10
	b:=20

	p, q:=getNumbers(a, b)
	fmt.Println(p)
	fmt.Println(q)
}

//30
//200
```

_5. Named Return Values_
Go allows you to name return values when defining a function. This makes your code more readable and can simplify error handling.

Syntax:

```go

func functionName(parameters) (namedReturnValue1, namedReturnValue2) {
    // some logic
    return value1, value2
}

```

Example:

```go
func divide(a int, b int) (result int, err error){
    if(b == 0){
        err = fmt.Error("Can not divide by zero");
        return
    }else{
        result = a / b
        return
    }
}
```

**The function divide returns two values: result (the quotient) and err (error, if any).**
Calling the function:

```go
result, err:= divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}
```

Output:

```bash
Result: 5
```

_⚙️ Variadic Functions (Function with Variable Arguments)_
A variadic function allows you to pass a variable number of arguments to it. This is very useful when you want a function to handle an unknown number of arguments.

Syntax:

```go

func functionName(parameters ...type) returnType

```

Example:

```go
func sum(nums ...int) int {
    total:= 0;
    for _, num := range nums {
        total+=num;
    }
    return total
}
```

**The function sum accepts a variable number of int arguments and returns their sum.**
Calling the function:

```go
total:= sum(1, 2, 3, 4, 5);
fmt.Println("Total: ", total);
```

_🔧 Anonymous Functions and IIFE_
Go also allows you to define anonymous functions — functions that do not have a name.

Syntax:

```go
func(parameters) returnType {
    // Function body
}
```

Example:

```go
func(){
    fmt.Println("Hello from anonymous function");
}()
```

**Here, the anonymous function prints a message and is immediately executed.(IIFE)**

```go
package main

import "fmt"

// standard or named function
// func add(a int, b int){
// 	fmt.Println(a+b)
// }

func main(){
	// anonymous functon and it should be called or it will give error
	//and if called then it is called immediate invoked function ---> IIFE
	func(a int, b int){
		c:= a + b

		fmt.Println(c)
	}(5, 6)
}

func init(){
	fmt.Println("I will be called first")
}

I will be called first
11
```

_🔹 What is init() in Go?_
The init() function is a special built-in function in Go that gets called automatically when your program starts — before the main() function is run.

You don't need to call it manually, and you can define multiple init() functions across different files or packages.

🧠 Key Points About init()

🔍 Feature 🔧 Description
Automatic Execution Called automatically before main()
No Parameters Cannot take any arguments
No Return Value Cannot return anything
Multiple init()s You can have more than one init() in the same package (in different files or the same file)
Used For Setup tasks: initializing variables, setting configs, connecting to DBs, etc.

✅ Syntax of init()

```go
func init() {
    // Initialization logic here
}
```

🧪 Example: Using init()

```go
package main

import "fmt"

func init() {
    fmt.Println("Initializing...") // runs before main
}

func main() {
    fmt.Println("Main function running")
}
```

🖥️ Output:

```bash
Initializing...
Main function running
✅ See? The init() function ran before main() without us calling it.
```

**🎯 Why Use init()?**
Here are some common use cases:

- 🔹 Set up global variables
- 🔹 Open connections (e.g., database)
- 🔹 Load configuration files
- 🔹 Initialize third-party libraries

**🔁 Multiple init() Functions**
You can define multiple init() functions in the same file or different files of the same package.
They are executed in the order they appear (top to bottom) and across files by compile order.
Example:

```go
func init() {
    fmt.Println("First init")
}

func init() {
    fmt.Println("Second init")
}
```

🖥️ Output:

```bash
First init
Second init
🔥 All init() functions run before main() starts, no matter how many you have.
```

**📦 init() in Imported Packages**
Let’s say you import a package in your Go file. If that package has its own init() function, it will also run — before your main package’s init() or main().

Example:
📁 utils/utils.go:

```go
package utils
import "fmt"

func init(){
    fmt.Println("utils package initializing");
}
```

📁 main.go:

```go
package main
import {
    "fmt"
    "example.com/utils"
}

func init(){
    fmt.Println("Main package init");
}

func main(){
    fmt.Println("Main Function");
}
```

🖥️ Output:

```bash
utils package initialized
main package init
main function
```

**_⚠️ Things to Remember_**

- ❗ init() cannot be called manually
- ❗ It must be declared with no parameters and no return values
- ❗ It’s run once per file/package before the program starts
- ❗ Avoid putting complex logic inside init() unless truly necessary

**🧩 Real-World Example: Loading Config**

```go
var config map[string]string

func init() {
    config = map[string]string{
        "app_name": "MyApp",
        "version":  "1.0.0",
    }
}
```

**Now any part of your code can use the initialized config variable without explicitly calling a setup function.**

_🧠 What is a Function Expression?_
In Go, unlike JavaScript or Python, we don’t typically assign anonymous functions to variables as often — but yes, you can do it.

This is what people often mean by "expression function" — when a function is:

- ✅ Created
- ✅ Assigned to a variable
- ✅ Then called using that variable

So basically:

```go
varName := func(...) {...}
```

```go
func main(){
    greet:= func(name string) string{
        return "Hi, " + name "👋"
    };

    message:= greet("Somel");
    fmt.Println(message);
}
```

```go
package main

import "fmt"

func main(){
	// Function expression or Assign Function in variable
	add:= func(a int, b int){
		c:= a + b
		fmt.Println(c)
	}

	add(2, 3)
}

func init(){
	fmt.Println("I will be called first")
}

I will be called first
5
```

```go
package main

import "fmt"

func main(){
	// Function expression or Assign Function in variable
	add(2, 3)

	add:= func(a int, b int){
		c:= a + b
		fmt.Println(c)
	}

}

func init(){
	fmt.Println("I will be called first")
}

undefined: addcompilerUndeclaredName
```

```go
package main

import "fmt"

func main(){
	// Function expression or Assign Function in variable
	add:= func(a int, b int){
		c:= a + b
		fmt.Println(c)
	}

	add(2, 3)
}

func init(){
	fmt.Println("I will be called first")
}

I will be called first
5
```

```go
package main

import "fmt"

func main(){
	// Function expression or Assign Function in variable
	add(2, 3)

	add:= func(a int, b int){
		c:= a + b
		fmt.Println(c)
	}

}

func init(){
	fmt.Println("I will be called first")
}

undefined: addcompilerUndeclaredName
```

_🧠 What is a Higher-Order Function?_
A Higher-Order Function (HOF) is a function that:

- ✅ Takes one or more functions as arguments, or
- ✅ Returns a function as a result

Yes! Functions can be treated like values in Go too! 💡

🟩 Example 1: Passing a Function as an Argument

```go
func greet(name string){
    fmt.Println("Hello, ", name);
}

func greetSomeone(action func(string)){
    action("Nobody");
}

func main(){
    greetSomeone(greet);
}
```

```bash
✅ Output:
Hello, NoobDev
```

🔍 Explanation:
greetSomeone takes a function action as an argument
Inside, it calls that function
We pass greet to it — like giving the function as a gift! 🎁

🟨 Example 2: Returning a Function

```go
func multiplier(factor int) func(int) int{
    return func(n int) int{
        return n * factor;
    }
}

func main(){
    double:= multiplier(2);
    triple:= multiplier(3);

    fmt.Println("Double of 5 is: ", double(5)) // 10
    fmt.Println("Tripple of 5 is: ", triple(5))// 15
}
```

💡 Output:

```bash
Double of 5: 10
Triple of 5: 15
```

🔍 Explanation:
multiplier returns a new function that multiplies by the factor
We call multiplier(2) and get a new function that doubles

```go
package main

import "fmt"

func preprocess(a int, b int, op func(int, int)){
	op(a, b);
}

func add(x int, y int){
	z:= x + y;
	fmt.Println(z);
}
func main(){
	preprocess(1, 2, add); // 3
}
```
```go
package main

import "fmt"

func call() func(int, int){
	return add;
}

func add(x int, y int){
	z:= x + y;
	fmt.Println(z);
}
func main(){
	sum:= call(); // add function
	sum(4, 5) // 9
}
```

**🧠 Functions as First-Class Citizens**
In Go, functions are first-class citizens. This means:
You can assign functions to variables.
You can pass functions as arguments.
You can return functions from other functions.