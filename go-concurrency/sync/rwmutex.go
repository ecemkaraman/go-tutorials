package sync

import (
	"fmt"
	"sync"
	"time"
)

var data int
var rwMutex sync.RWMutex

// Read function
func readData(id int) {
	rwMutex.RLock()
	fmt.Printf("📖 Reader %d sees value: %d\n", id, data)
	rwMutex.RUnlock()
}

// Write function
func writeData(id int, value int) {
	rwMutex.Lock()
	fmt.Printf("✍️ Writer %d updating value to: %d\n", id, value)
	data = value
	rwMutex.Unlock()
}

// Runs RWMutex Example
func RunRWMutex() {
	data = 10

	go readData(1)
	go readData(2)
	go writeData(1, 20)
	go readData(3)

	time.Sleep(1 * time.Second)
}

/*
🔹 Explanation:
- `rwMutex.RLock()`: Multiple Goroutines can **read** simultaneously.
- `rwMutex.Lock()`: Only one Goroutine can **write** at a time.
- More efficient than `Mutex` when **many reads, few writes**.
*/
