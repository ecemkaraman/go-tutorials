package errorhandling

import "fmt"

func Run() {
	fmt.Println("=== Standard Errors ===")
	RunErrors()

	fmt.Println("\n=== Custom Errors ===")
	RunCustomErrors()

	fmt.Println("\n=== Panic and Recover ===")
	RunPanic()
}
