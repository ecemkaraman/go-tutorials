package concurrency

import (
	"context"
	"fmt"
	"time"
)

// Runs Context Example
func RunContextExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch := make(chan string)

	// ✅ Start Goroutine with context
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("⏹️ Goroutine Stopped: Context Timeout")
				return
			default:
				ch <- "Processing..."
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	// ✅ Read from channel
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		case <-ctx.Done():
			fmt.Println("✅ Done Processing")
			return
		}
	}
}

/*
🔹 Explanation:
- `context.WithTimeout()`: Cancels Goroutine after timeout.
- `ctx.Done()`: Signals completion.
*/
