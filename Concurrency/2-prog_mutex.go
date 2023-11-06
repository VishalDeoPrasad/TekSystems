package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	count := 0
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			val := count
			time.Sleep(time.Millisecond)
			val = val + 1
			count = val
			mu.Unlock()
		}()
	}
	wg.Wait() //wait until all go rountine finish the execution
	fmt.Println("count:", count)
}
