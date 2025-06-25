package strings

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("🔤 Strings & Unicode in Go")

	// ✅ 1. Regular String (UTF-8 encoded by default)
	str := "résumé" // contains non-ASCII characters (é)

	// ✅ 2. Byte Length (each 'é' takes 2 bytes in UTF-8)
	fmt.Printf("Byte length: %d\n", len(str)) // Output: 8 (6 letters, but 2 of them are multi-byte)

	// ✅ 3. Iterate over bytes (not chars!)(not recommended for multi-byte chars)
	fmt.Println("🔸 Iterating over bytes:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("Byte %d: %v (%T)\n", i, str[i], str[i])
	}

	// ✅ 4. Convert string to rune slice for character-safe indexing
	runes := []rune(str)                          // Runes = Unicode code points
	fmt.Printf("\nRune length: %d\n", len(runes)) // Output: 6 (6 actual characters)

	// ✅ 5. Indexing a character using rune slice
	var indexed = runes[1]
	fmt.Printf("\nIndexed rune: %v, Type: %T\n", indexed, indexed) // prints 233 (U+00E9 for 'é')

	// ✅ 6. Range loop (automatically decodes UTF-8 into runes)
	fmt.Println("\n🔸 Iterating over runes (characters):")
	for i, r := range str {
		fmt.Printf("Index %d → Rune: %c → Code Point: U+%04X\n", i, r, r)
	}

	// ✅ 7. Show total rune length of the string
	fmt.Printf("\n✅ The length of '%s' in runes is %v\n", str, len(runes))

	// ✅ 8. String Building (efficient concat)
	var sb strings.Builder
	sb.WriteString("Hello")
	sb.WriteString(", ")
	sb.WriteString("world!")
	result := sb.String()
	fmt.Println("\n🔧 Built String:", result)

	// ✅ 9. Converting between string and []byte / []rune
	b := []byte("Go💙")         // UTF-8 encoded bytes
	r := []rune("Go💙")         // Unicode code points
	fmt.Println("\nBytes:", b) // May print [71 111 240 159 146 153]
	fmt.Println("Runes:", r)   // Prints code points: [71 111 128153]

	// ✅ 10. String from rune slice
	newStr := string([]rune{72, 101, 108, 108, 111}) // "Hello"
	fmt.Println("From runes:", newStr)
}
