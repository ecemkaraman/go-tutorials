package arrays

import "fmt"

func Run() {
	fmt.Println("\n🔹 Arrays in Go (Fixed-Size Storage)")
	// ✅ 1. Declaration & Initialization
	var arr1 [5]int                     //declares array of 5 ints with default 0s, no initialization
	fmt.Println("Default array:", arr1) // [0 0 0 0 0]

	var arr2 = [3]int{10, 20, 30}                      // Declare and initialize explicitly
	fmt.Println("Explicitly initialized array:", arr2) // [10 20 30]

	arr3 := [4]int{1, 2, 3, 4}              //Short declaration with `:=`
	fmt.Println("Short declaration:", arr3) // [1 2 3 4]

	arr4 := [...]int{5, 10, 15, 20, 25}    //size inferred using ...
	fmt.Println("Auto-sized array:", arr4) // [5 10 15 20 25]

	arr5 := [5]int{0: 100, 3: 200}               // Initializes specific indices only->only sets index 0 and 3
	fmt.Println("Partial initialization:", arr5) // [100 0 0 200 0]

	var arr6 [5]int // Initialize an array via loop
	for i := 0; i < len(arr6); i++ {
		arr6[i] = i * 10 // Assign values dynamically
	}
	fmt.Println("Loop-initialized array:", arr6) // [0 10 20 30 40]

	fmt.Print("Using range loop: ") //Iterating over an array using range
	for index, value := range arr6 {
		fmt.Printf("Index %d -> Value %d | ", index, value)
	}

	// ---------------------------------------------------------
	fmt.Println() // New line for formatting
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)

	// ✅ 2. Indexed Access
	fmt.Println("First Element:", arr[0])
	fmt.Println("Last Element:", arr[len(arr)-1])

	// ✅ 3. Iterating Over an Array
	for index, value := range arr {
		fmt.Printf("Index %d → Value %d\n", index, value)
	}

	// ✅ 4. Modifying Elements
	arr[2] = 99
	fmt.Println("Updated Array:", arr)

	// ✅ 5. Array Utility Functions
	fmt.Println("Array Length:", len(arr))

	slice := arr[1:4] //Array Slicing -> [2 3 4]
	fmt.Println("Array Slice:", slice)

	copyArr := arr // Array Copy (Shallow Copy, not a reference)

	fmt.Println("Are arrays equal?", arr == copyArr) //Element-wise comparison

	// ✅ 6. Advanced Use Cases
	// Multi-Dimensional Arrays
	matrix := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("Matrix:", matrix)

	// Array of Slices
	slices := [2][]int{{1, 2}, {3, 4}} //len(arr)=2, len(slice)=dynamic ([])
	fmt.Println("Array of Slices:", slices)

	// Array of Maps
	maps := [2]map[string]int{{"Alice": 30}, {"Bob": 25}}
	fmt.Println("Array of Maps:", maps)

	// Array of Structs
	type Person struct {
		Name string
		Age  int
	}
	people := [2]Person{{"Alice", 30}, {"Bob", 25}}
	fmt.Println("Array of Structs:", people)
}

// ✅ Declares a fixed-size array → arr := [5]int{1, 2, 3, 4, 5} (cannot resize).
// ✅ Accesses elements via index → arr[0], arr[len(arr)-1].
// ✅ Iterates over the array → Uses range to print index-value pairs.
// ✅ Modifies an array element → arr[2] = 99.

// How to declare an array in Go?

/*
🚀 Arrays are best for storing a fixed number of elements with direct index access.
✅ Declares a fixed-size array → arr := [5]int{1, 2, 3, 4, 5} (cannot resize).
✅ Accesses elements via index → arr[0], arr[len(arr)-1].
✅ Iterates over the array → Uses range to print index-value pairs.
✅ Modifies an array element → arr[2] = 99.
*/
