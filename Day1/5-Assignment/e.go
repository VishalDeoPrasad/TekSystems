package main

import "fmt"

func main() {
	numbers := []int{5, 20, 15, 1, 2}
	min := numbers[0]
	for i:=1; i<len(numbers); i++{
		if numbers[i] < min{
			min = numbers[i]
		}
	}
	fmt.Println("Minimum Number =", min)
}
