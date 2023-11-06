package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	c := make(chan int)
	list := []int{1, 2, 3, 4, 5}
	go func(){ //outer go rountine wait until inner go routine finished his task
		for _, v := range list {
			wg.Add(1)
			go func(v int) {
				defer wg.Done()
				result := Square(v)
				c <- result
			}(v)
		}
		wg.Wait()
		close(c)
	}()

	for v := range c {
		Process(v)
	}
}

func Square(i int) int {
	return i * i
}

func Process(n int) {
	ans := float64(n) / 2
	fmt.Println(ans)
}
