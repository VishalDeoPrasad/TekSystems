package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	even, odd := 0, 0

	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			even += numbers[i]
		} else {
			odd += numbers[i]
		}
	}
	fmt.Printf("Sum of Even Number : %d\nSum of Odd Number : %d", even, odd)

}
