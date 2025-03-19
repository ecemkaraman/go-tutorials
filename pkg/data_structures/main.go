package main

import (
	"fmt"

	"github.com/ecemkaraman/go-tutorials/go-data-structures/arrays"
	"github.com/ecemkaraman/go-tutorials/go-data-structures/maps"
	"github.com/ecemkaraman/go-tutorials/go-data-structures/slices"
	"github.com/ecemkaraman/go-tutorials/go-data-structures/structs"
)

func main() {
	fmt.Println("🚀 Go Data Structures Practice 🚀")

	// Arrays
	arrays.PracticeArrays()

	// Slices
	slices.PracticeSlices()

	// Maps
	maps.PracticeMaps()

	// Structs
	structs.PracticeStructs()
}

/*
✅ This is the central hub of the program!
1️⃣ Declares the main package → Every Go program must have a main package as the entry point.
2️⃣ Imports all custom packages (arrays, maps, slices, structs) → Allows calling their functions.
3️⃣ Defines main() → Prints a banner and calls each package's function.
4️⃣ Function calls (PracticeArrays(), PracticeSlices(), etc.) → Executes the learning functions for each data structure.
*/
