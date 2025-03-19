package testing

import (
	"testing"
)

// ✅ Function to benchmark
func addNumbers(a, b int) int {
	return a + b
}

// ✅ Benchmark Test
func BenchmarkAddNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addNumbers(10, 20)
	}
}

/*
🔹 Run with: `go test -bench .`
- `b.N`: Number of iterations (automatically set).
- `for i := 0; i < b.N; i++ {}`: Loops to measure execution time.
*/
