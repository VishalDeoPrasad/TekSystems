package main

import (
	"fmt"
	"sync"
)

func sender(ch chan int, val int, wg *sync.WaitGroup) {
	wg.Done()
	ch <- val
}

func receiver(ch chan int, wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println(<-ch)

}
func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int) //blank channel
	wg.Add(5)
	go sender(ch, 10) //fill up channel
	go sender(ch, 20)
	go sender(ch, 30)
	go receiver(ch) //receiving from channel
	go receiver(ch)

	// time.Sleep(1 * time.Millisecond)
	wg.Wait()

	fmt.Println("Successfully Printed:")

}
