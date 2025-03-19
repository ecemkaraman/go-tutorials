package arrays

import "fmt"

// PracticeArrays demonstrates common array usage scenarios
func PracticeArrays() {
	fmt.Println("\n🔹 Arrays in Go (Fixed-Size Storage)")

	// ✅ 1. Basic Declaration
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)

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
}

/*
🚀 Arrays are best for storing a fixed number of elements with direct index access.
✅ Declares a fixed-size array → arr := [5]int{1, 2, 3, 4, 5} (cannot resize).
✅ Accesses elements via index → arr[0], arr[len(arr)-1].
✅ Iterates over the array → Uses range to print index-value pairs.
✅ Modifies an array element → arr[2] = 99.
*/
