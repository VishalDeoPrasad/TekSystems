package main

import (
	"fmt"
)

func main() {
	numbers := []int{5, 4, 70, 45, 65, 34, 20, 15, 1, 2, 3}
	even, odd := 0, 0

	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			even += 1
		} else {
			odd += 1
		}
	}
	fmt.Printf("Count of Even Number : %d\nCount of Odd Number : %d", even, odd)

}
