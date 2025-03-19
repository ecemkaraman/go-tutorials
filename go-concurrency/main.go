package main

import (
	"fmt"

	"github.com/ecemkaraman/go-tutorials/go-concurrency/channels"
	"github.com/ecemkaraman/go-tutorials/go-concurrency/goroutines"
	"github.com/ecemkaraman/go-tutorials/go-concurrency/sync"
)

func main() {
	fmt.Println("🚀 Go Concurrency Practice")

	// ✅ Goroutines
	fmt.Println("\n🔹 Running Goroutines Example")
	goroutines.RunGoroutines()

	// ✅ Unbuffered Channels
	fmt.Println("\n🔹 Running Unbuffered Channel Example")
	channels.RunUnbufferedChannel()

	// ✅ Buffered Channels
	fmt.Println("\n🔹 Running Buffered Channel Example")
	channels.RunBufferedChannel()

	// ✅ WaitGroups
	fmt.Println("\n🔹 Running WaitGroup Example")
	sync.RunWaitGroup()

	// ✅ Mutex
	fmt.Println("\n🔹 Running Mutex Example")
	sync.RunMutex()

	// ✅ RWMutex
	fmt.Println("\n🔹 Running RWMutex Example")
	sync.RunRWMutex()
}
