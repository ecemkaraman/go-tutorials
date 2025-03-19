package reflection

import (
	"fmt"
	"reflect"
)

// Sample Struct
type Person struct {
	Name string
	Age  int
}

// Runs Reflection Example
func RunReflection() {
	p := Person{Name: "Ecem", Age: 28}

	// ✅ Get type info
	t := reflect.TypeOf(p)
	fmt.Println("✅ Type:", t.Name())

	// ✅ Iterate over struct fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("🔹 Field: %s, Type: %s\n", field.Name, field.Type)
	}

	// ✅ Modify values using Reflection
	v := reflect.ValueOf(&p).Elem()
	v.FieldByName("Name").SetString("Updated Name")
	fmt.Println("✅ Modified Struct:", p)
}

/*
🔹 Explanation:
- `reflect.TypeOf(obj)`: Gets type metadata.
- `reflect.ValueOf(obj)`: Gets actual value.
- `.FieldByName("Name").SetString()`: Modifies struct field dynamically.
*/
