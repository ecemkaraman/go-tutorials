package main

import (
	"fmt"

	"github.com/ecemkaraman/go-tutorials/go-advanced/concurrency"
	"github.com/ecemkaraman/go-tutorials/go-advanced/memory"
	"github.com/ecemkaraman/go-tutorials/go-advanced/modules"
	"github.com/ecemkaraman/go-tutorials/go-advanced/reflection"
)

func main() {
	fmt.Println("🚀 Go Advanced Topics Practice")

	// ✅ Run Reflection Example
	fmt.Println("\n🔹 Running Reflection Example")
	reflection.RunReflection()

	// ✅ Run Context Example
	fmt.Println("\n🔹 Running Context Management Example")
	concurrency.RunContextExample()

	// ✅ Run Memory Optimization with sync.Pool
	fmt.Println("\n🔹 Running sync.Pool Memory Optimization Example")
	memory.RunSyncPoolExample()

	// ✅ Run Benchmark Test (Manually use `go test -bench .`)
	fmt.Println("\n🔹 Run `go test -bench .` to execute Benchmark Test")

	// ✅ Run Go Modules & Dependency Fetching
	fmt.Println("\n🔹 Running Go Modules Example")
	modules.RunGoModules()
}
