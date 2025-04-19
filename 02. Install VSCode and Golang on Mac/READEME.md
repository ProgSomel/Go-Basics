🟦 Mac-এ Golang ইনস্টল করার ধাপ:
Golang ওয়েবসাইটে যান:
লিংক: https://go.dev/dl/

তোমার Mac অনুযায়ী সঠিক ফাইলটি ডাউনলোড করো:

যদি Intel Mac হয় → macOS (pkg)

যদি M1/M2 (Apple Silicon) হয় → macOS ARM64 (pkg)

.pkg ফাইলটি ডাবল ক্লিক করে ইনস্টল করো।
পুরো ইনস্টলেশন উইজার্ড ফলো করলেই হবে।

টার্মিনালে গিয়ে চেক করো Go ঠিকমতো ইনস্টল হয়েছে কিনা:


go version
যদি সব ঠিক থাকে, তাহলে কিছু এমন দেখাবে:


go version go1.22.1 darwin/arm64
🟪 Mac-এ Visual Studio Code ইনস্টল করার ধাপ:
VS Code এর অফিসিয়াল সাইটে যাও:
লিংক: https://code.visualstudio.com/

"Download for macOS" বাটনে ক্লিক করে VS Code ডাউনলোড করো।

ডাউনলোড শেষ হলে, ফাইলটাকে Applications ফোল্ডারে ড্র্যাগ করে রাখো।

🟨 Go Extension VS Code এ ইনস্টল করো:
VS Code ওপেন করো।

বাঁ পাশে এক্সটেনশন আইকনে ক্লিক করো (বা Cmd + Shift + X চাপো)।

সার্চ করো: Go

Go Team at Google এর যেই এক্সটেনশন আসবে, ওটাই ইনস্টল করে নাও।

✅ এখন চেক করে দেখো কাজ করছে কিনা:
ডেস্কটপে বা অন্য কোথাও একটা ফোল্ডার বানাও: golang-test

তার ভিতরে main.go নামে একটি ফাইল বানাও।


// main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Golang on Mac!")
}
টার্মিনাল ওপেন করে ঐ ফোল্ডারে যাও:


cd path/to/golang-test
go run main.go
দেখবে টার্মিনালে প্রিন্ট হবে:


Hello, Golang on Mac!
