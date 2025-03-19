package slices

import "fmt"

// PracticeSlices demonstrates common slice usage scenarios
func PracticeSlices() {
	fmt.Println("\n🔹 Slices in Go (Dynamic Storage)")

	// ✅ 1. Basic Slice Declaration
	slice := []int{10, 20, 30}
	fmt.Println("Slice before append:", slice)

	// ✅ 2. Appending Elements (Resizable)
	slice = append(slice, 40, 50)
	fmt.Println("Slice after append:", slice)

	// ✅ 3. Slicing an Existing Slice
	subSlice := slice[1:4] // Extract elements from index 1 to 3
	fmt.Println("Sub-Slice:", subSlice)

	// ✅ 4. Using `make()` for Preallocated Capacity
	slice2 := make([]int, 3, 5) // Length 3, Capacity 5
	fmt.Println("Preallocated Slice:", slice2, "Length:", len(slice2), "Capacity:", cap(slice2))

	// ✅ 5. Avoiding Excessive Reallocation (Preallocation Example)
	slice3 := make([]int, 0, 1000) // Avoids multiple reallocations
	for i := 0; i < 1000; i++ {
		slice3 = append(slice3, i)
	}
	fmt.Println("Optimized Slice:", len(slice3), "Capacity:", cap(slice3))
}
