package interfaces

import "fmt"

// ✅ Function with empty interface (can take any type)
func PrintAnything(v interface{}) {
	fmt.Println("Received:", v)
}

// Runs the empty interface example
func RunEmptyInterface() {
	PrintAnything(42)          // Integer
	PrintAnything("Hello")     // String
	PrintAnything([]int{1, 2}) // Slice
}

/*
🔹 Explanation:
- `interface{}` (empty interface): Accepts any type.
- Useful for generic functions or unknown input types.
*/
