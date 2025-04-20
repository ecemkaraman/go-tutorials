package main

import (
	"fmt"
)

// Run executes all pointer examples
func main() {
	fmt.Println("🔹 Pointers in Go: Demonstration")

	// 1️⃣ Basic Pointer Declaration
	var ptr *int                                           // Declaring a pointer to an int (initially nil)
	fmt.Println("Pointer declared but not assigned:", ptr) // Will print <nil>

	// 2️⃣ Creating a Pointer Using `new()`
	num := new(int) // Allocates memory for an int and returns a pointer
	*num = 42       // Assigning a value via dereferencing
	fmt.Println("Pointer created with `new()`, value:", *num)

	// 3️⃣ Assigning Pointer to an Existing Variable
	val := 10
	ptrToVal := &val // Get memory address of `val`
	fmt.Println("Memory Address of `val`:", ptrToVal)
	fmt.Println("Value accessed via pointer:", *ptrToVal)

	// 4️⃣ Modifying Value via Pointer
	*ptrToVal = 20 // Changing value at memory address
	fmt.Println("Updated `val` via pointer:", val)

	// 5️⃣ Passing Pointers to Functions (Pass by Reference)
	fmt.Println("\n🔹 Passing Pointers to Functions")
	number := 5
	fmt.Println("Before:", number)
	increment(&number) // Pass address to modify original variable
	fmt.Println("After:", number)

	// 6️⃣ Returning Pointers from Functions
	newPointer := createPointer(99)
	fmt.Println("\n🔹 Returning Pointers from Functions")
	fmt.Println("Value of pointer returned:", *newPointer)

	// 7️⃣ Nil Pointers & Checking for Nil
	var uninitializedPtr *string
	if uninitializedPtr == nil {
		fmt.Println("\n🔹 Nil Pointer Check: Pointer is nil, avoiding dereference error")
	}

	// 8️⃣ Pointer to a Struct
	fmt.Println("\n🔹 Pointers to Structs")
	type Person struct {
		Name string
		Age  int
	}
	p := Person{"Alice", 25}
	ptrToPerson := &p
	fmt.Println("Person's Name via Pointer:", ptrToPerson.Name) // Direct access via pointer

	// 9️⃣ Pointer to an Array
	fmt.Println("\n🔹 Pointer to Arrays")
	arr := [3]int{1, 2, 3}
	arrPtr := &arr
	fmt.Println("Array via pointer:", *arrPtr) // Dereferencing prints the whole array

	// 🔟 Pointer to a Slice
	fmt.Println("\n🔹 Pointer to Slices")
	slice := []int{10, 20, 30}
	slicePtr := &slice
	fmt.Println("Slice via pointer:", *slicePtr) // Works like normal slice reference

	// 🔹 Summary
	fmt.Println("\n✅ Pointers Demonstration Completed!")
}

// 📌 Function that takes a pointer parameter (Modifies original variable)
func increment(n *int) {
	*n += 1 // Dereference and modify
}

// 📌 Function that returns a pointer to an integer
func createPointer(val int) *int {
	p := val
	return &p // Return pointer to local variable (safe in this case)
}
