package cli

import (
	"flag"
	"fmt"
)

// ✅ CLI Tool using `flag`
func Run() {
	name := flag.String("name", "User", "Your name")
	age := flag.Int("age", 25, "Your age")
	flag.Parse()
	fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)
}

/*
🔹 Explanation:
- `flag.String()`, `flag.Int()`: Define CLI flags.
- `flag.Parse()`: Parse command-line arguments.
- `go run main.go -name="Ecem" -age=28`
*/
