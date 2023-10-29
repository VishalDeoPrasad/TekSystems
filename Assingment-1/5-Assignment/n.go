package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30, 40, 50, 70, 60}
	key := 60
	for i := 0; i<len(slice); i++{
		if slice[i] == key{
			fmt.Printf("%d -Found at index:%d", key, i)
			return
		}
	}
	fmt.Println("Element Not Found!")
		
	
}
