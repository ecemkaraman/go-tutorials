package filehandling

import (
	"fmt"
	"os"
)

// Runs file appending example
func Run() {
	content := "\nAppending new content."

	// ✅ Open file in append mode
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("❌ Error opening file:", err)
		return
	}
	defer file.Close()

	// ✅ Append data to file
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("❌ Error appending to file:", err)
		return
	}

	fmt.Println("✅ File Appended Successfully")
}

/*
🔹 Explanation:
- `os.OpenFile("file", os.O_APPEND|os.O_WRONLY, 0644)`: Opens file for appending.
- `file.WriteString("data")`: Appends text.
- `defer file.Close()`: Ensures file closure.
*/
