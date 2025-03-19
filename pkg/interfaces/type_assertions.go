package interfaces

import "fmt"

// ✅ Function using type assertion
func IdentifyType(i interface{}) {
	value, ok := i.(string)
	if ok {
		fmt.Println("It's a string:", value)
	} else {
		fmt.Println("Not a string")
	}
}

// Runs the type assertions example
func RunTypeAssertions() {
	IdentifyType("Golang") // "It's a string: Golang"
	IdentifyType(100)      // "Not a string"
}

/*
🔹 Explanation:
- `i.(string)`: Type assertion to check if `i` is a string.
- `ok`: Boolean result of assertion.
*/
