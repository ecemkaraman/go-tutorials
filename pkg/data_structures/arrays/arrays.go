package arrays

import "fmt"

// PracticeArrays demonstrates common array usage scenarios
func Run() {
	fmt.Println("\n🔹 Arrays in Go (Fixed-Size Storage)")
	// ✅ 1. Declaration & Initialization
	// 1️⃣ Declare an array without initializing (defaults to zero values)
	var arr1 [5]int
	fmt.Println("Default array:", arr1) // [0 0 0 0 0]

	// 2️⃣ Declare and initialize explicitly
	var arr2 = [3]int{10, 20, 30}
	fmt.Println("Explicitly initialized array:", arr2) // [10 20 30]

	// 3️⃣ Short declaration with `:=`
	arr3 := [4]int{1, 2, 3, 4}
	fmt.Println("Short declaration:", arr3) // [1 2 3 4]

	// 4️⃣ Let the compiler infer the size using `...`
	arr4 := [...]int{5, 10, 15, 20, 25}
	fmt.Println("Auto-sized array:", arr4) // [5 10 15 20 25]

	// 5️⃣ Initialize specific indexes
	arr5 := [5]int{0: 100, 3: 200}               // Only sets index 0 and 3, others remain 0
	fmt.Println("Partial initialization:", arr5) // [100 0 0 200 0]

	// 6️⃣ Using a loop to initialize an array
	var arr6 [5]int
	for i := 0; i < len(arr6); i++ {
		arr6[i] = i * 10 // Assign values dynamically
	}
	fmt.Println("Loop-initialized array:", arr6) // [0 10 20 30 40]

	// 7️⃣ Iterating over an array using range
	fmt.Print("Using range loop: ")
	for index, value := range arr6 {
		fmt.Printf("Index %d -> Value %d | ", index, value)
	}
	fmt.Println() // New line for formatting
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)
	// ---------------------------------------------------------
	// ✅ 2. Indexed Access
	fmt.Println("First Element:", arr[0])
	fmt.Println("Last Element:", arr[len(arr)-1])

	// ✅ 3. Iterating Over an Array
	for i, v := range arr {
		fmt.Printf("Index %d → Value %d\n", i, v)
	}

	// ✅ 4. Modifying Elements
	arr[2] = 99
	fmt.Println("Updated Array:", arr)

	// ✅ 5. Array Length
	fmt.Println("Array Length:", len(arr))

	// ✅ 6. Array Slicing
	slice := arr[1:4]
	fmt.Println("Array Slice:", slice)

	// ✅ 7. Array Copy
	copyArr := arr

	// ✅ 8. Array Comparison
	fmt.Println("Are arrays equal?", arr == copyArr)

	// ✅ 9. Multi-Dimensional Arrays
	matrix := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("Matrix:", matrix)

	// ✅ 10. Array of Arrays
	arrays := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("Array of Arrays:", arrays)

	// ✅ 11. Array of Slices
	slices := [2][]int{{1, 2}, {3, 4}}
	fmt.Println("Array of Slices:", slices)

	// ✅ 12. Array of Maps
	maps := [2]map[string]int{{"Alice": 30}, {"Bob": 25}}
	fmt.Println("Array of Maps:", maps)

	// ✅ 13. Array of Structs
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
