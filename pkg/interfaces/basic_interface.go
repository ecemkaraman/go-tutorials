package interfaces

import "fmt"

// ✅ Define an Interface
type Speaker interface {
	Speak() string
}

// ✅ Implement Interface on a Struct
type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

// Runs the basic interface example
func RunBasicInterface() {
	var s Speaker = Person{Name: "Ecem"}
	fmt.Println(s.Speak()) // "Hello, my name is Ecem"
}

/*
🔹 Explanation:
- `type Speaker interface { Speak() string }`: Defines an interface.
- `Person struct` implements `Speak() string` → satisfies `Speaker`.
- Can assign `Person` to `Speaker` since method matches.
*/
