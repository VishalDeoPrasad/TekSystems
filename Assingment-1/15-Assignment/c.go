package main

import (
	"fmt"
	"time"
)

func sender(ch chan int, val int) {
	ch <- val
}

func receiver(ch chan int) {
	fmt.Println(<-ch)

}
func main() {
	ch := make(chan int) //blank channel
	go sender(ch, 10)    //fill up channel
	go sender(ch, 20)
	go sender(ch, 30)
	go receiver(ch) //receiving from channel
	go receiver(ch)

	time.Sleep(1 * time.Millisecond)

	fmt.Println("Successfully Printed:")

}
