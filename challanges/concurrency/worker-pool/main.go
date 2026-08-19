package main

import (
	"fmt"
	"time"
)

// Job represents a unit of work.
type Job struct {
	ID int
}

// Result represents the outcome of a job.
type Result struct {
	JobID  int
	Worker int
}

// worker is a goroutine. It reads jobs, processes them, and sends results.
// Note the channel directions: jobs is read-only (<-chan), results is write-only (chan<-)
func worker(id int, jobs <-chan Job, results chan<- Result) {
	// This loop blocks and waits for a job.
	// It exits automatically once the jobs channel is closed and emptied.
	for job := range jobs {
		fmt.Printf("Worker %d: Started job %d\n", id, job.ID)

		time.Sleep(100 * time.Millisecond) // Simulating work

		fmt.Printf("Worker %d: Finished job %d\n", id, job.ID)

		// Send the outcome back to the main goroutine
		results <- Result{JobID: job.ID, Worker: id}
	}
}

func main() {
	numJobs := 5

	// Create buffered channels so we don't block sender goroutines immediately
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// 1. Spawn 3 workers (goroutines)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 2. Send the 5 jobs into the queue
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j}
	}

	// 3. CLOSE the jobs channel.
	// This tells the workers: "No more jobs are coming, finish what you have and exit."
	close(jobs)

	// 4. Collect results from the results channel
	for a := 1; a <= numJobs; a++ {
		res := <-results // Blocks until a worker sends a result
		fmt.Printf("Main: Job %d processed by Worker %d\n", res.JobID, res.Worker)
	}
}
