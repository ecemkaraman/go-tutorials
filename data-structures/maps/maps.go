package maps

import "fmt"

// PracticeMaps demonstrates common map usage scenarios
func PracticeMaps() {
	fmt.Println("\n🔹 Maps in Go (Key-Value Storage)")

	// ✅ 1. Basic Map Declaration
	ages := map[string]int{"Alice": 30, "Bob": 25}
	fmt.Println("Initial Map:", ages)

	// ✅ 2. Adding/Updating Values
	ages["Charlie"] = 40
	ages["Alice"] = 31
	fmt.Println("Updated Map:", ages)

	// ✅ 3. Deleting a Key
	delete(ages, "Bob")
	fmt.Println("After Deletion:", ages)

	// ✅ 4. Checking If a Key Exists
	age, exists := ages["Alice"]
	if exists {
		fmt.Println("Alice's age:", age)
	} else {
		fmt.Println("Alice not found")
	}

	// ✅ 5. Iterating Over a Map
	for key, value := range ages {
		fmt.Printf("Key: %s → Value: %d\n", key, value)
	}
}
