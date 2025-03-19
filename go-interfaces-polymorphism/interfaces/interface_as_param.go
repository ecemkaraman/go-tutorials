package interfaces

import "fmt"

// ✅ Define an Interface
type Animal interface {
	MakeSound() string
}

// ✅ Implement Interface
type Dog struct{}

func (d Dog) MakeSound() string { return "Woof!" }

type Cat struct{}

func (c Cat) MakeSound() string { return "Meow!" }

// ✅ Function that accepts an interface
func Speak(a Animal) {
	fmt.Println(a.MakeSound())
}

// Runs the interface as parameter example
func RunInterfaceAsParam() {
	dog := Dog{}
	cat := Cat{}

	Speak(dog) // Woof!
	Speak(cat) // Meow!
}

/*
🔹 Explanation:
- `type Animal interface { MakeSound() string }`: Defines behavior.
- `Dog` & `Cat` implement `MakeSound() string` → satisfy `Animal`.
- `Speak(Animal)`: Function accepts anything implementing `Animal`.
*/
