package channels

import "fmt"

// Runs Buffered Channel Example
func RunBufferedChannel() {
	buffered := make(chan int, 3) // Buffered channel with capacity 3

	buffered <- 10
	buffered <- 20
	buffered <- 30

	fmt.Println("📦 Buffered Received:", <-buffered, <-buffered, <-buffered)
}

/*
🔹 Explanation:
- Buffered channels allow multiple values before blocking.
- `make(chan int, 3)`: Creates a channel that holds 3 values.
- Values are received in FIFO order.
*/
