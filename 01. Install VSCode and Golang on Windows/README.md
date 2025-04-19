## 🧱 Step 1: Go (Golang) ইনস্টল করা

### 🔗 Go ডাউনড লিংক

[The Go Programming Language](https://go.dev/)

1. উপরের লিংকে যান
2. আপনার সিস্টেম অনুযায়ী Windows-এর `.msi` ফাইলটি ডাউনলোড করুন
3. ডাবল ক্লিক করে ইনস্টল করুন
4. ইনস্টল শেষে `Command Prompt` খুলুন, লিখুন:

```bash
bash
CopyEdit
go version

```

✅ যদি দেখতে পান `go version go1.x.x`, তাহলে ইনস্টল সফল হয়েছে।

---

## 🧱 Step 2: Visual Studio Code (VS Code) ইনস্টল করা

### 🔗 [VS Code ডাউনলোড লিংক](https://code.visualstudio.com/)

1. ওয়েবসাইট থেকে Windows এর জন্য `User Installer` বা `System Installer` ডাউনলোড করুন
2. ডাবল ক্লিক করে ইনস্টল করুন
3. ইনস্টলের সময় অপশন আসলে এইগুলো চেক করে নিন:
   - “Add to PATH”
   - “Register code as editor for supported file types”

✅ ইনস্টল হয়ে গেলে, Start Menu থেকে "Visual Studio Code" চালু করুন

---

## 🧱 Step 3: Go Extension ইনস্টল VS Code-এ

1. VS Code চালু করুন
2. বাঁ পাশে Extensions আইকনে ক্লিক করুন (বা Ctrl + Shift + X চাপুন)
3. খুঁজুন: `Go`
4. যেটার নাম `Go by Go Team at Google` — ওটা ইনস্টল করুন

👉 এতে অটো কমপ্লিশন, লিন্টিং, রান করার সুবিধা পাবেন।

---

## 🧪 Step 4: টেস্ট প্রোগ্রাম চালানো

1. VS Code-এ একটি নতুন ফোল্ডার ওপেন করুন
2. একটি ফাইল তৈরি করুন: `main.go`
3. নিচের কোড লিখুন:

```go
go
CopyEdit
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}

```

1. টার্মিনাল ওপেন করুন (Ctrl + ` )
2. নিচের কমান্ড রান করুন:

```bash
bash
CopyEdit
go run main.go

```

✅ আউটপুট আসবে: `Hello, Go!`
