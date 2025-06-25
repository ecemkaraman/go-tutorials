package basics

import (
	"fmt"
	"math"
)

const Pi = 3.14

func Run() {
	// -------- Variables --------
	var explicit int = 42
	implicit := "Go is fun!"
	var zeroValue string
	var x, y int = 10, 20
	const greeting = "Hello, World!"

	fmt.Println("explicit:", explicit)
	fmt.Println("implicit:", implicit)
	fmt.Println("zeroValue:", zeroValue)
	fmt.Println("multi vars:", x, y)
	fmt.Println("constant:", greeting)
	fmt.Println("Pi:", Pi)

	// -------- Arithmetic --------
	a, b := 5, 2
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)
	fmt.Println("float division:", float64(a)/float64(b))

	// -------- Logic --------
	fmt.Println("true && false:", true && false)
	fmt.Println("true || false:", true || false)
	fmt.Println("!true:", !true)

	// -------- Conversion --------
	num := 42
	converted := float64(num)
	fmt.Printf("Type conversion: %d -> %.2f\n", num, converted)

	// -------- Shadowing --------
	z := 100
	fmt.Println("Outer z:", z)
	{
		z := 200
		fmt.Println("Inner z:", z)
	}
	fmt.Println("Outer z again:", z)

	// -------- Functions --------
	sum := add(3, 4)
	fmt.Println("Sum:", sum)

	result, ok := divide(10, 2)
	if ok {
		fmt.Println("Divide:", result)
	} else {
		fmt.Println("Invalid division")
	}

	_, bad := divide(1, 0)
	if !bad {
		fmt.Println("Handled divide by zero gracefully.")
	}

	// -------- Math --------
	fmt.Println("Sqrt(16):", math.Sqrt(16))
	fmt.Println("Pow(2, 4):", math.Pow(2, 4))
}

func add(a int, b int) int {
	return a + b
}

func divide(a, b float64) (result float64, valid bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}
