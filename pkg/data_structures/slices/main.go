package main

import "fmt"

func main() {
	fmt.Println("🔹 Slices in Go (Dynamic, Resizable Views over Arrays)")

	// ✅ Empty slice (zero length, nil by default)
	var s1 []int
	fmt.Println("Empty slice (nil):", s1, "Length:", len(s1), "Cap:", cap(s1))

	// ✅ Initialized with values (slice literal)
	s2 := []int{1, 2, 3}
	fmt.Println("Slice literal:", s2)

	// ✅ Appending elements
	fmt.Printf("Initial ➤ Length = %v | Capacity = %v\n", len(s2), cap(s2))
	s2 = append(s2, 4, 5)
	fmt.Println("Appended slice:", s2)
	fmt.Printf("After Append ➤ Length = %v | Capacity = %v\n\n", len(s2), cap(s2))

	// ✅ Creating slice from an array
	arr := [5]int{10, 20, 30, 40, 50}
	s3 := arr[1:4] // [20, 30, 40]
	fmt.Println("Slice from array:", s3)

	// ✅ Creating slice from another slice
	s4 := s2[2:5] // [3, 4, 5]
	fmt.Println("Slice from slice:", s4)

	// ✅ Using make(): length = 3, capacity = 5 (optional) 
	s5 := make([]int, 3, 5)
	s5[0], s5[1], s5[2] = 7, 8, 9
	fmt.Println("Slice with make():", s5, "Len:", len(s5), "Cap:", cap(s5))

	// ✅ Copy slices
	dst := make([]int, len(s4))
	copy(dst, s4)
	fmt.Println("Copied slice:", dst)

	// ✅ Looping with range
	fmt.Println("Range loop:")
	for i, v := range s2 {
		fmt.Printf("Index %d → Value %d\n", i, v)
	}

	// ✅ Slice from full array (omit boundaries)
	full := arr[:]      // All elements
	startTo3 := arr[:3] // First 3
	from2 := arr[2:]    // From index 2 to end
	fmt.Println("Full slice:", full)
	fmt.Println("Start to 3:", startTo3)
	fmt.Println("From 2 to end:", from2)

	// ✅ Declare empty slice explicitly (not nil)
	empty := []int{}
	fmt.Println("Explicit empty slice:", empty, "Len:", len(empty), "Cap:", cap(empty))
}
