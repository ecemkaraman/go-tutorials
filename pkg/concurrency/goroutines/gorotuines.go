package goroutines

import (
	"fmt"
	"time"
)

// Function that runs as a Goroutine
func printMessage(msg string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(500 * time.Millisecond)
	}
}

// Runs Goroutine Example
func Run() {
	go printMessage("🔥 Goroutine 1")
	go printMessage("⚡ Goroutine 2")

	// Wait for Goroutines to finish
	time.Sleep(2 * time.Second)
	fmt.Println("✅ Main function done")
}

/*
🔹 Explanation:
- `go` keyword runs `printMessage` concurrently.
- `time.Sleep(2s)` prevents `main()` from exiting before Goroutines finish.
- Functions run asynchronously, making execution non-blocking.
*/
