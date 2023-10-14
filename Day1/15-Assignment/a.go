package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
}

func receiver(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Print(<-ch)
	}
	fmt.Println()

}
func main() {
	ch:=make(chan int) //blank channel
	go sender(ch) //fill up channel
	go receiver(ch) //receiving from channel
	time.Sleep(3*time.Second) 
	fmt.Println("Successfully Printed:")

}