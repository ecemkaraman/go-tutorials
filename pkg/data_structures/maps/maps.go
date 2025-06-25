package maps

import "fmt"

func Run() {
	fmt.Println("\n🔹 Maps in Go (Key-Value Storage)")

	// ✅ 1. Basic Map Declaration
	ages := map[string]int{"Alice": 30, "Bob": 25}
	fmt.Println("Initial Map:", ages)

	// ✅ 2. Adding/Updating Values
	ages["Charlie"] = 40 // Add new key-value
	ages["Alice"] = 31   // Update existing key
	fmt.Println("Updated Map:", ages)

	// ✅ 3. Deleting a Key
	delete(ages, "Bob")
	fmt.Println("After Deletion:", ages)

	// ✅ 4. Checking If a Key Exists
	age, exists := ages["Alice"]
	if exists {
		fmt.Println("✅ Alice's age:", age)
	} else {
		fmt.Println("❌ Alice not found")
	}

	// ❓ Check for a key that doesn't exist
	age2, ok := ages["Jason"]              //Try to access a non-existent key
	fmt.Println("Jason exists?", ok)       // false
	fmt.Println("Jason age (zero):", age2) // 0 (default value of int)
	if ok {
		fmt.Printf("The age is %v\n", age) // If key exists
	} else {
		fmt.Println("Invalid Name") // If key doesn't exist
	}

	// ✅ 5. Iterating Over a Map
	fmt.Println("\n🌀 Iterating Over Map:")
	for key, value := range ages {
		fmt.Printf("Key: %s → Value: %d\n", key, value)
	}
}
