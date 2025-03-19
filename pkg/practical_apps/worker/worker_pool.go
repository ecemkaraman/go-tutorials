package worker

import (
	"fmt"
	"sync"
	"time"
)

// ✅ Worker function
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second)
		results <- job * 2 // Example processing
	}
}

// ✅ Run Worker Pool
func RunWorkerPool() {
	numWorkers := 3
	numJobs := 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Create workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for all workers
	wg.Wait()
	close(results)

	// Collect results
	for result := range results {
		fmt.Println("Result:", result)
	}
}

/*
🔹 Explanation:
- `go worker()`: Launches workers concurrently.
- `jobs <- j`: Sends jobs into channel.
- `results <- job * 2`: Stores processed data.
- `sync.WaitGroup`: Ensures all workers finish.
*/
