package main

import "fmt"

func main() {
	fmt.Println("\n🔹 Control Flow Practice in Go")

	// ✅ If-Else
	age := 17
	if age > 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	// ✅ If with short declaration
	if x := 5; x > 3 {
		fmt.Println("x is greater than 3")
	}

	// ✅ Switch-Case
	day := "Friday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Weekend soon!")
	default:
		fmt.Println("Regular day")
	}

	// ✅ Expression-based Switch
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	default:
		fmt.Println("Grade: C")
	}

	// ✅ Standard For Loop
	fmt.Println("\nStandard For Loop:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// ✅ While-style Loop
	fmt.Println("\nWhile-style Loop:")
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	// ✅ Infinite Loop (exits after 1 iteration)-> for{}
	fmt.Println("\nInfinite Loop (with break):")
	for {
		fmt.Println("Running forever (not really)")
		break
	}

	// ✅ Range Loop - Slice
	nums := []int{10, 20, 30}
	fmt.Println("\nRange Loop - Slice:")
	for index, value := range nums {
		fmt.Printf("Index: %d → Value: %d\n", index, value)
	}

	// ✅ Range Loop - Map
	colors := map[string]string{"R": "Red", "G": "Green", "B": "Blue"}
	fmt.Println("\nRange Loop - Map:")
	for key, value := range colors {
		fmt.Printf("Key: %s → Value: %s\n", key, value)
	}
}

/*
🧠 Notes:
- Go does NOT have a ternary operator.
- Use switch for cleaner if-else replacement.
- `for` is the only loop keyword (acts as for/while/infinite).
*/
