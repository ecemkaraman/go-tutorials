package interfaces

import "fmt"

// ✅ Function using type switch
func CheckType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Boolean:", v)
	default:
		fmt.Println("Unknown Type")
	}
}

// Runs the type switch example
func RunTypeSwitch() {
	CheckType(42)       // Integer
	CheckType("GoLang") // String
	CheckType(true)     // Boolean
	CheckType(3.14)     // Unknown Type
}

/*
🔹 Explanation:
- `switch v := i.(type)`: Checks variable type dynamically.
- Cases match different types and handle accordingly.
*/
