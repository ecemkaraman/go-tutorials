package main

import "fmt"

// 🔹 Global variable declaration (outside any function)
var globalVar string = "I am a global variable"

func main() {
	fmt.Println("🚀 Go Syntax & Core Features Practice 🚀")

	// ✅ 1. Variables & Data Types
	declareVariables()

	// ✅ 2. Operators
	operatorsDemo()

	// ✅ 3. Control Flow (if-else, switch)
	controlFlowDemo(10)

	// ✅ 4. Loops (for loop, range, while-style)
	loopsDemo()

	// ✅ 5. Functions (with parameters, return values, multiple returns)
	sum, diff := mathOperations(15, 5)
	fmt.Println("Sum:", sum, "Difference:", diff)
}

// 🔹 1. Variables & Data Types
func declareVariables() {
	// Explicit variable declaration
	var age int = 25
	var height float64 = 5.9
	var isStudent bool = false
	var name string = "Alice"

	// Type Inference (using `:=`)
	weight := 60.5 // Go automatically infers float64

	fmt.Println("\n🔹 Variables & Data Types")
	fmt.Println("Name:", name, "Age:", age, "Height:", height, "Weight:", weight, "Is Student:", isStudent)

	// Constants
	const pi float64 = 3.1415
	fmt.Println("Constant PI:", pi)
}

// 🔹 2. Operators Demo
func operatorsDemo() {
	a, b := 10, 3
	fmt.Println("\n🔹 Operators Demo")

	// Arithmetic Operators
	fmt.Println("Addition:", a+b, "Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b, "Division:", a/b, "Modulus:", a%b)

	// Comparison Operators
	fmt.Println("a == b:", a == b, "a != b:", a != b, "a > b:", a > b, "a < b:", a < b)

	// Logical Operators
	fmt.Println("true && false:", true && false, "true || false:", true || false, "!true:", !true)
}

// 🔹 3. Control Flow (if-else, switch)
func controlFlowDemo(num int) {
	fmt.Println("\n🔹 Control Flow Demo")

	// If-Else
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// Switch Case
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Another day")
	}
}

// 🔹 4. Loops (for loop, range loop, while-style loop)
func loopsDemo() {
	fmt.Println("\n🔹 Loops Demo")

	// Standard For Loop
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Range Loop (iterating over a slice)
	nums := []int{10, 20, 30, 40}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// While-style Loop
	count := 0
	for count < 3 {
		fmt.Println("While-style loop iteration:", count)
		count++
	}
}

// 🔹 5. Functions (single & multiple returns)
func mathOperations(x, y int) (int, int) {
	sum := x + y
	diff := x - y
	return sum, diff
}
