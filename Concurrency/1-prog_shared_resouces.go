package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	count := 0
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val := count
			time.Sleep(time.Millisecond)
			val = val + 1
			count = val
		}()
	}
	wg.Wait()
	fmt.Println("count:", count)
}
