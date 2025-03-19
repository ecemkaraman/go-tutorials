package structs

import "fmt"

// Person struct
type Person struct {
	Name  string
	Age   int
	Email string
}

// Struct with an embedded field
type Employee struct {
	Person
	Salary int
}

// PracticeStructs demonstrates struct usage
func PracticeStructs() {
	fmt.Println("\n🔹 Structs in Go (Custom Data Types)")

	// ✅ 1. Basic Struct Usage
	p := Person{Name: "John Doe", Age: 28, Email: "john@example.com"}
	fmt.Println("Person Struct:", p)

	// ✅ 2. Accessing & Modifying Fields
	p.Age = 29
	fmt.Println("Updated Person Age:", p.Age)

	// ✅ 3. Using Structs with Methods
	fmt.Println("Greeting:", p.Greet())

	// ✅ 4. Struct Composition (Embedding)
	e := Employee{Person: Person{Name: "Alice", Age: 30, Email: "alice@example.com"}, Salary: 60000}
	fmt.Println("Employee Struct:", e)
}

// Greet method for Person struct
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}
