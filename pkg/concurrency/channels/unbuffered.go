package channels

import "fmt"

// Function that sends data to an unbuffered channel
func sendData(channel chan string) {
	channel <- "📦 Data Sent from Goroutine"
}

// Runs Unbuffered Channel Example
func RunUnbufferedChannel() {
	ch := make(chan string) // Unbuffered channel
	go sendData(ch)         // Start Goroutine

	msg := <-ch // Receive data
	fmt.Println("✅ Received:", msg)
}

/*
🔹 Explanation:
- `make(chan string)`: Creates an unbuffered channel.
- `channel <- "data"`: Sends data.
- `<-channel`: Blocks execution until data is received.
*/
