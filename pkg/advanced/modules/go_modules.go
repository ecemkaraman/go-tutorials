package modules

import (
	"fmt"
	_ "golang.org/x/text/language" // Example external package, need to install-> go get golang.org/x/text/language
)

// Runs Go Modules Example
func RunGoModules() {
	fmt.Println("✅ Go Modules Initialized")

	// ✅ Install package using `go get`
	fmt.Println("Run: go get golang.org/x/text/language")
}

/*
🔹 Explanation:
- `go mod init project-name`: Initializes module.
- `go get package-name`: Fetches dependency.
- `_ "import-path"`: Import-only for initialization.
*/
