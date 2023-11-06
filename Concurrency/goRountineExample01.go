package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	wg := new(sync.WaitGroup)
	go func() {
		for i := 1; i <= 10; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				c <- i
			}(i)
		}
		wg.Wait()
		close(c)
	}()
	for n := range c {
		fmt.Print(n, ", ")
	}

}
