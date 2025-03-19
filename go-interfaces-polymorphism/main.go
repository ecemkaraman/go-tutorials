package main

import (
	"fmt"

	"github.com/ecemkaraman/go-tutorials/go-interfaces-polymorphism/interfaces"
)

func main() {
	fmt.Println("🚀 Go Interfaces & Polymorphism Practice")

	// ✅ Basic Interface Implementation
	fmt.Println("\n🔹 Running Basic Interface Example")
	interfaces.RunBasicInterface()

	// ✅ Interface as Function Parameter
	fmt.Println("\n🔹 Running Interface as Parameter Example")
	interfaces.RunInterfaceAsParam()

	// ✅ Using Empty Interface (Generic Handling)
	fmt.Println("\n🔹 Running Empty Interface Example")
	interfaces.RunEmptyInterface()

	// ✅ Using Type Assertions
	fmt.Println("\n🔹 Running Type Assertions Example")
	interfaces.RunTypeAssertions()

	// ✅ Using Type Switch
	fmt.Println("\n🔹 Running Type Switch Example")
	interfaces.RunTypeSwitch()
}
