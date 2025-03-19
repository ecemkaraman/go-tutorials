package sync

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // Define WaitGroup

// Worker function
func worker(id int) {
	defer wg.Done() // Mark as done
	fmt.Printf("👷 Worker %d started\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("✅ Worker %d finished\n", id)
}

// Runs WaitGroup Example
func RunWaitGroup() {
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Add Goroutine
		go worker(i)
	}

	wg.Wait() // Wait for all Goroutines
	fmt.Println("All workers completed!")
}

/*
🔹 Explanation:
- `wg.Add(1)`: Adds Goroutine to WaitGroup.
- `defer wg.Done()`: Marks completion.
- `wg.Wait()`: Blocks execution until all Goroutines finish.
*/
