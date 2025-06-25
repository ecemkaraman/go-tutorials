package sync

import "fmt"

func Run() {
	fmt.Println("=== Mutex Demo ===")
	RunMutex()

	fmt.Println("\n=== RWMutex Demo ===")
	RunRWMutex()

	fmt.Println("\n=== WaitGroup Demo ===")
	RunWaitGroup()
}
