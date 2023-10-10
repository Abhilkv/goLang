package advanced

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Decrement the WaitGroup counter when done

    fmt.Printf("Worker %d is working\n", id)
    time.Sleep(time.Second) // Simulate some work
    fmt.Printf("Worker %d has finished\n", id)
}

func WorkerGroup() {
    var wg sync.WaitGroup // Create a Wait Group

    for i := 1; i <= 3; i++ {
        wg.Add(1)        // Increment the Wait Group counter for each worker
        go worker(i, &wg) // Start a worker goroutine
    }

    // Wait for all workers to finish
    wg.Wait()

    fmt.Println("All workers have completed")
}
