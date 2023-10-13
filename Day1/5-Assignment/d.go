package main

import "fmt"

func main() {
	numbers := []int{5, 10, 15, 20, 25}
	var sum int
	for i := range numbers {
		sum += numbers[i]
	}
	fmt.Println("Sum =", sum)
	fmt.Println("Average =",sum/len(numbers))
}
