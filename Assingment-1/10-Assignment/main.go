package main

import "fmt"

func sum(numbers ...int) int {
	var total int
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}
func main() {
	nums := []int{1, 2, 3, 4, 4, 5}
	fmt.Println("sum =", sum(nums...))
}
