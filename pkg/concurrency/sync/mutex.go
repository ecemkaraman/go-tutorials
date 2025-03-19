package sync

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mutex sync.Mutex // Mutex to avoid race conditions

// Safe Increment function
func safeIncrement(id int) {
	mutex.Lock()
	count++
	fmt.Printf("🔢 Goroutine %d incremented count to %d\n", id, count)
	mutex.Unlock()
}

// Runs Mutex Example
func RunMutex() {
	for i := 1; i <= 5; i++ {
		go safeIncrement(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("✅ Final Count:", count)
}

/*
🔹 Explanation:
- `mutex.Lock()`: Ensures only one Goroutine modifies `count` at a time.
- `mutex.Unlock()`: Releases lock for next Goroutine.
- Prevents race conditions when multiple Goroutines modify shared data.
*/
