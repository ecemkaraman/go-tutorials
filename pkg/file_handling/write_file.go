package filehandling

import (
	"fmt"
	"os"
)

// Runs file writing example
func RunWriteFile() {
	content := "Hello, Go! This is a new file."

	// ✅ Write content to file
	err := os.WriteFile("output.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("❌ Error writing file:", err)
		return
	}
	fmt.Println("✅ File Written Successfully")
}

/*
🔹 Explanation:
- `ioutil.WriteFile("file", []byte, perm)`: Creates & writes content.
- `0644`: File permissions.
*/
