// An example program demoing a worker pool concept

package main

import (
	"fmt"
	"strconv"
	"sync"
)

const (
	MAX_WORKERS = 3
	MAX_JOBS    = 100
)

func Worker(id int, jobs <-chan int, result chan<- string, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d online\n", id)
	for job := range jobs {
		fmt.Printf("Worker %d received job %d\n", id, job)
		// do something with the received job
		// currently just echoing it back as a string
		result <- strconv.Itoa(job)
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
	wg.Done()
}

func main() {
	jobs := make(chan int, MAX_JOBS)
	results := make(chan string)
	wg := sync.WaitGroup{}

	// create workers
	for i := 0; i < MAX_WORKERS; i++ {
		wg.Add(1)
		go Worker(i, jobs, results, &wg)
	}

	// create jobs
	for i := 0; i < MAX_JOBS; i++ {
		jobs <- i
	}
	close(jobs)

	// reading results
	for i := 0; i < MAX_JOBS; i++ {
		fmt.Printf("Received result : %s\n", <-results)
	}

	//bloacking main goroutine till workers return
	wg.Wait()

}
