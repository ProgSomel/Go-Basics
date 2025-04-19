### 🟦 Mac-এ Golang ইনস্টল করার ধাপ:
🔗 Golang ওয়েবসাইটে যাও:
### https://go.dev/dl/

💾 তোমার Mac অনুযায়ী সঠিক ফাইলটি ডাউনলোড করো:

✅ Intel Mac হলে → macOS (pkg)

✅ M1/M2 (Apple Silicon) হলে → macOS ARM64 (pkg)

📦 .pkg ফাইলটি ডাবল ক্লিক করে ইনস্টল করো।
📋 ইনস্টলেশন উইজার্ডের সব ধাপ ফলো করলেই হবে।

🧪 টার্মিনালে গিয়ে চেক করো Go ইনস্টল হয়েছে কিনা:

```bash

go version

```

🖥️ যদি ইনস্টল ঠিকমতো হয়, তাহলে নিচের মতো আউটপুট দেখাবে:

```bash
go version go1.22.1 darwin/arm64
```

### 🟪 Mac-এ Visual Studio Code ইনস্টল করার ধাপ:
🌐 VS Code এর অফিসিয়াল সাইটে যাও:
[https://code.visualstudio.com/]

⬇️ "Download for macOS" বাটনে ক্লিক করে VS Code ডাউনলোড করো।

🖱️ ডাউনলোড হয়ে গেলে .zip ফাইল আনজিপ করে Applications ফোল্ডারে ড্র্যাগ করো।

🟨 VS Code-এ Go Extension ইনস্টল করো:
🟣 VS Code ওপেন করো

🎯 বাঁ পাশে এক্সটেনশন আইকনে ক্লিক করো (বা Cmd + Shift + X চাপো)

🔍 সার্চ করো: Go

🧩 "Go" - by Go Team at Google এক্সটেনশনটি ইনস্টল করো

### ✅ এখন চেক করে দেখো সব ঠিকভাবে কাজ করছে কিনা:
📁 ডেস্কটপে বা যেকোনো জায়গায় একটি ফোল্ডার তৈরি করো: golang-test

📄 ফোল্ডারের ভিতরে main.go নামে একটি ফাইল তৈরি করো, নিচের কোডটি লেখো:

```go

// main.go
package main

import "fmt"

func main() {
fmt.Println("Hello, Golang on Mac!")
}
```

🖥️ টার্মিনাল খুলে ঐ ফোল্ডারে যাও:

```bash

cd path/to/golang-test
```
### 🚀 এরপর কোড রান করো:

```bash

go run main.go

```

### 📤 আউটপুট:

Hello, Golang on Mac!
