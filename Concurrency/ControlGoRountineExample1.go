package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// Create a channel with a buffer size of 5
	// This will limit the number of concurrent goroutines to 5
	semaphore := make(chan struct{}, 5)

	for i := 1; i <= 100; i++ {
		wg.Add(1)

		// Start a goroutine when there's room in the semaphore
		semaphore <- struct{}{}
		go func(num int) {
			defer wg.Done()
			defer func() { 
				<-semaphore 
				}() // Release the semaphore when the goroutine completes
			fmt.Println(num)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
