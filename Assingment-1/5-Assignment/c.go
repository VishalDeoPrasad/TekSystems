package main

import "fmt"

func main() {
	numbers := []int{}
	fmt.Println(numbers)
	numbers = append(numbers, 5)
	fmt.Println(numbers)
	numbers = append(numbers, 10)
	fmt.Println(numbers)
	numbers = append(numbers, 15)
	fmt.Println(numbers)
	numbers = append(numbers, 20)
	fmt.Println(numbers)
	numbers = append(numbers, 25)
	fmt.Println(numbers)
	fmt.Printf("length=%d capacity=%d\n", len(numbers), cap(numbers))
	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println(numbers)
}
