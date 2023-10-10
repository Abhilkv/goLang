package advanced

import (
    "fmt"
    "net/http"
    "sync"
)

func fetchURL(url string, wg *sync.WaitGroup) {
    defer wg.Done()			// Decrement the WaitGroup counter when the function finishes

    // Make an HTTP GET request
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Error fetching %s: %v\n", url, err)
        return
    }
    defer resp.Body.Close()   //resp.Body.Close() call, you ensure that the response body will be closed automatically when your function exits
								//  regardless of whether the function returns normally or due to an error. This is crucial for proper resource cleanup.

    fmt.Printf("Fetched %s, Status Code: %d\n", url, resp.StatusCode)
}

func ConcurrencyAGoroutines() {
    urls := []string{
        "https://www.example.com",
        "https://www.google.com",
        "https://www.openai.com",
    }

    var wg sync.WaitGroup

    for _, url := range urls {
        wg.Add(1) // Increment the WaitGroup counter for each URL
        go fetchURL(url, &wg)
    }

    // Wait for all goroutines to finish
    wg.Wait()

    fmt.Println("All downloads completed")
}
