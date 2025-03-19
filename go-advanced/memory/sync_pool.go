package memory

import (
	"fmt"
	"sync"
)

// Runs sync.Pool Example
func RunSyncPoolExample() {
	var pool = sync.Pool{
		New: func() interface{} {
			return "New Object"
		},
	}

	// ✅ Get object from pool
	obj := pool.Get()
	fmt.Println("✅ Object Retrieved:", obj)

	// ✅ Put object back in pool
	pool.Put("Reused Object")

	// ✅ Get reused object
	obj = pool.Get()
	fmt.Println("✅ Object Reused:", obj)
}

/*
🔹 Explanation:
- `sync.Pool{New: func()}`: Creates object pool.
- `.Get()`: Retrieves object.
- `.Put()`: Returns object for reuse.
*/
