package errorhandling

import (
	"fmt"
)

// ✅ Custom Error Type
type CustomError struct {
	Code    int
	Message string
}

// ✅ Implement Error() method for CustomError
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// Runs custom error example
func RunCustomErrors() {
	err := &CustomError{Code: 404, Message: "Resource Not Found"}
	fmt.Println(err) // Implicitly calls Error()
}

/*
🔹 Explanation:
- `type CustomError struct{}`: Defines custom error.
- `Error() string`: Implements built-in error interface.
*/
