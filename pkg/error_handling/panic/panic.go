package errorhandling

import "fmt"

// Runs panic & recover example
func RunPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("🛑 Recovered from Panic:", r)
		}
	}()

	fmt.Println("🚀 Executing function...")
	panic("Something went terribly wrong!") // 🔥 Trigger panic
	fmt.Println("This won't execute")
}

/*
🔹 Explanation:
- `panic()`: Halts execution & unwinds stack.
- `recover()`: Captures panic & prevents crash.
- `defer`: Ensures cleanup before function exits.
*/
