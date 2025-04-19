```go
package main

import "fmt"

func main(){

	fmt.Println("Hello, world!");
}
```

### 🔹 `package main`

- In Go, every file belongs to a **package**.
- `main` is a special package that tells Go:
  > "Start execution from here."
- If you don’t write `package main`, the program won’t run directly.

➡️ Remember: `package main` makes this an **executable program**.

### 🔹 `import "fmt"`

- To use built-in or external packages in Go, you need to **import** them.
- `fmt` is the package that handles **formatted I/O** — like printing text.

➡️ `fmt` stands for **format** (used for text formatting and output).

### 🔹 `func main() { ... }`

- This is the **entry point** of your Go program.
- Go will always start executing your code from the `main()` function.

➡️ `func` means you’re defining a **function**.

### 🔹 `fmt.Println("Hello, World!")`

- `Println` = Print + New Line
- This line prints the message to the screen.
- `"Hello, World!"` is the string you want to print.

➡️ Output will be:

```

Hello, World!

```

### 🔧 Steps to Run This Program

1. Create a folder, e.g., `hello-go`
2. Inside it, create a file: `main.go`
3. Paste and save the code above
4. Open your terminal and navigate to the folder
5. Run the program with:

```bash

go run main.go

```

## 🎉 Done!

Your Go program will successfully print:

```

Hello, World!

```
