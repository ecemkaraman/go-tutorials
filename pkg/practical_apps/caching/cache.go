package caching

import (
	"fmt"
	"sync"
	"time"
)

// ✅ Cache Struct with RWMutex for concurrency safety
type Cache struct {
	data map[string]string
	mu   sync.RWMutex
}

// ✅ Creates a new cache instance
func NewCache() *Cache {
	return &Cache{data: make(map[string]string)}
}

// ✅ Set cache key-value
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// ✅ Get cache value
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, exists := c.data[key]
	return val, exists
}

// ✅ Run Cache Example
func RunCache() {
	cache := NewCache()
	cache.Set("username", "Ecem")

	time.Sleep(1 * time.Second)

	if val, found := cache.Get("username"); found {
		fmt.Println("Cached value:", val)
	} else {
		fmt.Println("Key not found")
	}
}

/*
🔹 Explanation:
- `sync.RWMutex`: Controls read/write access.
- `Set()`: Locks while updating.
- `Get()`: Uses read lock to allow concurrent reads.
*/
