package errorhandling

import (
	"errors"
	"fmt"
)

// Runs basic error handling example
func RunErrors() {
	// ✅ Create a new error
	err := errors.New("this is a basic error")
	fmt.Println("Error:", err)

	// ✅ Format an error with context
	name := "Ecem"
	err = fmt.Errorf("user %s not found", name)
	fmt.Println("Formatted Error:", err)

	// ✅ Check for errors
	if err != nil {
		fmt.Println("❌ An error occurred:", err)
	}
}

/*
🔹 Explanation:
- `errors.New()`: Creates a basic error.
- `fmt.Errorf()`: Formats error messages dynamically.
- `if err != nil`: Standard Go error handling check.
*/
