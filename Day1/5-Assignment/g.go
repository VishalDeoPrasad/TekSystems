package main

import (
	"fmt"
)

func main() {
	numbers := []int{5, 20, 15, 1, 2, 3}
	fmt.Println("Before Reverse:", numbers)
	n := len(numbers)
	for i := 0; i<n/2; i++{
		numbers[i], numbers[n-i-1] = numbers[n-i-1], numbers[i]
	}
	fmt.Println("After Reverse:", numbers)

	
}
