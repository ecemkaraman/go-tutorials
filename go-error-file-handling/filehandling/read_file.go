package filehandling

import (
	"fmt"
	"os"
)

// Runs file reading example
func RunReadFile() {
	// ✅ Open and read file content
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("❌ Error reading file:", err)
		return
	}
	fmt.Println("✅ File Content:\n", string(data))
}

/*
🔹 Explanation:
- `ioutil.ReadFile("file")`: Reads entire file.
- `if err != nil`: Error handling if file missing.
*/
