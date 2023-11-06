package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	bc := make(chan struct{}, 5)
	wg := new(sync.WaitGroup)
	c := make(chan int)
	list := []int{}

	for i := 0; i < 1000; i++ {
		list = append(list, i)
	}
	go func() {
		for _, v := range list {
			wg.Add(1)
			bc <- struct{}{}
			go func(v int) {
				defer wg.Done()
				defer func() {
					<-bc
				}()
				result := Square(v)
				c <- result

				fmt.Println("Num GR:", runtime.NumGoroutine())
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
	v := n / 2
	fmt.Printf("processed result for %v is %v\n", n, v)
}
