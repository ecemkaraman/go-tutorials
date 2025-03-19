package main

import (
	"fmt"

	"github.com/ecemkaraman/go-tutorials/go-error-file-handling/errorhandling"
	"github.com/ecemkaraman/go-tutorials/go-error-file-handling/filehandling"
)

func main() {
	fmt.Println("🚀 Go Error Handling & File Handling Practice")

	// ✅ Run Basic Error Handling
	fmt.Println("\n🔹 Running Error Handling")
	errorhandling.RunErrors()

	// ✅ Run Panic & Recover Example
	fmt.Println("\n🔹 Running Panic & Recover Example")
	errorhandling.RunPanicRecover()

	// ✅ Run Custom Error Example
	fmt.Println("\n🔹 Running Custom Error Example")
	errorhandling.RunCustomError()

	// ✅ Run File Reading Example
	fmt.Println("\n🔹 Running File Reading Example")
	filehandling.RunReadFile()

	// ✅ Run File Writing Example
	fmt.Println("\n🔹 Running File Writing Example")
	filehandling.RunWriteFile()

	// ✅ Run File Appending Example
	fmt.Println("\n🔹 Running File Appending Example")
	filehandling.RunAppendFile()
}
